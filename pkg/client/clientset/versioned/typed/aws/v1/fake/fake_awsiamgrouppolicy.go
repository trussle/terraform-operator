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

// FakeAwsIamGroupPolicies implements AwsIamGroupPolicyInterface
type FakeAwsIamGroupPolicies struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsiamgrouppoliciesResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsiamgrouppolicies"}

var awsiamgrouppoliciesKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsIamGroupPolicy"}

// Get takes name of the awsIamGroupPolicy, and returns the corresponding awsIamGroupPolicy object, and an error if there is any.
func (c *FakeAwsIamGroupPolicies) Get(name string, options v1.GetOptions) (result *aws_v1.AwsIamGroupPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsiamgrouppoliciesResource, c.ns, name), &aws_v1.AwsIamGroupPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsIamGroupPolicy), err
}

// List takes label and field selectors, and returns the list of AwsIamGroupPolicies that match those selectors.
func (c *FakeAwsIamGroupPolicies) List(opts v1.ListOptions) (result *aws_v1.AwsIamGroupPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsiamgrouppoliciesResource, awsiamgrouppoliciesKind, c.ns, opts), &aws_v1.AwsIamGroupPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsIamGroupPolicyList{}
	for _, item := range obj.(*aws_v1.AwsIamGroupPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsIamGroupPolicies.
func (c *FakeAwsIamGroupPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsiamgrouppoliciesResource, c.ns, opts))

}

// Create takes the representation of a awsIamGroupPolicy and creates it.  Returns the server's representation of the awsIamGroupPolicy, and an error, if there is any.
func (c *FakeAwsIamGroupPolicies) Create(awsIamGroupPolicy *aws_v1.AwsIamGroupPolicy) (result *aws_v1.AwsIamGroupPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsiamgrouppoliciesResource, c.ns, awsIamGroupPolicy), &aws_v1.AwsIamGroupPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsIamGroupPolicy), err
}

// Update takes the representation of a awsIamGroupPolicy and updates it. Returns the server's representation of the awsIamGroupPolicy, and an error, if there is any.
func (c *FakeAwsIamGroupPolicies) Update(awsIamGroupPolicy *aws_v1.AwsIamGroupPolicy) (result *aws_v1.AwsIamGroupPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsiamgrouppoliciesResource, c.ns, awsIamGroupPolicy), &aws_v1.AwsIamGroupPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsIamGroupPolicy), err
}

// Delete takes name of the awsIamGroupPolicy and deletes it. Returns an error if one occurs.
func (c *FakeAwsIamGroupPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsiamgrouppoliciesResource, c.ns, name), &aws_v1.AwsIamGroupPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsIamGroupPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsiamgrouppoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsIamGroupPolicyList{})
	return err
}

// Patch applies the patch and returns the patched awsIamGroupPolicy.
func (c *FakeAwsIamGroupPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsIamGroupPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsiamgrouppoliciesResource, c.ns, name, data, subresources...), &aws_v1.AwsIamGroupPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsIamGroupPolicy), err
}
