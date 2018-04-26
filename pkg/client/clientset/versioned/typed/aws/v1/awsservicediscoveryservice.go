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

// AwsServiceDiscoveryServicesGetter has a method to return a AwsServiceDiscoveryServiceInterface.
// A group's client should implement this interface.
type AwsServiceDiscoveryServicesGetter interface {
	AwsServiceDiscoveryServices(namespace string) AwsServiceDiscoveryServiceInterface
}

// AwsServiceDiscoveryServiceInterface has methods to work with AwsServiceDiscoveryService resources.
type AwsServiceDiscoveryServiceInterface interface {
	Create(*v1.AwsServiceDiscoveryService) (*v1.AwsServiceDiscoveryService, error)
	Update(*v1.AwsServiceDiscoveryService) (*v1.AwsServiceDiscoveryService, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsServiceDiscoveryService, error)
	List(opts meta_v1.ListOptions) (*v1.AwsServiceDiscoveryServiceList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsServiceDiscoveryService, err error)
	AwsServiceDiscoveryServiceExpansion
}

// awsServiceDiscoveryServices implements AwsServiceDiscoveryServiceInterface
type awsServiceDiscoveryServices struct {
	client rest.Interface
	ns     string
}

// newAwsServiceDiscoveryServices returns a AwsServiceDiscoveryServices
func newAwsServiceDiscoveryServices(c *TrussleV1Client, namespace string) *awsServiceDiscoveryServices {
	return &awsServiceDiscoveryServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsServiceDiscoveryService, and returns the corresponding awsServiceDiscoveryService object, and an error if there is any.
func (c *awsServiceDiscoveryServices) Get(name string, options meta_v1.GetOptions) (result *v1.AwsServiceDiscoveryService, err error) {
	result = &v1.AwsServiceDiscoveryService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsservicediscoveryservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsServiceDiscoveryServices that match those selectors.
func (c *awsServiceDiscoveryServices) List(opts meta_v1.ListOptions) (result *v1.AwsServiceDiscoveryServiceList, err error) {
	result = &v1.AwsServiceDiscoveryServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsservicediscoveryservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsServiceDiscoveryServices.
func (c *awsServiceDiscoveryServices) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsservicediscoveryservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsServiceDiscoveryService and creates it.  Returns the server's representation of the awsServiceDiscoveryService, and an error, if there is any.
func (c *awsServiceDiscoveryServices) Create(awsServiceDiscoveryService *v1.AwsServiceDiscoveryService) (result *v1.AwsServiceDiscoveryService, err error) {
	result = &v1.AwsServiceDiscoveryService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsservicediscoveryservices").
		Body(awsServiceDiscoveryService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsServiceDiscoveryService and updates it. Returns the server's representation of the awsServiceDiscoveryService, and an error, if there is any.
func (c *awsServiceDiscoveryServices) Update(awsServiceDiscoveryService *v1.AwsServiceDiscoveryService) (result *v1.AwsServiceDiscoveryService, err error) {
	result = &v1.AwsServiceDiscoveryService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsservicediscoveryservices").
		Name(awsServiceDiscoveryService.Name).
		Body(awsServiceDiscoveryService).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsServiceDiscoveryService and deletes it. Returns an error if one occurs.
func (c *awsServiceDiscoveryServices) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsservicediscoveryservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsServiceDiscoveryServices) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsservicediscoveryservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsServiceDiscoveryService.
func (c *awsServiceDiscoveryServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsServiceDiscoveryService, err error) {
	result = &v1.AwsServiceDiscoveryService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsservicediscoveryservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}