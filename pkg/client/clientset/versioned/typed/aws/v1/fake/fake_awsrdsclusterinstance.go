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

// FakeAwsRdsClusterInstances implements AwsRdsClusterInstanceInterface
type FakeAwsRdsClusterInstances struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsrdsclusterinstancesResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsrdsclusterinstances"}

var awsrdsclusterinstancesKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsRdsClusterInstance"}

// Get takes name of the awsRdsClusterInstance, and returns the corresponding awsRdsClusterInstance object, and an error if there is any.
func (c *FakeAwsRdsClusterInstances) Get(name string, options v1.GetOptions) (result *aws_v1.AwsRdsClusterInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsrdsclusterinstancesResource, c.ns, name), &aws_v1.AwsRdsClusterInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsRdsClusterInstance), err
}

// List takes label and field selectors, and returns the list of AwsRdsClusterInstances that match those selectors.
func (c *FakeAwsRdsClusterInstances) List(opts v1.ListOptions) (result *aws_v1.AwsRdsClusterInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsrdsclusterinstancesResource, awsrdsclusterinstancesKind, c.ns, opts), &aws_v1.AwsRdsClusterInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsRdsClusterInstanceList{}
	for _, item := range obj.(*aws_v1.AwsRdsClusterInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsRdsClusterInstances.
func (c *FakeAwsRdsClusterInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsrdsclusterinstancesResource, c.ns, opts))

}

// Create takes the representation of a awsRdsClusterInstance and creates it.  Returns the server's representation of the awsRdsClusterInstance, and an error, if there is any.
func (c *FakeAwsRdsClusterInstances) Create(awsRdsClusterInstance *aws_v1.AwsRdsClusterInstance) (result *aws_v1.AwsRdsClusterInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsrdsclusterinstancesResource, c.ns, awsRdsClusterInstance), &aws_v1.AwsRdsClusterInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsRdsClusterInstance), err
}

// Update takes the representation of a awsRdsClusterInstance and updates it. Returns the server's representation of the awsRdsClusterInstance, and an error, if there is any.
func (c *FakeAwsRdsClusterInstances) Update(awsRdsClusterInstance *aws_v1.AwsRdsClusterInstance) (result *aws_v1.AwsRdsClusterInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsrdsclusterinstancesResource, c.ns, awsRdsClusterInstance), &aws_v1.AwsRdsClusterInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsRdsClusterInstance), err
}

// Delete takes name of the awsRdsClusterInstance and deletes it. Returns an error if one occurs.
func (c *FakeAwsRdsClusterInstances) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsrdsclusterinstancesResource, c.ns, name), &aws_v1.AwsRdsClusterInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsRdsClusterInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsrdsclusterinstancesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsRdsClusterInstanceList{})
	return err
}

// Patch applies the patch and returns the patched awsRdsClusterInstance.
func (c *FakeAwsRdsClusterInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsRdsClusterInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsrdsclusterinstancesResource, c.ns, name, data, subresources...), &aws_v1.AwsRdsClusterInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsRdsClusterInstance), err
}