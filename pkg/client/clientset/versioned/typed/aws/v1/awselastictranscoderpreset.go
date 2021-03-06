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

// AwsElastictranscoderPresetsGetter has a method to return a AwsElastictranscoderPresetInterface.
// A group's client should implement this interface.
type AwsElastictranscoderPresetsGetter interface {
	AwsElastictranscoderPresets(namespace string) AwsElastictranscoderPresetInterface
}

// AwsElastictranscoderPresetInterface has methods to work with AwsElastictranscoderPreset resources.
type AwsElastictranscoderPresetInterface interface {
	Create(*v1.AwsElastictranscoderPreset) (*v1.AwsElastictranscoderPreset, error)
	Update(*v1.AwsElastictranscoderPreset) (*v1.AwsElastictranscoderPreset, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsElastictranscoderPreset, error)
	List(opts meta_v1.ListOptions) (*v1.AwsElastictranscoderPresetList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsElastictranscoderPreset, err error)
	AwsElastictranscoderPresetExpansion
}

// awsElastictranscoderPresets implements AwsElastictranscoderPresetInterface
type awsElastictranscoderPresets struct {
	client rest.Interface
	ns     string
}

// newAwsElastictranscoderPresets returns a AwsElastictranscoderPresets
func newAwsElastictranscoderPresets(c *TrussleV1Client, namespace string) *awsElastictranscoderPresets {
	return &awsElastictranscoderPresets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsElastictranscoderPreset, and returns the corresponding awsElastictranscoderPreset object, and an error if there is any.
func (c *awsElastictranscoderPresets) Get(name string, options meta_v1.GetOptions) (result *v1.AwsElastictranscoderPreset, err error) {
	result = &v1.AwsElastictranscoderPreset{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awselastictranscoderpresets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsElastictranscoderPresets that match those selectors.
func (c *awsElastictranscoderPresets) List(opts meta_v1.ListOptions) (result *v1.AwsElastictranscoderPresetList, err error) {
	result = &v1.AwsElastictranscoderPresetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awselastictranscoderpresets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsElastictranscoderPresets.
func (c *awsElastictranscoderPresets) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awselastictranscoderpresets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsElastictranscoderPreset and creates it.  Returns the server's representation of the awsElastictranscoderPreset, and an error, if there is any.
func (c *awsElastictranscoderPresets) Create(awsElastictranscoderPreset *v1.AwsElastictranscoderPreset) (result *v1.AwsElastictranscoderPreset, err error) {
	result = &v1.AwsElastictranscoderPreset{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awselastictranscoderpresets").
		Body(awsElastictranscoderPreset).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsElastictranscoderPreset and updates it. Returns the server's representation of the awsElastictranscoderPreset, and an error, if there is any.
func (c *awsElastictranscoderPresets) Update(awsElastictranscoderPreset *v1.AwsElastictranscoderPreset) (result *v1.AwsElastictranscoderPreset, err error) {
	result = &v1.AwsElastictranscoderPreset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awselastictranscoderpresets").
		Name(awsElastictranscoderPreset.Name).
		Body(awsElastictranscoderPreset).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsElastictranscoderPreset and deletes it. Returns an error if one occurs.
func (c *awsElastictranscoderPresets) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awselastictranscoderpresets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsElastictranscoderPresets) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awselastictranscoderpresets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsElastictranscoderPreset.
func (c *awsElastictranscoderPresets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsElastictranscoderPreset, err error) {
	result = &v1.AwsElastictranscoderPreset{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awselastictranscoderpresets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
