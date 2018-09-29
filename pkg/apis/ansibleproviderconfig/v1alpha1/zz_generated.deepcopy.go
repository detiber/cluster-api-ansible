// +build !ignore_autogenerated

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
// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsibleClusterProviderConfig) DeepCopyInto(out *AnsibleClusterProviderConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsibleClusterProviderConfig.
func (in *AnsibleClusterProviderConfig) DeepCopy() *AnsibleClusterProviderConfig {
	if in == nil {
		return nil
	}
	out := new(AnsibleClusterProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnsibleClusterProviderConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsibleClusterProviderConfigList) DeepCopyInto(out *AnsibleClusterProviderConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AnsibleClusterProviderConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsibleClusterProviderConfigList.
func (in *AnsibleClusterProviderConfigList) DeepCopy() *AnsibleClusterProviderConfigList {
	if in == nil {
		return nil
	}
	out := new(AnsibleClusterProviderConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnsibleClusterProviderConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsibleClusterProviderConfigSpec) DeepCopyInto(out *AnsibleClusterProviderConfigSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsibleClusterProviderConfigSpec.
func (in *AnsibleClusterProviderConfigSpec) DeepCopy() *AnsibleClusterProviderConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AnsibleClusterProviderConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsibleClusterProviderConfigStatus) DeepCopyInto(out *AnsibleClusterProviderConfigStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsibleClusterProviderConfigStatus.
func (in *AnsibleClusterProviderConfigStatus) DeepCopy() *AnsibleClusterProviderConfigStatus {
	if in == nil {
		return nil
	}
	out := new(AnsibleClusterProviderConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsibleMachineProviderConfig) DeepCopyInto(out *AnsibleMachineProviderConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsibleMachineProviderConfig.
func (in *AnsibleMachineProviderConfig) DeepCopy() *AnsibleMachineProviderConfig {
	if in == nil {
		return nil
	}
	out := new(AnsibleMachineProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnsibleMachineProviderConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsibleMachineProviderConfigList) DeepCopyInto(out *AnsibleMachineProviderConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AnsibleMachineProviderConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsibleMachineProviderConfigList.
func (in *AnsibleMachineProviderConfigList) DeepCopy() *AnsibleMachineProviderConfigList {
	if in == nil {
		return nil
	}
	out := new(AnsibleMachineProviderConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnsibleMachineProviderConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsibleMachineProviderConfigSpec) DeepCopyInto(out *AnsibleMachineProviderConfigSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsibleMachineProviderConfigSpec.
func (in *AnsibleMachineProviderConfigSpec) DeepCopy() *AnsibleMachineProviderConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AnsibleMachineProviderConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsibleMachineProviderConfigStatus) DeepCopyInto(out *AnsibleMachineProviderConfigStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsibleMachineProviderConfigStatus.
func (in *AnsibleMachineProviderConfigStatus) DeepCopy() *AnsibleMachineProviderConfigStatus {
	if in == nil {
		return nil
	}
	out := new(AnsibleMachineProviderConfigStatus)
	in.DeepCopyInto(out)
	return out
}