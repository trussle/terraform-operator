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

package fake

import (
	aws_v1 "github.com/trussle/terraform-operator/pkg/apis/aws/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAwsEcsClusters implements AwsEcsClusterInterface
type FakeAwsEcsClusters struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsecsclustersResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsecsclusters"}

var awsecsclustersKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsEcsCluster"}

// Get takes name of the awsEcsCluster, and returns the corresponding awsEcsCluster object, and an error if there is any.
func (c *FakeAwsEcsClusters) Get(name string, options v1.GetOptions) (result *aws_v1.AwsEcsCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsecsclustersResource, c.ns, name), &aws_v1.AwsEcsCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsEcsCluster), err
}

// List takes label and field selectors, and returns the list of AwsEcsClusters that match those selectors.
func (c *FakeAwsEcsClusters) List(opts v1.ListOptions) (result *aws_v1.AwsEcsClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsecsclustersResource, awsecsclustersKind, c.ns, opts), &aws_v1.AwsEcsClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsEcsClusterList{}
	for _, item := range obj.(*aws_v1.AwsEcsClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsEcsClusters.
func (c *FakeAwsEcsClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsecsclustersResource, c.ns, opts))

}

// Create takes the representation of a awsEcsCluster and creates it.  Returns the server's representation of the awsEcsCluster, and an error, if there is any.
func (c *FakeAwsEcsClusters) Create(awsEcsCluster *aws_v1.AwsEcsCluster) (result *aws_v1.AwsEcsCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsecsclustersResource, c.ns, awsEcsCluster), &aws_v1.AwsEcsCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsEcsCluster), err
}

// Update takes the representation of a awsEcsCluster and updates it. Returns the server's representation of the awsEcsCluster, and an error, if there is any.
func (c *FakeAwsEcsClusters) Update(awsEcsCluster *aws_v1.AwsEcsCluster) (result *aws_v1.AwsEcsCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsecsclustersResource, c.ns, awsEcsCluster), &aws_v1.AwsEcsCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsEcsCluster), err
}

// Delete takes name of the awsEcsCluster and deletes it. Returns an error if one occurs.
func (c *FakeAwsEcsClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsecsclustersResource, c.ns, name), &aws_v1.AwsEcsCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsEcsClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsecsclustersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsEcsClusterList{})
	return err
}

// Patch applies the patch and returns the patched awsEcsCluster.
func (c *FakeAwsEcsClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsEcsCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsecsclustersResource, c.ns, name, data, subresources...), &aws_v1.AwsEcsCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsEcsCluster), err
}
