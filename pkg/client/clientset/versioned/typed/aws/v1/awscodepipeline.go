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

// AwsCodepipelinesGetter has a method to return a AwsCodepipelineInterface.
// A group's client should implement this interface.
type AwsCodepipelinesGetter interface {
	AwsCodepipelines(namespace string) AwsCodepipelineInterface
}

// AwsCodepipelineInterface has methods to work with AwsCodepipeline resources.
type AwsCodepipelineInterface interface {
	Create(*v1.AwsCodepipeline) (*v1.AwsCodepipeline, error)
	Update(*v1.AwsCodepipeline) (*v1.AwsCodepipeline, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsCodepipeline, error)
	List(opts meta_v1.ListOptions) (*v1.AwsCodepipelineList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsCodepipeline, err error)
	AwsCodepipelineExpansion
}

// awsCodepipelines implements AwsCodepipelineInterface
type awsCodepipelines struct {
	client rest.Interface
	ns     string
}

// newAwsCodepipelines returns a AwsCodepipelines
func newAwsCodepipelines(c *TrussleV1Client, namespace string) *awsCodepipelines {
	return &awsCodepipelines{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsCodepipeline, and returns the corresponding awsCodepipeline object, and an error if there is any.
func (c *awsCodepipelines) Get(name string, options meta_v1.GetOptions) (result *v1.AwsCodepipeline, err error) {
	result = &v1.AwsCodepipeline{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awscodepipelines").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsCodepipelines that match those selectors.
func (c *awsCodepipelines) List(opts meta_v1.ListOptions) (result *v1.AwsCodepipelineList, err error) {
	result = &v1.AwsCodepipelineList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awscodepipelines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsCodepipelines.
func (c *awsCodepipelines) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awscodepipelines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsCodepipeline and creates it.  Returns the server's representation of the awsCodepipeline, and an error, if there is any.
func (c *awsCodepipelines) Create(awsCodepipeline *v1.AwsCodepipeline) (result *v1.AwsCodepipeline, err error) {
	result = &v1.AwsCodepipeline{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awscodepipelines").
		Body(awsCodepipeline).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsCodepipeline and updates it. Returns the server's representation of the awsCodepipeline, and an error, if there is any.
func (c *awsCodepipelines) Update(awsCodepipeline *v1.AwsCodepipeline) (result *v1.AwsCodepipeline, err error) {
	result = &v1.AwsCodepipeline{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awscodepipelines").
		Name(awsCodepipeline.Name).
		Body(awsCodepipeline).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsCodepipeline and deletes it. Returns an error if one occurs.
func (c *awsCodepipelines) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awscodepipelines").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsCodepipelines) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awscodepipelines").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsCodepipeline.
func (c *awsCodepipelines) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsCodepipeline, err error) {
	result = &v1.AwsCodepipeline{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awscodepipelines").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
