/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ansible

import (
	"errors"

	"github.com/ghodss/yaml"
	"github.com/golang/glog"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/record"
	clusterv1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
	apierrors "sigs.k8s.io/cluster-api/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	ansibleconfigv1 "github.com/detiber/cluster-api-provider-ansible/pkg/apis/ansibleproviderconfig/v1alpha1"
)

const (
	ProviderName = "ansible"
)

const (
	createEventAction = "Create"
	deleteEventAction = "Delete"
	noEventAction     = ""
)

var MachineActuator *AnsibleClient

type AnsibleClient struct {
	client        client.Client
	eventRecorder record.EventRecorder
	scheme        *runtime.Scheme
}

type MachineActuatorParams struct {
	Client        client.Client
	EventRecorder record.EventRecorder
	Scheme        *runtime.Scheme
}

func NewMachineActuator(params MachineActuatorParams) (*AnsibleClient, error) {
	return &AnsibleClient{
		client:        params.Client,
		eventRecorder: params.EventRecorder,
		scheme:        params.Scheme,
	}, nil
}

// TODO move the following four functions to a separate file?
func clusterProviderFromProviderConfig(providerConfig clusterv1.ProviderConfig) (*ansibleconfigv1.AnsibleClusterProviderConfig, error) {
	var config ansibleconfigv1.AnsibleClusterProviderConfig
	if err := yaml.Unmarshal(providerConfig.Value.Raw, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func machineProviderFromProviderConfig(providerConfig clusterv1.ProviderConfig) (*ansibleconfigv1.AnsibleMachineProviderConfig, error) {
	var config ansibleconfigv1.AnsibleMachineProviderConfig
	if err := yaml.Unmarshal(providerConfig.Value.Raw, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

// TODO these two funcs shouldn't be exported, but need to be for testing...
func ProviderConfigFromCluster(in *ansibleconfigv1.AnsibleClusterProviderConfig) (*clusterv1.ProviderConfig, error) {
	mpc, err := yaml.Marshal(in)
	if err != nil {
		return nil, err
	}
	return &clusterv1.ProviderConfig{
		Value: &runtime.RawExtension{Raw: mpc},
	}, nil
}

func ProviderConfigFromMachine(in *ansibleconfigv1.AnsibleMachineProviderConfig) (*clusterv1.ProviderConfig, error) {
	mpc, err := yaml.Marshal(in)
	if err != nil {
		return nil, err
	}
	return &clusterv1.ProviderConfig{
		Value: &runtime.RawExtension{Raw: mpc},
	}, nil
}

func (ans *AnsibleClient) CreateMachineController(cluster *clusterv1.Cluster, initialMachines []*clusterv1.Machine, clientSet kubernetes.Clientset) error {
	return nil
}

func (ans *AnsibleClient) Create(cluster *clusterv1.Cluster, machine *clusterv1.Machine) error {
	return errors.New("not yet implemented")
}

func (ans *AnsibleClient) Delete(cluster *clusterv1.Cluster, machine *clusterv1.Machine) error {
	return errors.New("not yet implemented")
}

func (ans *AnsibleClient) Update(cluster *clusterv1.Cluster, machine *clusterv1.Machine) error {
	return errors.New("not yet implemented")
}

func (ans *AnsibleClient) Exists(cluster *clusterv1.Cluster, machine *clusterv1.Machine) (bool, error) {
	return false, errors.New("not yet implemented")
}

// If the AnsibleClient has a client for updating Machine objects, this will set
// the appropriate reason/message on the Machine.Status. If not, such as during
// cluster installation, it will operate as a no-op. It also returns the
// original error for convenience, so callers can do "return handleMachineError(...)".
func (ans *AnsibleClient) handleMachineError(machine *clusterv1.Machine, err *apierrors.MachineError, eventAction string) error {
	if ans.client != nil {
		reason := err.Reason
		message := err.Message
		machine.Status.ErrorReason = &reason
		machine.Status.ErrorMessage = &message
		panic("UpdateStatus not implemented")
	}

	if eventAction != noEventAction {
		ans.eventRecorder.Eventf(machine, corev1.EventTypeWarning, "Failed"+eventAction, "%v", err.Reason)
	}

	glog.Errorf("Machine error: %v", err.Message)
	return err
}
