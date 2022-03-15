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
	v1alpha1 "github.com/karmada-io/api/policy/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFederatedResourceQuotas implements FederatedResourceQuotaInterface
type FakeFederatedResourceQuotas struct {
	Fake *FakePolicyV1alpha1
	ns   string
}

var federatedresourcequotasResource = schema.GroupVersionResource{Group: "policy.karmada.io", Version: "v1alpha1", Resource: "federatedresourcequotas"}

var federatedresourcequotasKind = schema.GroupVersionKind{Group: "policy.karmada.io", Version: "v1alpha1", Kind: "FederatedResourceQuota"}

// Get takes name of the federatedResourceQuota, and returns the corresponding federatedResourceQuota object, and an error if there is any.
func (c *FakeFederatedResourceQuotas) Get(name string, options v1.GetOptions) (result *v1alpha1.FederatedResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(federatedresourcequotasResource, c.ns, name), &v1alpha1.FederatedResourceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedResourceQuota), err
}

// List takes label and field selectors, and returns the list of FederatedResourceQuotas that match those selectors.
func (c *FakeFederatedResourceQuotas) List(opts v1.ListOptions) (result *v1alpha1.FederatedResourceQuotaList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(federatedresourcequotasResource, federatedresourcequotasKind, c.ns, opts), &v1alpha1.FederatedResourceQuotaList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FederatedResourceQuotaList{}
	for _, item := range obj.(*v1alpha1.FederatedResourceQuotaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested federatedResourceQuotas.
func (c *FakeFederatedResourceQuotas) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(federatedresourcequotasResource, c.ns, opts))

}

// Create takes the representation of a federatedResourceQuota and creates it.  Returns the server's representation of the federatedResourceQuota, and an error, if there is any.
func (c *FakeFederatedResourceQuotas) Create(federatedResourceQuota *v1alpha1.FederatedResourceQuota) (result *v1alpha1.FederatedResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(federatedresourcequotasResource, c.ns, federatedResourceQuota), &v1alpha1.FederatedResourceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedResourceQuota), err
}

// Update takes the representation of a federatedResourceQuota and updates it. Returns the server's representation of the federatedResourceQuota, and an error, if there is any.
func (c *FakeFederatedResourceQuotas) Update(federatedResourceQuota *v1alpha1.FederatedResourceQuota) (result *v1alpha1.FederatedResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(federatedresourcequotasResource, c.ns, federatedResourceQuota), &v1alpha1.FederatedResourceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedResourceQuota), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFederatedResourceQuotas) UpdateStatus(federatedResourceQuota *v1alpha1.FederatedResourceQuota) (*v1alpha1.FederatedResourceQuota, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(federatedresourcequotasResource, "status", c.ns, federatedResourceQuota), &v1alpha1.FederatedResourceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedResourceQuota), err
}

// Delete takes name of the federatedResourceQuota and deletes it. Returns an error if one occurs.
func (c *FakeFederatedResourceQuotas) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(federatedresourcequotasResource, c.ns, name), &v1alpha1.FederatedResourceQuota{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFederatedResourceQuotas) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(federatedresourcequotasResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.FederatedResourceQuotaList{})
	return err
}

// Patch applies the patch and returns the patched federatedResourceQuota.
func (c *FakeFederatedResourceQuotas) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FederatedResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(federatedresourcequotasResource, c.ns, name, data, subresources...), &v1alpha1.FederatedResourceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedResourceQuota), err
}
