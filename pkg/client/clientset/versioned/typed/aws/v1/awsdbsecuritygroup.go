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

// AwsDbSecurityGroupsGetter has a method to return a AwsDbSecurityGroupInterface.
// A group's client should implement this interface.
type AwsDbSecurityGroupsGetter interface {
	AwsDbSecurityGroups(namespace string) AwsDbSecurityGroupInterface
}

// AwsDbSecurityGroupInterface has methods to work with AwsDbSecurityGroup resources.
type AwsDbSecurityGroupInterface interface {
	Create(*v1.AwsDbSecurityGroup) (*v1.AwsDbSecurityGroup, error)
	Update(*v1.AwsDbSecurityGroup) (*v1.AwsDbSecurityGroup, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsDbSecurityGroup, error)
	List(opts meta_v1.ListOptions) (*v1.AwsDbSecurityGroupList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsDbSecurityGroup, err error)
	AwsDbSecurityGroupExpansion
}

// awsDbSecurityGroups implements AwsDbSecurityGroupInterface
type awsDbSecurityGroups struct {
	client rest.Interface
	ns     string
}

// newAwsDbSecurityGroups returns a AwsDbSecurityGroups
func newAwsDbSecurityGroups(c *TrussleV1Client, namespace string) *awsDbSecurityGroups {
	return &awsDbSecurityGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsDbSecurityGroup, and returns the corresponding awsDbSecurityGroup object, and an error if there is any.
func (c *awsDbSecurityGroups) Get(name string, options meta_v1.GetOptions) (result *v1.AwsDbSecurityGroup, err error) {
	result = &v1.AwsDbSecurityGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsdbsecuritygroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsDbSecurityGroups that match those selectors.
func (c *awsDbSecurityGroups) List(opts meta_v1.ListOptions) (result *v1.AwsDbSecurityGroupList, err error) {
	result = &v1.AwsDbSecurityGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsdbsecuritygroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsDbSecurityGroups.
func (c *awsDbSecurityGroups) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsdbsecuritygroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsDbSecurityGroup and creates it.  Returns the server's representation of the awsDbSecurityGroup, and an error, if there is any.
func (c *awsDbSecurityGroups) Create(awsDbSecurityGroup *v1.AwsDbSecurityGroup) (result *v1.AwsDbSecurityGroup, err error) {
	result = &v1.AwsDbSecurityGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsdbsecuritygroups").
		Body(awsDbSecurityGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsDbSecurityGroup and updates it. Returns the server's representation of the awsDbSecurityGroup, and an error, if there is any.
func (c *awsDbSecurityGroups) Update(awsDbSecurityGroup *v1.AwsDbSecurityGroup) (result *v1.AwsDbSecurityGroup, err error) {
	result = &v1.AwsDbSecurityGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsdbsecuritygroups").
		Name(awsDbSecurityGroup.Name).
		Body(awsDbSecurityGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsDbSecurityGroup and deletes it. Returns an error if one occurs.
func (c *awsDbSecurityGroups) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsdbsecuritygroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsDbSecurityGroups) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsdbsecuritygroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsDbSecurityGroup.
func (c *awsDbSecurityGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsDbSecurityGroup, err error) {
	result = &v1.AwsDbSecurityGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsdbsecuritygroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
