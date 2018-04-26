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

// FakeAwsDynamodbGlobalTables implements AwsDynamodbGlobalTableInterface
type FakeAwsDynamodbGlobalTables struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsdynamodbglobaltablesResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsdynamodbglobaltables"}

var awsdynamodbglobaltablesKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsDynamodbGlobalTable"}

// Get takes name of the awsDynamodbGlobalTable, and returns the corresponding awsDynamodbGlobalTable object, and an error if there is any.
func (c *FakeAwsDynamodbGlobalTables) Get(name string, options v1.GetOptions) (result *aws_v1.AwsDynamodbGlobalTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsdynamodbglobaltablesResource, c.ns, name), &aws_v1.AwsDynamodbGlobalTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsDynamodbGlobalTable), err
}

// List takes label and field selectors, and returns the list of AwsDynamodbGlobalTables that match those selectors.
func (c *FakeAwsDynamodbGlobalTables) List(opts v1.ListOptions) (result *aws_v1.AwsDynamodbGlobalTableList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsdynamodbglobaltablesResource, awsdynamodbglobaltablesKind, c.ns, opts), &aws_v1.AwsDynamodbGlobalTableList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsDynamodbGlobalTableList{}
	for _, item := range obj.(*aws_v1.AwsDynamodbGlobalTableList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsDynamodbGlobalTables.
func (c *FakeAwsDynamodbGlobalTables) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsdynamodbglobaltablesResource, c.ns, opts))

}

// Create takes the representation of a awsDynamodbGlobalTable and creates it.  Returns the server's representation of the awsDynamodbGlobalTable, and an error, if there is any.
func (c *FakeAwsDynamodbGlobalTables) Create(awsDynamodbGlobalTable *aws_v1.AwsDynamodbGlobalTable) (result *aws_v1.AwsDynamodbGlobalTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsdynamodbglobaltablesResource, c.ns, awsDynamodbGlobalTable), &aws_v1.AwsDynamodbGlobalTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsDynamodbGlobalTable), err
}

// Update takes the representation of a awsDynamodbGlobalTable and updates it. Returns the server's representation of the awsDynamodbGlobalTable, and an error, if there is any.
func (c *FakeAwsDynamodbGlobalTables) Update(awsDynamodbGlobalTable *aws_v1.AwsDynamodbGlobalTable) (result *aws_v1.AwsDynamodbGlobalTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsdynamodbglobaltablesResource, c.ns, awsDynamodbGlobalTable), &aws_v1.AwsDynamodbGlobalTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsDynamodbGlobalTable), err
}

// Delete takes name of the awsDynamodbGlobalTable and deletes it. Returns an error if one occurs.
func (c *FakeAwsDynamodbGlobalTables) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsdynamodbglobaltablesResource, c.ns, name), &aws_v1.AwsDynamodbGlobalTable{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsDynamodbGlobalTables) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsdynamodbglobaltablesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsDynamodbGlobalTableList{})
	return err
}

// Patch applies the patch and returns the patched awsDynamodbGlobalTable.
func (c *FakeAwsDynamodbGlobalTables) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsDynamodbGlobalTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsdynamodbglobaltablesResource, c.ns, name, data, subresources...), &aws_v1.AwsDynamodbGlobalTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsDynamodbGlobalTable), err
}
