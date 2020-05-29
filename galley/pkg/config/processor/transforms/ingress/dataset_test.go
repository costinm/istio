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

package ingress

import (
	"bytes"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"

	"istio.io/api/mesh/v1alpha1"
	"istio.io/api/networking/v1alpha3"

	"istio.io/istio/galley/pkg/config/mesh"
	"istio.io/istio/galley/pkg/config/source/kube/rt"
	"istio.io/istio/pkg/config/resource"
)

func ingress1() *resource.Instance {
	return toIngressResource(`
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: foo
  namespace: ns
  annotations:
    kubernetes.io/ingress.class: "cls"
  resourceVersion: v1
spec:
  rules:
  - host: foohost.bar.com
    http:
      paths:
      - path: /foopath
        backend:
          serviceName: service1
          servicePort: 4200
`)
}

func ingress1v2() *resource.Instance {
	return toIngressResource(`
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: foo
  namespace: ns
  annotations:
    kubernetes.io/ingress.class: "cls"
  resourceVersion: v2
spec:
  rules:
  - host: foohost.bar.com
    http:
      paths:
      - path: /foopath
        backend:
          serviceName: service2
          servicePort: 2400
`)
}

func gw1() *resource.Instance {
	return &resource.Instance{
		Metadata: resource.Metadata{
			FullName:    resource.NewFullName("istio-system", "foo-istio-autogenerated-k8s-ingress"),
			Version:     "$ing_O/wmlZTvZJIo6adLqwDwQu/JHVrMb77jGjgugNQjiP4",
			Annotations: map[string]string{},
		},
		Message: parseGateway(`
 {
        "selector": {
          "istio": "ingress"
        },
        "servers": [
          {
            "hosts": [
              "*"
            ],
            "port": {
              "name": "http-80-i-foo-ns",
              "number": 80,
              "protocol": "HTTP"
            }
          }
        ]
      },
`),
		Origin: (*rt.Origin)(nil),
	}
}

func gw1v2() *resource.Instance {
	return &resource.Instance{
		Metadata: resource.Metadata{
			FullName:    resource.NewFullName("istio-system", "foo-istio-autogenerated-k8s-ingress"),
			Version:     "$ing_+wTctpcOTD0Yc95R/VpQ17tGszgxE2AmZcNQ7EC1+ZA",
			Annotations: map[string]string{},
		},
		Message: parseGateway(`
 {
        "selector": {
          "istio": "ingress"
        },
        "servers": [
          {
            "hosts": [
              "*"
            ],
            "port": {
              "name": "http-80-i-foo-ns",
              "number": 80,
              "protocol": "HTTP"
            }
          }
        ]
      },
`),
		Origin: (*rt.Origin)(nil),
	}
}

func vs1() *resource.Instance {
	return &resource.Instance{
		Metadata: resource.Metadata{
			FullName:    resource.NewFullName("istio-system", "foohost-bar-com-foo-istio-autogenerated-k8s-ingress"),
			Version:     "$ing_zW/HWlEZ6+A8Z2HIpAsaRVskHx9AgXAyTvL7UNl5vuU",
			Annotations: map[string]string{},
		},
		Message: &v1alpha3.VirtualService{
			Hosts: []string{
				"foohost.bar.com",
			},
			Gateways: []string{"istio-autogenerated-k8s-ingress"},
			Http: []*v1alpha3.HTTPRoute{
				{
					Match: []*v1alpha3.HTTPMatchRequest{
						{
							Uri: &v1alpha3.StringMatch{
								MatchType: &v1alpha3.StringMatch_Exact{
									Exact: "/foopath",
								},
							},
						},
					},

					Route: []*v1alpha3.HTTPRouteDestination{
						{
							Destination: &v1alpha3.Destination{
								Host: "service1.ns.svc.cluster.local",
								Port: &v1alpha3.PortSelector{
									Number: 4200,
								},
							},
							Weight: 100,
						},
					},
				},
			},
		},
		// Rationale: Gomega will insist on typed nil, but output only 'nil' on failure.
		Origin: (*rt.Origin)(nil),
	}
}

func vs1v2() *resource.Instance {
	return &resource.Instance{
		Metadata: resource.Metadata{
			FullName:    resource.NewFullName("istio-system", "foohost-bar-com-foo-istio-autogenerated-k8s-ingress"),
			Version:     "$ing_HWr/Pv0tKjRCWxF3pL8DhUuXlBRbnBgfI7EsEMVXuSY",
			Annotations: map[string]string{},
		},
		Message: &v1alpha3.VirtualService{
			Hosts: []string{
				"foohost.bar.com",
			},
			Gateways: []string{"istio-autogenerated-k8s-ingress"},
			Http: []*v1alpha3.HTTPRoute{
				{
					Match: []*v1alpha3.HTTPMatchRequest{
						{
							Uri: &v1alpha3.StringMatch{
								MatchType: &v1alpha3.StringMatch_Exact{
									Exact: "/foopath",
								},
							},
						},
					},

					Route: []*v1alpha3.HTTPRouteDestination{
						{
							Destination: &v1alpha3.Destination{
								Host: "service2.ns.svc.cluster.local",
								Port: &v1alpha3.PortSelector{
									Number: 2400,
								},
							},
							Weight: 100,
						},
					},
				},
			},
		},
		Origin: (*rt.Origin)(nil),
	}
}

func meshConfig() *v1alpha1.MeshConfig {
	m := mesh.DefaultMeshConfig()
	m.IngressClass = "cls"
	m.IngressControllerMode = v1alpha1.MeshConfig_STRICT
	return m
}

func toIngressResource(s string) *resource.Instance {
	r, err := ingressAdapter.JSONToEntry(s)
	if err != nil {
		panic(err)
	}
	return r
}

func parseGateway(s string) proto.Message {
	p := &v1alpha3.Gateway{}
	b := bytes.NewReader([]byte(s))
	err := jsonpb.Unmarshal(b, p)
	if err != nil {
		panic(err)
	}
	return p
}
