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

// AwsOpsworksMysqlLayersGetter has a method to return a AwsOpsworksMysqlLayerInterface.
// A group's client should implement this interface.
type AwsOpsworksMysqlLayersGetter interface {
	AwsOpsworksMysqlLayers(namespace string) AwsOpsworksMysqlLayerInterface
}

// AwsOpsworksMysqlLayerInterface has methods to work with AwsOpsworksMysqlLayer resources.
type AwsOpsworksMysqlLayerInterface interface {
	Create(*v1.AwsOpsworksMysqlLayer) (*v1.AwsOpsworksMysqlLayer, error)
	Update(*v1.AwsOpsworksMysqlLayer) (*v1.AwsOpsworksMysqlLayer, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsOpsworksMysqlLayer, error)
	List(opts meta_v1.ListOptions) (*v1.AwsOpsworksMysqlLayerList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsOpsworksMysqlLayer, err error)
	AwsOpsworksMysqlLayerExpansion
}

// awsOpsworksMysqlLayers implements AwsOpsworksMysqlLayerInterface
type awsOpsworksMysqlLayers struct {
	client rest.Interface
	ns     string
}

// newAwsOpsworksMysqlLayers returns a AwsOpsworksMysqlLayers
func newAwsOpsworksMysqlLayers(c *TrussleV1Client, namespace string) *awsOpsworksMysqlLayers {
	return &awsOpsworksMysqlLayers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsOpsworksMysqlLayer, and returns the corresponding awsOpsworksMysqlLayer object, and an error if there is any.
func (c *awsOpsworksMysqlLayers) Get(name string, options meta_v1.GetOptions) (result *v1.AwsOpsworksMysqlLayer, err error) {
	result = &v1.AwsOpsworksMysqlLayer{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsopsworksmysqllayers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsOpsworksMysqlLayers that match those selectors.
func (c *awsOpsworksMysqlLayers) List(opts meta_v1.ListOptions) (result *v1.AwsOpsworksMysqlLayerList, err error) {
	result = &v1.AwsOpsworksMysqlLayerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsopsworksmysqllayers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsOpsworksMysqlLayers.
func (c *awsOpsworksMysqlLayers) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsopsworksmysqllayers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsOpsworksMysqlLayer and creates it.  Returns the server's representation of the awsOpsworksMysqlLayer, and an error, if there is any.
func (c *awsOpsworksMysqlLayers) Create(awsOpsworksMysqlLayer *v1.AwsOpsworksMysqlLayer) (result *v1.AwsOpsworksMysqlLayer, err error) {
	result = &v1.AwsOpsworksMysqlLayer{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsopsworksmysqllayers").
		Body(awsOpsworksMysqlLayer).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsOpsworksMysqlLayer and updates it. Returns the server's representation of the awsOpsworksMysqlLayer, and an error, if there is any.
func (c *awsOpsworksMysqlLayers) Update(awsOpsworksMysqlLayer *v1.AwsOpsworksMysqlLayer) (result *v1.AwsOpsworksMysqlLayer, err error) {
	result = &v1.AwsOpsworksMysqlLayer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsopsworksmysqllayers").
		Name(awsOpsworksMysqlLayer.Name).
		Body(awsOpsworksMysqlLayer).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsOpsworksMysqlLayer and deletes it. Returns an error if one occurs.
func (c *awsOpsworksMysqlLayers) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsopsworksmysqllayers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsOpsworksMysqlLayers) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsopsworksmysqllayers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsOpsworksMysqlLayer.
func (c *awsOpsworksMysqlLayers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsOpsworksMysqlLayer, err error) {
	result = &v1.AwsOpsworksMysqlLayer{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsopsworksmysqllayers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
