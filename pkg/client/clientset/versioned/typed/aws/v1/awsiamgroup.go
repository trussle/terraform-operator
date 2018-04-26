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

// AwsIamGroupsGetter has a method to return a AwsIamGroupInterface.
// A group's client should implement this interface.
type AwsIamGroupsGetter interface {
	AwsIamGroups(namespace string) AwsIamGroupInterface
}

// AwsIamGroupInterface has methods to work with AwsIamGroup resources.
type AwsIamGroupInterface interface {
	Create(*v1.AwsIamGroup) (*v1.AwsIamGroup, error)
	Update(*v1.AwsIamGroup) (*v1.AwsIamGroup, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsIamGroup, error)
	List(opts meta_v1.ListOptions) (*v1.AwsIamGroupList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsIamGroup, err error)
	AwsIamGroupExpansion
}

// awsIamGroups implements AwsIamGroupInterface
type awsIamGroups struct {
	client rest.Interface
	ns     string
}

// newAwsIamGroups returns a AwsIamGroups
func newAwsIamGroups(c *TrussleV1Client, namespace string) *awsIamGroups {
	return &awsIamGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsIamGroup, and returns the corresponding awsIamGroup object, and an error if there is any.
func (c *awsIamGroups) Get(name string, options meta_v1.GetOptions) (result *v1.AwsIamGroup, err error) {
	result = &v1.AwsIamGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsiamgroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsIamGroups that match those selectors.
func (c *awsIamGroups) List(opts meta_v1.ListOptions) (result *v1.AwsIamGroupList, err error) {
	result = &v1.AwsIamGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsiamgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsIamGroups.
func (c *awsIamGroups) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsiamgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsIamGroup and creates it.  Returns the server's representation of the awsIamGroup, and an error, if there is any.
func (c *awsIamGroups) Create(awsIamGroup *v1.AwsIamGroup) (result *v1.AwsIamGroup, err error) {
	result = &v1.AwsIamGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsiamgroups").
		Body(awsIamGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsIamGroup and updates it. Returns the server's representation of the awsIamGroup, and an error, if there is any.
func (c *awsIamGroups) Update(awsIamGroup *v1.AwsIamGroup) (result *v1.AwsIamGroup, err error) {
	result = &v1.AwsIamGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsiamgroups").
		Name(awsIamGroup.Name).
		Body(awsIamGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsIamGroup and deletes it. Returns an error if one occurs.
func (c *awsIamGroups) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsiamgroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsIamGroups) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsiamgroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsIamGroup.
func (c *awsIamGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsIamGroup, err error) {
	result = &v1.AwsIamGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsiamgroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}