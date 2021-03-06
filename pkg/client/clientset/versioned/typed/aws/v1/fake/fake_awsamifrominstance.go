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

// FakeAwsAmiFromInstances implements AwsAmiFromInstanceInterface
type FakeAwsAmiFromInstances struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsamifrominstancesResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsamifrominstances"}

var awsamifrominstancesKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsAmiFromInstance"}

// Get takes name of the awsAmiFromInstance, and returns the corresponding awsAmiFromInstance object, and an error if there is any.
func (c *FakeAwsAmiFromInstances) Get(name string, options v1.GetOptions) (result *aws_v1.AwsAmiFromInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsamifrominstancesResource, c.ns, name), &aws_v1.AwsAmiFromInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsAmiFromInstance), err
}

// List takes label and field selectors, and returns the list of AwsAmiFromInstances that match those selectors.
func (c *FakeAwsAmiFromInstances) List(opts v1.ListOptions) (result *aws_v1.AwsAmiFromInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsamifrominstancesResource, awsamifrominstancesKind, c.ns, opts), &aws_v1.AwsAmiFromInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsAmiFromInstanceList{}
	for _, item := range obj.(*aws_v1.AwsAmiFromInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsAmiFromInstances.
func (c *FakeAwsAmiFromInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsamifrominstancesResource, c.ns, opts))

}

// Create takes the representation of a awsAmiFromInstance and creates it.  Returns the server's representation of the awsAmiFromInstance, and an error, if there is any.
func (c *FakeAwsAmiFromInstances) Create(awsAmiFromInstance *aws_v1.AwsAmiFromInstance) (result *aws_v1.AwsAmiFromInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsamifrominstancesResource, c.ns, awsAmiFromInstance), &aws_v1.AwsAmiFromInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsAmiFromInstance), err
}

// Update takes the representation of a awsAmiFromInstance and updates it. Returns the server's representation of the awsAmiFromInstance, and an error, if there is any.
func (c *FakeAwsAmiFromInstances) Update(awsAmiFromInstance *aws_v1.AwsAmiFromInstance) (result *aws_v1.AwsAmiFromInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsamifrominstancesResource, c.ns, awsAmiFromInstance), &aws_v1.AwsAmiFromInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsAmiFromInstance), err
}

// Delete takes name of the awsAmiFromInstance and deletes it. Returns an error if one occurs.
func (c *FakeAwsAmiFromInstances) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsamifrominstancesResource, c.ns, name), &aws_v1.AwsAmiFromInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsAmiFromInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsamifrominstancesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsAmiFromInstanceList{})
	return err
}

// Patch applies the patch and returns the patched awsAmiFromInstance.
func (c *FakeAwsAmiFromInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsAmiFromInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsamifrominstancesResource, c.ns, name, data, subresources...), &aws_v1.AwsAmiFromInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsAmiFromInstance), err
}
