package kube_test

import (
	"istio.io/istio/tests/util"
	"testing"
	"time"

	"istio.io/istio/pilot/pkg/serviceregistry"
	"istio.io/istio/pilot/pkg/serviceregistry/aggregate"
	"istio.io/istio/pilot/pkg/serviceregistry/kube"
	"istio.io/istio/tests/k8s"
	"k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Test EDS incremental against real local apiserver.
// Will create 2 registries, one using the mc controller (both backed by same
// apiserver, but using different namespaces)
func TestK8SInc(t *testing.T) {
	util.EnsureTestServer()

	addRemoteCluster(t)

}

func k8sInit(t *testing.T, ki kubernetes.Interface, ns string) {
	_, _ = ki.CoreV1().Namespaces().Create(&v1.Namespace{
		ObjectMeta: meta_v1.ObjectMeta{
			Name: ns,
		},
	})
	makeService("s1", ns, ki, t)
}

// Use the kube registry controller - uses the local apiserver in circle or local.
func addRemoteCluster(t *testing.T) {
	kconf := k8s.Kubeconfig("/../../../../../.circleci/config")
	ki, err := kube.CreateInterface(kconf)
	if err != nil {
		t.Log("Skipping k8s test, no local apiserver")
		return
	}
	ns := "edstest"
	k8sInit(t, ki, ns)
	k8sInit(t, ki, "edsremote")
	clusterID := "remote1"

	stop := make(chan struct{})
	kubectl := kube.NewController(ki, kube.ControllerOptions{
		WatchedNamespace: ns,
		ResyncPeriod:     10 * time.Second,
		DomainSuffix:     "cluster.local",
	})

	kubectl2 := kube.NewController(ki, kube.ControllerOptions{
		WatchedNamespace: "edsremote",
		ResyncPeriod:     10 * time.Second,
		DomainSuffix:     "cluster.local",
		ClusterID:        clusterID,
	})

	agg, _ := util.MockTestServer.EnvoyXdsServer.Env.ServiceDiscovery.(*aggregate.Controller)
	agg.AddRegistry(aggregate.Registry{
		Name:             serviceregistry.KubernetesRegistry, // name used by bootstrap
		ClusterID:        string(serviceregistry.KubernetesRegistry),
		ServiceDiscovery: kubectl,
		ServiceAccounts:  kubectl,
		Controller:       kubectl,
	})

	agg.AddRegistry(aggregate.Registry{
		Name:             serviceregistry.KubernetesRegistry, // name used by bootstrap
		ClusterID:        string(serviceregistry.KubernetesRegistry),
		ServiceDiscovery: kubectl2,
		ServiceAccounts:  kubectl2,
		Controller:       kubectl2,
	})

	go kubectl.Run(stop)
	go kubectl2.Run(stop)

}

func makeService(n, ns string, cl kubernetes.Interface, t *testing.T) {
	_, _ = cl.CoreV1().Services(ns).Create(&v1.Service{
		ObjectMeta: meta_v1.ObjectMeta{Name: n},
		Spec: v1.ServiceSpec{
			Ports: []v1.ServicePort{
				{
					Port: 80,
					Name: "http-main",
				},
			},
		},
	})
	// ignore error - service already created.
}
