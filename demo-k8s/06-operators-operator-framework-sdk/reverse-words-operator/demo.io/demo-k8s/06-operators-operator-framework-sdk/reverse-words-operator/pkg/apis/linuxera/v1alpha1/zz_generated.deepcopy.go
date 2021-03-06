// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReverseWordsApp) DeepCopyInto(out *ReverseWordsApp) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReverseWordsApp.
func (in *ReverseWordsApp) DeepCopy() *ReverseWordsApp {
	if in == nil {
		return nil
	}
	out := new(ReverseWordsApp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReverseWordsApp) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReverseWordsAppList) DeepCopyInto(out *ReverseWordsAppList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReverseWordsApp, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReverseWordsAppList.
func (in *ReverseWordsAppList) DeepCopy() *ReverseWordsAppList {
	if in == nil {
		return nil
	}
	out := new(ReverseWordsAppList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReverseWordsAppList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReverseWordsAppSpec) DeepCopyInto(out *ReverseWordsAppSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReverseWordsAppSpec.
func (in *ReverseWordsAppSpec) DeepCopy() *ReverseWordsAppSpec {
	if in == nil {
		return nil
	}
	out := new(ReverseWordsAppSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReverseWordsAppStatus) DeepCopyInto(out *ReverseWordsAppStatus) {
	*out = *in
	if in.AppPods != nil {
		in, out := &in.AppPods, &out.AppPods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReverseWordsAppStatus.
func (in *ReverseWordsAppStatus) DeepCopy() *ReverseWordsAppStatus {
	if in == nil {
		return nil
	}
	out := new(ReverseWordsAppStatus)
	in.DeepCopyInto(out)
	return out
}
