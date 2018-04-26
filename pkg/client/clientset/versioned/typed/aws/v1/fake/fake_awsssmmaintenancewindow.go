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

// FakeAwsSsmMaintenanceWindows implements AwsSsmMaintenanceWindowInterface
type FakeAwsSsmMaintenanceWindows struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsssmmaintenancewindowsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsssmmaintenancewindows"}

var awsssmmaintenancewindowsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsSsmMaintenanceWindow"}

// Get takes name of the awsSsmMaintenanceWindow, and returns the corresponding awsSsmMaintenanceWindow object, and an error if there is any.
func (c *FakeAwsSsmMaintenanceWindows) Get(name string, options v1.GetOptions) (result *aws_v1.AwsSsmMaintenanceWindow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsssmmaintenancewindowsResource, c.ns, name), &aws_v1.AwsSsmMaintenanceWindow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSsmMaintenanceWindow), err
}

// List takes label and field selectors, and returns the list of AwsSsmMaintenanceWindows that match those selectors.
func (c *FakeAwsSsmMaintenanceWindows) List(opts v1.ListOptions) (result *aws_v1.AwsSsmMaintenanceWindowList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsssmmaintenancewindowsResource, awsssmmaintenancewindowsKind, c.ns, opts), &aws_v1.AwsSsmMaintenanceWindowList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsSsmMaintenanceWindowList{}
	for _, item := range obj.(*aws_v1.AwsSsmMaintenanceWindowList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSsmMaintenanceWindows.
func (c *FakeAwsSsmMaintenanceWindows) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsssmmaintenancewindowsResource, c.ns, opts))

}

// Create takes the representation of a awsSsmMaintenanceWindow and creates it.  Returns the server's representation of the awsSsmMaintenanceWindow, and an error, if there is any.
func (c *FakeAwsSsmMaintenanceWindows) Create(awsSsmMaintenanceWindow *aws_v1.AwsSsmMaintenanceWindow) (result *aws_v1.AwsSsmMaintenanceWindow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsssmmaintenancewindowsResource, c.ns, awsSsmMaintenanceWindow), &aws_v1.AwsSsmMaintenanceWindow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSsmMaintenanceWindow), err
}

// Update takes the representation of a awsSsmMaintenanceWindow and updates it. Returns the server's representation of the awsSsmMaintenanceWindow, and an error, if there is any.
func (c *FakeAwsSsmMaintenanceWindows) Update(awsSsmMaintenanceWindow *aws_v1.AwsSsmMaintenanceWindow) (result *aws_v1.AwsSsmMaintenanceWindow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsssmmaintenancewindowsResource, c.ns, awsSsmMaintenanceWindow), &aws_v1.AwsSsmMaintenanceWindow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSsmMaintenanceWindow), err
}

// Delete takes name of the awsSsmMaintenanceWindow and deletes it. Returns an error if one occurs.
func (c *FakeAwsSsmMaintenanceWindows) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsssmmaintenancewindowsResource, c.ns, name), &aws_v1.AwsSsmMaintenanceWindow{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSsmMaintenanceWindows) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsssmmaintenancewindowsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsSsmMaintenanceWindowList{})
	return err
}

// Patch applies the patch and returns the patched awsSsmMaintenanceWindow.
func (c *FakeAwsSsmMaintenanceWindows) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsSsmMaintenanceWindow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsssmmaintenancewindowsResource, c.ns, name, data, subresources...), &aws_v1.AwsSsmMaintenanceWindow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSsmMaintenanceWindow), err
}
