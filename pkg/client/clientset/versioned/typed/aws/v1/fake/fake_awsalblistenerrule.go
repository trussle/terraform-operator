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

// FakeAwsAlbListenerRules implements AwsAlbListenerRuleInterface
type FakeAwsAlbListenerRules struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsalblistenerrulesResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsalblistenerrules"}

var awsalblistenerrulesKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsAlbListenerRule"}

// Get takes name of the awsAlbListenerRule, and returns the corresponding awsAlbListenerRule object, and an error if there is any.
func (c *FakeAwsAlbListenerRules) Get(name string, options v1.GetOptions) (result *aws_v1.AwsAlbListenerRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsalblistenerrulesResource, c.ns, name), &aws_v1.AwsAlbListenerRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsAlbListenerRule), err
}

// List takes label and field selectors, and returns the list of AwsAlbListenerRules that match those selectors.
func (c *FakeAwsAlbListenerRules) List(opts v1.ListOptions) (result *aws_v1.AwsAlbListenerRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsalblistenerrulesResource, awsalblistenerrulesKind, c.ns, opts), &aws_v1.AwsAlbListenerRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsAlbListenerRuleList{}
	for _, item := range obj.(*aws_v1.AwsAlbListenerRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsAlbListenerRules.
func (c *FakeAwsAlbListenerRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsalblistenerrulesResource, c.ns, opts))

}

// Create takes the representation of a awsAlbListenerRule and creates it.  Returns the server's representation of the awsAlbListenerRule, and an error, if there is any.
func (c *FakeAwsAlbListenerRules) Create(awsAlbListenerRule *aws_v1.AwsAlbListenerRule) (result *aws_v1.AwsAlbListenerRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsalblistenerrulesResource, c.ns, awsAlbListenerRule), &aws_v1.AwsAlbListenerRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsAlbListenerRule), err
}

// Update takes the representation of a awsAlbListenerRule and updates it. Returns the server's representation of the awsAlbListenerRule, and an error, if there is any.
func (c *FakeAwsAlbListenerRules) Update(awsAlbListenerRule *aws_v1.AwsAlbListenerRule) (result *aws_v1.AwsAlbListenerRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsalblistenerrulesResource, c.ns, awsAlbListenerRule), &aws_v1.AwsAlbListenerRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsAlbListenerRule), err
}

// Delete takes name of the awsAlbListenerRule and deletes it. Returns an error if one occurs.
func (c *FakeAwsAlbListenerRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsalblistenerrulesResource, c.ns, name), &aws_v1.AwsAlbListenerRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsAlbListenerRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsalblistenerrulesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsAlbListenerRuleList{})
	return err
}

// Patch applies the patch and returns the patched awsAlbListenerRule.
func (c *FakeAwsAlbListenerRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsAlbListenerRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsalblistenerrulesResource, c.ns, name, data, subresources...), &aws_v1.AwsAlbListenerRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsAlbListenerRule), err
}
