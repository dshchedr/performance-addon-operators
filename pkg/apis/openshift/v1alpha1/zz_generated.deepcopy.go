// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CpuPerformanceProfile) DeepCopyInto(out *CpuPerformanceProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CpuPerformanceProfile.
func (in *CpuPerformanceProfile) DeepCopy() *CpuPerformanceProfile {
	if in == nil {
		return nil
	}
	out := new(CpuPerformanceProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CpuPerformanceProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CpuPerformanceProfileList) DeepCopyInto(out *CpuPerformanceProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CpuPerformanceProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CpuPerformanceProfileList.
func (in *CpuPerformanceProfileList) DeepCopy() *CpuPerformanceProfileList {
	if in == nil {
		return nil
	}
	out := new(CpuPerformanceProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CpuPerformanceProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CpuPerformanceProfileSpec) DeepCopyInto(out *CpuPerformanceProfileSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CpuPerformanceProfileSpec.
func (in *CpuPerformanceProfileSpec) DeepCopy() *CpuPerformanceProfileSpec {
	if in == nil {
		return nil
	}
	out := new(CpuPerformanceProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CpuPerformanceProfileStatus) DeepCopyInto(out *CpuPerformanceProfileStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CpuPerformanceProfileStatus.
func (in *CpuPerformanceProfileStatus) DeepCopy() *CpuPerformanceProfileStatus {
	if in == nil {
		return nil
	}
	out := new(CpuPerformanceProfileStatus)
	in.DeepCopyInto(out)
	return out
}
