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

// AwsLbTargetGroupAttachmentsGetter has a method to return a AwsLbTargetGroupAttachmentInterface.
// A group's client should implement this interface.
type AwsLbTargetGroupAttachmentsGetter interface {
	AwsLbTargetGroupAttachments(namespace string) AwsLbTargetGroupAttachmentInterface
}

// AwsLbTargetGroupAttachmentInterface has methods to work with AwsLbTargetGroupAttachment resources.
type AwsLbTargetGroupAttachmentInterface interface {
	Create(*v1.AwsLbTargetGroupAttachment) (*v1.AwsLbTargetGroupAttachment, error)
	Update(*v1.AwsLbTargetGroupAttachment) (*v1.AwsLbTargetGroupAttachment, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsLbTargetGroupAttachment, error)
	List(opts meta_v1.ListOptions) (*v1.AwsLbTargetGroupAttachmentList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsLbTargetGroupAttachment, err error)
	AwsLbTargetGroupAttachmentExpansion
}

// awsLbTargetGroupAttachments implements AwsLbTargetGroupAttachmentInterface
type awsLbTargetGroupAttachments struct {
	client rest.Interface
	ns     string
}

// newAwsLbTargetGroupAttachments returns a AwsLbTargetGroupAttachments
func newAwsLbTargetGroupAttachments(c *TrussleV1Client, namespace string) *awsLbTargetGroupAttachments {
	return &awsLbTargetGroupAttachments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsLbTargetGroupAttachment, and returns the corresponding awsLbTargetGroupAttachment object, and an error if there is any.
func (c *awsLbTargetGroupAttachments) Get(name string, options meta_v1.GetOptions) (result *v1.AwsLbTargetGroupAttachment, err error) {
	result = &v1.AwsLbTargetGroupAttachment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awslbtargetgroupattachments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsLbTargetGroupAttachments that match those selectors.
func (c *awsLbTargetGroupAttachments) List(opts meta_v1.ListOptions) (result *v1.AwsLbTargetGroupAttachmentList, err error) {
	result = &v1.AwsLbTargetGroupAttachmentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awslbtargetgroupattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsLbTargetGroupAttachments.
func (c *awsLbTargetGroupAttachments) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awslbtargetgroupattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsLbTargetGroupAttachment and creates it.  Returns the server's representation of the awsLbTargetGroupAttachment, and an error, if there is any.
func (c *awsLbTargetGroupAttachments) Create(awsLbTargetGroupAttachment *v1.AwsLbTargetGroupAttachment) (result *v1.AwsLbTargetGroupAttachment, err error) {
	result = &v1.AwsLbTargetGroupAttachment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awslbtargetgroupattachments").
		Body(awsLbTargetGroupAttachment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsLbTargetGroupAttachment and updates it. Returns the server's representation of the awsLbTargetGroupAttachment, and an error, if there is any.
func (c *awsLbTargetGroupAttachments) Update(awsLbTargetGroupAttachment *v1.AwsLbTargetGroupAttachment) (result *v1.AwsLbTargetGroupAttachment, err error) {
	result = &v1.AwsLbTargetGroupAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awslbtargetgroupattachments").
		Name(awsLbTargetGroupAttachment.Name).
		Body(awsLbTargetGroupAttachment).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsLbTargetGroupAttachment and deletes it. Returns an error if one occurs.
func (c *awsLbTargetGroupAttachments) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awslbtargetgroupattachments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsLbTargetGroupAttachments) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awslbtargetgroupattachments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsLbTargetGroupAttachment.
func (c *awsLbTargetGroupAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsLbTargetGroupAttachment, err error) {
	result = &v1.AwsLbTargetGroupAttachment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awslbtargetgroupattachments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
