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

// AwsEcrRepositoryPoliciesGetter has a method to return a AwsEcrRepositoryPolicyInterface.
// A group's client should implement this interface.
type AwsEcrRepositoryPoliciesGetter interface {
	AwsEcrRepositoryPolicies(namespace string) AwsEcrRepositoryPolicyInterface
}

// AwsEcrRepositoryPolicyInterface has methods to work with AwsEcrRepositoryPolicy resources.
type AwsEcrRepositoryPolicyInterface interface {
	Create(*v1.AwsEcrRepositoryPolicy) (*v1.AwsEcrRepositoryPolicy, error)
	Update(*v1.AwsEcrRepositoryPolicy) (*v1.AwsEcrRepositoryPolicy, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsEcrRepositoryPolicy, error)
	List(opts meta_v1.ListOptions) (*v1.AwsEcrRepositoryPolicyList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsEcrRepositoryPolicy, err error)
	AwsEcrRepositoryPolicyExpansion
}

// awsEcrRepositoryPolicies implements AwsEcrRepositoryPolicyInterface
type awsEcrRepositoryPolicies struct {
	client rest.Interface
	ns     string
}

// newAwsEcrRepositoryPolicies returns a AwsEcrRepositoryPolicies
func newAwsEcrRepositoryPolicies(c *TrussleV1Client, namespace string) *awsEcrRepositoryPolicies {
	return &awsEcrRepositoryPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsEcrRepositoryPolicy, and returns the corresponding awsEcrRepositoryPolicy object, and an error if there is any.
func (c *awsEcrRepositoryPolicies) Get(name string, options meta_v1.GetOptions) (result *v1.AwsEcrRepositoryPolicy, err error) {
	result = &v1.AwsEcrRepositoryPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsecrrepositorypolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsEcrRepositoryPolicies that match those selectors.
func (c *awsEcrRepositoryPolicies) List(opts meta_v1.ListOptions) (result *v1.AwsEcrRepositoryPolicyList, err error) {
	result = &v1.AwsEcrRepositoryPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsecrrepositorypolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsEcrRepositoryPolicies.
func (c *awsEcrRepositoryPolicies) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsecrrepositorypolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsEcrRepositoryPolicy and creates it.  Returns the server's representation of the awsEcrRepositoryPolicy, and an error, if there is any.
func (c *awsEcrRepositoryPolicies) Create(awsEcrRepositoryPolicy *v1.AwsEcrRepositoryPolicy) (result *v1.AwsEcrRepositoryPolicy, err error) {
	result = &v1.AwsEcrRepositoryPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsecrrepositorypolicies").
		Body(awsEcrRepositoryPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsEcrRepositoryPolicy and updates it. Returns the server's representation of the awsEcrRepositoryPolicy, and an error, if there is any.
func (c *awsEcrRepositoryPolicies) Update(awsEcrRepositoryPolicy *v1.AwsEcrRepositoryPolicy) (result *v1.AwsEcrRepositoryPolicy, err error) {
	result = &v1.AwsEcrRepositoryPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsecrrepositorypolicies").
		Name(awsEcrRepositoryPolicy.Name).
		Body(awsEcrRepositoryPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsEcrRepositoryPolicy and deletes it. Returns an error if one occurs.
func (c *awsEcrRepositoryPolicies) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsecrrepositorypolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsEcrRepositoryPolicies) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsecrrepositorypolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsEcrRepositoryPolicy.
func (c *awsEcrRepositoryPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsEcrRepositoryPolicy, err error) {
	result = &v1.AwsEcrRepositoryPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsecrrepositorypolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
