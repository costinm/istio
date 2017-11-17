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

// THIS FILE IS AUTOMATICALLY GENERATED.

package svcctrlreport

import (
	"context"
	"time"

	"istio.io/istio/mixer/pkg/adapter"
)

// Fully qualified name of the template
const TemplateName = "svcctrlreport"

// Instance is constructed by Mixer for the 'svcctrlreport' template.
//
// A template used by Google Service Control (svcctrl) adapter. The adapter
// generates metrics and logentry for each request based on the data point
// defined by this template.
//
// Config example:
// ```
// apiVersion: "config.istio.io/v1alpha2"
// kind: svcctrlreport
// metadata:
//   name: report
//   namespace: istio-system
// spec:
//   api_version : api.version | ""
//   api_operation : api.operation | ""
//   api_protocol : api.protocol | ""
//   api_service : api.service | ""
//   api_key : api.key | ""
//   request_time : request.time
//   request_method : request.method
//   request_path : request.path
//   request_bytes: request.size
//   response_time : response.time
//   response_code : response.code | 520
//   response_bytes : response.size | 0
//   response_latency : response.duration | "0ms"
// ```
type Instance struct {
	// Name of the instance as specified in configuration.
	Name string

	ApiVersion string

	ApiOperation string

	ApiProtocol string

	ApiService string

	ApiKey string

	RequestTime time.Time

	RequestMethod string

	RequestPath string

	RequestBytes int64

	ResponseTime time.Time

	ResponseCode int64

	ResponseBytes int64

	ResponseLatency time.Duration
}

// HandlerBuilder must be implemented by adapters if they want to
// process data associated with the 'svcctrlreport' template.
//
// Mixer uses this interface to call into the adapter at configuration time to configure
// it with adapter-specific configuration as well as all template-specific type information.
type HandlerBuilder interface {
	adapter.HandlerBuilder

	// SetSvcctrlReportTypes is invoked by Mixer to pass the template-specific Type information for instances that an adapter
	// may receive at runtime. The type information describes the shape of the instance.
	SetSvcctrlReportTypes(map[string]*Type /*Instance name -> Type*/)
}

// Handler must be implemented by adapter code if it wants to
// process data associated with the 'svcctrlreport' template.
//
// Mixer uses this interface to call into the adapter at request time in order to dispatch
// created instances to the adapter. Adapters take the incoming instances and do what they
// need to achieve their primary function.
//
// The name of each instance can be used as a key into the Type map supplied to the adapter
// at configuration time via the method 'SetSvcctrlReportTypes'.
// These Type associated with an instance describes the shape of the instance
type Handler interface {
	adapter.Handler

	// HandleSvcctrlReport is called by Mixer at request time to deliver instances to
	// to an adapter.
	HandleSvcctrlReport(context.Context, []*Instance) error
}
