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

// FakeAwsApiGatewayMethods implements AwsApiGatewayMethodInterface
type FakeAwsApiGatewayMethods struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsapigatewaymethodsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsapigatewaymethods"}

var awsapigatewaymethodsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsApiGatewayMethod"}

// Get takes name of the awsApiGatewayMethod, and returns the corresponding awsApiGatewayMethod object, and an error if there is any.
func (c *FakeAwsApiGatewayMethods) Get(name string, options v1.GetOptions) (result *aws_v1.AwsApiGatewayMethod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsapigatewaymethodsResource, c.ns, name), &aws_v1.AwsApiGatewayMethod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsApiGatewayMethod), err
}

// List takes label and field selectors, and returns the list of AwsApiGatewayMethods that match those selectors.
func (c *FakeAwsApiGatewayMethods) List(opts v1.ListOptions) (result *aws_v1.AwsApiGatewayMethodList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsapigatewaymethodsResource, awsapigatewaymethodsKind, c.ns, opts), &aws_v1.AwsApiGatewayMethodList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsApiGatewayMethodList{}
	for _, item := range obj.(*aws_v1.AwsApiGatewayMethodList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsApiGatewayMethods.
func (c *FakeAwsApiGatewayMethods) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsapigatewaymethodsResource, c.ns, opts))

}

// Create takes the representation of a awsApiGatewayMethod and creates it.  Returns the server's representation of the awsApiGatewayMethod, and an error, if there is any.
func (c *FakeAwsApiGatewayMethods) Create(awsApiGatewayMethod *aws_v1.AwsApiGatewayMethod) (result *aws_v1.AwsApiGatewayMethod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsapigatewaymethodsResource, c.ns, awsApiGatewayMethod), &aws_v1.AwsApiGatewayMethod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsApiGatewayMethod), err
}

// Update takes the representation of a awsApiGatewayMethod and updates it. Returns the server's representation of the awsApiGatewayMethod, and an error, if there is any.
func (c *FakeAwsApiGatewayMethods) Update(awsApiGatewayMethod *aws_v1.AwsApiGatewayMethod) (result *aws_v1.AwsApiGatewayMethod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsapigatewaymethodsResource, c.ns, awsApiGatewayMethod), &aws_v1.AwsApiGatewayMethod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsApiGatewayMethod), err
}

// Delete takes name of the awsApiGatewayMethod and deletes it. Returns an error if one occurs.
func (c *FakeAwsApiGatewayMethods) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsapigatewaymethodsResource, c.ns, name), &aws_v1.AwsApiGatewayMethod{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsApiGatewayMethods) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsapigatewaymethodsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsApiGatewayMethodList{})
	return err
}

// Patch applies the patch and returns the patched awsApiGatewayMethod.
func (c *FakeAwsApiGatewayMethods) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsApiGatewayMethod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsapigatewaymethodsResource, c.ns, name, data, subresources...), &aws_v1.AwsApiGatewayMethod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsApiGatewayMethod), err
}
