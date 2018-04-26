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

// FakeAwsSsmActivations implements AwsSsmActivationInterface
type FakeAwsSsmActivations struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsssmactivationsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsssmactivations"}

var awsssmactivationsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsSsmActivation"}

// Get takes name of the awsSsmActivation, and returns the corresponding awsSsmActivation object, and an error if there is any.
func (c *FakeAwsSsmActivations) Get(name string, options v1.GetOptions) (result *aws_v1.AwsSsmActivation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsssmactivationsResource, c.ns, name), &aws_v1.AwsSsmActivation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSsmActivation), err
}

// List takes label and field selectors, and returns the list of AwsSsmActivations that match those selectors.
func (c *FakeAwsSsmActivations) List(opts v1.ListOptions) (result *aws_v1.AwsSsmActivationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsssmactivationsResource, awsssmactivationsKind, c.ns, opts), &aws_v1.AwsSsmActivationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsSsmActivationList{}
	for _, item := range obj.(*aws_v1.AwsSsmActivationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSsmActivations.
func (c *FakeAwsSsmActivations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsssmactivationsResource, c.ns, opts))

}

// Create takes the representation of a awsSsmActivation and creates it.  Returns the server's representation of the awsSsmActivation, and an error, if there is any.
func (c *FakeAwsSsmActivations) Create(awsSsmActivation *aws_v1.AwsSsmActivation) (result *aws_v1.AwsSsmActivation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsssmactivationsResource, c.ns, awsSsmActivation), &aws_v1.AwsSsmActivation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSsmActivation), err
}

// Update takes the representation of a awsSsmActivation and updates it. Returns the server's representation of the awsSsmActivation, and an error, if there is any.
func (c *FakeAwsSsmActivations) Update(awsSsmActivation *aws_v1.AwsSsmActivation) (result *aws_v1.AwsSsmActivation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsssmactivationsResource, c.ns, awsSsmActivation), &aws_v1.AwsSsmActivation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSsmActivation), err
}

// Delete takes name of the awsSsmActivation and deletes it. Returns an error if one occurs.
func (c *FakeAwsSsmActivations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsssmactivationsResource, c.ns, name), &aws_v1.AwsSsmActivation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSsmActivations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsssmactivationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsSsmActivationList{})
	return err
}

// Patch applies the patch and returns the patched awsSsmActivation.
func (c *FakeAwsSsmActivations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsSsmActivation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsssmactivationsResource, c.ns, name, data, subresources...), &aws_v1.AwsSsmActivation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSsmActivation), err
}
