package ansible

import (
	"fmt"

	clusterv1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
)

type AnsibleClient struct {
}

func (ans *AnsibleClient) Create(cluster *clusterv1.Cluster, machine *clusterv1.Machine) error {
	return fmt.Errorf("Not implemented yet.")
}

func (ans *AnsibleClient) Delete(cluster *clusterv1.Cluster, machine *clusterv1.Machine) error {
	return fmt.Errorf("Not implemented yet.")
}

func (ans *AnsibleClient) Update(cluster *clusterv1.Cluster, goalMachine *clusterv1.Machine) error {
	return fmt.Errorf("Not implemented yet.")
}

func (ans *AnsibleClient) Exists(cluster *clusterv1.Cluster, machine *clusterv1.Machine) (bool, error) {
	return false, fmt.Errorf("Not implemented yet.")
}

func NewMachineActuator() (*AnsibleClient, error) {
	return nil, fmt.Errorf("Not implemented yet.")
}
