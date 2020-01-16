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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type BundleView struct {
	metav1.TypeMeta `json:",inline"`
	PackageMeta     `json:",inline" protobuf:"bytes,1,opt,name=packageMeta"`

	Packages []PackageCard `json:"packages" protobuf:"bytes,2,rep,name=packages"`
	Addons   []AddonView   `json:"addons" protobuf:"bytes,3,rep,name=addons"`
}

type AddonView struct {
	Feature string             `json:"feature" protobuf:"bytes,1,opt,name=feature"`
	Bundle  *BundleOptionView  `json:"bundle,omitempty" protobuf:"bytes,2,opt,name=bundle"`
	OneOf   []BundleOptionView `json:"oneOf,omitempty" protobuf:"bytes,3,rep,name=oneOf"`
}

type BundleOptionView struct {
	BundleView `json:",inline" protobuf:"bytes,1,opt,name=bundleView"`
	Versions   []VersionOption `json:"versions" protobuf:"bytes,2,rep,name=versions"`
}
