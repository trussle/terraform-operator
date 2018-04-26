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

// FakeAwsInstances implements AwsInstanceInterface
type FakeAwsInstances struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsinstancesResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsinstances"}

var awsinstancesKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsInstance"}

// Get takes name of the awsInstance, and returns the corresponding awsInstance object, and an error if there is any.
func (c *FakeAwsInstances) Get(name string, options v1.GetOptions) (result *aws_v1.AwsInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsinstancesResource, c.ns, name), &aws_v1.AwsInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsInstance), err
}

// List takes label and field selectors, and returns the list of AwsInstances that match those selectors.
func (c *FakeAwsInstances) List(opts v1.ListOptions) (result *aws_v1.AwsInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsinstancesResource, awsinstancesKind, c.ns, opts), &aws_v1.AwsInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsInstanceList{}
	for _, item := range obj.(*aws_v1.AwsInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsInstances.
func (c *FakeAwsInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsinstancesResource, c.ns, opts))

}

// Create takes the representation of a awsInstance and creates it.  Returns the server's representation of the awsInstance, and an error, if there is any.
func (c *FakeAwsInstances) Create(awsInstance *aws_v1.AwsInstance) (result *aws_v1.AwsInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsinstancesResource, c.ns, awsInstance), &aws_v1.AwsInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsInstance), err
}

// Update takes the representation of a awsInstance and updates it. Returns the server's representation of the awsInstance, and an error, if there is any.
func (c *FakeAwsInstances) Update(awsInstance *aws_v1.AwsInstance) (result *aws_v1.AwsInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsinstancesResource, c.ns, awsInstance), &aws_v1.AwsInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsInstance), err
}

// Delete takes name of the awsInstance and deletes it. Returns an error if one occurs.
func (c *FakeAwsInstances) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsinstancesResource, c.ns, name), &aws_v1.AwsInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsinstancesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsInstanceList{})
	return err
}

// Patch applies the patch and returns the patched awsInstance.
func (c *FakeAwsInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsinstancesResource, c.ns, name, data, subresources...), &aws_v1.AwsInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsInstance), err
}
