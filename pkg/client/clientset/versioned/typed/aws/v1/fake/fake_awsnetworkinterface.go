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

// FakeAwsNetworkInterfaces implements AwsNetworkInterfaceInterface
type FakeAwsNetworkInterfaces struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsnetworkinterfacesResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsnetworkinterfaces"}

var awsnetworkinterfacesKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsNetworkInterface"}

// Get takes name of the awsNetworkInterface, and returns the corresponding awsNetworkInterface object, and an error if there is any.
func (c *FakeAwsNetworkInterfaces) Get(name string, options v1.GetOptions) (result *aws_v1.AwsNetworkInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsnetworkinterfacesResource, c.ns, name), &aws_v1.AwsNetworkInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsNetworkInterface), err
}

// List takes label and field selectors, and returns the list of AwsNetworkInterfaces that match those selectors.
func (c *FakeAwsNetworkInterfaces) List(opts v1.ListOptions) (result *aws_v1.AwsNetworkInterfaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsnetworkinterfacesResource, awsnetworkinterfacesKind, c.ns, opts), &aws_v1.AwsNetworkInterfaceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsNetworkInterfaceList{}
	for _, item := range obj.(*aws_v1.AwsNetworkInterfaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsNetworkInterfaces.
func (c *FakeAwsNetworkInterfaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsnetworkinterfacesResource, c.ns, opts))

}

// Create takes the representation of a awsNetworkInterface and creates it.  Returns the server's representation of the awsNetworkInterface, and an error, if there is any.
func (c *FakeAwsNetworkInterfaces) Create(awsNetworkInterface *aws_v1.AwsNetworkInterface) (result *aws_v1.AwsNetworkInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsnetworkinterfacesResource, c.ns, awsNetworkInterface), &aws_v1.AwsNetworkInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsNetworkInterface), err
}

// Update takes the representation of a awsNetworkInterface and updates it. Returns the server's representation of the awsNetworkInterface, and an error, if there is any.
func (c *FakeAwsNetworkInterfaces) Update(awsNetworkInterface *aws_v1.AwsNetworkInterface) (result *aws_v1.AwsNetworkInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsnetworkinterfacesResource, c.ns, awsNetworkInterface), &aws_v1.AwsNetworkInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsNetworkInterface), err
}

// Delete takes name of the awsNetworkInterface and deletes it. Returns an error if one occurs.
func (c *FakeAwsNetworkInterfaces) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsnetworkinterfacesResource, c.ns, name), &aws_v1.AwsNetworkInterface{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsNetworkInterfaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsnetworkinterfacesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsNetworkInterfaceList{})
	return err
}

// Patch applies the patch and returns the patched awsNetworkInterface.
func (c *FakeAwsNetworkInterfaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsNetworkInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsnetworkinterfacesResource, c.ns, name, data, subresources...), &aws_v1.AwsNetworkInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsNetworkInterface), err
}
