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

// AwsCodebuildProjectsGetter has a method to return a AwsCodebuildProjectInterface.
// A group's client should implement this interface.
type AwsCodebuildProjectsGetter interface {
	AwsCodebuildProjects(namespace string) AwsCodebuildProjectInterface
}

// AwsCodebuildProjectInterface has methods to work with AwsCodebuildProject resources.
type AwsCodebuildProjectInterface interface {
	Create(*v1.AwsCodebuildProject) (*v1.AwsCodebuildProject, error)
	Update(*v1.AwsCodebuildProject) (*v1.AwsCodebuildProject, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsCodebuildProject, error)
	List(opts meta_v1.ListOptions) (*v1.AwsCodebuildProjectList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsCodebuildProject, err error)
	AwsCodebuildProjectExpansion
}

// awsCodebuildProjects implements AwsCodebuildProjectInterface
type awsCodebuildProjects struct {
	client rest.Interface
	ns     string
}

// newAwsCodebuildProjects returns a AwsCodebuildProjects
func newAwsCodebuildProjects(c *TrussleV1Client, namespace string) *awsCodebuildProjects {
	return &awsCodebuildProjects{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsCodebuildProject, and returns the corresponding awsCodebuildProject object, and an error if there is any.
func (c *awsCodebuildProjects) Get(name string, options meta_v1.GetOptions) (result *v1.AwsCodebuildProject, err error) {
	result = &v1.AwsCodebuildProject{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awscodebuildprojects").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsCodebuildProjects that match those selectors.
func (c *awsCodebuildProjects) List(opts meta_v1.ListOptions) (result *v1.AwsCodebuildProjectList, err error) {
	result = &v1.AwsCodebuildProjectList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awscodebuildprojects").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsCodebuildProjects.
func (c *awsCodebuildProjects) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awscodebuildprojects").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsCodebuildProject and creates it.  Returns the server's representation of the awsCodebuildProject, and an error, if there is any.
func (c *awsCodebuildProjects) Create(awsCodebuildProject *v1.AwsCodebuildProject) (result *v1.AwsCodebuildProject, err error) {
	result = &v1.AwsCodebuildProject{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awscodebuildprojects").
		Body(awsCodebuildProject).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsCodebuildProject and updates it. Returns the server's representation of the awsCodebuildProject, and an error, if there is any.
func (c *awsCodebuildProjects) Update(awsCodebuildProject *v1.AwsCodebuildProject) (result *v1.AwsCodebuildProject, err error) {
	result = &v1.AwsCodebuildProject{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awscodebuildprojects").
		Name(awsCodebuildProject.Name).
		Body(awsCodebuildProject).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsCodebuildProject and deletes it. Returns an error if one occurs.
func (c *awsCodebuildProjects) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awscodebuildprojects").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsCodebuildProjects) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awscodebuildprojects").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsCodebuildProject.
func (c *awsCodebuildProjects) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsCodebuildProject, err error) {
	result = &v1.AwsCodebuildProject{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awscodebuildprojects").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
