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

// AwsNetworkAclsGetter has a method to return a AwsNetworkAclInterface.
// A group's client should implement this interface.
type AwsNetworkAclsGetter interface {
	AwsNetworkAcls(namespace string) AwsNetworkAclInterface
}

// AwsNetworkAclInterface has methods to work with AwsNetworkAcl resources.
type AwsNetworkAclInterface interface {
	Create(*v1.AwsNetworkAcl) (*v1.AwsNetworkAcl, error)
	Update(*v1.AwsNetworkAcl) (*v1.AwsNetworkAcl, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsNetworkAcl, error)
	List(opts meta_v1.ListOptions) (*v1.AwsNetworkAclList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsNetworkAcl, err error)
	AwsNetworkAclExpansion
}

// awsNetworkAcls implements AwsNetworkAclInterface
type awsNetworkAcls struct {
	client rest.Interface
	ns     string
}

// newAwsNetworkAcls returns a AwsNetworkAcls
func newAwsNetworkAcls(c *TrussleV1Client, namespace string) *awsNetworkAcls {
	return &awsNetworkAcls{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsNetworkAcl, and returns the corresponding awsNetworkAcl object, and an error if there is any.
func (c *awsNetworkAcls) Get(name string, options meta_v1.GetOptions) (result *v1.AwsNetworkAcl, err error) {
	result = &v1.AwsNetworkAcl{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsnetworkacls").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsNetworkAcls that match those selectors.
func (c *awsNetworkAcls) List(opts meta_v1.ListOptions) (result *v1.AwsNetworkAclList, err error) {
	result = &v1.AwsNetworkAclList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsnetworkacls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsNetworkAcls.
func (c *awsNetworkAcls) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsnetworkacls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsNetworkAcl and creates it.  Returns the server's representation of the awsNetworkAcl, and an error, if there is any.
func (c *awsNetworkAcls) Create(awsNetworkAcl *v1.AwsNetworkAcl) (result *v1.AwsNetworkAcl, err error) {
	result = &v1.AwsNetworkAcl{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsnetworkacls").
		Body(awsNetworkAcl).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsNetworkAcl and updates it. Returns the server's representation of the awsNetworkAcl, and an error, if there is any.
func (c *awsNetworkAcls) Update(awsNetworkAcl *v1.AwsNetworkAcl) (result *v1.AwsNetworkAcl, err error) {
	result = &v1.AwsNetworkAcl{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsnetworkacls").
		Name(awsNetworkAcl.Name).
		Body(awsNetworkAcl).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsNetworkAcl and deletes it. Returns an error if one occurs.
func (c *awsNetworkAcls) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsnetworkacls").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsNetworkAcls) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsnetworkacls").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsNetworkAcl.
func (c *awsNetworkAcls) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsNetworkAcl, err error) {
	result = &v1.AwsNetworkAcl{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsnetworkacls").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}