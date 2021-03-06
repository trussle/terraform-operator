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

// FakeAwsWafregionalWebAclAssociations implements AwsWafregionalWebAclAssociationInterface
type FakeAwsWafregionalWebAclAssociations struct {
	Fake *FakeTrussleV1
	ns   string
}

var awswafregionalwebaclassociationsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awswafregionalwebaclassociations"}

var awswafregionalwebaclassociationsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsWafregionalWebAclAssociation"}

// Get takes name of the awsWafregionalWebAclAssociation, and returns the corresponding awsWafregionalWebAclAssociation object, and an error if there is any.
func (c *FakeAwsWafregionalWebAclAssociations) Get(name string, options v1.GetOptions) (result *aws_v1.AwsWafregionalWebAclAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awswafregionalwebaclassociationsResource, c.ns, name), &aws_v1.AwsWafregionalWebAclAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsWafregionalWebAclAssociation), err
}

// List takes label and field selectors, and returns the list of AwsWafregionalWebAclAssociations that match those selectors.
func (c *FakeAwsWafregionalWebAclAssociations) List(opts v1.ListOptions) (result *aws_v1.AwsWafregionalWebAclAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awswafregionalwebaclassociationsResource, awswafregionalwebaclassociationsKind, c.ns, opts), &aws_v1.AwsWafregionalWebAclAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsWafregionalWebAclAssociationList{}
	for _, item := range obj.(*aws_v1.AwsWafregionalWebAclAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsWafregionalWebAclAssociations.
func (c *FakeAwsWafregionalWebAclAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awswafregionalwebaclassociationsResource, c.ns, opts))

}

// Create takes the representation of a awsWafregionalWebAclAssociation and creates it.  Returns the server's representation of the awsWafregionalWebAclAssociation, and an error, if there is any.
func (c *FakeAwsWafregionalWebAclAssociations) Create(awsWafregionalWebAclAssociation *aws_v1.AwsWafregionalWebAclAssociation) (result *aws_v1.AwsWafregionalWebAclAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awswafregionalwebaclassociationsResource, c.ns, awsWafregionalWebAclAssociation), &aws_v1.AwsWafregionalWebAclAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsWafregionalWebAclAssociation), err
}

// Update takes the representation of a awsWafregionalWebAclAssociation and updates it. Returns the server's representation of the awsWafregionalWebAclAssociation, and an error, if there is any.
func (c *FakeAwsWafregionalWebAclAssociations) Update(awsWafregionalWebAclAssociation *aws_v1.AwsWafregionalWebAclAssociation) (result *aws_v1.AwsWafregionalWebAclAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awswafregionalwebaclassociationsResource, c.ns, awsWafregionalWebAclAssociation), &aws_v1.AwsWafregionalWebAclAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsWafregionalWebAclAssociation), err
}

// Delete takes name of the awsWafregionalWebAclAssociation and deletes it. Returns an error if one occurs.
func (c *FakeAwsWafregionalWebAclAssociations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awswafregionalwebaclassociationsResource, c.ns, name), &aws_v1.AwsWafregionalWebAclAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsWafregionalWebAclAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awswafregionalwebaclassociationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsWafregionalWebAclAssociationList{})
	return err
}

// Patch applies the patch and returns the patched awsWafregionalWebAclAssociation.
func (c *FakeAwsWafregionalWebAclAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsWafregionalWebAclAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awswafregionalwebaclassociationsResource, c.ns, name, data, subresources...), &aws_v1.AwsWafregionalWebAclAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsWafregionalWebAclAssociation), err
}
