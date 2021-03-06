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

// FakeAwsNetworkAcls implements AwsNetworkAclInterface
type FakeAwsNetworkAcls struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsnetworkaclsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsnetworkacls"}

var awsnetworkaclsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsNetworkAcl"}

// Get takes name of the awsNetworkAcl, and returns the corresponding awsNetworkAcl object, and an error if there is any.
func (c *FakeAwsNetworkAcls) Get(name string, options v1.GetOptions) (result *aws_v1.AwsNetworkAcl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsnetworkaclsResource, c.ns, name), &aws_v1.AwsNetworkAcl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsNetworkAcl), err
}

// List takes label and field selectors, and returns the list of AwsNetworkAcls that match those selectors.
func (c *FakeAwsNetworkAcls) List(opts v1.ListOptions) (result *aws_v1.AwsNetworkAclList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsnetworkaclsResource, awsnetworkaclsKind, c.ns, opts), &aws_v1.AwsNetworkAclList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsNetworkAclList{}
	for _, item := range obj.(*aws_v1.AwsNetworkAclList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsNetworkAcls.
func (c *FakeAwsNetworkAcls) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsnetworkaclsResource, c.ns, opts))

}

// Create takes the representation of a awsNetworkAcl and creates it.  Returns the server's representation of the awsNetworkAcl, and an error, if there is any.
func (c *FakeAwsNetworkAcls) Create(awsNetworkAcl *aws_v1.AwsNetworkAcl) (result *aws_v1.AwsNetworkAcl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsnetworkaclsResource, c.ns, awsNetworkAcl), &aws_v1.AwsNetworkAcl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsNetworkAcl), err
}

// Update takes the representation of a awsNetworkAcl and updates it. Returns the server's representation of the awsNetworkAcl, and an error, if there is any.
func (c *FakeAwsNetworkAcls) Update(awsNetworkAcl *aws_v1.AwsNetworkAcl) (result *aws_v1.AwsNetworkAcl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsnetworkaclsResource, c.ns, awsNetworkAcl), &aws_v1.AwsNetworkAcl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsNetworkAcl), err
}

// Delete takes name of the awsNetworkAcl and deletes it. Returns an error if one occurs.
func (c *FakeAwsNetworkAcls) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsnetworkaclsResource, c.ns, name), &aws_v1.AwsNetworkAcl{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsNetworkAcls) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsnetworkaclsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsNetworkAclList{})
	return err
}

// Patch applies the patch and returns the patched awsNetworkAcl.
func (c *FakeAwsNetworkAcls) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsNetworkAcl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsnetworkaclsResource, c.ns, name, data, subresources...), &aws_v1.AwsNetworkAcl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsNetworkAcl), err
}
