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

// FakeAwsNatGateways implements AwsNatGatewayInterface
type FakeAwsNatGateways struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsnatgatewaysResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsnatgateways"}

var awsnatgatewaysKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsNatGateway"}

// Get takes name of the awsNatGateway, and returns the corresponding awsNatGateway object, and an error if there is any.
func (c *FakeAwsNatGateways) Get(name string, options v1.GetOptions) (result *aws_v1.AwsNatGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsnatgatewaysResource, c.ns, name), &aws_v1.AwsNatGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsNatGateway), err
}

// List takes label and field selectors, and returns the list of AwsNatGateways that match those selectors.
func (c *FakeAwsNatGateways) List(opts v1.ListOptions) (result *aws_v1.AwsNatGatewayList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsnatgatewaysResource, awsnatgatewaysKind, c.ns, opts), &aws_v1.AwsNatGatewayList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsNatGatewayList{}
	for _, item := range obj.(*aws_v1.AwsNatGatewayList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsNatGateways.
func (c *FakeAwsNatGateways) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsnatgatewaysResource, c.ns, opts))

}

// Create takes the representation of a awsNatGateway and creates it.  Returns the server's representation of the awsNatGateway, and an error, if there is any.
func (c *FakeAwsNatGateways) Create(awsNatGateway *aws_v1.AwsNatGateway) (result *aws_v1.AwsNatGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsnatgatewaysResource, c.ns, awsNatGateway), &aws_v1.AwsNatGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsNatGateway), err
}

// Update takes the representation of a awsNatGateway and updates it. Returns the server's representation of the awsNatGateway, and an error, if there is any.
func (c *FakeAwsNatGateways) Update(awsNatGateway *aws_v1.AwsNatGateway) (result *aws_v1.AwsNatGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsnatgatewaysResource, c.ns, awsNatGateway), &aws_v1.AwsNatGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsNatGateway), err
}

// Delete takes name of the awsNatGateway and deletes it. Returns an error if one occurs.
func (c *FakeAwsNatGateways) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsnatgatewaysResource, c.ns, name), &aws_v1.AwsNatGateway{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsNatGateways) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsnatgatewaysResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsNatGatewayList{})
	return err
}

// Patch applies the patch and returns the patched awsNatGateway.
func (c *FakeAwsNatGateways) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsNatGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsnatgatewaysResource, c.ns, name, data, subresources...), &aws_v1.AwsNatGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsNatGateway), err
}
