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

// AwsKmsGrantsGetter has a method to return a AwsKmsGrantInterface.
// A group's client should implement this interface.
type AwsKmsGrantsGetter interface {
	AwsKmsGrants(namespace string) AwsKmsGrantInterface
}

// AwsKmsGrantInterface has methods to work with AwsKmsGrant resources.
type AwsKmsGrantInterface interface {
	Create(*v1.AwsKmsGrant) (*v1.AwsKmsGrant, error)
	Update(*v1.AwsKmsGrant) (*v1.AwsKmsGrant, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsKmsGrant, error)
	List(opts meta_v1.ListOptions) (*v1.AwsKmsGrantList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsKmsGrant, err error)
	AwsKmsGrantExpansion
}

// awsKmsGrants implements AwsKmsGrantInterface
type awsKmsGrants struct {
	client rest.Interface
	ns     string
}

// newAwsKmsGrants returns a AwsKmsGrants
func newAwsKmsGrants(c *TrussleV1Client, namespace string) *awsKmsGrants {
	return &awsKmsGrants{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsKmsGrant, and returns the corresponding awsKmsGrant object, and an error if there is any.
func (c *awsKmsGrants) Get(name string, options meta_v1.GetOptions) (result *v1.AwsKmsGrant, err error) {
	result = &v1.AwsKmsGrant{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awskmsgrants").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsKmsGrants that match those selectors.
func (c *awsKmsGrants) List(opts meta_v1.ListOptions) (result *v1.AwsKmsGrantList, err error) {
	result = &v1.AwsKmsGrantList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awskmsgrants").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsKmsGrants.
func (c *awsKmsGrants) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awskmsgrants").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsKmsGrant and creates it.  Returns the server's representation of the awsKmsGrant, and an error, if there is any.
func (c *awsKmsGrants) Create(awsKmsGrant *v1.AwsKmsGrant) (result *v1.AwsKmsGrant, err error) {
	result = &v1.AwsKmsGrant{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awskmsgrants").
		Body(awsKmsGrant).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsKmsGrant and updates it. Returns the server's representation of the awsKmsGrant, and an error, if there is any.
func (c *awsKmsGrants) Update(awsKmsGrant *v1.AwsKmsGrant) (result *v1.AwsKmsGrant, err error) {
	result = &v1.AwsKmsGrant{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awskmsgrants").
		Name(awsKmsGrant.Name).
		Body(awsKmsGrant).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsKmsGrant and deletes it. Returns an error if one occurs.
func (c *awsKmsGrants) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awskmsgrants").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsKmsGrants) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awskmsgrants").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsKmsGrant.
func (c *awsKmsGrants) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsKmsGrant, err error) {
	result = &v1.AwsKmsGrant{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awskmsgrants").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
