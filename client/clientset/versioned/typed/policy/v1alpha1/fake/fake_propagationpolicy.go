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

// FakePropagationPolicies implements PropagationPolicyInterface
type FakePropagationPolicies struct {
	Fake *FakePolicyV1alpha1
	ns   string
}

var propagationpoliciesResource = schema.GroupVersionResource{Group: "policy.karmada.io", Version: "v1alpha1", Resource: "propagationpolicies"}

var propagationpoliciesKind = schema.GroupVersionKind{Group: "policy.karmada.io", Version: "v1alpha1", Kind: "PropagationPolicy"}

// Get takes name of the propagationPolicy, and returns the corresponding propagationPolicy object, and an error if there is any.
func (c *FakePropagationPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.PropagationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(propagationpoliciesResource, c.ns, name), &v1alpha1.PropagationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PropagationPolicy), err
}

// List takes label and field selectors, and returns the list of PropagationPolicies that match those selectors.
func (c *FakePropagationPolicies) List(opts v1.ListOptions) (result *v1alpha1.PropagationPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(propagationpoliciesResource, propagationpoliciesKind, c.ns, opts), &v1alpha1.PropagationPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PropagationPolicyList{}
	for _, item := range obj.(*v1alpha1.PropagationPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested propagationPolicies.
func (c *FakePropagationPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(propagationpoliciesResource, c.ns, opts))

}

// Create takes the representation of a propagationPolicy and creates it.  Returns the server's representation of the propagationPolicy, and an error, if there is any.
func (c *FakePropagationPolicies) Create(propagationPolicy *v1alpha1.PropagationPolicy) (result *v1alpha1.PropagationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(propagationpoliciesResource, c.ns, propagationPolicy), &v1alpha1.PropagationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PropagationPolicy), err
}

// Update takes the representation of a propagationPolicy and updates it. Returns the server's representation of the propagationPolicy, and an error, if there is any.
func (c *FakePropagationPolicies) Update(propagationPolicy *v1alpha1.PropagationPolicy) (result *v1alpha1.PropagationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(propagationpoliciesResource, c.ns, propagationPolicy), &v1alpha1.PropagationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PropagationPolicy), err
}

// Delete takes name of the propagationPolicy and deletes it. Returns an error if one occurs.
func (c *FakePropagationPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(propagationpoliciesResource, c.ns, name), &v1alpha1.PropagationPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePropagationPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(propagationpoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PropagationPolicyList{})
	return err
}

// Patch applies the patch and returns the patched propagationPolicy.
func (c *FakePropagationPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PropagationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(propagationpoliciesResource, c.ns, name, data, subresources...), &v1alpha1.PropagationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PropagationPolicy), err
}
