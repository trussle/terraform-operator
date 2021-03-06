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

// AwsAutoscalingNotificationsGetter has a method to return a AwsAutoscalingNotificationInterface.
// A group's client should implement this interface.
type AwsAutoscalingNotificationsGetter interface {
	AwsAutoscalingNotifications(namespace string) AwsAutoscalingNotificationInterface
}

// AwsAutoscalingNotificationInterface has methods to work with AwsAutoscalingNotification resources.
type AwsAutoscalingNotificationInterface interface {
	Create(*v1.AwsAutoscalingNotification) (*v1.AwsAutoscalingNotification, error)
	Update(*v1.AwsAutoscalingNotification) (*v1.AwsAutoscalingNotification, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsAutoscalingNotification, error)
	List(opts meta_v1.ListOptions) (*v1.AwsAutoscalingNotificationList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsAutoscalingNotification, err error)
	AwsAutoscalingNotificationExpansion
}

// awsAutoscalingNotifications implements AwsAutoscalingNotificationInterface
type awsAutoscalingNotifications struct {
	client rest.Interface
	ns     string
}

// newAwsAutoscalingNotifications returns a AwsAutoscalingNotifications
func newAwsAutoscalingNotifications(c *TrussleV1Client, namespace string) *awsAutoscalingNotifications {
	return &awsAutoscalingNotifications{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsAutoscalingNotification, and returns the corresponding awsAutoscalingNotification object, and an error if there is any.
func (c *awsAutoscalingNotifications) Get(name string, options meta_v1.GetOptions) (result *v1.AwsAutoscalingNotification, err error) {
	result = &v1.AwsAutoscalingNotification{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsautoscalingnotifications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsAutoscalingNotifications that match those selectors.
func (c *awsAutoscalingNotifications) List(opts meta_v1.ListOptions) (result *v1.AwsAutoscalingNotificationList, err error) {
	result = &v1.AwsAutoscalingNotificationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsautoscalingnotifications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsAutoscalingNotifications.
func (c *awsAutoscalingNotifications) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsautoscalingnotifications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsAutoscalingNotification and creates it.  Returns the server's representation of the awsAutoscalingNotification, and an error, if there is any.
func (c *awsAutoscalingNotifications) Create(awsAutoscalingNotification *v1.AwsAutoscalingNotification) (result *v1.AwsAutoscalingNotification, err error) {
	result = &v1.AwsAutoscalingNotification{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsautoscalingnotifications").
		Body(awsAutoscalingNotification).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsAutoscalingNotification and updates it. Returns the server's representation of the awsAutoscalingNotification, and an error, if there is any.
func (c *awsAutoscalingNotifications) Update(awsAutoscalingNotification *v1.AwsAutoscalingNotification) (result *v1.AwsAutoscalingNotification, err error) {
	result = &v1.AwsAutoscalingNotification{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsautoscalingnotifications").
		Name(awsAutoscalingNotification.Name).
		Body(awsAutoscalingNotification).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsAutoscalingNotification and deletes it. Returns an error if one occurs.
func (c *awsAutoscalingNotifications) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsautoscalingnotifications").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsAutoscalingNotifications) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsautoscalingnotifications").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsAutoscalingNotification.
func (c *awsAutoscalingNotifications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsAutoscalingNotification, err error) {
	result = &v1.AwsAutoscalingNotification{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsautoscalingnotifications").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
