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

	"github.com/golang/glog"
	clusterv1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

type AnsibleClusterClient struct {
	client client.Client
}

type ClusterActuatorParams struct {
}

func NewClusterActuator(m manager.Manager, params ClusterActuatorParams) (*AnsibleClusterClient, error) {
	return &AnsibleClusterClient{
		client: m.GetClient(),
	}, nil
}

func (ans *AnsibleClusterClient) Reconcile(cluster *clusterv1.Cluster) error {
	glog.Infof("Reconciling cluster %v.", cluster.Name)
	return errors.New("not yet implemented")
}

func (ans *AnsibleClusterClient) Delete(cluster *clusterv1.Cluster) error {
	return errors.New("not yet implemented")
}
