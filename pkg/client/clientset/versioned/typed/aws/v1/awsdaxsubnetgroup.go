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

// AwsDaxSubnetGroupsGetter has a method to return a AwsDaxSubnetGroupInterface.
// A group's client should implement this interface.
type AwsDaxSubnetGroupsGetter interface {
	AwsDaxSubnetGroups(namespace string) AwsDaxSubnetGroupInterface
}

// AwsDaxSubnetGroupInterface has methods to work with AwsDaxSubnetGroup resources.
type AwsDaxSubnetGroupInterface interface {
	Create(*v1.AwsDaxSubnetGroup) (*v1.AwsDaxSubnetGroup, error)
	Update(*v1.AwsDaxSubnetGroup) (*v1.AwsDaxSubnetGroup, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsDaxSubnetGroup, error)
	List(opts meta_v1.ListOptions) (*v1.AwsDaxSubnetGroupList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsDaxSubnetGroup, err error)
	AwsDaxSubnetGroupExpansion
}

// awsDaxSubnetGroups implements AwsDaxSubnetGroupInterface
type awsDaxSubnetGroups struct {
	client rest.Interface
	ns     string
}

// newAwsDaxSubnetGroups returns a AwsDaxSubnetGroups
func newAwsDaxSubnetGroups(c *TrussleV1Client, namespace string) *awsDaxSubnetGroups {
	return &awsDaxSubnetGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsDaxSubnetGroup, and returns the corresponding awsDaxSubnetGroup object, and an error if there is any.
func (c *awsDaxSubnetGroups) Get(name string, options meta_v1.GetOptions) (result *v1.AwsDaxSubnetGroup, err error) {
	result = &v1.AwsDaxSubnetGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsdaxsubnetgroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsDaxSubnetGroups that match those selectors.
func (c *awsDaxSubnetGroups) List(opts meta_v1.ListOptions) (result *v1.AwsDaxSubnetGroupList, err error) {
	result = &v1.AwsDaxSubnetGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsdaxsubnetgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsDaxSubnetGroups.
func (c *awsDaxSubnetGroups) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsdaxsubnetgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsDaxSubnetGroup and creates it.  Returns the server's representation of the awsDaxSubnetGroup, and an error, if there is any.
func (c *awsDaxSubnetGroups) Create(awsDaxSubnetGroup *v1.AwsDaxSubnetGroup) (result *v1.AwsDaxSubnetGroup, err error) {
	result = &v1.AwsDaxSubnetGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsdaxsubnetgroups").
		Body(awsDaxSubnetGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsDaxSubnetGroup and updates it. Returns the server's representation of the awsDaxSubnetGroup, and an error, if there is any.
func (c *awsDaxSubnetGroups) Update(awsDaxSubnetGroup *v1.AwsDaxSubnetGroup) (result *v1.AwsDaxSubnetGroup, err error) {
	result = &v1.AwsDaxSubnetGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsdaxsubnetgroups").
		Name(awsDaxSubnetGroup.Name).
		Body(awsDaxSubnetGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsDaxSubnetGroup and deletes it. Returns an error if one occurs.
func (c *awsDaxSubnetGroups) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsdaxsubnetgroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsDaxSubnetGroups) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsdaxsubnetgroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsDaxSubnetGroup.
func (c *awsDaxSubnetGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsDaxSubnetGroup, err error) {
	result = &v1.AwsDaxSubnetGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsdaxsubnetgroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
