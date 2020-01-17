// +build !ignore_autogenerated

/*
Copyright The Kubepack Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Application) DeepCopyInto(out *Application) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Application.
func (in *Application) DeepCopy() *Application {
	if in == nil {
		return nil
	}
	out := new(Application)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Application) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationList) DeepCopyInto(out *ApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Application, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationList.
func (in *ApplicationList) DeepCopy() *ApplicationList {
	if in == nil {
		return nil
	}
	out := new(ApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSpec) DeepCopyInto(out *ApplicationSpec) {
	*out = *in
	if in.ComponentGroupKinds != nil {
		in, out := &in.ComponentGroupKinds, &out.ComponentGroupKinds
		*out = make([]v1.GroupKind, len(*in))
		copy(*out, *in)
	}
	in.Description.DeepCopyInto(&out.Description)
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Info != nil {
		in, out := &in.Info, &out.Info
		*out = make([]InfoItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSpec.
func (in *ApplicationSpec) DeepCopy() *ApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationStatus) DeepCopyInto(out *ApplicationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.ComponentList.DeepCopyInto(&out.ComponentList)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationStatus.
func (in *ApplicationStatus) DeepCopy() *ApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bundle) DeepCopyInto(out *Bundle) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bundle.
func (in *Bundle) DeepCopy() *Bundle {
	if in == nil {
		return nil
	}
	out := new(Bundle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Bundle) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleList) DeepCopyInto(out *BundleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Bundle, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleList.
func (in *BundleList) DeepCopy() *BundleList {
	if in == nil {
		return nil
	}
	out := new(BundleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BundleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleOption) DeepCopyInto(out *BundleOption) {
	*out = *in
	out.BundleRef = in.BundleRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleOption.
func (in *BundleOption) DeepCopy() *BundleOption {
	if in == nil {
		return nil
	}
	out := new(BundleOption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleOptionView) DeepCopyInto(out *BundleOptionView) {
	*out = *in
	in.PackageMeta.DeepCopyInto(&out.PackageMeta)
	if in.Packages != nil {
		in, out := &in.Packages, &out.Packages
		*out = make([]PackageCard, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleOptionView.
func (in *BundleOptionView) DeepCopy() *BundleOptionView {
	if in == nil {
		return nil
	}
	out := new(BundleOptionView)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleRef) DeepCopyInto(out *BundleRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleRef.
func (in *BundleRef) DeepCopy() *BundleRef {
	if in == nil {
		return nil
	}
	out := new(BundleRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleSpec) DeepCopyInto(out *BundleSpec) {
	*out = *in
	in.PackageDescriptor.DeepCopyInto(&out.PackageDescriptor)
	if in.Packages != nil {
		in, out := &in.Packages, &out.Packages
		*out = make([]PackageRef, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleSpec.
func (in *BundleSpec) DeepCopy() *BundleSpec {
	if in == nil {
		return nil
	}
	out := new(BundleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleStatus) DeepCopyInto(out *BundleStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleStatus.
func (in *BundleStatus) DeepCopy() *BundleStatus {
	if in == nil {
		return nil
	}
	out := new(BundleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleView) DeepCopyInto(out *BundleView) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.BundleOptionView.DeepCopyInto(&out.BundleOptionView)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleView.
func (in *BundleView) DeepCopy() *BundleView {
	if in == nil {
		return nil
	}
	out := new(BundleView)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BundleView) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartCard) DeepCopyInto(out *ChartCard) {
	*out = *in
	in.ChartRef.DeepCopyInto(&out.ChartRef)
	in.PackageDescriptor.DeepCopyInto(&out.PackageDescriptor)
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]VersionOption, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartCard.
func (in *ChartCard) DeepCopy() *ChartCard {
	if in == nil {
		return nil
	}
	out := new(ChartCard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartOption) DeepCopyInto(out *ChartOption) {
	*out = *in
	in.ChartRef.DeepCopyInto(&out.ChartRef)
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]VersionDetail, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartOption.
func (in *ChartOption) DeepCopy() *ChartOption {
	if in == nil {
		return nil
	}
	out := new(ChartOption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartRef) DeepCopyInto(out *ChartRef) {
	*out = *in
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartRef.
func (in *ChartRef) DeepCopy() *ChartRef {
	if in == nil {
		return nil
	}
	out := new(ChartRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartSelection) DeepCopyInto(out *ChartSelection) {
	*out = *in
	in.ChartRef.DeepCopyInto(&out.ChartRef)
	if in.ValuesPatch != nil {
		in, out := &in.ValuesPatch, &out.ValuesPatch
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(ResourceDefinitions)
		(*in).DeepCopyInto(*out)
	}
	if in.WaitFors != nil {
		in, out := &in.WaitFors, &out.WaitFors
		*out = make([]WaitOptions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartSelection.
func (in *ChartSelection) DeepCopy() *ChartSelection {
	if in == nil {
		return nil
	}
	out := new(ChartSelection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartVersionRef) DeepCopyInto(out *ChartVersionRef) {
	*out = *in
	in.ChartRef.DeepCopyInto(&out.ChartRef)
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartVersionRef.
func (in *ChartVersionRef) DeepCopy() *ChartVersionRef {
	if in == nil {
		return nil
	}
	out := new(ChartVersionRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentList) DeepCopyInto(out *ComponentList) {
	*out = *in
	if in.Objects != nil {
		in, out := &in.Objects, &out.Objects
		*out = make([]ObjectStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentList.
func (in *ComponentList) DeepCopy() *ComponentList {
	if in == nil {
		return nil
	}
	out := new(ComponentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapKeySelector) DeepCopyInto(out *ConfigMapKeySelector) {
	*out = *in
	out.ObjectReference = in.ObjectReference
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapKeySelector.
func (in *ConfigMapKeySelector) DeepCopy() *ConfigMapKeySelector {
	if in == nil {
		return nil
	}
	out := new(ConfigMapKeySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactData) DeepCopyInto(out *ContactData) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactData.
func (in *ContactData) DeepCopy() *ContactData {
	if in == nil {
		return nil
	}
	out := new(ContactData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Descriptor) DeepCopyInto(out *Descriptor) {
	*out = *in
	if in.Icons != nil {
		in, out := &in.Icons, &out.Icons
		*out = make([]ImageSpec, len(*in))
		copy(*out, *in)
	}
	if in.Maintainers != nil {
		in, out := &in.Maintainers, &out.Maintainers
		*out = make([]ContactData, len(*in))
		copy(*out, *in)
	}
	if in.Owners != nil {
		in, out := &in.Owners, &out.Owners
		*out = make([]ContactData, len(*in))
		copy(*out, *in)
	}
	if in.Keywords != nil {
		in, out := &in.Keywords, &out.Keywords
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]Link, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Descriptor.
func (in *Descriptor) DeepCopy() *Descriptor {
	if in == nil {
		return nil
	}
	out := new(Descriptor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupVersionResource) DeepCopyInto(out *GroupVersionResource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupVersionResource.
func (in *GroupVersionResource) DeepCopy() *GroupVersionResource {
	if in == nil {
		return nil
	}
	out := new(GroupVersionResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfoItem) DeepCopyInto(out *InfoItem) {
	*out = *in
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(InfoItemSource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfoItem.
func (in *InfoItem) DeepCopy() *InfoItem {
	if in == nil {
		return nil
	}
	out := new(InfoItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfoItemSource) DeepCopyInto(out *InfoItemSource) {
	*out = *in
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(SecretKeySelector)
		**out = **in
	}
	if in.ConfigMapKeyRef != nil {
		in, out := &in.ConfigMapKeyRef, &out.ConfigMapKeyRef
		*out = new(ConfigMapKeySelector)
		**out = **in
	}
	if in.ServiceRef != nil {
		in, out := &in.ServiceRef, &out.ServiceRef
		*out = new(ServiceSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.IngressRef != nil {
		in, out := &in.IngressRef, &out.IngressRef
		*out = new(IngressSelector)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfoItemSource.
func (in *InfoItemSource) DeepCopy() *InfoItemSource {
	if in == nil {
		return nil
	}
	out := new(InfoItemSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressSelector) DeepCopyInto(out *IngressSelector) {
	*out = *in
	out.ObjectReference = in.ObjectReference
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressSelector.
func (in *IngressSelector) DeepCopy() *IngressSelector {
	if in == nil {
		return nil
	}
	out := new(IngressSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Link) DeepCopyInto(out *Link) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Link.
func (in *Link) DeepCopy() *Link {
	if in == nil {
		return nil
	}
	out := new(Link)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStatus) DeepCopyInto(out *ObjectStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStatus.
func (in *ObjectStatus) DeepCopy() *ObjectStatus {
	if in == nil {
		return nil
	}
	out := new(ObjectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Order) DeepCopyInto(out *Order) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Order.
func (in *Order) DeepCopy() *Order {
	if in == nil {
		return nil
	}
	out := new(Order)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Order) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderList) DeepCopyInto(out *OrderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Order, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderList.
func (in *OrderList) DeepCopy() *OrderList {
	if in == nil {
		return nil
	}
	out := new(OrderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderSpec) DeepCopyInto(out *OrderSpec) {
	*out = *in
	if in.Packages != nil {
		in, out := &in.Packages, &out.Packages
		*out = make([]PackageSelection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderSpec.
func (in *OrderSpec) DeepCopy() *OrderSpec {
	if in == nil {
		return nil
	}
	out := new(OrderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderStatus) DeepCopyInto(out *OrderStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderStatus.
func (in *OrderStatus) DeepCopy() *OrderStatus {
	if in == nil {
		return nil
	}
	out := new(OrderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pack) DeepCopyInto(out *Pack) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pack.
func (in *Pack) DeepCopy() *Pack {
	if in == nil {
		return nil
	}
	out := new(Pack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pack) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackList) DeepCopyInto(out *PackList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pack, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackList.
func (in *PackList) DeepCopy() *PackList {
	if in == nil {
		return nil
	}
	out := new(PackList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PackList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackSpec) DeepCopyInto(out *PackSpec) {
	*out = *in
	if in.Chart != nil {
		in, out := &in.Chart, &out.Chart
		*out = new(ChartRef)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(ResourceDefinitions)
		(*in).DeepCopyInto(*out)
	}
	if in.WaitFors != nil {
		in, out := &in.WaitFors, &out.WaitFors
		*out = make([]WaitOptions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackSpec.
func (in *PackSpec) DeepCopy() *PackSpec {
	if in == nil {
		return nil
	}
	out := new(PackSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackStatus) DeepCopyInto(out *PackStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackStatus.
func (in *PackStatus) DeepCopy() *PackStatus {
	if in == nil {
		return nil
	}
	out := new(PackStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageCard) DeepCopyInto(out *PackageCard) {
	*out = *in
	if in.Chart != nil {
		in, out := &in.Chart, &out.Chart
		*out = new(ChartCard)
		(*in).DeepCopyInto(*out)
	}
	if in.Bundle != nil {
		in, out := &in.Bundle, &out.Bundle
		*out = new(BundleOptionView)
		(*in).DeepCopyInto(*out)
	}
	if in.OneOf != nil {
		in, out := &in.OneOf, &out.OneOf
		*out = make([]*BundleOptionView, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(BundleOptionView)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageCard.
func (in *PackageCard) DeepCopy() *PackageCard {
	if in == nil {
		return nil
	}
	out := new(PackageCard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageDescriptor) DeepCopyInto(out *PackageDescriptor) {
	*out = *in
	if in.Icons != nil {
		in, out := &in.Icons, &out.Icons
		*out = make([]ImageSpec, len(*in))
		copy(*out, *in)
	}
	if in.Maintainers != nil {
		in, out := &in.Maintainers, &out.Maintainers
		*out = make([]ContactData, len(*in))
		copy(*out, *in)
	}
	if in.Owners != nil {
		in, out := &in.Owners, &out.Owners
		*out = make([]ContactData, len(*in))
		copy(*out, *in)
	}
	if in.Keywords != nil {
		in, out := &in.Keywords, &out.Keywords
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]Link, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageDescriptor.
func (in *PackageDescriptor) DeepCopy() *PackageDescriptor {
	if in == nil {
		return nil
	}
	out := new(PackageDescriptor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageMeta) DeepCopyInto(out *PackageMeta) {
	*out = *in
	in.PackageDescriptor.DeepCopyInto(&out.PackageDescriptor)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageMeta.
func (in *PackageMeta) DeepCopy() *PackageMeta {
	if in == nil {
		return nil
	}
	out := new(PackageMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageRef) DeepCopyInto(out *PackageRef) {
	*out = *in
	if in.Chart != nil {
		in, out := &in.Chart, &out.Chart
		*out = new(ChartOption)
		(*in).DeepCopyInto(*out)
	}
	if in.Bundle != nil {
		in, out := &in.Bundle, &out.Bundle
		*out = new(BundleOption)
		**out = **in
	}
	if in.OneOf != nil {
		in, out := &in.OneOf, &out.OneOf
		*out = make([]*BundleOption, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(BundleOption)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageRef.
func (in *PackageRef) DeepCopy() *PackageRef {
	if in == nil {
		return nil
	}
	out := new(PackageRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageSelection) DeepCopyInto(out *PackageSelection) {
	*out = *in
	if in.Chart != nil {
		in, out := &in.Chart, &out.Chart
		*out = new(ChartSelection)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageSelection.
func (in *PackageSelection) DeepCopy() *PackageSelection {
	if in == nil {
		return nil
	}
	out := new(PackageSelection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageView) DeepCopyInto(out *PackageView) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.PackageMeta.DeepCopyInto(&out.PackageMeta)
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.Validation != nil {
		in, out := &in.Validation, &out.Validation
		*out = new(v1beta1.CustomResourceValidation)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageView.
func (in *PackageView) DeepCopy() *PackageView {
	if in == nil {
		return nil
	}
	out := new(PackageView)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PackageView) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Repo) DeepCopyInto(out *Repo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Repo.
func (in *Repo) DeepCopy() *Repo {
	if in == nil {
		return nil
	}
	out := new(Repo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceDefinitions) DeepCopyInto(out *ResourceDefinitions) {
	*out = *in
	if in.Owned != nil {
		in, out := &in.Owned, &out.Owned
		*out = make([]ResourceID, len(*in))
		copy(*out, *in)
	}
	if in.Required != nil {
		in, out := &in.Required, &out.Required
		*out = make([]ResourceID, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceDefinitions.
func (in *ResourceDefinitions) DeepCopy() *ResourceDefinitions {
	if in == nil {
		return nil
	}
	out := new(ResourceDefinitions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceID) DeepCopyInto(out *ResourceID) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceID.
func (in *ResourceID) DeepCopy() *ResourceID {
	if in == nil {
		return nil
	}
	out := new(ResourceID)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeySelector) DeepCopyInto(out *SecretKeySelector) {
	*out = *in
	out.ObjectReference = in.ObjectReference
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeySelector.
func (in *SecretKeySelector) DeepCopy() *SecretKeySelector {
	if in == nil {
		return nil
	}
	out := new(SecretKeySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSelector) DeepCopyInto(out *ServiceSelector) {
	*out = *in
	out.ObjectReference = in.ObjectReference
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSelector.
func (in *ServiceSelector) DeepCopy() *ServiceSelector {
	if in == nil {
		return nil
	}
	out := new(ServiceSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionDetail) DeepCopyInto(out *VersionDetail) {
	*out = *in
	in.VersionOption.DeepCopyInto(&out.VersionOption)
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(ResourceDefinitions)
		(*in).DeepCopyInto(*out)
	}
	if in.WaitFors != nil {
		in, out := &in.WaitFors, &out.WaitFors
		*out = make([]WaitOptions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionDetail.
func (in *VersionDetail) DeepCopy() *VersionDetail {
	if in == nil {
		return nil
	}
	out := new(VersionDetail)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionOption) DeepCopyInto(out *VersionOption) {
	*out = *in
	if in.ValuesPatch != nil {
		in, out := &in.ValuesPatch, &out.ValuesPatch
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionOption.
func (in *VersionOption) DeepCopy() *VersionOption {
	if in == nil {
		return nil
	}
	out := new(VersionOption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WaitOptions) DeepCopyInto(out *WaitOptions) {
	*out = *in
	out.Resource = in.Resource
	in.Labels.DeepCopyInto(&out.Labels)
	out.Timeout = in.Timeout
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WaitOptions.
func (in *WaitOptions) DeepCopy() *WaitOptions {
	if in == nil {
		return nil
	}
	out := new(WaitOptions)
	in.DeepCopyInto(out)
	return out
}
