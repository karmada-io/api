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

package v1alpha1

import (
	scheme "github.com/karmada-io/api/client/clientset/versioned/scheme"
	v1alpha1 "github.com/karmada-io/api/policy/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FederatedResourceQuotasGetter has a method to return a FederatedResourceQuotaInterface.
// A group's client should implement this interface.
type FederatedResourceQuotasGetter interface {
	FederatedResourceQuotas(namespace string) FederatedResourceQuotaInterface
}

// FederatedResourceQuotaInterface has methods to work with FederatedResourceQuota resources.
type FederatedResourceQuotaInterface interface {
	Create(*v1alpha1.FederatedResourceQuota) (*v1alpha1.FederatedResourceQuota, error)
	Update(*v1alpha1.FederatedResourceQuota) (*v1alpha1.FederatedResourceQuota, error)
	UpdateStatus(*v1alpha1.FederatedResourceQuota) (*v1alpha1.FederatedResourceQuota, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.FederatedResourceQuota, error)
	List(opts v1.ListOptions) (*v1alpha1.FederatedResourceQuotaList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FederatedResourceQuota, err error)
	FederatedResourceQuotaExpansion
}

// federatedResourceQuotas implements FederatedResourceQuotaInterface
type federatedResourceQuotas struct {
	client rest.Interface
	ns     string
}

// newFederatedResourceQuotas returns a FederatedResourceQuotas
func newFederatedResourceQuotas(c *PolicyV1alpha1Client, namespace string) *federatedResourceQuotas {
	return &federatedResourceQuotas{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the federatedResourceQuota, and returns the corresponding federatedResourceQuota object, and an error if there is any.
func (c *federatedResourceQuotas) Get(name string, options v1.GetOptions) (result *v1alpha1.FederatedResourceQuota, err error) {
	result = &v1alpha1.FederatedResourceQuota{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedresourcequotas").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FederatedResourceQuotas that match those selectors.
func (c *federatedResourceQuotas) List(opts v1.ListOptions) (result *v1alpha1.FederatedResourceQuotaList, err error) {
	result = &v1alpha1.FederatedResourceQuotaList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedresourcequotas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested federatedResourceQuotas.
func (c *federatedResourceQuotas) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("federatedresourcequotas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a federatedResourceQuota and creates it.  Returns the server's representation of the federatedResourceQuota, and an error, if there is any.
func (c *federatedResourceQuotas) Create(federatedResourceQuota *v1alpha1.FederatedResourceQuota) (result *v1alpha1.FederatedResourceQuota, err error) {
	result = &v1alpha1.FederatedResourceQuota{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("federatedresourcequotas").
		Body(federatedResourceQuota).
		Do().
		Into(result)
	return
}

// Update takes the representation of a federatedResourceQuota and updates it. Returns the server's representation of the federatedResourceQuota, and an error, if there is any.
func (c *federatedResourceQuotas) Update(federatedResourceQuota *v1alpha1.FederatedResourceQuota) (result *v1alpha1.FederatedResourceQuota, err error) {
	result = &v1alpha1.FederatedResourceQuota{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedresourcequotas").
		Name(federatedResourceQuota.Name).
		Body(federatedResourceQuota).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *federatedResourceQuotas) UpdateStatus(federatedResourceQuota *v1alpha1.FederatedResourceQuota) (result *v1alpha1.FederatedResourceQuota, err error) {
	result = &v1alpha1.FederatedResourceQuota{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedresourcequotas").
		Name(federatedResourceQuota.Name).
		SubResource("status").
		Body(federatedResourceQuota).
		Do().
		Into(result)
	return
}

// Delete takes name of the federatedResourceQuota and deletes it. Returns an error if one occurs.
func (c *federatedResourceQuotas) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedresourcequotas").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *federatedResourceQuotas) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedresourcequotas").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched federatedResourceQuota.
func (c *federatedResourceQuotas) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FederatedResourceQuota, err error) {
	result = &v1alpha1.FederatedResourceQuota{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("federatedresourcequotas").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
