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

// AwsGlueCatalogDatabasesGetter has a method to return a AwsGlueCatalogDatabaseInterface.
// A group's client should implement this interface.
type AwsGlueCatalogDatabasesGetter interface {
	AwsGlueCatalogDatabases(namespace string) AwsGlueCatalogDatabaseInterface
}

// AwsGlueCatalogDatabaseInterface has methods to work with AwsGlueCatalogDatabase resources.
type AwsGlueCatalogDatabaseInterface interface {
	Create(*v1.AwsGlueCatalogDatabase) (*v1.AwsGlueCatalogDatabase, error)
	Update(*v1.AwsGlueCatalogDatabase) (*v1.AwsGlueCatalogDatabase, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsGlueCatalogDatabase, error)
	List(opts meta_v1.ListOptions) (*v1.AwsGlueCatalogDatabaseList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsGlueCatalogDatabase, err error)
	AwsGlueCatalogDatabaseExpansion
}

// awsGlueCatalogDatabases implements AwsGlueCatalogDatabaseInterface
type awsGlueCatalogDatabases struct {
	client rest.Interface
	ns     string
}

// newAwsGlueCatalogDatabases returns a AwsGlueCatalogDatabases
func newAwsGlueCatalogDatabases(c *TrussleV1Client, namespace string) *awsGlueCatalogDatabases {
	return &awsGlueCatalogDatabases{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsGlueCatalogDatabase, and returns the corresponding awsGlueCatalogDatabase object, and an error if there is any.
func (c *awsGlueCatalogDatabases) Get(name string, options meta_v1.GetOptions) (result *v1.AwsGlueCatalogDatabase, err error) {
	result = &v1.AwsGlueCatalogDatabase{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsgluecatalogdatabases").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsGlueCatalogDatabases that match those selectors.
func (c *awsGlueCatalogDatabases) List(opts meta_v1.ListOptions) (result *v1.AwsGlueCatalogDatabaseList, err error) {
	result = &v1.AwsGlueCatalogDatabaseList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsgluecatalogdatabases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsGlueCatalogDatabases.
func (c *awsGlueCatalogDatabases) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsgluecatalogdatabases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsGlueCatalogDatabase and creates it.  Returns the server's representation of the awsGlueCatalogDatabase, and an error, if there is any.
func (c *awsGlueCatalogDatabases) Create(awsGlueCatalogDatabase *v1.AwsGlueCatalogDatabase) (result *v1.AwsGlueCatalogDatabase, err error) {
	result = &v1.AwsGlueCatalogDatabase{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsgluecatalogdatabases").
		Body(awsGlueCatalogDatabase).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsGlueCatalogDatabase and updates it. Returns the server's representation of the awsGlueCatalogDatabase, and an error, if there is any.
func (c *awsGlueCatalogDatabases) Update(awsGlueCatalogDatabase *v1.AwsGlueCatalogDatabase) (result *v1.AwsGlueCatalogDatabase, err error) {
	result = &v1.AwsGlueCatalogDatabase{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsgluecatalogdatabases").
		Name(awsGlueCatalogDatabase.Name).
		Body(awsGlueCatalogDatabase).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsGlueCatalogDatabase and deletes it. Returns an error if one occurs.
func (c *awsGlueCatalogDatabases) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsgluecatalogdatabases").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsGlueCatalogDatabases) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsgluecatalogdatabases").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsGlueCatalogDatabase.
func (c *awsGlueCatalogDatabases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsGlueCatalogDatabase, err error) {
	result = &v1.AwsGlueCatalogDatabase{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsgluecatalogdatabases").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
