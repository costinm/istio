// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gcpmonitoring

import (
	"context"
	"errors"
	"io/ioutil"
	"strings"
	"sync"
	"time"

	ocprom "contrib.go.opencensus.io/exporter/prometheus"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/monitoredresource"
	"go.opencensus.io/stats/view"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/oauth"

	"istio.io/istio/pilot/pkg/security/model"
	"istio.io/istio/pkg/bootstrap/platform"
	"istio.io/istio/security/pkg/stsservice/tokenmanager"
	"istio.io/pkg/env"
	"istio.io/pkg/log"
	"istio.io/pkg/version"
)

const (
	authScope = "https://www.googleapis.com/auth/cloud-platform"
)

var (
	trustDomain = ""

	// env vars that are used to set labels/monitored resource of control plane metrics, which are added by ASM profile.
	// `ASM_CONTROL_PLANE` prefix is added to avoid collision with env var from OSS code base.
	// TODO(bianpengyuan): POD_NAME and POD_NAMESPACE are also defined in pilot bootstrap package. Remove these after refactoring pilot code
	// to make env vars accessible any where in the code base.
	podNameVar      = env.RegisterStringVar("ASM_CONTROL_PLANE_POD_NAME", "unknown", "istiod pod name, specified by GCP installation profile.")
	podNamespaceVar = env.RegisterStringVar("ASM_CONTROL_PLANE_POD_NAMESPACE", "istio-system", "istiod pod namespace, specified by GCP installation profile.")
	meshIDVar       = env.RegisterStringVar("ASM_CONTROL_PLANE_MESH_ID", "", "mesh id, specified by GCP installation profile.")
	cloudRunServiceVar  = env.RegisterStringVar("K_SERVICE", "", "cloud run service name")
	cloudRunRevisionVar = env.RegisterStringVar("K_REVISION", "", "cloud run revision")
	cloudRunConfigVar   = env.RegisterStringVar("K_CONFIGURATION", "", "name of cloud run configuration")
	)
type cloudRunRevision struct {
	service       string
		revision      string
		location      string
		configuration string
		projectID     string
}
// MonitoredResource returns the resource type and resource labels.
// Implements monitoredresource.Interface
func (c *cloudRunRevision) MonitoredResource() (resType string, labels map[string]string) {
		labels = map[string]string{
				"project_id":         c.projectID,
				"location":           c.location,
				"service_name":       c.service,
				"revision_name":      c.revision,
				"configuration_name": c.configuration,
			}
		return "cloud_run_revision", labels
	}


// ASMExporter is a stats exporter used for ASM control plane metrics.
// It wraps a prometheus exporter and a stackdriver exporter, and exports two types of views.
type ASMExporter struct {
	PromExporter *ocprom.Exporter
	sdExporter   *stackdriver.Exporter
}

// SetTrustDomain sets GCP trust domain, which is used to fetch GCP metrics.
// Use this function instead of passing trust domain string around to avoid conflicting with OSS changes.
func SetTrustDomain(td string) {
	trustDomain = td
}

