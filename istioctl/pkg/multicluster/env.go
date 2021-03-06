// Copyright 2019 Istio Authors.
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

package multicluster

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd/api"

	"istio.io/istio/pkg/kube"
)

type Environment interface {
	GetConfig() *api.Config
	CreateClientSet(context string) (kubernetes.Interface, error)
	Stdout() io.Writer
	Stderr() io.Writer
	ReadFile(filename string) ([]byte, error)
	Printf(format string, a ...interface{})
	Errorf(format string, a ...interface{})
}

type KubeEnvironment struct {
	config     *api.Config
	stdout     io.Writer
	stderr     io.Writer
	kubeconfig string
}

func (e *KubeEnvironment) CreateClientSet(context string) (kubernetes.Interface, error) {
	return kube.CreateClientset(e.kubeconfig, context)
}

func (e *KubeEnvironment) Printf(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(e.stdout, format, a...)
}
func (e *KubeEnvironment) Errorf(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(e.stderr, format, a...)
}

func (e *KubeEnvironment) GetConfig() *api.Config                   { return e.config }
func (e *KubeEnvironment) Stdout() io.Writer                        { return e.stdout }
func (e *KubeEnvironment) Stderr() io.Writer                        { return e.stderr }
func (e *KubeEnvironment) ReadFile(filename string) ([]byte, error) { return ioutil.ReadFile(filename) }

var _ Environment = (*KubeEnvironment)(nil)

func NewEnvironment(kubeconfig, context string, stdout, stderr io.Writer) (*KubeEnvironment, error) {
	config, err := kube.BuildClientCmd(kubeconfig, context).ConfigAccess().GetStartingConfig()
	if err != nil {
		return nil, err
	}

	return &KubeEnvironment{
		config:     config,
		stdout:     stdout,
		stderr:     stderr,
		kubeconfig: kubeconfig,
	}, nil
}

func NewEnvironmentFromCobra(kubeconfig, context string, cmd *cobra.Command) (Environment, error) {
	return NewEnvironment(kubeconfig, context, cmd.OutOrStdout(), cmd.OutOrStderr())
}
