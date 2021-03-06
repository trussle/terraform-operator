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

// AwsSqsQueuesGetter has a method to return a AwsSqsQueueInterface.
// A group's client should implement this interface.
type AwsSqsQueuesGetter interface {
	AwsSqsQueues(namespace string) AwsSqsQueueInterface
}

// AwsSqsQueueInterface has methods to work with AwsSqsQueue resources.
type AwsSqsQueueInterface interface {
	Create(*v1.AwsSqsQueue) (*v1.AwsSqsQueue, error)
	Update(*v1.AwsSqsQueue) (*v1.AwsSqsQueue, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsSqsQueue, error)
	List(opts meta_v1.ListOptions) (*v1.AwsSqsQueueList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsSqsQueue, err error)
	AwsSqsQueueExpansion
}

// awsSqsQueues implements AwsSqsQueueInterface
type awsSqsQueues struct {
	client rest.Interface
	ns     string
}

// newAwsSqsQueues returns a AwsSqsQueues
func newAwsSqsQueues(c *TrussleV1Client, namespace string) *awsSqsQueues {
	return &awsSqsQueues{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsSqsQueue, and returns the corresponding awsSqsQueue object, and an error if there is any.
func (c *awsSqsQueues) Get(name string, options meta_v1.GetOptions) (result *v1.AwsSqsQueue, err error) {
	result = &v1.AwsSqsQueue{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awssqsqueues").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsSqsQueues that match those selectors.
func (c *awsSqsQueues) List(opts meta_v1.ListOptions) (result *v1.AwsSqsQueueList, err error) {
	result = &v1.AwsSqsQueueList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awssqsqueues").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsSqsQueues.
func (c *awsSqsQueues) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awssqsqueues").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsSqsQueue and creates it.  Returns the server's representation of the awsSqsQueue, and an error, if there is any.
func (c *awsSqsQueues) Create(awsSqsQueue *v1.AwsSqsQueue) (result *v1.AwsSqsQueue, err error) {
	result = &v1.AwsSqsQueue{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awssqsqueues").
		Body(awsSqsQueue).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsSqsQueue and updates it. Returns the server's representation of the awsSqsQueue, and an error, if there is any.
func (c *awsSqsQueues) Update(awsSqsQueue *v1.AwsSqsQueue) (result *v1.AwsSqsQueue, err error) {
	result = &v1.AwsSqsQueue{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awssqsqueues").
		Name(awsSqsQueue.Name).
		Body(awsSqsQueue).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsSqsQueue and deletes it. Returns an error if one occurs.
func (c *awsSqsQueues) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awssqsqueues").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsSqsQueues) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awssqsqueues").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsSqsQueue.
func (c *awsSqsQueues) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsSqsQueue, err error) {
	result = &v1.AwsSqsQueue{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awssqsqueues").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
