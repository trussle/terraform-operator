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

// FakeAwsAlbListeners implements AwsAlbListenerInterface
type FakeAwsAlbListeners struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsalblistenersResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsalblisteners"}

var awsalblistenersKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsAlbListener"}

// Get takes name of the awsAlbListener, and returns the corresponding awsAlbListener object, and an error if there is any.
func (c *FakeAwsAlbListeners) Get(name string, options v1.GetOptions) (result *aws_v1.AwsAlbListener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsalblistenersResource, c.ns, name), &aws_v1.AwsAlbListener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsAlbListener), err
}

// List takes label and field selectors, and returns the list of AwsAlbListeners that match those selectors.
func (c *FakeAwsAlbListeners) List(opts v1.ListOptions) (result *aws_v1.AwsAlbListenerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsalblistenersResource, awsalblistenersKind, c.ns, opts), &aws_v1.AwsAlbListenerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsAlbListenerList{}
	for _, item := range obj.(*aws_v1.AwsAlbListenerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsAlbListeners.
func (c *FakeAwsAlbListeners) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsalblistenersResource, c.ns, opts))

}

// Create takes the representation of a awsAlbListener and creates it.  Returns the server's representation of the awsAlbListener, and an error, if there is any.
func (c *FakeAwsAlbListeners) Create(awsAlbListener *aws_v1.AwsAlbListener) (result *aws_v1.AwsAlbListener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsalblistenersResource, c.ns, awsAlbListener), &aws_v1.AwsAlbListener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsAlbListener), err
}

// Update takes the representation of a awsAlbListener and updates it. Returns the server's representation of the awsAlbListener, and an error, if there is any.
func (c *FakeAwsAlbListeners) Update(awsAlbListener *aws_v1.AwsAlbListener) (result *aws_v1.AwsAlbListener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsalblistenersResource, c.ns, awsAlbListener), &aws_v1.AwsAlbListener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsAlbListener), err
}

// Delete takes name of the awsAlbListener and deletes it. Returns an error if one occurs.
func (c *FakeAwsAlbListeners) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsalblistenersResource, c.ns, name), &aws_v1.AwsAlbListener{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsAlbListeners) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsalblistenersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsAlbListenerList{})
	return err
}

// Patch applies the patch and returns the patched awsAlbListener.
func (c *FakeAwsAlbListeners) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsAlbListener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsalblistenersResource, c.ns, name, data, subresources...), &aws_v1.AwsAlbListener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsAlbListener), err
}
