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

// AwsLaunchTemplatesGetter has a method to return a AwsLaunchTemplateInterface.
// A group's client should implement this interface.
type AwsLaunchTemplatesGetter interface {
	AwsLaunchTemplates(namespace string) AwsLaunchTemplateInterface
}

// AwsLaunchTemplateInterface has methods to work with AwsLaunchTemplate resources.
type AwsLaunchTemplateInterface interface {
	Create(*v1.AwsLaunchTemplate) (*v1.AwsLaunchTemplate, error)
	Update(*v1.AwsLaunchTemplate) (*v1.AwsLaunchTemplate, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsLaunchTemplate, error)
	List(opts meta_v1.ListOptions) (*v1.AwsLaunchTemplateList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsLaunchTemplate, err error)
	AwsLaunchTemplateExpansion
}

// awsLaunchTemplates implements AwsLaunchTemplateInterface
type awsLaunchTemplates struct {
	client rest.Interface
	ns     string
}

// newAwsLaunchTemplates returns a AwsLaunchTemplates
func newAwsLaunchTemplates(c *TrussleV1Client, namespace string) *awsLaunchTemplates {
	return &awsLaunchTemplates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsLaunchTemplate, and returns the corresponding awsLaunchTemplate object, and an error if there is any.
func (c *awsLaunchTemplates) Get(name string, options meta_v1.GetOptions) (result *v1.AwsLaunchTemplate, err error) {
	result = &v1.AwsLaunchTemplate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awslaunchtemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsLaunchTemplates that match those selectors.
func (c *awsLaunchTemplates) List(opts meta_v1.ListOptions) (result *v1.AwsLaunchTemplateList, err error) {
	result = &v1.AwsLaunchTemplateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awslaunchtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsLaunchTemplates.
func (c *awsLaunchTemplates) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awslaunchtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsLaunchTemplate and creates it.  Returns the server's representation of the awsLaunchTemplate, and an error, if there is any.
func (c *awsLaunchTemplates) Create(awsLaunchTemplate *v1.AwsLaunchTemplate) (result *v1.AwsLaunchTemplate, err error) {
	result = &v1.AwsLaunchTemplate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awslaunchtemplates").
		Body(awsLaunchTemplate).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsLaunchTemplate and updates it. Returns the server's representation of the awsLaunchTemplate, and an error, if there is any.
func (c *awsLaunchTemplates) Update(awsLaunchTemplate *v1.AwsLaunchTemplate) (result *v1.AwsLaunchTemplate, err error) {
	result = &v1.AwsLaunchTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awslaunchtemplates").
		Name(awsLaunchTemplate.Name).
		Body(awsLaunchTemplate).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsLaunchTemplate and deletes it. Returns an error if one occurs.
func (c *awsLaunchTemplates) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awslaunchtemplates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsLaunchTemplates) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awslaunchtemplates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsLaunchTemplate.
func (c *awsLaunchTemplates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsLaunchTemplate, err error) {
	result = &v1.AwsLaunchTemplate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awslaunchtemplates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}