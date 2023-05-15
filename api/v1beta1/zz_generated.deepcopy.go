//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/openstack-k8s-operators/lib-common/modules/common/condition"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUCountReq) DeepCopyInto(out *CPUCountReq) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUCountReq.
func (in *CPUCountReq) DeepCopy() *CPUCountReq {
	if in == nil {
		return nil
	}
	out := new(CPUCountReq)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUMhzReq) DeepCopyInto(out *CPUMhzReq) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUMhzReq.
func (in *CPUMhzReq) DeepCopy() *CPUMhzReq {
	if in == nil {
		return nil
	}
	out := new(CPUMhzReq)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUReqs) DeepCopyInto(out *CPUReqs) {
	*out = *in
	out.CountReq = in.CountReq
	out.MhzReq = in.MhzReq
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUReqs.
func (in *CPUReqs) DeepCopy() *CPUReqs {
	if in == nil {
		return nil
	}
	out := new(CPUReqs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskGbReq) DeepCopyInto(out *DiskGbReq) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskGbReq.
func (in *DiskGbReq) DeepCopy() *DiskGbReq {
	if in == nil {
		return nil
	}
	out := new(DiskGbReq)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskReqs) DeepCopyInto(out *DiskReqs) {
	*out = *in
	out.GbReq = in.GbReq
	out.SSDReq = in.SSDReq
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskReqs.
func (in *DiskReqs) DeepCopy() *DiskReqs {
	if in == nil {
		return nil
	}
	out := new(DiskReqs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskSSDReq) DeepCopyInto(out *DiskSSDReq) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskSSDReq.
func (in *DiskSSDReq) DeepCopy() *DiskSSDReq {
	if in == nil {
		return nil
	}
	out := new(DiskSSDReq)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HardwareReqs) DeepCopyInto(out *HardwareReqs) {
	*out = *in
	out.CPUReqs = in.CPUReqs
	out.MemReqs = in.MemReqs
	out.DiskReqs = in.DiskReqs
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HardwareReqs.
func (in *HardwareReqs) DeepCopy() *HardwareReqs {
	if in == nil {
		return nil
	}
	out := new(HardwareReqs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostStatus) DeepCopyInto(out *HostStatus) {
	*out = *in
	in.IPStatus.DeepCopyInto(&out.IPStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostStatus.
func (in *HostStatus) DeepCopy() *HostStatus {
	if in == nil {
		return nil
	}
	out := new(HostStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPStatus) DeepCopyInto(out *IPStatus) {
	*out = *in
	if in.IPAddresses != nil {
		in, out := &in.IPAddresses, &out.IPAddresses
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPStatus.
func (in *IPStatus) DeepCopy() *IPStatus {
	if in == nil {
		return nil
	}
	out := new(IPStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpec) DeepCopyInto(out *InstanceSpec) {
	*out = *in
	if in.UserData != nil {
		in, out := &in.UserData, &out.UserData
		*out = new(v1.SecretReference)
		**out = **in
	}
	if in.NetworkData != nil {
		in, out := &in.NetworkData, &out.NetworkData
		*out = new(v1.SecretReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpec.
func (in *InstanceSpec) DeepCopy() *InstanceSpec {
	if in == nil {
		return nil
	}
	out := new(InstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemGbReq) DeepCopyInto(out *MemGbReq) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemGbReq.
func (in *MemGbReq) DeepCopy() *MemGbReq {
	if in == nil {
		return nil
	}
	out := new(MemGbReq)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemReqs) DeepCopyInto(out *MemReqs) {
	*out = *in
	out.GbReq = in.GbReq
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemReqs.
func (in *MemReqs) DeepCopy() *MemReqs {
	if in == nil {
		return nil
	}
	out := new(MemReqs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackBaremetalSet) DeepCopyInto(out *OpenStackBaremetalSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackBaremetalSet.
func (in *OpenStackBaremetalSet) DeepCopy() *OpenStackBaremetalSet {
	if in == nil {
		return nil
	}
	out := new(OpenStackBaremetalSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackBaremetalSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackBaremetalSetList) DeepCopyInto(out *OpenStackBaremetalSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackBaremetalSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackBaremetalSetList.
func (in *OpenStackBaremetalSetList) DeepCopy() *OpenStackBaremetalSetList {
	if in == nil {
		return nil
	}
	out := new(OpenStackBaremetalSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackBaremetalSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackBaremetalSetSpec) DeepCopyInto(out *OpenStackBaremetalSetSpec) {
	*out = *in
	if in.BaremetalHosts != nil {
		in, out := &in.BaremetalHosts, &out.BaremetalHosts
		*out = make(map[string]InstanceSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.UserData != nil {
		in, out := &in.UserData, &out.UserData
		*out = new(v1.SecretReference)
		**out = **in
	}
	if in.NetworkData != nil {
		in, out := &in.NetworkData, &out.NetworkData
		*out = new(v1.SecretReference)
		**out = **in
	}
	if in.BmhLabelSelector != nil {
		in, out := &in.BmhLabelSelector, &out.BmhLabelSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.HardwareReqs = in.HardwareReqs
	if in.PasswordSecret != nil {
		in, out := &in.PasswordSecret, &out.PasswordSecret
		*out = new(v1.SecretReference)
		**out = **in
	}
	if in.BootstrapDNS != nil {
		in, out := &in.BootstrapDNS, &out.BootstrapDNS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DNSSearchDomains != nil {
		in, out := &in.DNSSearchDomains, &out.DNSSearchDomains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackBaremetalSetSpec.
func (in *OpenStackBaremetalSetSpec) DeepCopy() *OpenStackBaremetalSetSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackBaremetalSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackBaremetalSetStatus) DeepCopyInto(out *OpenStackBaremetalSetStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.BaremetalHosts != nil {
		in, out := &in.BaremetalHosts, &out.BaremetalHosts
		*out = make(map[string]HostStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackBaremetalSetStatus.
func (in *OpenStackBaremetalSetStatus) DeepCopy() *OpenStackBaremetalSetStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackBaremetalSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackProvisionServer) DeepCopyInto(out *OpenStackProvisionServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackProvisionServer.
func (in *OpenStackProvisionServer) DeepCopy() *OpenStackProvisionServer {
	if in == nil {
		return nil
	}
	out := new(OpenStackProvisionServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackProvisionServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackProvisionServerDefaults) DeepCopyInto(out *OpenStackProvisionServerDefaults) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackProvisionServerDefaults.
func (in *OpenStackProvisionServerDefaults) DeepCopy() *OpenStackProvisionServerDefaults {
	if in == nil {
		return nil
	}
	out := new(OpenStackProvisionServerDefaults)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackProvisionServerList) DeepCopyInto(out *OpenStackProvisionServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackProvisionServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackProvisionServerList.
func (in *OpenStackProvisionServerList) DeepCopy() *OpenStackProvisionServerList {
	if in == nil {
		return nil
	}
	out := new(OpenStackProvisionServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackProvisionServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackProvisionServerSpec) DeepCopyInto(out *OpenStackProvisionServerSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackProvisionServerSpec.
func (in *OpenStackProvisionServerSpec) DeepCopy() *OpenStackProvisionServerSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackProvisionServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackProvisionServerStatus) DeepCopyInto(out *OpenStackProvisionServerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackProvisionServerStatus.
func (in *OpenStackProvisionServerStatus) DeepCopy() *OpenStackProvisionServerStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackProvisionServerStatus)
	in.DeepCopyInto(out)
	return out
}
