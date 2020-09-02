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

package bugreport

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"

	config2 "istio.io/istio/tools/bug-report/pkg/config"
	"istio.io/pkg/log"
)

var (
	startTime, endTime, configFile, tempDir string
	included, excluded                      []string
	commandTimeout, since                   time.Duration
	gConfig                                 = &config2.BugReportConfig{}
)

func addFlags(cmd *cobra.Command, args *config2.BugReportConfig) {
	// k8s client config
	cmd.PersistentFlags().StringVarP(&args.KubeConfigPath, "kubeconfig", "c", "",
		"Path to kube config.")
	cmd.PersistentFlags().StringVar(&args.Context, "context", "",
		"Name of the kubeconfig Context to use.")

	// input config
	cmd.PersistentFlags().StringVarP(&configFile, "filename", "f", "",
		"Path to a file containing configuration in YAML format.")

	// dry run
	cmd.PersistentFlags().BoolVarP(&args.DryRun, "dry-run", "", false,
		"Only log commands that would be run, don't fetch or write.")

	// full secrets
	cmd.PersistentFlags().BoolVarP(&args.FullSecrets, "full-secrets", "", false,
		"If set, secret contents are included in output.")

	// istio namespaces
	cmd.PersistentFlags().StringVar(&args.IstioNamespace, "istio-namespace", bugReportDefaultIstioNamespace,
		"List of comma-separated namespaces where Istio control planes are installed.")

	// timeouts and max sizes
	cmd.PersistentFlags().DurationVar(&commandTimeout, "timeout", bugReportDefaultTimeout,
		"Maximum amount of time to spend fetching logs. When timeout is reached "+
			"only the logs captured so far are saved to the archive.")
	cmd.PersistentFlags().Int32Var(&args.MaxArchiveSizeMb, "max-size", bugReportDefaultMaxSizeMb,
		"Maximum size of the compressed archive in Mb. Logs are prioritized"+
			"according to importance heuristics.")

	// include / exclude specs
	cmd.PersistentFlags().StringSliceVar(&included, "include", bugReportDefaultInclude,
		"Spec for which pods' proxy logs to include in the archive. See 'help' for examples.")
	cmd.PersistentFlags().StringSliceVar(&excluded, "exclude", bugReportDefaultExclude,
		"Spec for which pods' proxy logs to exclude from the archive, after the include spec "+
			"is processed. See 'help' for examples.")

	// log time ranges
	cmd.PersistentFlags().StringVar(&startTime, "start-time", "",
		"Start time for the range of log entries to include in the archive. "+
			"Default is the infinite past. If set, Since must be unset.")
	cmd.PersistentFlags().StringVar(&endTime, "end-time", "",
		"End time for the range of log entries to include in the archive. Default is now.")
	cmd.PersistentFlags().DurationVar(&since, "duration", 0,
		"How far to go back in time from end-time for log entries to include in the archive. "+
			"Default is infinity. If set, start-time must be unset.")

	// log error control
	cmd.PersistentFlags().StringSliceVar(&args.CriticalErrors, "critical-errs", nil,
		"List of comma separated glob patters to match against log error strings. "+
			"If any pattern matches an error in the log, the logs is given the highest priority for archive inclusion.")
	cmd.PersistentFlags().StringSliceVar(&args.IgnoredErrors, "ignore-errs", nil,
		"List of comma separated glob patters to match against log error strings. "+
			"Any error matching these patters is ignored when calculating the log importance heuristic.")

	// archive and upload control
	cmd.PersistentFlags().StringVar(&args.GCSURL, "gcs-url", "",
		"URL of the GCS bucket where the archive is uploaded.")
	cmd.PersistentFlags().BoolVar(&args.UploadToGCS, "upload", false,
		"Upload archive to GCS bucket. If gcs-url is unset, a new bucket is created.")

	// output/working dir
	cmd.PersistentFlags().StringVar(&tempDir, "dir", "",
		"Set a specific directory for temporary artifact storage.")
}

func parseConfig() (*config2.BugReportConfig, error) {
	if err := parseTimes(gConfig, startTime, endTime); err != nil {
		log.Fatal(err.Error())
	}
	for _, s := range included {
		ss := &config2.SelectionSpec{}
		if err := ss.UnmarshalJSON([]byte(s)); err != nil {
			return nil, err
		}
		gConfig.Include = append(gConfig.Include, ss)
	}
	for _, s := range excluded {
		ss := &config2.SelectionSpec{}
		if err := ss.UnmarshalJSON([]byte(s)); err != nil {
			return nil, err
		}
		gConfig.Exclude = append(gConfig.Exclude, ss)
	}

	fileConfig := &config2.BugReportConfig{}
	if configFile != "" {
		b, err := ioutil.ReadFile(configFile)
		if err != nil {
			return nil, err
		}
		if err := yaml.Unmarshal(b, fileConfig); err != nil {
			return nil, err
		}
		gConfig, err = overlayConfig(gConfig, fileConfig)
		if err != nil {
			return nil, err
		}
	}

	return gConfig, nil
}

func parseTimes(config *config2.BugReportConfig, startTime, endTime string) error {
	config.EndTime = time.Now()
	if endTime != "" {
		var err error
		config.EndTime, err = time.Parse(time.RFC3339, endTime)
		if err != nil {
			return fmt.Errorf("bad format for end-time: %s, expect RFC3339 e.g. %s", endTime, time.RFC3339)
		}
	}
	if config.Since != 0 {
		if startTime != "" {
			return fmt.Errorf("only one --start-time or --Since may be set")
		}
		config.StartTime = config.EndTime.Add(-1 * time.Duration(config.Since))
	} else {
		var err error
		if startTime == "" {
			config.StartTime = time.Time{}
		} else {
			config.StartTime, err = time.Parse(time.RFC3339, startTime)
			if err != nil {
				return fmt.Errorf("bad format for start-time: %s, expect RFC3339 e.g. %s", startTime, time.RFC3339)
			}
		}
	}
	return nil
}

func overlayConfig(base, overlay *config2.BugReportConfig) (*config2.BugReportConfig, error) {
	bj, err := json.Marshal(base)
	if err != nil {
		return nil, err
	}
	oj, err := json.Marshal(overlay)
	if err != nil {
		return nil, err
	}

	mj, err := jsonpatch.MergePatch(bj, oj)
	if err != nil {
		return nil, fmt.Errorf("json merge error (%s) for base object: \n%s\n override object: \n%s", err, bj, oj)
	}

	out := &config2.BugReportConfig{}
	err = json.Unmarshal(mj, out)
	return out, err
}
