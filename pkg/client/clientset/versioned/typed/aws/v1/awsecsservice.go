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

// AwsEcsServicesGetter has a method to return a AwsEcsServiceInterface.
// A group's client should implement this interface.
type AwsEcsServicesGetter interface {
	AwsEcsServices(namespace string) AwsEcsServiceInterface
}

// AwsEcsServiceInterface has methods to work with AwsEcsService resources.
type AwsEcsServiceInterface interface {
	Create(*v1.AwsEcsService) (*v1.AwsEcsService, error)
	Update(*v1.AwsEcsService) (*v1.AwsEcsService, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsEcsService, error)
	List(opts meta_v1.ListOptions) (*v1.AwsEcsServiceList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsEcsService, err error)
	AwsEcsServiceExpansion
}

// awsEcsServices implements AwsEcsServiceInterface
type awsEcsServices struct {
	client rest.Interface
	ns     string
}

// newAwsEcsServices returns a AwsEcsServices
func newAwsEcsServices(c *TrussleV1Client, namespace string) *awsEcsServices {
	return &awsEcsServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsEcsService, and returns the corresponding awsEcsService object, and an error if there is any.
func (c *awsEcsServices) Get(name string, options meta_v1.GetOptions) (result *v1.AwsEcsService, err error) {
	result = &v1.AwsEcsService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsecsservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsEcsServices that match those selectors.
func (c *awsEcsServices) List(opts meta_v1.ListOptions) (result *v1.AwsEcsServiceList, err error) {
	result = &v1.AwsEcsServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsecsservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsEcsServices.
func (c *awsEcsServices) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsecsservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsEcsService and creates it.  Returns the server's representation of the awsEcsService, and an error, if there is any.
func (c *awsEcsServices) Create(awsEcsService *v1.AwsEcsService) (result *v1.AwsEcsService, err error) {
	result = &v1.AwsEcsService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsecsservices").
		Body(awsEcsService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsEcsService and updates it. Returns the server's representation of the awsEcsService, and an error, if there is any.
func (c *awsEcsServices) Update(awsEcsService *v1.AwsEcsService) (result *v1.AwsEcsService, err error) {
	result = &v1.AwsEcsService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsecsservices").
		Name(awsEcsService.Name).
		Body(awsEcsService).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsEcsService and deletes it. Returns an error if one occurs.
func (c *awsEcsServices) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsecsservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsEcsServices) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsecsservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsEcsService.
func (c *awsEcsServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsEcsService, err error) {
	result = &v1.AwsEcsService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsecsservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}