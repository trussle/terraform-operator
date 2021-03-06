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

// FakeAwsGlacierVaults implements AwsGlacierVaultInterface
type FakeAwsGlacierVaults struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsglaciervaultsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsglaciervaults"}

var awsglaciervaultsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsGlacierVault"}

// Get takes name of the awsGlacierVault, and returns the corresponding awsGlacierVault object, and an error if there is any.
func (c *FakeAwsGlacierVaults) Get(name string, options v1.GetOptions) (result *aws_v1.AwsGlacierVault, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsglaciervaultsResource, c.ns, name), &aws_v1.AwsGlacierVault{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsGlacierVault), err
}

// List takes label and field selectors, and returns the list of AwsGlacierVaults that match those selectors.
func (c *FakeAwsGlacierVaults) List(opts v1.ListOptions) (result *aws_v1.AwsGlacierVaultList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsglaciervaultsResource, awsglaciervaultsKind, c.ns, opts), &aws_v1.AwsGlacierVaultList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsGlacierVaultList{}
	for _, item := range obj.(*aws_v1.AwsGlacierVaultList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsGlacierVaults.
func (c *FakeAwsGlacierVaults) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsglaciervaultsResource, c.ns, opts))

}

// Create takes the representation of a awsGlacierVault and creates it.  Returns the server's representation of the awsGlacierVault, and an error, if there is any.
func (c *FakeAwsGlacierVaults) Create(awsGlacierVault *aws_v1.AwsGlacierVault) (result *aws_v1.AwsGlacierVault, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsglaciervaultsResource, c.ns, awsGlacierVault), &aws_v1.AwsGlacierVault{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsGlacierVault), err
}

// Update takes the representation of a awsGlacierVault and updates it. Returns the server's representation of the awsGlacierVault, and an error, if there is any.
func (c *FakeAwsGlacierVaults) Update(awsGlacierVault *aws_v1.AwsGlacierVault) (result *aws_v1.AwsGlacierVault, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsglaciervaultsResource, c.ns, awsGlacierVault), &aws_v1.AwsGlacierVault{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsGlacierVault), err
}

// Delete takes name of the awsGlacierVault and deletes it. Returns an error if one occurs.
func (c *FakeAwsGlacierVaults) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsglaciervaultsResource, c.ns, name), &aws_v1.AwsGlacierVault{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsGlacierVaults) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsglaciervaultsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsGlacierVaultList{})
	return err
}

// Patch applies the patch and returns the patched awsGlacierVault.
func (c *FakeAwsGlacierVaults) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsGlacierVault, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsglaciervaultsResource, c.ns, name, data, subresources...), &aws_v1.AwsGlacierVault{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsGlacierVault), err
}
