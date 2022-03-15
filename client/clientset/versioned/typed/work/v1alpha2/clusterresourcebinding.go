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

package v1alpha2

import (
	scheme "github.com/karmada-io/api/client/clientset/versioned/scheme"
	v1alpha2 "github.com/karmada-io/api/work/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterResourceBindingsGetter has a method to return a ClusterResourceBindingInterface.
// A group's client should implement this interface.
type ClusterResourceBindingsGetter interface {
	ClusterResourceBindings() ClusterResourceBindingInterface
}

// ClusterResourceBindingInterface has methods to work with ClusterResourceBinding resources.
type ClusterResourceBindingInterface interface {
	Create(*v1alpha2.ClusterResourceBinding) (*v1alpha2.ClusterResourceBinding, error)
	Update(*v1alpha2.ClusterResourceBinding) (*v1alpha2.ClusterResourceBinding, error)
	UpdateStatus(*v1alpha2.ClusterResourceBinding) (*v1alpha2.ClusterResourceBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.ClusterResourceBinding, error)
	List(opts v1.ListOptions) (*v1alpha2.ClusterResourceBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.ClusterResourceBinding, err error)
	ClusterResourceBindingExpansion
}

// clusterResourceBindings implements ClusterResourceBindingInterface
type clusterResourceBindings struct {
	client rest.Interface
}

// newClusterResourceBindings returns a ClusterResourceBindings
func newClusterResourceBindings(c *WorkV1alpha2Client) *clusterResourceBindings {
	return &clusterResourceBindings{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterResourceBinding, and returns the corresponding clusterResourceBinding object, and an error if there is any.
func (c *clusterResourceBindings) Get(name string, options v1.GetOptions) (result *v1alpha2.ClusterResourceBinding, err error) {
	result = &v1alpha2.ClusterResourceBinding{}
	err = c.client.Get().
		Resource("clusterresourcebindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterResourceBindings that match those selectors.
func (c *clusterResourceBindings) List(opts v1.ListOptions) (result *v1alpha2.ClusterResourceBindingList, err error) {
	result = &v1alpha2.ClusterResourceBindingList{}
	err = c.client.Get().
		Resource("clusterresourcebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterResourceBindings.
func (c *clusterResourceBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("clusterresourcebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a clusterResourceBinding and creates it.  Returns the server's representation of the clusterResourceBinding, and an error, if there is any.
func (c *clusterResourceBindings) Create(clusterResourceBinding *v1alpha2.ClusterResourceBinding) (result *v1alpha2.ClusterResourceBinding, err error) {
	result = &v1alpha2.ClusterResourceBinding{}
	err = c.client.Post().
		Resource("clusterresourcebindings").
		Body(clusterResourceBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a clusterResourceBinding and updates it. Returns the server's representation of the clusterResourceBinding, and an error, if there is any.
func (c *clusterResourceBindings) Update(clusterResourceBinding *v1alpha2.ClusterResourceBinding) (result *v1alpha2.ClusterResourceBinding, err error) {
	result = &v1alpha2.ClusterResourceBinding{}
	err = c.client.Put().
		Resource("clusterresourcebindings").
		Name(clusterResourceBinding.Name).
		Body(clusterResourceBinding).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *clusterResourceBindings) UpdateStatus(clusterResourceBinding *v1alpha2.ClusterResourceBinding) (result *v1alpha2.ClusterResourceBinding, err error) {
	result = &v1alpha2.ClusterResourceBinding{}
	err = c.client.Put().
		Resource("clusterresourcebindings").
		Name(clusterResourceBinding.Name).
		SubResource("status").
		Body(clusterResourceBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the clusterResourceBinding and deletes it. Returns an error if one occurs.
func (c *clusterResourceBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusterresourcebindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterResourceBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("clusterresourcebindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched clusterResourceBinding.
func (c *clusterResourceBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.ClusterResourceBinding, err error) {
	result = &v1alpha2.ClusterResourceBinding{}
	err = c.client.Patch(pt).
		Resource("clusterresourcebindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
