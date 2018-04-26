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

// FakeAwsRedshiftSecurityGroups implements AwsRedshiftSecurityGroupInterface
type FakeAwsRedshiftSecurityGroups struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsredshiftsecuritygroupsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsredshiftsecuritygroups"}

var awsredshiftsecuritygroupsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsRedshiftSecurityGroup"}

// Get takes name of the awsRedshiftSecurityGroup, and returns the corresponding awsRedshiftSecurityGroup object, and an error if there is any.
func (c *FakeAwsRedshiftSecurityGroups) Get(name string, options v1.GetOptions) (result *aws_v1.AwsRedshiftSecurityGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsredshiftsecuritygroupsResource, c.ns, name), &aws_v1.AwsRedshiftSecurityGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsRedshiftSecurityGroup), err
}

// List takes label and field selectors, and returns the list of AwsRedshiftSecurityGroups that match those selectors.
func (c *FakeAwsRedshiftSecurityGroups) List(opts v1.ListOptions) (result *aws_v1.AwsRedshiftSecurityGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsredshiftsecuritygroupsResource, awsredshiftsecuritygroupsKind, c.ns, opts), &aws_v1.AwsRedshiftSecurityGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsRedshiftSecurityGroupList{}
	for _, item := range obj.(*aws_v1.AwsRedshiftSecurityGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsRedshiftSecurityGroups.
func (c *FakeAwsRedshiftSecurityGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsredshiftsecuritygroupsResource, c.ns, opts))

}

// Create takes the representation of a awsRedshiftSecurityGroup and creates it.  Returns the server's representation of the awsRedshiftSecurityGroup, and an error, if there is any.
func (c *FakeAwsRedshiftSecurityGroups) Create(awsRedshiftSecurityGroup *aws_v1.AwsRedshiftSecurityGroup) (result *aws_v1.AwsRedshiftSecurityGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsredshiftsecuritygroupsResource, c.ns, awsRedshiftSecurityGroup), &aws_v1.AwsRedshiftSecurityGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsRedshiftSecurityGroup), err
}

// Update takes the representation of a awsRedshiftSecurityGroup and updates it. Returns the server's representation of the awsRedshiftSecurityGroup, and an error, if there is any.
func (c *FakeAwsRedshiftSecurityGroups) Update(awsRedshiftSecurityGroup *aws_v1.AwsRedshiftSecurityGroup) (result *aws_v1.AwsRedshiftSecurityGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsredshiftsecuritygroupsResource, c.ns, awsRedshiftSecurityGroup), &aws_v1.AwsRedshiftSecurityGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsRedshiftSecurityGroup), err
}

// Delete takes name of the awsRedshiftSecurityGroup and deletes it. Returns an error if one occurs.
func (c *FakeAwsRedshiftSecurityGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsredshiftsecuritygroupsResource, c.ns, name), &aws_v1.AwsRedshiftSecurityGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsRedshiftSecurityGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsredshiftsecuritygroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsRedshiftSecurityGroupList{})
	return err
}

// Patch applies the patch and returns the patched awsRedshiftSecurityGroup.
func (c *FakeAwsRedshiftSecurityGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsRedshiftSecurityGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsredshiftsecuritygroupsResource, c.ns, name, data, subresources...), &aws_v1.AwsRedshiftSecurityGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsRedshiftSecurityGroup), err
}
