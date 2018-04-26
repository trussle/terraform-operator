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

// FakeAwsOpsworksInstances implements AwsOpsworksInstanceInterface
type FakeAwsOpsworksInstances struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsopsworksinstancesResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsopsworksinstances"}

var awsopsworksinstancesKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsOpsworksInstance"}

// Get takes name of the awsOpsworksInstance, and returns the corresponding awsOpsworksInstance object, and an error if there is any.
func (c *FakeAwsOpsworksInstances) Get(name string, options v1.GetOptions) (result *aws_v1.AwsOpsworksInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsopsworksinstancesResource, c.ns, name), &aws_v1.AwsOpsworksInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsOpsworksInstance), err
}

// List takes label and field selectors, and returns the list of AwsOpsworksInstances that match those selectors.
func (c *FakeAwsOpsworksInstances) List(opts v1.ListOptions) (result *aws_v1.AwsOpsworksInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsopsworksinstancesResource, awsopsworksinstancesKind, c.ns, opts), &aws_v1.AwsOpsworksInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsOpsworksInstanceList{}
	for _, item := range obj.(*aws_v1.AwsOpsworksInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsOpsworksInstances.
func (c *FakeAwsOpsworksInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsopsworksinstancesResource, c.ns, opts))

}

// Create takes the representation of a awsOpsworksInstance and creates it.  Returns the server's representation of the awsOpsworksInstance, and an error, if there is any.
func (c *FakeAwsOpsworksInstances) Create(awsOpsworksInstance *aws_v1.AwsOpsworksInstance) (result *aws_v1.AwsOpsworksInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsopsworksinstancesResource, c.ns, awsOpsworksInstance), &aws_v1.AwsOpsworksInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsOpsworksInstance), err
}

// Update takes the representation of a awsOpsworksInstance and updates it. Returns the server's representation of the awsOpsworksInstance, and an error, if there is any.
func (c *FakeAwsOpsworksInstances) Update(awsOpsworksInstance *aws_v1.AwsOpsworksInstance) (result *aws_v1.AwsOpsworksInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsopsworksinstancesResource, c.ns, awsOpsworksInstance), &aws_v1.AwsOpsworksInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsOpsworksInstance), err
}

// Delete takes name of the awsOpsworksInstance and deletes it. Returns an error if one occurs.
func (c *FakeAwsOpsworksInstances) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsopsworksinstancesResource, c.ns, name), &aws_v1.AwsOpsworksInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsOpsworksInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsopsworksinstancesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsOpsworksInstanceList{})
	return err
}

// Patch applies the patch and returns the patched awsOpsworksInstance.
func (c *FakeAwsOpsworksInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsOpsworksInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsopsworksinstancesResource, c.ns, name, data, subresources...), &aws_v1.AwsOpsworksInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsOpsworksInstance), err
}
