// Copyright 2019 Istio Authors
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
package sidecar

import (
	"istio.io/api/networking/v1alpha3"
	"istio.io/istio/galley/pkg/config/analysis"
	"istio.io/istio/galley/pkg/config/analysis/msg"
	"istio.io/istio/galley/pkg/config/meta/metadata"
	"istio.io/istio/galley/pkg/config/meta/schema/collection"
	"istio.io/istio/galley/pkg/config/resource"
)

// DefaultSelectorAnalyzer validates, per namespace, that there aren't multiple Sidecar resources that have no selector
// This is distinct from SelectorAnalyzer because it does not require pods, so it can run even if that collection is unavailable
type DefaultSelectorAnalyzer struct{}

var _ analysis.Analyzer = &DefaultSelectorAnalyzer{}

// Metadata implements Analyzer
func (a *DefaultSelectorAnalyzer) Metadata() analysis.Metadata {
	return analysis.Metadata{
		Name: "sidecar.DefaultSelectorAnalyzer",
		Inputs: collection.Names{
			metadata.IstioNetworkingV1Alpha3Sidecars,
		},
	}
}

// Analyze implements Analyzer
func (a *DefaultSelectorAnalyzer) Analyze(c analysis.Context) {
	nsToSidecars := make(map[string][]*resource.Entry)

	c.ForEach(metadata.IstioNetworkingV1Alpha3Sidecars, func(r *resource.Entry) bool {
		s := r.Item.(*v1alpha3.Sidecar)

		ns, _ := r.Metadata.Name.InterpretAsNamespaceAndName()

		if s.WorkloadSelector == nil {
			nsToSidecars[ns] = append(nsToSidecars[ns], r)
		}
		return true
	})

	// Check for more than one selector-less sidecar instance, per namespace
	for ns, sList := range nsToSidecars {
		if len(sList) > 1 {
			sNames := getNames(sList)
			for _, r := range sList {
				c.Report(metadata.IstioNetworkingV1Alpha3Sidecars, msg.NewMultipleSidecarsWithoutWorkloadSelectors(r, sNames, ns))
			}
		}
	}
}
