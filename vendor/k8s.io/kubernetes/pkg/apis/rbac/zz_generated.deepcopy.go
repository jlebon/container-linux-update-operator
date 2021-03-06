// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package rbac

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

// Deprecated: register deep-copy functions.
func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// Deprecated: RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterRole).DeepCopyInto(out.(*ClusterRole))
			return nil
		}, InType: reflect.TypeOf(&ClusterRole{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterRoleBinding).DeepCopyInto(out.(*ClusterRoleBinding))
			return nil
		}, InType: reflect.TypeOf(&ClusterRoleBinding{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterRoleBindingList).DeepCopyInto(out.(*ClusterRoleBindingList))
			return nil
		}, InType: reflect.TypeOf(&ClusterRoleBindingList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterRoleList).DeepCopyInto(out.(*ClusterRoleList))
			return nil
		}, InType: reflect.TypeOf(&ClusterRoleList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PolicyRule).DeepCopyInto(out.(*PolicyRule))
			return nil
		}, InType: reflect.TypeOf(&PolicyRule{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Role).DeepCopyInto(out.(*Role))
			return nil
		}, InType: reflect.TypeOf(&Role{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*RoleBinding).DeepCopyInto(out.(*RoleBinding))
			return nil
		}, InType: reflect.TypeOf(&RoleBinding{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*RoleBindingList).DeepCopyInto(out.(*RoleBindingList))
			return nil
		}, InType: reflect.TypeOf(&RoleBindingList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*RoleList).DeepCopyInto(out.(*RoleList))
			return nil
		}, InType: reflect.TypeOf(&RoleList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*RoleRef).DeepCopyInto(out.(*RoleRef))
			return nil
		}, InType: reflect.TypeOf(&RoleRef{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Subject).DeepCopyInto(out.(*Subject))
			return nil
		}, InType: reflect.TypeOf(&Subject{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRole) DeepCopyInto(out *ClusterRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]PolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRole.
func (x *ClusterRole) DeepCopy() *ClusterRole {
	if x == nil {
		return nil
	}
	out := new(ClusterRole)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *ClusterRole) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRoleBinding) DeepCopyInto(out *ClusterRoleBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]Subject, len(*in))
		copy(*out, *in)
	}
	out.RoleRef = in.RoleRef
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRoleBinding.
func (x *ClusterRoleBinding) DeepCopy() *ClusterRoleBinding {
	if x == nil {
		return nil
	}
	out := new(ClusterRoleBinding)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *ClusterRoleBinding) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRoleBindingList) DeepCopyInto(out *ClusterRoleBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterRoleBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRoleBindingList.
func (x *ClusterRoleBindingList) DeepCopy() *ClusterRoleBindingList {
	if x == nil {
		return nil
	}
	out := new(ClusterRoleBindingList)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *ClusterRoleBindingList) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRoleList) DeepCopyInto(out *ClusterRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRoleList.
func (x *ClusterRoleList) DeepCopy() *ClusterRoleList {
	if x == nil {
		return nil
	}
	out := new(ClusterRoleList)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *ClusterRoleList) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyRule) DeepCopyInto(out *PolicyRule) {
	*out = *in
	if in.Verbs != nil {
		in, out := &in.Verbs, &out.Verbs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.APIGroups != nil {
		in, out := &in.APIGroups, &out.APIGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ResourceNames != nil {
		in, out := &in.ResourceNames, &out.ResourceNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NonResourceURLs != nil {
		in, out := &in.NonResourceURLs, &out.NonResourceURLs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new PolicyRule.
func (x *PolicyRule) DeepCopy() *PolicyRule {
	if x == nil {
		return nil
	}
	out := new(PolicyRule)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Role) DeepCopyInto(out *Role) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]PolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new Role.
func (x *Role) DeepCopy() *Role {
	if x == nil {
		return nil
	}
	out := new(Role)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *Role) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleBinding) DeepCopyInto(out *RoleBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]Subject, len(*in))
		copy(*out, *in)
	}
	out.RoleRef = in.RoleRef
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new RoleBinding.
func (x *RoleBinding) DeepCopy() *RoleBinding {
	if x == nil {
		return nil
	}
	out := new(RoleBinding)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *RoleBinding) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleBindingList) DeepCopyInto(out *RoleBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RoleBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new RoleBindingList.
func (x *RoleBindingList) DeepCopy() *RoleBindingList {
	if x == nil {
		return nil
	}
	out := new(RoleBindingList)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *RoleBindingList) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleList) DeepCopyInto(out *RoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Role, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new RoleList.
func (x *RoleList) DeepCopy() *RoleList {
	if x == nil {
		return nil
	}
	out := new(RoleList)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *RoleList) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleRef) DeepCopyInto(out *RoleRef) {
	*out = *in
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new RoleRef.
func (x *RoleRef) DeepCopy() *RoleRef {
	if x == nil {
		return nil
	}
	out := new(RoleRef)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subject) DeepCopyInto(out *Subject) {
	*out = *in
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new Subject.
func (x *Subject) DeepCopy() *Subject {
	if x == nil {
		return nil
	}
	out := new(Subject)
	x.DeepCopyInto(out)
	return out
}
