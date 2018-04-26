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

// AwsSesReceiptFiltersGetter has a method to return a AwsSesReceiptFilterInterface.
// A group's client should implement this interface.
type AwsSesReceiptFiltersGetter interface {
	AwsSesReceiptFilters(namespace string) AwsSesReceiptFilterInterface
}

// AwsSesReceiptFilterInterface has methods to work with AwsSesReceiptFilter resources.
type AwsSesReceiptFilterInterface interface {
	Create(*v1.AwsSesReceiptFilter) (*v1.AwsSesReceiptFilter, error)
	Update(*v1.AwsSesReceiptFilter) (*v1.AwsSesReceiptFilter, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsSesReceiptFilter, error)
	List(opts meta_v1.ListOptions) (*v1.AwsSesReceiptFilterList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsSesReceiptFilter, err error)
	AwsSesReceiptFilterExpansion
}

// awsSesReceiptFilters implements AwsSesReceiptFilterInterface
type awsSesReceiptFilters struct {
	client rest.Interface
	ns     string
}

// newAwsSesReceiptFilters returns a AwsSesReceiptFilters
func newAwsSesReceiptFilters(c *TrussleV1Client, namespace string) *awsSesReceiptFilters {
	return &awsSesReceiptFilters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsSesReceiptFilter, and returns the corresponding awsSesReceiptFilter object, and an error if there is any.
func (c *awsSesReceiptFilters) Get(name string, options meta_v1.GetOptions) (result *v1.AwsSesReceiptFilter, err error) {
	result = &v1.AwsSesReceiptFilter{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awssesreceiptfilters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsSesReceiptFilters that match those selectors.
func (c *awsSesReceiptFilters) List(opts meta_v1.ListOptions) (result *v1.AwsSesReceiptFilterList, err error) {
	result = &v1.AwsSesReceiptFilterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awssesreceiptfilters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsSesReceiptFilters.
func (c *awsSesReceiptFilters) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awssesreceiptfilters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsSesReceiptFilter and creates it.  Returns the server's representation of the awsSesReceiptFilter, and an error, if there is any.
func (c *awsSesReceiptFilters) Create(awsSesReceiptFilter *v1.AwsSesReceiptFilter) (result *v1.AwsSesReceiptFilter, err error) {
	result = &v1.AwsSesReceiptFilter{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awssesreceiptfilters").
		Body(awsSesReceiptFilter).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsSesReceiptFilter and updates it. Returns the server's representation of the awsSesReceiptFilter, and an error, if there is any.
func (c *awsSesReceiptFilters) Update(awsSesReceiptFilter *v1.AwsSesReceiptFilter) (result *v1.AwsSesReceiptFilter, err error) {
	result = &v1.AwsSesReceiptFilter{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awssesreceiptfilters").
		Name(awsSesReceiptFilter.Name).
		Body(awsSesReceiptFilter).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsSesReceiptFilter and deletes it. Returns an error if one occurs.
func (c *awsSesReceiptFilters) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awssesreceiptfilters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsSesReceiptFilters) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awssesreceiptfilters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsSesReceiptFilter.
func (c *awsSesReceiptFilters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsSesReceiptFilter, err error) {
	result = &v1.AwsSesReceiptFilter{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awssesreceiptfilters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
