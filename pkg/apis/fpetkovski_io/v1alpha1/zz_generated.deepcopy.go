// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRule) DeepCopyInto(out *ResourceRule) {
	*out = *in
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.MatchLabels != nil {
		in, out := &in.MatchLabels, &out.MatchLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRule.
func (in *ResourceRule) DeepCopy() *ResourceRule {
	if in == nil {
		return nil
	}
	out := new(ResourceRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TTLPolicy) DeepCopyInto(out *TTLPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TTLPolicy.
func (in *TTLPolicy) DeepCopy() *TTLPolicy {
	if in == nil {
		return nil
	}
	out := new(TTLPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TTLPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TTLPolicyList) DeepCopyInto(out *TTLPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TTLPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TTLPolicyList.
func (in *TTLPolicyList) DeepCopy() *TTLPolicyList {
	if in == nil {
		return nil
	}
	out := new(TTLPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TTLPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TTLPolicySpec) DeepCopyInto(out *TTLPolicySpec) {
	*out = *in
	in.ResourceRule.DeepCopyInto(&out.ResourceRule)
	if in.ExpirationFrom != nil {
		in, out := &in.ExpirationFrom, &out.ExpirationFrom
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TTLPolicySpec.
func (in *TTLPolicySpec) DeepCopy() *TTLPolicySpec {
	if in == nil {
		return nil
	}
	out := new(TTLPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TTLPolicyStatus) DeepCopyInto(out *TTLPolicyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TTLPolicyStatus.
func (in *TTLPolicyStatus) DeepCopy() *TTLPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(TTLPolicyStatus)
	in.DeepCopyInto(out)
	return out
}
