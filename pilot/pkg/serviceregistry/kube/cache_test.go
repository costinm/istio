// Copyright 2017 Istio Authors
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

package kube

import (
	"k8s.io/client-go/kubernetes"
	"reflect"
	"testing"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"istio.io/istio/pilot/pkg/model"
)

// Create a set of realistic test pod. This can be used in multiple tests, to
// avoid duplicating creation. It can be used with the fake or standalone apiserver.
func initTestPods(ki kubernetes.Interface) {

}

func TestPodCache(t *testing.T) {
	t.Run("localApiserver", func(t *testing.T) {
		c, fx := newLocalController(t)
		defer c.Stop()
		testPodCache(t, c, fx)
	})
	t.Run("fakeApiserver", func(t *testing.T) {
		c, fx := newFakeController(t)
		defer c.Stop()
		testPodCache(t, c, fx)
	})
}

func testPodCache(t *testing.T, c *Controller, fx *FakeXdsUpdater) {
	pods:= []*v1.Pod{
		generatePod("128.0.0.1", "pod1", "nsA", "", "", map[string]string{"app": "test-app"}, map[string]string{}),
		generatePod("128.0.0.2", "pod2", "nsA", "", "", map[string]string{"app": "prod-app-1"}, map[string]string{}),
		generatePod("128.0.0.3", "pod3", "nsB", "", "", map[string]string{"app": "prod-app-2"}, map[string]string{}),
	}

	// Populate podCache
	for _, pod := range pods {
		_, err := c.client.CoreV1().Pods(pod.Namespace).Create(pod)
		//if err := controller.pods.informer.GetStore().Add(pod); err != nil {
		if err != nil {
			t.Errorf("Cannot create %s in namespace %s (error: %v)", pod.ObjectMeta.Name, pod.ObjectMeta.Namespace, err)
		}
		ev := <-fx.Events
		if ev.Id != pod.Status.PodIP {
			t.Error("Workload event expected ", pod.Status.PodIP, "got", ev.Id)
		}
	}

	// Verify podCache
	wantLabels:= map[string]model.Labels{
		"128.0.0.1": {"app": "test-app"},
		"128.0.0.2": {"app": "prod-app-1"},
		"128.0.0.3": {"app": "prod-app-2"},
	}
	for addr, wantTag := range wantLabels {
		tag, found := c.pods.labelsByIP(addr)
		if !found {
			t.Error("Not found ", addr)
		}
		if !reflect.DeepEqual(wantTag, tag) {
			t.Errorf("Expected %v got %v", wantTag, tag)
		}
	}

	// Former 'wantNotFound' test. A pod not in the cache results in found = false
	_, found := c.pods.labelsByIP("128.0.0.4")
	if found {
		t.Error("Expected not found but was found")
	}
}

// Checks that events from the watcher create the proper internal structures
// Deprecated: the current structs are not efficient.
func TestPodCacheEvents(t *testing.T) {
	handler := &ChainHandler{}
	cache := newPodCache(cacheHandler{handler: handler}, nil)

	f := cache.event

	ns := "default"
	ip := "172.0.3.35"
	pod1 := metav1.ObjectMeta{Name: "pod1", Namespace: ns}
	if err := f(&v1.Pod{ObjectMeta: pod1}, model.EventAdd); err != nil {
		t.Error(err)
	}
	if err := f(&v1.Pod{ObjectMeta: pod1, Status: v1.PodStatus{PodIP: ip, Phase: v1.PodPending}}, model.EventUpdate); err != nil {
		t.Error(err)
	}

	if pod, exists := cache.getPodKey(ip); !exists || pod != "default/pod1" {
		t.Errorf("getPodKey => got %s, pod1 not found or incorrect", pod)
	}

	pod2 := metav1.ObjectMeta{Name: "pod2", Namespace: ns}
	if err := f(&v1.Pod{ObjectMeta: pod1, Status: v1.PodStatus{PodIP: ip, Phase: v1.PodFailed}}, model.EventUpdate); err != nil {
		t.Error(err)
	}
	if err := f(&v1.Pod{ObjectMeta: pod2, Status: v1.PodStatus{PodIP: ip, Phase: v1.PodRunning}}, model.EventAdd); err != nil {
		t.Error(err)
	}

	if pod, exists := cache.getPodKey(ip); !exists || pod != "default/pod2" {
		t.Errorf("getPodKey => got %s, pod2 not found or incorrect", pod)
	}

	if err := f(&v1.Pod{ObjectMeta: pod1, Status: v1.PodStatus{PodIP: ip, Phase: v1.PodFailed}}, model.EventDelete); err != nil {
		t.Error(err)
	}

	if pod, exists := cache.getPodKey(ip); !exists || pod != "default/pod2" {
		t.Errorf("getPodKey => got %s, pod2 not found or incorrect", pod)
	}

	if err := f(&v1.Pod{ObjectMeta: pod2, Status: v1.PodStatus{PodIP: ip, Phase: v1.PodFailed}}, model.EventDelete); err != nil {
		t.Error(err)
	}

	if pod, exists := cache.getPodKey(ip); exists {
		t.Errorf("getPodKey => got %s, want none", pod)
	}
}
