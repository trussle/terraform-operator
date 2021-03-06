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

// AwsAthenaDatabasesGetter has a method to return a AwsAthenaDatabaseInterface.
// A group's client should implement this interface.
type AwsAthenaDatabasesGetter interface {
	AwsAthenaDatabases(namespace string) AwsAthenaDatabaseInterface
}

// AwsAthenaDatabaseInterface has methods to work with AwsAthenaDatabase resources.
type AwsAthenaDatabaseInterface interface {
	Create(*v1.AwsAthenaDatabase) (*v1.AwsAthenaDatabase, error)
	Update(*v1.AwsAthenaDatabase) (*v1.AwsAthenaDatabase, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsAthenaDatabase, error)
	List(opts meta_v1.ListOptions) (*v1.AwsAthenaDatabaseList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsAthenaDatabase, err error)
	AwsAthenaDatabaseExpansion
}

// awsAthenaDatabases implements AwsAthenaDatabaseInterface
type awsAthenaDatabases struct {
	client rest.Interface
	ns     string
}

// newAwsAthenaDatabases returns a AwsAthenaDatabases
func newAwsAthenaDatabases(c *TrussleV1Client, namespace string) *awsAthenaDatabases {
	return &awsAthenaDatabases{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsAthenaDatabase, and returns the corresponding awsAthenaDatabase object, and an error if there is any.
func (c *awsAthenaDatabases) Get(name string, options meta_v1.GetOptions) (result *v1.AwsAthenaDatabase, err error) {
	result = &v1.AwsAthenaDatabase{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsathenadatabases").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsAthenaDatabases that match those selectors.
func (c *awsAthenaDatabases) List(opts meta_v1.ListOptions) (result *v1.AwsAthenaDatabaseList, err error) {
	result = &v1.AwsAthenaDatabaseList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsathenadatabases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsAthenaDatabases.
func (c *awsAthenaDatabases) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsathenadatabases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsAthenaDatabase and creates it.  Returns the server's representation of the awsAthenaDatabase, and an error, if there is any.
func (c *awsAthenaDatabases) Create(awsAthenaDatabase *v1.AwsAthenaDatabase) (result *v1.AwsAthenaDatabase, err error) {
	result = &v1.AwsAthenaDatabase{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsathenadatabases").
		Body(awsAthenaDatabase).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsAthenaDatabase and updates it. Returns the server's representation of the awsAthenaDatabase, and an error, if there is any.
func (c *awsAthenaDatabases) Update(awsAthenaDatabase *v1.AwsAthenaDatabase) (result *v1.AwsAthenaDatabase, err error) {
	result = &v1.AwsAthenaDatabase{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsathenadatabases").
		Name(awsAthenaDatabase.Name).
		Body(awsAthenaDatabase).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsAthenaDatabase and deletes it. Returns an error if one occurs.
func (c *awsAthenaDatabases) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsathenadatabases").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsAthenaDatabases) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsathenadatabases").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsAthenaDatabase.
func (c *awsAthenaDatabases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsAthenaDatabase, err error) {
	result = &v1.AwsAthenaDatabase{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsathenadatabases").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
