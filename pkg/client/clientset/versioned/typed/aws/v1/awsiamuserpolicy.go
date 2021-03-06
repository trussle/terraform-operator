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

// AwsIamUserPoliciesGetter has a method to return a AwsIamUserPolicyInterface.
// A group's client should implement this interface.
type AwsIamUserPoliciesGetter interface {
	AwsIamUserPolicies(namespace string) AwsIamUserPolicyInterface
}

// AwsIamUserPolicyInterface has methods to work with AwsIamUserPolicy resources.
type AwsIamUserPolicyInterface interface {
	Create(*v1.AwsIamUserPolicy) (*v1.AwsIamUserPolicy, error)
	Update(*v1.AwsIamUserPolicy) (*v1.AwsIamUserPolicy, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsIamUserPolicy, error)
	List(opts meta_v1.ListOptions) (*v1.AwsIamUserPolicyList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsIamUserPolicy, err error)
	AwsIamUserPolicyExpansion
}

// awsIamUserPolicies implements AwsIamUserPolicyInterface
type awsIamUserPolicies struct {
	client rest.Interface
	ns     string
}

// newAwsIamUserPolicies returns a AwsIamUserPolicies
func newAwsIamUserPolicies(c *TrussleV1Client, namespace string) *awsIamUserPolicies {
	return &awsIamUserPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsIamUserPolicy, and returns the corresponding awsIamUserPolicy object, and an error if there is any.
func (c *awsIamUserPolicies) Get(name string, options meta_v1.GetOptions) (result *v1.AwsIamUserPolicy, err error) {
	result = &v1.AwsIamUserPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsiamuserpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsIamUserPolicies that match those selectors.
func (c *awsIamUserPolicies) List(opts meta_v1.ListOptions) (result *v1.AwsIamUserPolicyList, err error) {
	result = &v1.AwsIamUserPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsiamuserpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsIamUserPolicies.
func (c *awsIamUserPolicies) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsiamuserpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsIamUserPolicy and creates it.  Returns the server's representation of the awsIamUserPolicy, and an error, if there is any.
func (c *awsIamUserPolicies) Create(awsIamUserPolicy *v1.AwsIamUserPolicy) (result *v1.AwsIamUserPolicy, err error) {
	result = &v1.AwsIamUserPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsiamuserpolicies").
		Body(awsIamUserPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsIamUserPolicy and updates it. Returns the server's representation of the awsIamUserPolicy, and an error, if there is any.
func (c *awsIamUserPolicies) Update(awsIamUserPolicy *v1.AwsIamUserPolicy) (result *v1.AwsIamUserPolicy, err error) {
	result = &v1.AwsIamUserPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsiamuserpolicies").
		Name(awsIamUserPolicy.Name).
		Body(awsIamUserPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsIamUserPolicy and deletes it. Returns an error if one occurs.
func (c *awsIamUserPolicies) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsiamuserpolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsIamUserPolicies) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsiamuserpolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsIamUserPolicy.
func (c *awsIamUserPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsIamUserPolicy, err error) {
	result = &v1.AwsIamUserPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsiamuserpolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
