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

// AwsDbOptionGroupsGetter has a method to return a AwsDbOptionGroupInterface.
// A group's client should implement this interface.
type AwsDbOptionGroupsGetter interface {
	AwsDbOptionGroups(namespace string) AwsDbOptionGroupInterface
}

// AwsDbOptionGroupInterface has methods to work with AwsDbOptionGroup resources.
type AwsDbOptionGroupInterface interface {
	Create(*v1.AwsDbOptionGroup) (*v1.AwsDbOptionGroup, error)
	Update(*v1.AwsDbOptionGroup) (*v1.AwsDbOptionGroup, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsDbOptionGroup, error)
	List(opts meta_v1.ListOptions) (*v1.AwsDbOptionGroupList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsDbOptionGroup, err error)
	AwsDbOptionGroupExpansion
}

// awsDbOptionGroups implements AwsDbOptionGroupInterface
type awsDbOptionGroups struct {
	client rest.Interface
	ns     string
}

// newAwsDbOptionGroups returns a AwsDbOptionGroups
func newAwsDbOptionGroups(c *TrussleV1Client, namespace string) *awsDbOptionGroups {
	return &awsDbOptionGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsDbOptionGroup, and returns the corresponding awsDbOptionGroup object, and an error if there is any.
func (c *awsDbOptionGroups) Get(name string, options meta_v1.GetOptions) (result *v1.AwsDbOptionGroup, err error) {
	result = &v1.AwsDbOptionGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsdboptiongroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsDbOptionGroups that match those selectors.
func (c *awsDbOptionGroups) List(opts meta_v1.ListOptions) (result *v1.AwsDbOptionGroupList, err error) {
	result = &v1.AwsDbOptionGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsdboptiongroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsDbOptionGroups.
func (c *awsDbOptionGroups) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsdboptiongroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsDbOptionGroup and creates it.  Returns the server's representation of the awsDbOptionGroup, and an error, if there is any.
func (c *awsDbOptionGroups) Create(awsDbOptionGroup *v1.AwsDbOptionGroup) (result *v1.AwsDbOptionGroup, err error) {
	result = &v1.AwsDbOptionGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsdboptiongroups").
		Body(awsDbOptionGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsDbOptionGroup and updates it. Returns the server's representation of the awsDbOptionGroup, and an error, if there is any.
func (c *awsDbOptionGroups) Update(awsDbOptionGroup *v1.AwsDbOptionGroup) (result *v1.AwsDbOptionGroup, err error) {
	result = &v1.AwsDbOptionGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsdboptiongroups").
		Name(awsDbOptionGroup.Name).
		Body(awsDbOptionGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsDbOptionGroup and deletes it. Returns an error if one occurs.
func (c *awsDbOptionGroups) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsdboptiongroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsDbOptionGroups) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsdboptiongroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsDbOptionGroup.
func (c *awsDbOptionGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsDbOptionGroup, err error) {
	result = &v1.AwsDbOptionGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsdboptiongroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
