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

// AwsIotCertificatesGetter has a method to return a AwsIotCertificateInterface.
// A group's client should implement this interface.
type AwsIotCertificatesGetter interface {
	AwsIotCertificates(namespace string) AwsIotCertificateInterface
}

// AwsIotCertificateInterface has methods to work with AwsIotCertificate resources.
type AwsIotCertificateInterface interface {
	Create(*v1.AwsIotCertificate) (*v1.AwsIotCertificate, error)
	Update(*v1.AwsIotCertificate) (*v1.AwsIotCertificate, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsIotCertificate, error)
	List(opts meta_v1.ListOptions) (*v1.AwsIotCertificateList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsIotCertificate, err error)
	AwsIotCertificateExpansion
}

// awsIotCertificates implements AwsIotCertificateInterface
type awsIotCertificates struct {
	client rest.Interface
	ns     string
}

// newAwsIotCertificates returns a AwsIotCertificates
func newAwsIotCertificates(c *TrussleV1Client, namespace string) *awsIotCertificates {
	return &awsIotCertificates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsIotCertificate, and returns the corresponding awsIotCertificate object, and an error if there is any.
func (c *awsIotCertificates) Get(name string, options meta_v1.GetOptions) (result *v1.AwsIotCertificate, err error) {
	result = &v1.AwsIotCertificate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsiotcertificates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsIotCertificates that match those selectors.
func (c *awsIotCertificates) List(opts meta_v1.ListOptions) (result *v1.AwsIotCertificateList, err error) {
	result = &v1.AwsIotCertificateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsiotcertificates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsIotCertificates.
func (c *awsIotCertificates) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsiotcertificates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsIotCertificate and creates it.  Returns the server's representation of the awsIotCertificate, and an error, if there is any.
func (c *awsIotCertificates) Create(awsIotCertificate *v1.AwsIotCertificate) (result *v1.AwsIotCertificate, err error) {
	result = &v1.AwsIotCertificate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsiotcertificates").
		Body(awsIotCertificate).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsIotCertificate and updates it. Returns the server's representation of the awsIotCertificate, and an error, if there is any.
func (c *awsIotCertificates) Update(awsIotCertificate *v1.AwsIotCertificate) (result *v1.AwsIotCertificate, err error) {
	result = &v1.AwsIotCertificate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsiotcertificates").
		Name(awsIotCertificate.Name).
		Body(awsIotCertificate).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsIotCertificate and deletes it. Returns an error if one occurs.
func (c *awsIotCertificates) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsiotcertificates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsIotCertificates) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsiotcertificates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsIotCertificate.
func (c *awsIotCertificates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsIotCertificate, err error) {
	result = &v1.AwsIotCertificate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsiotcertificates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
