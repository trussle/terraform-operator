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

// AwsInspectorAssessmentTargetsGetter has a method to return a AwsInspectorAssessmentTargetInterface.
// A group's client should implement this interface.
type AwsInspectorAssessmentTargetsGetter interface {
	AwsInspectorAssessmentTargets(namespace string) AwsInspectorAssessmentTargetInterface
}

// AwsInspectorAssessmentTargetInterface has methods to work with AwsInspectorAssessmentTarget resources.
type AwsInspectorAssessmentTargetInterface interface {
	Create(*v1.AwsInspectorAssessmentTarget) (*v1.AwsInspectorAssessmentTarget, error)
	Update(*v1.AwsInspectorAssessmentTarget) (*v1.AwsInspectorAssessmentTarget, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsInspectorAssessmentTarget, error)
	List(opts meta_v1.ListOptions) (*v1.AwsInspectorAssessmentTargetList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsInspectorAssessmentTarget, err error)
	AwsInspectorAssessmentTargetExpansion
}

// awsInspectorAssessmentTargets implements AwsInspectorAssessmentTargetInterface
type awsInspectorAssessmentTargets struct {
	client rest.Interface
	ns     string
}

// newAwsInspectorAssessmentTargets returns a AwsInspectorAssessmentTargets
func newAwsInspectorAssessmentTargets(c *TrussleV1Client, namespace string) *awsInspectorAssessmentTargets {
	return &awsInspectorAssessmentTargets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsInspectorAssessmentTarget, and returns the corresponding awsInspectorAssessmentTarget object, and an error if there is any.
func (c *awsInspectorAssessmentTargets) Get(name string, options meta_v1.GetOptions) (result *v1.AwsInspectorAssessmentTarget, err error) {
	result = &v1.AwsInspectorAssessmentTarget{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsinspectorassessmenttargets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsInspectorAssessmentTargets that match those selectors.
func (c *awsInspectorAssessmentTargets) List(opts meta_v1.ListOptions) (result *v1.AwsInspectorAssessmentTargetList, err error) {
	result = &v1.AwsInspectorAssessmentTargetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsinspectorassessmenttargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsInspectorAssessmentTargets.
func (c *awsInspectorAssessmentTargets) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsinspectorassessmenttargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsInspectorAssessmentTarget and creates it.  Returns the server's representation of the awsInspectorAssessmentTarget, and an error, if there is any.
func (c *awsInspectorAssessmentTargets) Create(awsInspectorAssessmentTarget *v1.AwsInspectorAssessmentTarget) (result *v1.AwsInspectorAssessmentTarget, err error) {
	result = &v1.AwsInspectorAssessmentTarget{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsinspectorassessmenttargets").
		Body(awsInspectorAssessmentTarget).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsInspectorAssessmentTarget and updates it. Returns the server's representation of the awsInspectorAssessmentTarget, and an error, if there is any.
func (c *awsInspectorAssessmentTargets) Update(awsInspectorAssessmentTarget *v1.AwsInspectorAssessmentTarget) (result *v1.AwsInspectorAssessmentTarget, err error) {
	result = &v1.AwsInspectorAssessmentTarget{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsinspectorassessmenttargets").
		Name(awsInspectorAssessmentTarget.Name).
		Body(awsInspectorAssessmentTarget).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsInspectorAssessmentTarget and deletes it. Returns an error if one occurs.
func (c *awsInspectorAssessmentTargets) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsinspectorassessmenttargets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsInspectorAssessmentTargets) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsinspectorassessmenttargets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsInspectorAssessmentTarget.
func (c *awsInspectorAssessmentTargets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsInspectorAssessmentTarget, err error) {
	result = &v1.AwsInspectorAssessmentTarget{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsinspectorassessmenttargets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
