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

package ambient

import (
	"net/netip"
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	networking "istio.io/api/networking/v1alpha3"
	networkingclient "istio.io/client-go/pkg/apis/networking/v1alpha3"
	"istio.io/istio/pilot/pkg/model"
	"istio.io/istio/pkg/kube/krt"
	"istio.io/istio/pkg/kube/krt/krttest"
	"istio.io/istio/pkg/slices"
	"istio.io/istio/pkg/test/util/assert"
	"istio.io/istio/pkg/workloadapi"
)

func TestServiceEntryServices(t *testing.T) {
	cases := []struct {
		name   string
		inputs []any
		se     *networkingclient.ServiceEntry
		result []*workloadapi.Service
	}{
		{
			name:   "DNS service entry with address",
			inputs: []any{},
			se: &networkingclient.ServiceEntry{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "name",
					Namespace: "ns",
				},
				Spec: networking.ServiceEntry{
					Addresses: []string{"1.2.3.4"},
					Hosts:     []string{"a.example.com", "b.example.com"},
					Ports: []*networking.ServicePort{{
						Number: 80,
						Name:   "http",
					}},
					SubjectAltNames: []string{"san1"},
					Resolution:      networking.ServiceEntry_DNS,
				},
			},
			result: []*workloadapi.Service{
				{
					Name:      "name",
					Namespace: "ns",
					Hostname:  "a.example.com",
					Addresses: []*workloadapi.NetworkAddress{{
						Network: testNW,
						Address: netip.AddrFrom4([4]byte{1, 2, 3, 4}).AsSlice(),
					}},
					Ports: []*workloadapi.Port{{
						ServicePort: 80,
						TargetPort:  80,
					}},
					SubjectAltNames: []string{"san1"},
				},
				{
					Name:      "name",
					Namespace: "ns",
					Hostname:  "b.example.com",
					Addresses: []*workloadapi.NetworkAddress{{
						Network: testNW,
						Address: netip.AddrFrom4([4]byte{1, 2, 3, 4}).AsSlice(),
					}},
					Ports: []*workloadapi.Port{{
						ServicePort: 80,
						TargetPort:  80,
					}},
					SubjectAltNames: []string{"san1"},
				},
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			mock := krttest.NewMock(t, tt.inputs)
			a := newAmbientUnitTest()
			builder := a.serviceEntryServiceBuilder(
				krttest.GetMockCollection[Waypoint](mock),
				krttest.GetMockCollection[*v1.Namespace](mock),
			)
			wrapper := builder(krt.TestingDummyContext{}, tt.se)
			res := slices.Map(wrapper, func(e model.ServiceInfo) *workloadapi.Service {
				return e.Service
			})
			assert.Equal(t, res, tt.result)
		})
	}
}
