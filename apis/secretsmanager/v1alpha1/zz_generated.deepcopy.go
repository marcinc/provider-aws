// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Secret) DeepCopyInto(out *Secret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Secret.
func (in *Secret) DeepCopy() *Secret {
	if in == nil {
		return nil
	}
	out := new(Secret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Secret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretList) DeepCopyInto(out *SecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Secret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretList.
func (in *SecretList) DeepCopy() *SecretList {
	if in == nil {
		return nil
	}
	out := new(SecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretObservation) DeepCopyInto(out *SecretObservation) {
	*out = *in
	if in.CreatedDate != nil {
		in, out := &in.CreatedDate, &out.CreatedDate
		*out = (*in).DeepCopy()
	}
	if in.DeletedDate != nil {
		in, out := &in.DeletedDate, &out.DeletedDate
		*out = (*in).DeepCopy()
	}
	if in.DeletionDate != nil {
		in, out := &in.DeletionDate, &out.DeletionDate
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretObservation.
func (in *SecretObservation) DeepCopy() *SecretObservation {
	if in == nil {
		return nil
	}
	out := new(SecretObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretParameters) DeepCopyInto(out *SecretParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.KmsKeyID != nil {
		in, out := &in.KmsKeyID, &out.KmsKeyID
		*out = new(string)
		**out = **in
	}
	if in.KmsKeyRef != nil {
		in, out := &in.KmsKeyRef, &out.KmsKeyRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.KmsKeySelector != nil {
		in, out := &in.KmsKeySelector, &out.KmsKeySelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(SecretSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ForceDeleteWithoutRecovery != nil {
		in, out := &in.ForceDeleteWithoutRecovery, &out.ForceDeleteWithoutRecovery
		*out = new(bool)
		**out = **in
	}
	if in.RecoveryWindowInDays != nil {
		in, out := &in.RecoveryWindowInDays, &out.RecoveryWindowInDays
		*out = new(int64)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]Tag, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretParameters.
func (in *SecretParameters) DeepCopy() *SecretParameters {
	if in == nil {
		return nil
	}
	out := new(SecretParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretSelector) DeepCopyInto(out *SecretSelector) {
	*out = *in
	if in.SecretReference != nil {
		in, out := &in.SecretReference, &out.SecretReference
		*out = new(v1.SecretReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretSelector.
func (in *SecretSelector) DeepCopy() *SecretSelector {
	if in == nil {
		return nil
	}
	out := new(SecretSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretSpec) DeepCopyInto(out *SecretSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretSpec.
func (in *SecretSpec) DeepCopy() *SecretSpec {
	if in == nil {
		return nil
	}
	out := new(SecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretStatus) DeepCopyInto(out *SecretStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretStatus.
func (in *SecretStatus) DeepCopy() *SecretStatus {
	if in == nil {
		return nil
	}
	out := new(SecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}
