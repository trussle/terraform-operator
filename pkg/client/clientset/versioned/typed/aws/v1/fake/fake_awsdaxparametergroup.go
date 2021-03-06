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

// FakeAwsDaxParameterGroups implements AwsDaxParameterGroupInterface
type FakeAwsDaxParameterGroups struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsdaxparametergroupsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsdaxparametergroups"}

var awsdaxparametergroupsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsDaxParameterGroup"}

// Get takes name of the awsDaxParameterGroup, and returns the corresponding awsDaxParameterGroup object, and an error if there is any.
func (c *FakeAwsDaxParameterGroups) Get(name string, options v1.GetOptions) (result *aws_v1.AwsDaxParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsdaxparametergroupsResource, c.ns, name), &aws_v1.AwsDaxParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsDaxParameterGroup), err
}

// List takes label and field selectors, and returns the list of AwsDaxParameterGroups that match those selectors.
func (c *FakeAwsDaxParameterGroups) List(opts v1.ListOptions) (result *aws_v1.AwsDaxParameterGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsdaxparametergroupsResource, awsdaxparametergroupsKind, c.ns, opts), &aws_v1.AwsDaxParameterGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsDaxParameterGroupList{}
	for _, item := range obj.(*aws_v1.AwsDaxParameterGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsDaxParameterGroups.
func (c *FakeAwsDaxParameterGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsdaxparametergroupsResource, c.ns, opts))

}

// Create takes the representation of a awsDaxParameterGroup and creates it.  Returns the server's representation of the awsDaxParameterGroup, and an error, if there is any.
func (c *FakeAwsDaxParameterGroups) Create(awsDaxParameterGroup *aws_v1.AwsDaxParameterGroup) (result *aws_v1.AwsDaxParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsdaxparametergroupsResource, c.ns, awsDaxParameterGroup), &aws_v1.AwsDaxParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsDaxParameterGroup), err
}

// Update takes the representation of a awsDaxParameterGroup and updates it. Returns the server's representation of the awsDaxParameterGroup, and an error, if there is any.
func (c *FakeAwsDaxParameterGroups) Update(awsDaxParameterGroup *aws_v1.AwsDaxParameterGroup) (result *aws_v1.AwsDaxParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsdaxparametergroupsResource, c.ns, awsDaxParameterGroup), &aws_v1.AwsDaxParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsDaxParameterGroup), err
}

// Delete takes name of the awsDaxParameterGroup and deletes it. Returns an error if one occurs.
func (c *FakeAwsDaxParameterGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsdaxparametergroupsResource, c.ns, name), &aws_v1.AwsDaxParameterGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsDaxParameterGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsdaxparametergroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsDaxParameterGroupList{})
	return err
}

// Patch applies the patch and returns the patched awsDaxParameterGroup.
func (c *FakeAwsDaxParameterGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsDaxParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsdaxparametergroupsResource, c.ns, name, data, subresources...), &aws_v1.AwsDaxParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsDaxParameterGroup), err
}
