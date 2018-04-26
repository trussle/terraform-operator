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

// FakeAwsSesReceiptRuleSets implements AwsSesReceiptRuleSetInterface
type FakeAwsSesReceiptRuleSets struct {
	Fake *FakeTrussleV1
	ns   string
}

var awssesreceiptrulesetsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awssesreceiptrulesets"}

var awssesreceiptrulesetsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsSesReceiptRuleSet"}

// Get takes name of the awsSesReceiptRuleSet, and returns the corresponding awsSesReceiptRuleSet object, and an error if there is any.
func (c *FakeAwsSesReceiptRuleSets) Get(name string, options v1.GetOptions) (result *aws_v1.AwsSesReceiptRuleSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awssesreceiptrulesetsResource, c.ns, name), &aws_v1.AwsSesReceiptRuleSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSesReceiptRuleSet), err
}

// List takes label and field selectors, and returns the list of AwsSesReceiptRuleSets that match those selectors.
func (c *FakeAwsSesReceiptRuleSets) List(opts v1.ListOptions) (result *aws_v1.AwsSesReceiptRuleSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awssesreceiptrulesetsResource, awssesreceiptrulesetsKind, c.ns, opts), &aws_v1.AwsSesReceiptRuleSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsSesReceiptRuleSetList{}
	for _, item := range obj.(*aws_v1.AwsSesReceiptRuleSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSesReceiptRuleSets.
func (c *FakeAwsSesReceiptRuleSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awssesreceiptrulesetsResource, c.ns, opts))

}

// Create takes the representation of a awsSesReceiptRuleSet and creates it.  Returns the server's representation of the awsSesReceiptRuleSet, and an error, if there is any.
func (c *FakeAwsSesReceiptRuleSets) Create(awsSesReceiptRuleSet *aws_v1.AwsSesReceiptRuleSet) (result *aws_v1.AwsSesReceiptRuleSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awssesreceiptrulesetsResource, c.ns, awsSesReceiptRuleSet), &aws_v1.AwsSesReceiptRuleSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSesReceiptRuleSet), err
}

// Update takes the representation of a awsSesReceiptRuleSet and updates it. Returns the server's representation of the awsSesReceiptRuleSet, and an error, if there is any.
func (c *FakeAwsSesReceiptRuleSets) Update(awsSesReceiptRuleSet *aws_v1.AwsSesReceiptRuleSet) (result *aws_v1.AwsSesReceiptRuleSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awssesreceiptrulesetsResource, c.ns, awsSesReceiptRuleSet), &aws_v1.AwsSesReceiptRuleSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSesReceiptRuleSet), err
}

// Delete takes name of the awsSesReceiptRuleSet and deletes it. Returns an error if one occurs.
func (c *FakeAwsSesReceiptRuleSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awssesreceiptrulesetsResource, c.ns, name), &aws_v1.AwsSesReceiptRuleSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSesReceiptRuleSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awssesreceiptrulesetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsSesReceiptRuleSetList{})
	return err
}

// Patch applies the patch and returns the patched awsSesReceiptRuleSet.
func (c *FakeAwsSesReceiptRuleSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsSesReceiptRuleSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awssesreceiptrulesetsResource, c.ns, name, data, subresources...), &aws_v1.AwsSesReceiptRuleSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSesReceiptRuleSet), err
}
