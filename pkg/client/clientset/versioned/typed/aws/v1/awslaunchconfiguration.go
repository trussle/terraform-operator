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

// AwsLaunchConfigurationsGetter has a method to return a AwsLaunchConfigurationInterface.
// A group's client should implement this interface.
type AwsLaunchConfigurationsGetter interface {
	AwsLaunchConfigurations(namespace string) AwsLaunchConfigurationInterface
}

// AwsLaunchConfigurationInterface has methods to work with AwsLaunchConfiguration resources.
type AwsLaunchConfigurationInterface interface {
	Create(*v1.AwsLaunchConfiguration) (*v1.AwsLaunchConfiguration, error)
	Update(*v1.AwsLaunchConfiguration) (*v1.AwsLaunchConfiguration, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsLaunchConfiguration, error)
	List(opts meta_v1.ListOptions) (*v1.AwsLaunchConfigurationList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsLaunchConfiguration, err error)
	AwsLaunchConfigurationExpansion
}

// awsLaunchConfigurations implements AwsLaunchConfigurationInterface
type awsLaunchConfigurations struct {
	client rest.Interface
	ns     string
}

// newAwsLaunchConfigurations returns a AwsLaunchConfigurations
func newAwsLaunchConfigurations(c *TrussleV1Client, namespace string) *awsLaunchConfigurations {
	return &awsLaunchConfigurations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsLaunchConfiguration, and returns the corresponding awsLaunchConfiguration object, and an error if there is any.
func (c *awsLaunchConfigurations) Get(name string, options meta_v1.GetOptions) (result *v1.AwsLaunchConfiguration, err error) {
	result = &v1.AwsLaunchConfiguration{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awslaunchconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsLaunchConfigurations that match those selectors.
func (c *awsLaunchConfigurations) List(opts meta_v1.ListOptions) (result *v1.AwsLaunchConfigurationList, err error) {
	result = &v1.AwsLaunchConfigurationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awslaunchconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsLaunchConfigurations.
func (c *awsLaunchConfigurations) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awslaunchconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsLaunchConfiguration and creates it.  Returns the server's representation of the awsLaunchConfiguration, and an error, if there is any.
func (c *awsLaunchConfigurations) Create(awsLaunchConfiguration *v1.AwsLaunchConfiguration) (result *v1.AwsLaunchConfiguration, err error) {
	result = &v1.AwsLaunchConfiguration{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awslaunchconfigurations").
		Body(awsLaunchConfiguration).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsLaunchConfiguration and updates it. Returns the server's representation of the awsLaunchConfiguration, and an error, if there is any.
func (c *awsLaunchConfigurations) Update(awsLaunchConfiguration *v1.AwsLaunchConfiguration) (result *v1.AwsLaunchConfiguration, err error) {
	result = &v1.AwsLaunchConfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awslaunchconfigurations").
		Name(awsLaunchConfiguration.Name).
		Body(awsLaunchConfiguration).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsLaunchConfiguration and deletes it. Returns an error if one occurs.
func (c *awsLaunchConfigurations) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awslaunchconfigurations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsLaunchConfigurations) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awslaunchconfigurations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsLaunchConfiguration.
func (c *awsLaunchConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsLaunchConfiguration, err error) {
	result = &v1.AwsLaunchConfiguration{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awslaunchconfigurations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
