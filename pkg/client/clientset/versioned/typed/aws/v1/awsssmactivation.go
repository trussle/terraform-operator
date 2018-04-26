/*
Copyright The Kubernetes Authors.

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

package v1

import (
	v1 "github.com/trussle/terraform-operator/pkg/apis/aws/v1"
	scheme "github.com/trussle/terraform-operator/pkg/client/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AwsSsmActivationsGetter has a method to return a AwsSsmActivationInterface.
// A group's client should implement this interface.
type AwsSsmActivationsGetter interface {
	AwsSsmActivations(namespace string) AwsSsmActivationInterface
}

// AwsSsmActivationInterface has methods to work with AwsSsmActivation resources.
type AwsSsmActivationInterface interface {
	Create(*v1.AwsSsmActivation) (*v1.AwsSsmActivation, error)
	Update(*v1.AwsSsmActivation) (*v1.AwsSsmActivation, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsSsmActivation, error)
	List(opts meta_v1.ListOptions) (*v1.AwsSsmActivationList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsSsmActivation, err error)
	AwsSsmActivationExpansion
}

// awsSsmActivations implements AwsSsmActivationInterface
type awsSsmActivations struct {
	client rest.Interface
	ns     string
}

// newAwsSsmActivations returns a AwsSsmActivations
func newAwsSsmActivations(c *TrussleV1Client, namespace string) *awsSsmActivations {
	return &awsSsmActivations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsSsmActivation, and returns the corresponding awsSsmActivation object, and an error if there is any.
func (c *awsSsmActivations) Get(name string, options meta_v1.GetOptions) (result *v1.AwsSsmActivation, err error) {
	result = &v1.AwsSsmActivation{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsssmactivations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsSsmActivations that match those selectors.
func (c *awsSsmActivations) List(opts meta_v1.ListOptions) (result *v1.AwsSsmActivationList, err error) {
	result = &v1.AwsSsmActivationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsssmactivations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsSsmActivations.
func (c *awsSsmActivations) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsssmactivations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsSsmActivation and creates it.  Returns the server's representation of the awsSsmActivation, and an error, if there is any.
func (c *awsSsmActivations) Create(awsSsmActivation *v1.AwsSsmActivation) (result *v1.AwsSsmActivation, err error) {
	result = &v1.AwsSsmActivation{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsssmactivations").
		Body(awsSsmActivation).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsSsmActivation and updates it. Returns the server's representation of the awsSsmActivation, and an error, if there is any.
func (c *awsSsmActivations) Update(awsSsmActivation *v1.AwsSsmActivation) (result *v1.AwsSsmActivation, err error) {
	result = &v1.AwsSsmActivation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsssmactivations").
		Name(awsSsmActivation.Name).
		Body(awsSsmActivation).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsSsmActivation and deletes it. Returns an error if one occurs.
func (c *awsSsmActivations) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsssmactivations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsSsmActivations) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsssmactivations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsSsmActivation.
func (c *awsSsmActivations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsSsmActivation, err error) {
	result = &v1.AwsSsmActivation{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsssmactivations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}