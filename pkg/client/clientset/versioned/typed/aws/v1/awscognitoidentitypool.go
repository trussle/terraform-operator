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

// AwsCognitoIdentityPoolsGetter has a method to return a AwsCognitoIdentityPoolInterface.
// A group's client should implement this interface.
type AwsCognitoIdentityPoolsGetter interface {
	AwsCognitoIdentityPools(namespace string) AwsCognitoIdentityPoolInterface
}

// AwsCognitoIdentityPoolInterface has methods to work with AwsCognitoIdentityPool resources.
type AwsCognitoIdentityPoolInterface interface {
	Create(*v1.AwsCognitoIdentityPool) (*v1.AwsCognitoIdentityPool, error)
	Update(*v1.AwsCognitoIdentityPool) (*v1.AwsCognitoIdentityPool, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsCognitoIdentityPool, error)
	List(opts meta_v1.ListOptions) (*v1.AwsCognitoIdentityPoolList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsCognitoIdentityPool, err error)
	AwsCognitoIdentityPoolExpansion
}

// awsCognitoIdentityPools implements AwsCognitoIdentityPoolInterface
type awsCognitoIdentityPools struct {
	client rest.Interface
	ns     string
}

// newAwsCognitoIdentityPools returns a AwsCognitoIdentityPools
func newAwsCognitoIdentityPools(c *TrussleV1Client, namespace string) *awsCognitoIdentityPools {
	return &awsCognitoIdentityPools{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsCognitoIdentityPool, and returns the corresponding awsCognitoIdentityPool object, and an error if there is any.
func (c *awsCognitoIdentityPools) Get(name string, options meta_v1.GetOptions) (result *v1.AwsCognitoIdentityPool, err error) {
	result = &v1.AwsCognitoIdentityPool{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awscognitoidentitypools").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsCognitoIdentityPools that match those selectors.
func (c *awsCognitoIdentityPools) List(opts meta_v1.ListOptions) (result *v1.AwsCognitoIdentityPoolList, err error) {
	result = &v1.AwsCognitoIdentityPoolList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awscognitoidentitypools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsCognitoIdentityPools.
func (c *awsCognitoIdentityPools) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awscognitoidentitypools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsCognitoIdentityPool and creates it.  Returns the server's representation of the awsCognitoIdentityPool, and an error, if there is any.
func (c *awsCognitoIdentityPools) Create(awsCognitoIdentityPool *v1.AwsCognitoIdentityPool) (result *v1.AwsCognitoIdentityPool, err error) {
	result = &v1.AwsCognitoIdentityPool{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awscognitoidentitypools").
		Body(awsCognitoIdentityPool).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsCognitoIdentityPool and updates it. Returns the server's representation of the awsCognitoIdentityPool, and an error, if there is any.
func (c *awsCognitoIdentityPools) Update(awsCognitoIdentityPool *v1.AwsCognitoIdentityPool) (result *v1.AwsCognitoIdentityPool, err error) {
	result = &v1.AwsCognitoIdentityPool{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awscognitoidentitypools").
		Name(awsCognitoIdentityPool.Name).
		Body(awsCognitoIdentityPool).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsCognitoIdentityPool and deletes it. Returns an error if one occurs.
func (c *awsCognitoIdentityPools) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awscognitoidentitypools").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsCognitoIdentityPools) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awscognitoidentitypools").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsCognitoIdentityPool.
func (c *awsCognitoIdentityPools) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsCognitoIdentityPool, err error) {
	result = &v1.AwsCognitoIdentityPool{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awscognitoidentitypools").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