// NewASMExporter creates an ASM opencensus exporter.
func NewASMExporter(pe *ocprom.Exporter) (*ASMExporter, error) {
	if !enableSDVar.Get() {
		// Stackdriver monitoring is not enabled, return early with only prometheus exporter initialized.
		return &ASMExporter{
			PromExporter: pe,
		}, nil
	}
	labels := &stackdriver.Labels{}
	labels.Set("mesh_uid", meshIDVar.Get(), "ID for Mesh")
	labels.Set("revision", version.Info.Version, "Control plane revision")
	gcpMetadata := platform.NewGCP().Metadata()
	clientOptions := []option.ClientOption{}


	var mr monitoredresource.Interface
		if svc := cloudRunServiceVar.Get(); svc != "" && false {
				mr = &cloudRunRevision{
						service:       svc,
						location:      gcpMetadata[platform.GCPLocation],
						projectID:     gcpMetadata[platform.GCPProject],
						revision:      cloudRunRevisionVar.Get(),
						configuration: cloudRunConfigVar.Get(),
					}
			} else {
				pod := podNameVar.Get()
				ns := podNamespaceVar.Get()
				mr = &monitoredresource.GKEContainer{
						ProjectID:                  gcpMetadata[platform.GCPProject],
						ClusterName:                gcpMetadata[platform.GCPCluster],
						Zone:                       gcpMetadata[platform.GCPLocation],
						NamespaceID:                ns,
						PodID:                      pod,
						ContainerName:              "discovery",
						LoggingMonitoringV2Enabled: true,
					}
			}

	if strings.HasSuffix(trustDomain, "svc.id.goog") {
		// Workload identity is enabled and P4SA access token is used.
		if subjectToken, err := ioutil.ReadFile(model.K8sSATrustworthyJwtFileName); err == nil {
			ts := tokenmanager.NewTokenSource(trustDomain, string(subjectToken), authScope)
			clientOptions = append(clientOptions, option.WithTokenSource(ts), option.WithQuotaProject(gcpMetadata[platform.GCPProject]))
			// Set up goroutine to read token file periodically and refresh subject token with new expiry.
			go func() {
				for range time.Tick(5 * time.Minute) {
					if subjectToken, err := ioutil.ReadFile(model.K8sSATrustworthyJwtFileName); err == nil {
						ts.RefreshSubjectToken(string(subjectToken))
					} else {
						log.Debugf("Cannot refresh subject token for sts token source: %v", err)
					}
				}
			}()
		} else {
			log.Errorf("Cannot read third party jwt token file, using default credentials %v", err)
			gcred, _ := oauth.NewApplicationDefault(context.Background())
			clientOptions = append(clientOptions, option.WithGRPCDialOption(grpc.WithPerRPCCredentials(gcred)), option.WithQuotaProject(gcpMetadata[platform.GCPProject]))
		}
	}
	cnt := 0
	se, err := stackdriver.NewExporter(stackdriver.Options{
		MetricPrefix:            "istio.io/control",
		MonitoringClientOptions: clientOptions,
		GetMetricType: func(view *view.View) string {
			return "istio.io/control/" + view.Name
		},
		MonitoredResource: mr,
		DefaultMonitoringLabels: labels,
		ReportingInterval:       60 * time.Second,
		OnError: func(err error) {
			if strings.Contains(err.Error(), "One or more TimeSeries could not be written") {
				if cnt % 100 == 0 {
					log.Warnf("Stackdriver error %v", err)
				}
				cnt++
				return
			}
			log.Warnf("Stackdriver error %v", err)
		},
	})

	if err != nil {
		return nil, errors.New("fail to initialize Stackdriver exporter")
	}

	return &ASMExporter{
		PromExporter: pe,
		sdExporter:   se,
	}, nil
}

// ExportView exports all views collected by control plane process.
// This function distinguished views for Stackdriver and views for Prometheus and exporting them separately.
func (e *ASMExporter) ExportView(vd *view.Data) {
	if _, ok := viewMap[vd.View.Name]; ok && e.sdExporter != nil {
		// This indicates that this is a stackdriver view
		e.sdExporter.ExportView(vd)
	} else {
		e.PromExporter.ExportView(vd)
	}
}

// TestExporter is used for GCP monitoring test.
type TestExporter struct {
	sync.Mutex

	Rows        map[string][]*view.Row
	invalidTags bool
}

// ExportView exports test views.
func (t *TestExporter) ExportView(d *view.Data) {
	t.Lock()
	defer t.Unlock()
	for _, tk := range d.View.TagKeys {
		if len(tk.Name()) < 1 {
			t.invalidTags = true
		}
	}
	t.Rows[d.View.Name] = append(t.Rows[d.View.Name], d.Rows...)
}
