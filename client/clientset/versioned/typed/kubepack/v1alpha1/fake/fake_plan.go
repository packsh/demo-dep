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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "kubepack.dev/kubepack/apis/kubepack/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePlans implements PlanInterface
type FakePlans struct {
	Fake *FakeKubepackV1alpha1
}

var plansResource = schema.GroupVersionResource{Group: "kubepack.com", Version: "v1alpha1", Resource: "plans"}

var plansKind = schema.GroupVersionKind{Group: "kubepack.com", Version: "v1alpha1", Kind: "Plan"}

// Get takes name of the plan, and returns the corresponding plan object, and an error if there is any.
func (c *FakePlans) Get(name string, options v1.GetOptions) (result *v1alpha1.Plan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(plansResource, name), &v1alpha1.Plan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Plan), err
}

// List takes label and field selectors, and returns the list of Plans that match those selectors.
func (c *FakePlans) List(opts v1.ListOptions) (result *v1alpha1.PlanList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(plansResource, plansKind, opts), &v1alpha1.PlanList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PlanList{ListMeta: obj.(*v1alpha1.PlanList).ListMeta}
	for _, item := range obj.(*v1alpha1.PlanList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested plans.
func (c *FakePlans) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(plansResource, opts))
}

// Create takes the representation of a plan and creates it.  Returns the server's representation of the plan, and an error, if there is any.
func (c *FakePlans) Create(plan *v1alpha1.Plan) (result *v1alpha1.Plan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(plansResource, plan), &v1alpha1.Plan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Plan), err
}

// Update takes the representation of a plan and updates it. Returns the server's representation of the plan, and an error, if there is any.
func (c *FakePlans) Update(plan *v1alpha1.Plan) (result *v1alpha1.Plan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(plansResource, plan), &v1alpha1.Plan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Plan), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePlans) UpdateStatus(plan *v1alpha1.Plan) (*v1alpha1.Plan, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(plansResource, "status", plan), &v1alpha1.Plan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Plan), err
}

// Delete takes name of the plan and deletes it. Returns an error if one occurs.
func (c *FakePlans) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(plansResource, name), &v1alpha1.Plan{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePlans) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(plansResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PlanList{})
	return err
}

// Patch applies the patch and returns the patched plan.
func (c *FakePlans) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Plan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(plansResource, name, pt, data, subresources...), &v1alpha1.Plan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Plan), err
}
