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

// AwsOpsworksMemcachedLayersGetter has a method to return a AwsOpsworksMemcachedLayerInterface.
// A group's client should implement this interface.
type AwsOpsworksMemcachedLayersGetter interface {
	AwsOpsworksMemcachedLayers(namespace string) AwsOpsworksMemcachedLayerInterface
}

// AwsOpsworksMemcachedLayerInterface has methods to work with AwsOpsworksMemcachedLayer resources.
type AwsOpsworksMemcachedLayerInterface interface {
	Create(*v1.AwsOpsworksMemcachedLayer) (*v1.AwsOpsworksMemcachedLayer, error)
	Update(*v1.AwsOpsworksMemcachedLayer) (*v1.AwsOpsworksMemcachedLayer, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsOpsworksMemcachedLayer, error)
	List(opts meta_v1.ListOptions) (*v1.AwsOpsworksMemcachedLayerList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsOpsworksMemcachedLayer, err error)
	AwsOpsworksMemcachedLayerExpansion
}

// awsOpsworksMemcachedLayers implements AwsOpsworksMemcachedLayerInterface
type awsOpsworksMemcachedLayers struct {
	client rest.Interface
	ns     string
}

// newAwsOpsworksMemcachedLayers returns a AwsOpsworksMemcachedLayers
func newAwsOpsworksMemcachedLayers(c *TrussleV1Client, namespace string) *awsOpsworksMemcachedLayers {
	return &awsOpsworksMemcachedLayers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsOpsworksMemcachedLayer, and returns the corresponding awsOpsworksMemcachedLayer object, and an error if there is any.
func (c *awsOpsworksMemcachedLayers) Get(name string, options meta_v1.GetOptions) (result *v1.AwsOpsworksMemcachedLayer, err error) {
	result = &v1.AwsOpsworksMemcachedLayer{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsopsworksmemcachedlayers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsOpsworksMemcachedLayers that match those selectors.
func (c *awsOpsworksMemcachedLayers) List(opts meta_v1.ListOptions) (result *v1.AwsOpsworksMemcachedLayerList, err error) {
	result = &v1.AwsOpsworksMemcachedLayerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsopsworksmemcachedlayers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsOpsworksMemcachedLayers.
func (c *awsOpsworksMemcachedLayers) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsopsworksmemcachedlayers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsOpsworksMemcachedLayer and creates it.  Returns the server's representation of the awsOpsworksMemcachedLayer, and an error, if there is any.
func (c *awsOpsworksMemcachedLayers) Create(awsOpsworksMemcachedLayer *v1.AwsOpsworksMemcachedLayer) (result *v1.AwsOpsworksMemcachedLayer, err error) {
	result = &v1.AwsOpsworksMemcachedLayer{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsopsworksmemcachedlayers").
		Body(awsOpsworksMemcachedLayer).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsOpsworksMemcachedLayer and updates it. Returns the server's representation of the awsOpsworksMemcachedLayer, and an error, if there is any.
func (c *awsOpsworksMemcachedLayers) Update(awsOpsworksMemcachedLayer *v1.AwsOpsworksMemcachedLayer) (result *v1.AwsOpsworksMemcachedLayer, err error) {
	result = &v1.AwsOpsworksMemcachedLayer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsopsworksmemcachedlayers").
		Name(awsOpsworksMemcachedLayer.Name).
		Body(awsOpsworksMemcachedLayer).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsOpsworksMemcachedLayer and deletes it. Returns an error if one occurs.
func (c *awsOpsworksMemcachedLayers) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsopsworksmemcachedlayers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsOpsworksMemcachedLayers) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsopsworksmemcachedlayers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsOpsworksMemcachedLayer.
func (c *awsOpsworksMemcachedLayers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsOpsworksMemcachedLayer, err error) {
	result = &v1.AwsOpsworksMemcachedLayer{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsopsworksmemcachedlayers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
