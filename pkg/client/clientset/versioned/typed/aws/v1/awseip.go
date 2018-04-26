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

// AwsEipsGetter has a method to return a AwsEipInterface.
// A group's client should implement this interface.
type AwsEipsGetter interface {
	AwsEips(namespace string) AwsEipInterface
}

// AwsEipInterface has methods to work with AwsEip resources.
type AwsEipInterface interface {
	Create(*v1.AwsEip) (*v1.AwsEip, error)
	Update(*v1.AwsEip) (*v1.AwsEip, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsEip, error)
	List(opts meta_v1.ListOptions) (*v1.AwsEipList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsEip, err error)
	AwsEipExpansion
}

// awsEips implements AwsEipInterface
type awsEips struct {
	client rest.Interface
	ns     string
}

// newAwsEips returns a AwsEips
func newAwsEips(c *TrussleV1Client, namespace string) *awsEips {
	return &awsEips{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsEip, and returns the corresponding awsEip object, and an error if there is any.
func (c *awsEips) Get(name string, options meta_v1.GetOptions) (result *v1.AwsEip, err error) {
	result = &v1.AwsEip{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awseips").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsEips that match those selectors.
func (c *awsEips) List(opts meta_v1.ListOptions) (result *v1.AwsEipList, err error) {
	result = &v1.AwsEipList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awseips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsEips.
func (c *awsEips) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awseips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsEip and creates it.  Returns the server's representation of the awsEip, and an error, if there is any.
func (c *awsEips) Create(awsEip *v1.AwsEip) (result *v1.AwsEip, err error) {
	result = &v1.AwsEip{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awseips").
		Body(awsEip).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsEip and updates it. Returns the server's representation of the awsEip, and an error, if there is any.
func (c *awsEips) Update(awsEip *v1.AwsEip) (result *v1.AwsEip, err error) {
	result = &v1.AwsEip{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awseips").
		Name(awsEip.Name).
		Body(awsEip).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsEip and deletes it. Returns an error if one occurs.
func (c *awsEips) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awseips").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsEips) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awseips").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsEip.
func (c *awsEips) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsEip, err error) {
	result = &v1.AwsEip{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awseips").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
