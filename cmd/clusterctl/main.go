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

package main

import (
	"github.com/golang/glog"
	"sigs.k8s.io/cluster-api/cmd/clusterctl/cmd"
	clustercommon "sigs.k8s.io/cluster-api/pkg/apis/cluster/common"

	"github.com/detiber/cluster-api-provider-ansible/pkg/cloud/ansible"
)

func main() {
	var err error
	ansible.MachineActuator, err = ansible.NewMachineActuator(ansible.MachineActuatorParams{})
	if err != nil {
		glog.Fatalf("Error creating cluster provisioner for ansible : %v", err)
	}
	clustercommon.RegisterClusterProvisioner(ansible.ProviderName, ansible.MachineActuator)
	cmd.Execute()
}
