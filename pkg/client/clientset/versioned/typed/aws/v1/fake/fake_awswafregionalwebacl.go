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

// FakeAwsWafregionalWebAcls implements AwsWafregionalWebAclInterface
type FakeAwsWafregionalWebAcls struct {
	Fake *FakeTrussleV1
	ns   string
}

var awswafregionalwebaclsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awswafregionalwebacls"}

var awswafregionalwebaclsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsWafregionalWebAcl"}

// Get takes name of the awsWafregionalWebAcl, and returns the corresponding awsWafregionalWebAcl object, and an error if there is any.
func (c *FakeAwsWafregionalWebAcls) Get(name string, options v1.GetOptions) (result *aws_v1.AwsWafregionalWebAcl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awswafregionalwebaclsResource, c.ns, name), &aws_v1.AwsWafregionalWebAcl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsWafregionalWebAcl), err
}

// List takes label and field selectors, and returns the list of AwsWafregionalWebAcls that match those selectors.
func (c *FakeAwsWafregionalWebAcls) List(opts v1.ListOptions) (result *aws_v1.AwsWafregionalWebAclList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awswafregionalwebaclsResource, awswafregionalwebaclsKind, c.ns, opts), &aws_v1.AwsWafregionalWebAclList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsWafregionalWebAclList{}
	for _, item := range obj.(*aws_v1.AwsWafregionalWebAclList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsWafregionalWebAcls.
func (c *FakeAwsWafregionalWebAcls) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awswafregionalwebaclsResource, c.ns, opts))

}

// Create takes the representation of a awsWafregionalWebAcl and creates it.  Returns the server's representation of the awsWafregionalWebAcl, and an error, if there is any.
func (c *FakeAwsWafregionalWebAcls) Create(awsWafregionalWebAcl *aws_v1.AwsWafregionalWebAcl) (result *aws_v1.AwsWafregionalWebAcl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awswafregionalwebaclsResource, c.ns, awsWafregionalWebAcl), &aws_v1.AwsWafregionalWebAcl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsWafregionalWebAcl), err
}

// Update takes the representation of a awsWafregionalWebAcl and updates it. Returns the server's representation of the awsWafregionalWebAcl, and an error, if there is any.
func (c *FakeAwsWafregionalWebAcls) Update(awsWafregionalWebAcl *aws_v1.AwsWafregionalWebAcl) (result *aws_v1.AwsWafregionalWebAcl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awswafregionalwebaclsResource, c.ns, awsWafregionalWebAcl), &aws_v1.AwsWafregionalWebAcl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsWafregionalWebAcl), err
}

// Delete takes name of the awsWafregionalWebAcl and deletes it. Returns an error if one occurs.
func (c *FakeAwsWafregionalWebAcls) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awswafregionalwebaclsResource, c.ns, name), &aws_v1.AwsWafregionalWebAcl{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsWafregionalWebAcls) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awswafregionalwebaclsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsWafregionalWebAclList{})
	return err
}

// Patch applies the patch and returns the patched awsWafregionalWebAcl.
func (c *FakeAwsWafregionalWebAcls) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsWafregionalWebAcl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awswafregionalwebaclsResource, c.ns, name, data, subresources...), &aws_v1.AwsWafregionalWebAcl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsWafregionalWebAcl), err
}
