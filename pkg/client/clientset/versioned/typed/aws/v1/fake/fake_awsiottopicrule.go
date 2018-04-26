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

// FakeAwsIotTopicRules implements AwsIotTopicRuleInterface
type FakeAwsIotTopicRules struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsiottopicrulesResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsiottopicrules"}

var awsiottopicrulesKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsIotTopicRule"}

// Get takes name of the awsIotTopicRule, and returns the corresponding awsIotTopicRule object, and an error if there is any.
func (c *FakeAwsIotTopicRules) Get(name string, options v1.GetOptions) (result *aws_v1.AwsIotTopicRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsiottopicrulesResource, c.ns, name), &aws_v1.AwsIotTopicRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsIotTopicRule), err
}

// List takes label and field selectors, and returns the list of AwsIotTopicRules that match those selectors.
func (c *FakeAwsIotTopicRules) List(opts v1.ListOptions) (result *aws_v1.AwsIotTopicRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsiottopicrulesResource, awsiottopicrulesKind, c.ns, opts), &aws_v1.AwsIotTopicRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsIotTopicRuleList{}
	for _, item := range obj.(*aws_v1.AwsIotTopicRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsIotTopicRules.
func (c *FakeAwsIotTopicRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsiottopicrulesResource, c.ns, opts))

}

// Create takes the representation of a awsIotTopicRule and creates it.  Returns the server's representation of the awsIotTopicRule, and an error, if there is any.
func (c *FakeAwsIotTopicRules) Create(awsIotTopicRule *aws_v1.AwsIotTopicRule) (result *aws_v1.AwsIotTopicRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsiottopicrulesResource, c.ns, awsIotTopicRule), &aws_v1.AwsIotTopicRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsIotTopicRule), err
}

// Update takes the representation of a awsIotTopicRule and updates it. Returns the server's representation of the awsIotTopicRule, and an error, if there is any.
func (c *FakeAwsIotTopicRules) Update(awsIotTopicRule *aws_v1.AwsIotTopicRule) (result *aws_v1.AwsIotTopicRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsiottopicrulesResource, c.ns, awsIotTopicRule), &aws_v1.AwsIotTopicRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsIotTopicRule), err
}

// Delete takes name of the awsIotTopicRule and deletes it. Returns an error if one occurs.
func (c *FakeAwsIotTopicRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsiottopicrulesResource, c.ns, name), &aws_v1.AwsIotTopicRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsIotTopicRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsiottopicrulesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsIotTopicRuleList{})
	return err
}

// Patch applies the patch and returns the patched awsIotTopicRule.
func (c *FakeAwsIotTopicRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsIotTopicRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsiottopicrulesResource, c.ns, name, data, subresources...), &aws_v1.AwsIotTopicRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsIotTopicRule), err
}
