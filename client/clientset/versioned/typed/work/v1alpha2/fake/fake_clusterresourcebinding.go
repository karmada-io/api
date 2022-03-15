/*
Copyright 2022 The Karmada Authors.

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
	v1alpha2 "github.com/karmada-io/api/work/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterResourceBindings implements ClusterResourceBindingInterface
type FakeClusterResourceBindings struct {
	Fake *FakeWorkV1alpha2
}

var clusterresourcebindingsResource = schema.GroupVersionResource{Group: "work.karmada.io", Version: "v1alpha2", Resource: "clusterresourcebindings"}

var clusterresourcebindingsKind = schema.GroupVersionKind{Group: "work.karmada.io", Version: "v1alpha2", Kind: "ClusterResourceBinding"}

// Get takes name of the clusterResourceBinding, and returns the corresponding clusterResourceBinding object, and an error if there is any.
func (c *FakeClusterResourceBindings) Get(name string, options v1.GetOptions) (result *v1alpha2.ClusterResourceBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterresourcebindingsResource, name), &v1alpha2.ClusterResourceBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ClusterResourceBinding), err
}

// List takes label and field selectors, and returns the list of ClusterResourceBindings that match those selectors.
func (c *FakeClusterResourceBindings) List(opts v1.ListOptions) (result *v1alpha2.ClusterResourceBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterresourcebindingsResource, clusterresourcebindingsKind, opts), &v1alpha2.ClusterResourceBindingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.ClusterResourceBindingList{}
	for _, item := range obj.(*v1alpha2.ClusterResourceBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterResourceBindings.
func (c *FakeClusterResourceBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterresourcebindingsResource, opts))
}

// Create takes the representation of a clusterResourceBinding and creates it.  Returns the server's representation of the clusterResourceBinding, and an error, if there is any.
func (c *FakeClusterResourceBindings) Create(clusterResourceBinding *v1alpha2.ClusterResourceBinding) (result *v1alpha2.ClusterResourceBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterresourcebindingsResource, clusterResourceBinding), &v1alpha2.ClusterResourceBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ClusterResourceBinding), err
}

// Update takes the representation of a clusterResourceBinding and updates it. Returns the server's representation of the clusterResourceBinding, and an error, if there is any.
func (c *FakeClusterResourceBindings) Update(clusterResourceBinding *v1alpha2.ClusterResourceBinding) (result *v1alpha2.ClusterResourceBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterresourcebindingsResource, clusterResourceBinding), &v1alpha2.ClusterResourceBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ClusterResourceBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterResourceBindings) UpdateStatus(clusterResourceBinding *v1alpha2.ClusterResourceBinding) (*v1alpha2.ClusterResourceBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterresourcebindingsResource, "status", clusterResourceBinding), &v1alpha2.ClusterResourceBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ClusterResourceBinding), err
}

// Delete takes name of the clusterResourceBinding and deletes it. Returns an error if one occurs.
func (c *FakeClusterResourceBindings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clusterresourcebindingsResource, name), &v1alpha2.ClusterResourceBinding{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterResourceBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterresourcebindingsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha2.ClusterResourceBindingList{})
	return err
}

// Patch applies the patch and returns the patched clusterResourceBinding.
func (c *FakeClusterResourceBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.ClusterResourceBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterresourcebindingsResource, name, data, subresources...), &v1alpha2.ClusterResourceBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ClusterResourceBinding), err
}
