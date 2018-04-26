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

// AwsKinesisStreamsGetter has a method to return a AwsKinesisStreamInterface.
// A group's client should implement this interface.
type AwsKinesisStreamsGetter interface {
	AwsKinesisStreams(namespace string) AwsKinesisStreamInterface
}

// AwsKinesisStreamInterface has methods to work with AwsKinesisStream resources.
type AwsKinesisStreamInterface interface {
	Create(*v1.AwsKinesisStream) (*v1.AwsKinesisStream, error)
	Update(*v1.AwsKinesisStream) (*v1.AwsKinesisStream, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsKinesisStream, error)
	List(opts meta_v1.ListOptions) (*v1.AwsKinesisStreamList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsKinesisStream, err error)
	AwsKinesisStreamExpansion
}

// awsKinesisStreams implements AwsKinesisStreamInterface
type awsKinesisStreams struct {
	client rest.Interface
	ns     string
}

// newAwsKinesisStreams returns a AwsKinesisStreams
func newAwsKinesisStreams(c *TrussleV1Client, namespace string) *awsKinesisStreams {
	return &awsKinesisStreams{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsKinesisStream, and returns the corresponding awsKinesisStream object, and an error if there is any.
func (c *awsKinesisStreams) Get(name string, options meta_v1.GetOptions) (result *v1.AwsKinesisStream, err error) {
	result = &v1.AwsKinesisStream{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awskinesisstreams").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsKinesisStreams that match those selectors.
func (c *awsKinesisStreams) List(opts meta_v1.ListOptions) (result *v1.AwsKinesisStreamList, err error) {
	result = &v1.AwsKinesisStreamList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awskinesisstreams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsKinesisStreams.
func (c *awsKinesisStreams) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awskinesisstreams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsKinesisStream and creates it.  Returns the server's representation of the awsKinesisStream, and an error, if there is any.
func (c *awsKinesisStreams) Create(awsKinesisStream *v1.AwsKinesisStream) (result *v1.AwsKinesisStream, err error) {
	result = &v1.AwsKinesisStream{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awskinesisstreams").
		Body(awsKinesisStream).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsKinesisStream and updates it. Returns the server's representation of the awsKinesisStream, and an error, if there is any.
func (c *awsKinesisStreams) Update(awsKinesisStream *v1.AwsKinesisStream) (result *v1.AwsKinesisStream, err error) {
	result = &v1.AwsKinesisStream{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awskinesisstreams").
		Name(awsKinesisStream.Name).
		Body(awsKinesisStream).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsKinesisStream and deletes it. Returns an error if one occurs.
func (c *awsKinesisStreams) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awskinesisstreams").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsKinesisStreams) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awskinesisstreams").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsKinesisStream.
func (c *awsKinesisStreams) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsKinesisStream, err error) {
	result = &v1.AwsKinesisStream{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awskinesisstreams").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}