// Copyright 2020 Istio Authors
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

package apigen_test

import (
	"testing"
	"time"

	xdsapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/gogo/protobuf/proto"

	"istio.io/istio/pilot/pkg/config/memory"
	"istio.io/istio/pilot/pkg/model"
	"istio.io/istio/pilot/pkg/networking/apigen"
	envoyv2 "istio.io/istio/pilot/pkg/proxy/envoy/v2"
	"istio.io/istio/pilot/pkg/proxy/envoy/xds"
	"istio.io/istio/pkg/adsc"
	"istio.io/istio/pkg/config/schema/collections"

	_ "google.golang.org/grpc/xds/experimental" // To install the xds resolvers and balancers.
)

var (
	istiodDNSAddr = "127.0.0.1:14053"
	agentDNSAddr  = "127.0.0.1:14054"

	grpcAddr         = "127.0.0.1:14056"
	grpcUpstreamAddr = grpcAddr
	// Address the tests are connecting to - normally the mock in-process server
	// Can be changed to point to a real server, so tests validate a real deployment.
	//	grpcUpstreamAddr = "127.0.0.1:15010"

	// Address of the Istiod gRPC service, used in tests.
	istiodSvcAddr = "istiod.istio-system.svc.cluster.local:14056"
)

// Creates an in-process discovery server, using the same code as Istiod, but
// backed by an in-memory config and endpoint store.
func initDS() *xds.Server {
	ds := xds.NewXDS()

	sd := ds.DiscoveryServer.MemRegistry
	sd.AddHTTPService("fortio1.fortio.svc.cluster.local", "10.10.10.1", 8081)
	sd.SetEndpoints("fortio1.fortio.svc.cluster.local", "", []*model.IstioEndpoint{
		{
			Address:         "127.0.0.1",
			EndpointPort:    uint32(14056),
			ServicePortName: "http-main",
		},
	})
	return ds
}

// Test using resolving DNS over GRPC. This uses XDS protocol, and Listener resources
// to represent the names. The protocol is based on GRPC resolution of XDS resources.
func TestAPIGen(t *testing.T) {
	ds := initDS()
	ds.DiscoveryServer.Generators["api"] = &apigen.ApiGenerator{}
	epGen := &envoyv2.EdsGenerator{ds.DiscoveryServer}
	ds.DiscoveryServer.Generators["api/"+envoyv2.EndpointType] = epGen

	err := ds.StartGRPC(grpcAddr)
	if err != nil {
		t.Fatal(err)
	}
	defer ds.GRPCListener.Close()

	// Verify we can receive the DNS cluster IPs using XDS
	t.Run("adsc", func(t *testing.T) {
		adscConn, err := adsc.Dial(grpcUpstreamAddr, "", &adsc.Config{
			IP: "1.2.3.4",
			Meta: model.NodeMetadata{
				Generator: "api",
			}.ToStruct(),
		})
		if err != nil {
			t.Fatal("Error connecting ", err)
		}
		store := memory.Make(collections.Pilot)

		configController := memory.NewController(store)
		adscConn.Store = model.MakeIstioStore(configController)

		adscConn.Send(&xdsapi.DiscoveryRequest{
			TypeUrl: adsc.ListenerType,
		})

		adscConn.WatchConfig()

		data, err := adscConn.WaitVersion(10*time.Second, adsc.ListenerType, "")
		if err != nil {
			t.Fatal("Failed to receive lds", err)
		}

		for _, rs := range data.Resources {
			l := &xdsapi.Listener{}
			err = proto.Unmarshal(rs.Value, l)
			if err != nil {
				t.Fatal("Unmarshall error ", err)
			}

			t.Log("LDS: ", l)
		}
		data, err = adscConn.WaitVersion(10*time.Second, collections.IstioNetworkingV1Alpha3Serviceentries.Resource().GroupVersionKind().String(), "")
		if err != nil {
			t.Fatal("Failed to receive lds", err)
		}

		ses := adscConn.Store.ServiceEntries()
		for _, se := range ses {
			t.Log(se)
		}
		sec, _ := adscConn.Store.List(collections.IstioNetworkingV1Alpha3Envoyfilters.Resource().GroupVersionKind(), "")
		for _, se := range sec {
			t.Log(se)
		}

	})

	t.Run("adsc-gen1", func(t *testing.T) {
		adscConn, err := adsc.Dial(grpcUpstreamAddr, "", &adsc.Config{
			IP:   "1.2.3.5",
			Meta: model.NodeMetadata{}.ToStruct(),
		})
		if err != nil {
			t.Fatal("Error connecting ", err)
		}

		adscConn.Send(&xdsapi.DiscoveryRequest{
			TypeUrl: adsc.ClusterType,
		})

		got, err := adscConn.Wait(10*time.Second, "cds")
		if err != nil {
			t.Fatal("Failed to receive lds", err)
		}
		if len(got) == 0 {
			t.Fatal("No LDS response")
		}
		data := adscConn.Received[adsc.ClusterType]
		for _, rs := range data.Resources {
			l := &xdsapi.Cluster{}
			err = proto.Unmarshal(rs.Value, l)
			if err != nil {
				t.Fatal("Unmarshall error ", err)
			}

			t.Log("CDS: ", l)
		}
	})
}
