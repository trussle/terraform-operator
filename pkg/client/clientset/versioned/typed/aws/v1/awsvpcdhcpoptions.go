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

// AwsVpcDhcpOptionsesGetter has a method to return a AwsVpcDhcpOptionsInterface.
// A group's client should implement this interface.
type AwsVpcDhcpOptionsesGetter interface {
	AwsVpcDhcpOptionses(namespace string) AwsVpcDhcpOptionsInterface
}

// AwsVpcDhcpOptionsInterface has methods to work with AwsVpcDhcpOptions resources.
type AwsVpcDhcpOptionsInterface interface {
	Create(*v1.AwsVpcDhcpOptions) (*v1.AwsVpcDhcpOptions, error)
	Update(*v1.AwsVpcDhcpOptions) (*v1.AwsVpcDhcpOptions, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsVpcDhcpOptions, error)
	List(opts meta_v1.ListOptions) (*v1.AwsVpcDhcpOptionsList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsVpcDhcpOptions, err error)
	AwsVpcDhcpOptionsExpansion
}

// awsVpcDhcpOptionses implements AwsVpcDhcpOptionsInterface
type awsVpcDhcpOptionses struct {
	client rest.Interface
	ns     string
}

// newAwsVpcDhcpOptionses returns a AwsVpcDhcpOptionses
func newAwsVpcDhcpOptionses(c *TrussleV1Client, namespace string) *awsVpcDhcpOptionses {
	return &awsVpcDhcpOptionses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsVpcDhcpOptions, and returns the corresponding awsVpcDhcpOptions object, and an error if there is any.
func (c *awsVpcDhcpOptionses) Get(name string, options meta_v1.GetOptions) (result *v1.AwsVpcDhcpOptions, err error) {
	result = &v1.AwsVpcDhcpOptions{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsvpcdhcpoptionses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsVpcDhcpOptionses that match those selectors.
func (c *awsVpcDhcpOptionses) List(opts meta_v1.ListOptions) (result *v1.AwsVpcDhcpOptionsList, err error) {
	result = &v1.AwsVpcDhcpOptionsList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsvpcdhcpoptionses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsVpcDhcpOptionses.
func (c *awsVpcDhcpOptionses) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsvpcdhcpoptionses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsVpcDhcpOptions and creates it.  Returns the server's representation of the awsVpcDhcpOptions, and an error, if there is any.
func (c *awsVpcDhcpOptionses) Create(awsVpcDhcpOptions *v1.AwsVpcDhcpOptions) (result *v1.AwsVpcDhcpOptions, err error) {
	result = &v1.AwsVpcDhcpOptions{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsvpcdhcpoptionses").
		Body(awsVpcDhcpOptions).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsVpcDhcpOptions and updates it. Returns the server's representation of the awsVpcDhcpOptions, and an error, if there is any.
func (c *awsVpcDhcpOptionses) Update(awsVpcDhcpOptions *v1.AwsVpcDhcpOptions) (result *v1.AwsVpcDhcpOptions, err error) {
	result = &v1.AwsVpcDhcpOptions{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsvpcdhcpoptionses").
		Name(awsVpcDhcpOptions.Name).
		Body(awsVpcDhcpOptions).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsVpcDhcpOptions and deletes it. Returns an error if one occurs.
func (c *awsVpcDhcpOptionses) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsvpcdhcpoptionses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsVpcDhcpOptionses) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsvpcdhcpoptionses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsVpcDhcpOptions.
func (c *awsVpcDhcpOptionses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsVpcDhcpOptions, err error) {
	result = &v1.AwsVpcDhcpOptions{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsvpcdhcpoptionses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
