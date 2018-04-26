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

// AwsOpsworksUserProfilesGetter has a method to return a AwsOpsworksUserProfileInterface.
// A group's client should implement this interface.
type AwsOpsworksUserProfilesGetter interface {
	AwsOpsworksUserProfiles(namespace string) AwsOpsworksUserProfileInterface
}

// AwsOpsworksUserProfileInterface has methods to work with AwsOpsworksUserProfile resources.
type AwsOpsworksUserProfileInterface interface {
	Create(*v1.AwsOpsworksUserProfile) (*v1.AwsOpsworksUserProfile, error)
	Update(*v1.AwsOpsworksUserProfile) (*v1.AwsOpsworksUserProfile, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsOpsworksUserProfile, error)
	List(opts meta_v1.ListOptions) (*v1.AwsOpsworksUserProfileList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsOpsworksUserProfile, err error)
	AwsOpsworksUserProfileExpansion
}

// awsOpsworksUserProfiles implements AwsOpsworksUserProfileInterface
type awsOpsworksUserProfiles struct {
	client rest.Interface
	ns     string
}

// newAwsOpsworksUserProfiles returns a AwsOpsworksUserProfiles
func newAwsOpsworksUserProfiles(c *TrussleV1Client, namespace string) *awsOpsworksUserProfiles {
	return &awsOpsworksUserProfiles{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsOpsworksUserProfile, and returns the corresponding awsOpsworksUserProfile object, and an error if there is any.
func (c *awsOpsworksUserProfiles) Get(name string, options meta_v1.GetOptions) (result *v1.AwsOpsworksUserProfile, err error) {
	result = &v1.AwsOpsworksUserProfile{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsopsworksuserprofiles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsOpsworksUserProfiles that match those selectors.
func (c *awsOpsworksUserProfiles) List(opts meta_v1.ListOptions) (result *v1.AwsOpsworksUserProfileList, err error) {
	result = &v1.AwsOpsworksUserProfileList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsopsworksuserprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsOpsworksUserProfiles.
func (c *awsOpsworksUserProfiles) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsopsworksuserprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsOpsworksUserProfile and creates it.  Returns the server's representation of the awsOpsworksUserProfile, and an error, if there is any.
func (c *awsOpsworksUserProfiles) Create(awsOpsworksUserProfile *v1.AwsOpsworksUserProfile) (result *v1.AwsOpsworksUserProfile, err error) {
	result = &v1.AwsOpsworksUserProfile{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsopsworksuserprofiles").
		Body(awsOpsworksUserProfile).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsOpsworksUserProfile and updates it. Returns the server's representation of the awsOpsworksUserProfile, and an error, if there is any.
func (c *awsOpsworksUserProfiles) Update(awsOpsworksUserProfile *v1.AwsOpsworksUserProfile) (result *v1.AwsOpsworksUserProfile, err error) {
	result = &v1.AwsOpsworksUserProfile{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsopsworksuserprofiles").
		Name(awsOpsworksUserProfile.Name).
		Body(awsOpsworksUserProfile).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsOpsworksUserProfile and deletes it. Returns an error if one occurs.
func (c *awsOpsworksUserProfiles) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsopsworksuserprofiles").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsOpsworksUserProfiles) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsopsworksuserprofiles").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsOpsworksUserProfile.
func (c *awsOpsworksUserProfiles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsOpsworksUserProfile, err error) {
	result = &v1.AwsOpsworksUserProfile{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsopsworksuserprofiles").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
