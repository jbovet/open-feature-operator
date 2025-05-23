//go:build !ignore_autogenerated

/*
Copyright 2022.

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
	"encoding/json"
	"k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apisv1 "sigs.k8s.io/gateway-api/apis/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureFlag) DeepCopyInto(out *FeatureFlag) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureFlag.
func (in *FeatureFlag) DeepCopy() *FeatureFlag {
	if in == nil {
		return nil
	}
	out := new(FeatureFlag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FeatureFlag) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureFlagList) DeepCopyInto(out *FeatureFlagList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FeatureFlag, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureFlagList.
func (in *FeatureFlagList) DeepCopy() *FeatureFlagList {
	if in == nil {
		return nil
	}
	out := new(FeatureFlagList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FeatureFlagList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureFlagSource) DeepCopyInto(out *FeatureFlagSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureFlagSource.
func (in *FeatureFlagSource) DeepCopy() *FeatureFlagSource {
	if in == nil {
		return nil
	}
	out := new(FeatureFlagSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FeatureFlagSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureFlagSourceList) DeepCopyInto(out *FeatureFlagSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FeatureFlagSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureFlagSourceList.
func (in *FeatureFlagSourceList) DeepCopy() *FeatureFlagSourceList {
	if in == nil {
		return nil
	}
	out := new(FeatureFlagSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FeatureFlagSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureFlagSourceSpec) DeepCopyInto(out *FeatureFlagSourceSpec) {
	*out = *in
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make([]Source, len(*in))
		copy(*out, *in)
	}
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SyncProviderArgs != nil {
		in, out := &in.SyncProviderArgs, &out.SyncProviderArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RolloutOnChange != nil {
		in, out := &in.RolloutOnChange, &out.RolloutOnChange
		*out = new(bool)
		**out = **in
	}
	if in.ProbesEnabled != nil {
		in, out := &in.ProbesEnabled, &out.ProbesEnabled
		*out = new(bool)
		**out = **in
	}
	if in.DebugLogging != nil {
		in, out := &in.DebugLogging, &out.DebugLogging
		*out = new(bool)
		**out = **in
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureFlagSourceSpec.
func (in *FeatureFlagSourceSpec) DeepCopy() *FeatureFlagSourceSpec {
	if in == nil {
		return nil
	}
	out := new(FeatureFlagSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureFlagSourceStatus) DeepCopyInto(out *FeatureFlagSourceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureFlagSourceStatus.
func (in *FeatureFlagSourceStatus) DeepCopy() *FeatureFlagSourceStatus {
	if in == nil {
		return nil
	}
	out := new(FeatureFlagSourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureFlagSpec) DeepCopyInto(out *FeatureFlagSpec) {
	*out = *in
	in.FlagSpec.DeepCopyInto(&out.FlagSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureFlagSpec.
func (in *FeatureFlagSpec) DeepCopy() *FeatureFlagSpec {
	if in == nil {
		return nil
	}
	out := new(FeatureFlagSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureFlagStatus) DeepCopyInto(out *FeatureFlagStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureFlagStatus.
func (in *FeatureFlagStatus) DeepCopy() *FeatureFlagStatus {
	if in == nil {
		return nil
	}
	out := new(FeatureFlagStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Flag) DeepCopyInto(out *Flag) {
	*out = *in
	if in.Variants != nil {
		in, out := &in.Variants, &out.Variants
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	if in.Targeting != nil {
		in, out := &in.Targeting, &out.Targeting
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Flag.
func (in *Flag) DeepCopy() *Flag {
	if in == nil {
		return nil
	}
	out := new(Flag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlagSpec) DeepCopyInto(out *FlagSpec) {
	*out = *in
	in.Flags.DeepCopyInto(&out.Flags)
	if in.Evaluators != nil {
		in, out := &in.Evaluators, &out.Evaluators
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlagSpec.
func (in *FlagSpec) DeepCopy() *FlagSpec {
	if in == nil {
		return nil
	}
	out := new(FlagSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Flagd) DeepCopyInto(out *Flagd) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Flagd.
func (in *Flagd) DeepCopy() *Flagd {
	if in == nil {
		return nil
	}
	out := new(Flagd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Flagd) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlagdList) DeepCopyInto(out *FlagdList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Flagd, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlagdList.
func (in *FlagdList) DeepCopy() *FlagdList {
	if in == nil {
		return nil
	}
	out := new(FlagdList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlagdList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlagdSpec) DeepCopyInto(out *FlagdSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	in.Ingress.DeepCopyInto(&out.Ingress)
	in.GatewayApiRoutes.DeepCopyInto(&out.GatewayApiRoutes)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlagdSpec.
func (in *FlagdSpec) DeepCopy() *FlagdSpec {
	if in == nil {
		return nil
	}
	out := new(FlagdSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlagdStatus) DeepCopyInto(out *FlagdStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlagdStatus.
func (in *FlagdStatus) DeepCopy() *FlagdStatus {
	if in == nil {
		return nil
	}
	out := new(FlagdStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Flags) DeepCopyInto(out *Flags) {
	*out = *in
	if in.FlagsMap != nil {
		in, out := &in.FlagsMap, &out.FlagsMap
		*out = make(map[string]Flag, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Flags.
func (in *Flags) DeepCopy() *Flags {
	if in == nil {
		return nil
	}
	out := new(Flags)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayApiSpec) DeepCopyInto(out *GatewayApiSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ParentRefs != nil {
		in, out := &in.ParentRefs, &out.ParentRefs
		*out = make([]apisv1.ParentReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayApiSpec.
func (in *GatewayApiSpec) DeepCopy() *GatewayApiSpec {
	if in == nil {
		return nil
	}
	out := new(GatewayApiSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InProcessConfiguration) DeepCopyInto(out *InProcessConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InProcessConfiguration.
func (in *InProcessConfiguration) DeepCopy() *InProcessConfiguration {
	if in == nil {
		return nil
	}
	out := new(InProcessConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InProcessConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InProcessConfigurationList) DeepCopyInto(out *InProcessConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InProcessConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InProcessConfigurationList.
func (in *InProcessConfigurationList) DeepCopy() *InProcessConfigurationList {
	if in == nil {
		return nil
	}
	out := new(InProcessConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InProcessConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InProcessConfigurationSpec) DeepCopyInto(out *InProcessConfigurationSpec) {
	*out = *in
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InProcessConfigurationSpec.
func (in *InProcessConfigurationSpec) DeepCopy() *InProcessConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(InProcessConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InProcessConfigurationStatus) DeepCopyInto(out *InProcessConfigurationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InProcessConfigurationStatus.
func (in *InProcessConfigurationStatus) DeepCopy() *InProcessConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(InProcessConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressSpec) DeepCopyInto(out *IngressSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = make([]networkingv1.IngressTLS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IngressClassName != nil {
		in, out := &in.IngressClassName, &out.IngressClassName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressSpec.
func (in *IngressSpec) DeepCopy() *IngressSpec {
	if in == nil {
		return nil
	}
	out := new(IngressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Source) DeepCopyInto(out *Source) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Source.
func (in *Source) DeepCopy() *Source {
	if in == nil {
		return nil
	}
	out := new(Source)
	in.DeepCopyInto(out)
	return out
}
