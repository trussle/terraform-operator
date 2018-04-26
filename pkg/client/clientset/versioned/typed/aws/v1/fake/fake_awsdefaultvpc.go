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

// FakeAwsDefaultVpcs implements AwsDefaultVpcInterface
type FakeAwsDefaultVpcs struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsdefaultvpcsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsdefaultvpcs"}

var awsdefaultvpcsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsDefaultVpc"}

// Get takes name of the awsDefaultVpc, and returns the corresponding awsDefaultVpc object, and an error if there is any.
func (c *FakeAwsDefaultVpcs) Get(name string, options v1.GetOptions) (result *aws_v1.AwsDefaultVpc, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsdefaultvpcsResource, c.ns, name), &aws_v1.AwsDefaultVpc{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsDefaultVpc), err
}

// List takes label and field selectors, and returns the list of AwsDefaultVpcs that match those selectors.
func (c *FakeAwsDefaultVpcs) List(opts v1.ListOptions) (result *aws_v1.AwsDefaultVpcList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsdefaultvpcsResource, awsdefaultvpcsKind, c.ns, opts), &aws_v1.AwsDefaultVpcList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsDefaultVpcList{}
	for _, item := range obj.(*aws_v1.AwsDefaultVpcList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsDefaultVpcs.
func (c *FakeAwsDefaultVpcs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsdefaultvpcsResource, c.ns, opts))

}

// Create takes the representation of a awsDefaultVpc and creates it.  Returns the server's representation of the awsDefaultVpc, and an error, if there is any.
func (c *FakeAwsDefaultVpcs) Create(awsDefaultVpc *aws_v1.AwsDefaultVpc) (result *aws_v1.AwsDefaultVpc, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsdefaultvpcsResource, c.ns, awsDefaultVpc), &aws_v1.AwsDefaultVpc{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsDefaultVpc), err
}

// Update takes the representation of a awsDefaultVpc and updates it. Returns the server's representation of the awsDefaultVpc, and an error, if there is any.
func (c *FakeAwsDefaultVpcs) Update(awsDefaultVpc *aws_v1.AwsDefaultVpc) (result *aws_v1.AwsDefaultVpc, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsdefaultvpcsResource, c.ns, awsDefaultVpc), &aws_v1.AwsDefaultVpc{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsDefaultVpc), err
}

// Delete takes name of the awsDefaultVpc and deletes it. Returns an error if one occurs.
func (c *FakeAwsDefaultVpcs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsdefaultvpcsResource, c.ns, name), &aws_v1.AwsDefaultVpc{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsDefaultVpcs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsdefaultvpcsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsDefaultVpcList{})
	return err
}

// Patch applies the patch and returns the patched awsDefaultVpc.
func (c *FakeAwsDefaultVpcs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsDefaultVpc, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsdefaultvpcsResource, c.ns, name, data, subresources...), &aws_v1.AwsDefaultVpc{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsDefaultVpc), err
}
