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

package checkproducer

import (
	"context"

	"istio.io/istio/mixer/pkg/adapter"
)

// Fully qualified name of the template
const TemplateName = "checkproducer"

// Instance is constructed by Mixer for the 'checkproducer' template.
//
// input template
type Instance struct {
	// Name of the instance as specified in configuration.
	Name string

	StringPrimitive string
}

// Output struct is returned by the attribute producing adapters that handle this template.
//
// output template
type Output struct {
	fieldsSet map[string]bool

	Int64Primitive int64

	BoolPrimitive bool

	DoublePrimitive float64

	StringPrimitive string

	StringMap map[string]string
}

func NewOutput() *Output {
	return &Output{fieldsSet: make(map[string]bool)}
}

func (o *Output) SetInt64Primitive(val int64) {
	o.fieldsSet["int64Primitive"] = true
	o.Int64Primitive = val
}

func (o *Output) SetBoolPrimitive(val bool) {
	o.fieldsSet["boolPrimitive"] = true
	o.BoolPrimitive = val
}

func (o *Output) SetDoublePrimitive(val float64) {
	o.fieldsSet["doublePrimitive"] = true
	o.DoublePrimitive = val
}

func (o *Output) SetStringPrimitive(val string) {
	o.fieldsSet["stringPrimitive"] = true
	o.StringPrimitive = val
}

func (o *Output) SetStringMap(val map[string]string) {
	o.fieldsSet["stringMap"] = true
	o.StringMap = val
}

func (o *Output) WasSet(field string) bool {
	_, found := o.fieldsSet[field]
	return found
}

// HandlerBuilder must be implemented by adapters if they want to
// process data associated with the 'checkproducer' template.
//
// Mixer uses this interface to call into the adapter at configuration time to configure
// it with adapter-specific configuration as well as all template-specific type information.
type HandlerBuilder interface {
	adapter.HandlerBuilder

	// SetCheckProducerTypes is invoked by Mixer to pass the template-specific Type information for instances that an adapter
	// may receive at runtime. The type information describes the shape of the instance.
	SetCheckProducerTypes(map[string]*Type /*Instance name -> Type*/)
}

// Handler must be implemented by adapter code if it wants to
// process data associated with the 'checkproducer' template.
//
// Mixer uses this interface to call into the adapter at request time in order to dispatch
// created instances to the adapter. Adapters take the incoming instances and do what they
// need to achieve their primary function.
//
// The name of each instance can be used as a key into the Type map supplied to the adapter
// at configuration time via the method 'SetCheckProducerTypes'.
// These Type associated with an instance describes the shape of the instance
type Handler interface {
	adapter.Handler

	// HandleCheckProducer is called by Mixer at request time to deliver instances to
	// to an adapter.
	HandleCheckProducer(context.Context, *Instance) (adapter.CheckResult, *Output, error)
}
