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

// FakeAwsLambdaEventSourceMappings implements AwsLambdaEventSourceMappingInterface
type FakeAwsLambdaEventSourceMappings struct {
	Fake *FakeTrussleV1
	ns   string
}

var awslambdaeventsourcemappingsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awslambdaeventsourcemappings"}

var awslambdaeventsourcemappingsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsLambdaEventSourceMapping"}

// Get takes name of the awsLambdaEventSourceMapping, and returns the corresponding awsLambdaEventSourceMapping object, and an error if there is any.
func (c *FakeAwsLambdaEventSourceMappings) Get(name string, options v1.GetOptions) (result *aws_v1.AwsLambdaEventSourceMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awslambdaeventsourcemappingsResource, c.ns, name), &aws_v1.AwsLambdaEventSourceMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLambdaEventSourceMapping), err
}

// List takes label and field selectors, and returns the list of AwsLambdaEventSourceMappings that match those selectors.
func (c *FakeAwsLambdaEventSourceMappings) List(opts v1.ListOptions) (result *aws_v1.AwsLambdaEventSourceMappingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awslambdaeventsourcemappingsResource, awslambdaeventsourcemappingsKind, c.ns, opts), &aws_v1.AwsLambdaEventSourceMappingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsLambdaEventSourceMappingList{}
	for _, item := range obj.(*aws_v1.AwsLambdaEventSourceMappingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsLambdaEventSourceMappings.
func (c *FakeAwsLambdaEventSourceMappings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awslambdaeventsourcemappingsResource, c.ns, opts))

}

// Create takes the representation of a awsLambdaEventSourceMapping and creates it.  Returns the server's representation of the awsLambdaEventSourceMapping, and an error, if there is any.
func (c *FakeAwsLambdaEventSourceMappings) Create(awsLambdaEventSourceMapping *aws_v1.AwsLambdaEventSourceMapping) (result *aws_v1.AwsLambdaEventSourceMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awslambdaeventsourcemappingsResource, c.ns, awsLambdaEventSourceMapping), &aws_v1.AwsLambdaEventSourceMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLambdaEventSourceMapping), err
}

// Update takes the representation of a awsLambdaEventSourceMapping and updates it. Returns the server's representation of the awsLambdaEventSourceMapping, and an error, if there is any.
func (c *FakeAwsLambdaEventSourceMappings) Update(awsLambdaEventSourceMapping *aws_v1.AwsLambdaEventSourceMapping) (result *aws_v1.AwsLambdaEventSourceMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awslambdaeventsourcemappingsResource, c.ns, awsLambdaEventSourceMapping), &aws_v1.AwsLambdaEventSourceMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLambdaEventSourceMapping), err
}

// Delete takes name of the awsLambdaEventSourceMapping and deletes it. Returns an error if one occurs.
func (c *FakeAwsLambdaEventSourceMappings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awslambdaeventsourcemappingsResource, c.ns, name), &aws_v1.AwsLambdaEventSourceMapping{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsLambdaEventSourceMappings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awslambdaeventsourcemappingsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsLambdaEventSourceMappingList{})
	return err
}

// Patch applies the patch and returns the patched awsLambdaEventSourceMapping.
func (c *FakeAwsLambdaEventSourceMappings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsLambdaEventSourceMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awslambdaeventsourcemappingsResource, c.ns, name, data, subresources...), &aws_v1.AwsLambdaEventSourceMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLambdaEventSourceMapping), err
}
