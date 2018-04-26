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

// AwsEcrLifecyclePoliciesGetter has a method to return a AwsEcrLifecyclePolicyInterface.
// A group's client should implement this interface.
type AwsEcrLifecyclePoliciesGetter interface {
	AwsEcrLifecyclePolicies(namespace string) AwsEcrLifecyclePolicyInterface
}

// AwsEcrLifecyclePolicyInterface has methods to work with AwsEcrLifecyclePolicy resources.
type AwsEcrLifecyclePolicyInterface interface {
	Create(*v1.AwsEcrLifecyclePolicy) (*v1.AwsEcrLifecyclePolicy, error)
	Update(*v1.AwsEcrLifecyclePolicy) (*v1.AwsEcrLifecyclePolicy, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsEcrLifecyclePolicy, error)
	List(opts meta_v1.ListOptions) (*v1.AwsEcrLifecyclePolicyList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsEcrLifecyclePolicy, err error)
	AwsEcrLifecyclePolicyExpansion
}

// awsEcrLifecyclePolicies implements AwsEcrLifecyclePolicyInterface
type awsEcrLifecyclePolicies struct {
	client rest.Interface
	ns     string
}

// newAwsEcrLifecyclePolicies returns a AwsEcrLifecyclePolicies
func newAwsEcrLifecyclePolicies(c *TrussleV1Client, namespace string) *awsEcrLifecyclePolicies {
	return &awsEcrLifecyclePolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsEcrLifecyclePolicy, and returns the corresponding awsEcrLifecyclePolicy object, and an error if there is any.
func (c *awsEcrLifecyclePolicies) Get(name string, options meta_v1.GetOptions) (result *v1.AwsEcrLifecyclePolicy, err error) {
	result = &v1.AwsEcrLifecyclePolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsecrlifecyclepolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsEcrLifecyclePolicies that match those selectors.
func (c *awsEcrLifecyclePolicies) List(opts meta_v1.ListOptions) (result *v1.AwsEcrLifecyclePolicyList, err error) {
	result = &v1.AwsEcrLifecyclePolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsecrlifecyclepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsEcrLifecyclePolicies.
func (c *awsEcrLifecyclePolicies) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsecrlifecyclepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsEcrLifecyclePolicy and creates it.  Returns the server's representation of the awsEcrLifecyclePolicy, and an error, if there is any.
func (c *awsEcrLifecyclePolicies) Create(awsEcrLifecyclePolicy *v1.AwsEcrLifecyclePolicy) (result *v1.AwsEcrLifecyclePolicy, err error) {
	result = &v1.AwsEcrLifecyclePolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsecrlifecyclepolicies").
		Body(awsEcrLifecyclePolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsEcrLifecyclePolicy and updates it. Returns the server's representation of the awsEcrLifecyclePolicy, and an error, if there is any.
func (c *awsEcrLifecyclePolicies) Update(awsEcrLifecyclePolicy *v1.AwsEcrLifecyclePolicy) (result *v1.AwsEcrLifecyclePolicy, err error) {
	result = &v1.AwsEcrLifecyclePolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsecrlifecyclepolicies").
		Name(awsEcrLifecyclePolicy.Name).
		Body(awsEcrLifecyclePolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsEcrLifecyclePolicy and deletes it. Returns an error if one occurs.
func (c *awsEcrLifecyclePolicies) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsecrlifecyclepolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsEcrLifecyclePolicies) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsecrlifecyclepolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsEcrLifecyclePolicy.
func (c *awsEcrLifecyclePolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsEcrLifecyclePolicy, err error) {
	result = &v1.AwsEcrLifecyclePolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsecrlifecyclepolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
