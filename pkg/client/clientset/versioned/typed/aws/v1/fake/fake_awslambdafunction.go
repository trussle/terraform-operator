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

// FakeAwsLambdaFunctions implements AwsLambdaFunctionInterface
type FakeAwsLambdaFunctions struct {
	Fake *FakeTrussleV1
	ns   string
}

var awslambdafunctionsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awslambdafunctions"}

var awslambdafunctionsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsLambdaFunction"}

// Get takes name of the awsLambdaFunction, and returns the corresponding awsLambdaFunction object, and an error if there is any.
func (c *FakeAwsLambdaFunctions) Get(name string, options v1.GetOptions) (result *aws_v1.AwsLambdaFunction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awslambdafunctionsResource, c.ns, name), &aws_v1.AwsLambdaFunction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLambdaFunction), err
}

// List takes label and field selectors, and returns the list of AwsLambdaFunctions that match those selectors.
func (c *FakeAwsLambdaFunctions) List(opts v1.ListOptions) (result *aws_v1.AwsLambdaFunctionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awslambdafunctionsResource, awslambdafunctionsKind, c.ns, opts), &aws_v1.AwsLambdaFunctionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsLambdaFunctionList{}
	for _, item := range obj.(*aws_v1.AwsLambdaFunctionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsLambdaFunctions.
func (c *FakeAwsLambdaFunctions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awslambdafunctionsResource, c.ns, opts))

}

// Create takes the representation of a awsLambdaFunction and creates it.  Returns the server's representation of the awsLambdaFunction, and an error, if there is any.
func (c *FakeAwsLambdaFunctions) Create(awsLambdaFunction *aws_v1.AwsLambdaFunction) (result *aws_v1.AwsLambdaFunction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awslambdafunctionsResource, c.ns, awsLambdaFunction), &aws_v1.AwsLambdaFunction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLambdaFunction), err
}

// Update takes the representation of a awsLambdaFunction and updates it. Returns the server's representation of the awsLambdaFunction, and an error, if there is any.
func (c *FakeAwsLambdaFunctions) Update(awsLambdaFunction *aws_v1.AwsLambdaFunction) (result *aws_v1.AwsLambdaFunction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awslambdafunctionsResource, c.ns, awsLambdaFunction), &aws_v1.AwsLambdaFunction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLambdaFunction), err
}

// Delete takes name of the awsLambdaFunction and deletes it. Returns an error if one occurs.
func (c *FakeAwsLambdaFunctions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awslambdafunctionsResource, c.ns, name), &aws_v1.AwsLambdaFunction{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsLambdaFunctions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awslambdafunctionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsLambdaFunctionList{})
	return err
}

// Patch applies the patch and returns the patched awsLambdaFunction.
func (c *FakeAwsLambdaFunctions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsLambdaFunction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awslambdafunctionsResource, c.ns, name, data, subresources...), &aws_v1.AwsLambdaFunction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLambdaFunction), err
}
