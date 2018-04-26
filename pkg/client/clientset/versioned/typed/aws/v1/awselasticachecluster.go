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

// AwsElasticacheClustersGetter has a method to return a AwsElasticacheClusterInterface.
// A group's client should implement this interface.
type AwsElasticacheClustersGetter interface {
	AwsElasticacheClusters(namespace string) AwsElasticacheClusterInterface
}

// AwsElasticacheClusterInterface has methods to work with AwsElasticacheCluster resources.
type AwsElasticacheClusterInterface interface {
	Create(*v1.AwsElasticacheCluster) (*v1.AwsElasticacheCluster, error)
	Update(*v1.AwsElasticacheCluster) (*v1.AwsElasticacheCluster, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsElasticacheCluster, error)
	List(opts meta_v1.ListOptions) (*v1.AwsElasticacheClusterList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsElasticacheCluster, err error)
	AwsElasticacheClusterExpansion
}

// awsElasticacheClusters implements AwsElasticacheClusterInterface
type awsElasticacheClusters struct {
	client rest.Interface
	ns     string
}

// newAwsElasticacheClusters returns a AwsElasticacheClusters
func newAwsElasticacheClusters(c *TrussleV1Client, namespace string) *awsElasticacheClusters {
	return &awsElasticacheClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsElasticacheCluster, and returns the corresponding awsElasticacheCluster object, and an error if there is any.
func (c *awsElasticacheClusters) Get(name string, options meta_v1.GetOptions) (result *v1.AwsElasticacheCluster, err error) {
	result = &v1.AwsElasticacheCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awselasticacheclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsElasticacheClusters that match those selectors.
func (c *awsElasticacheClusters) List(opts meta_v1.ListOptions) (result *v1.AwsElasticacheClusterList, err error) {
	result = &v1.AwsElasticacheClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awselasticacheclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsElasticacheClusters.
func (c *awsElasticacheClusters) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awselasticacheclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsElasticacheCluster and creates it.  Returns the server's representation of the awsElasticacheCluster, and an error, if there is any.
func (c *awsElasticacheClusters) Create(awsElasticacheCluster *v1.AwsElasticacheCluster) (result *v1.AwsElasticacheCluster, err error) {
	result = &v1.AwsElasticacheCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awselasticacheclusters").
		Body(awsElasticacheCluster).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsElasticacheCluster and updates it. Returns the server's representation of the awsElasticacheCluster, and an error, if there is any.
func (c *awsElasticacheClusters) Update(awsElasticacheCluster *v1.AwsElasticacheCluster) (result *v1.AwsElasticacheCluster, err error) {
	result = &v1.AwsElasticacheCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awselasticacheclusters").
		Name(awsElasticacheCluster.Name).
		Body(awsElasticacheCluster).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsElasticacheCluster and deletes it. Returns an error if one occurs.
func (c *awsElasticacheClusters) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awselasticacheclusters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsElasticacheClusters) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awselasticacheclusters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsElasticacheCluster.
func (c *awsElasticacheClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsElasticacheCluster, err error) {
	result = &v1.AwsElasticacheCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awselasticacheclusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
