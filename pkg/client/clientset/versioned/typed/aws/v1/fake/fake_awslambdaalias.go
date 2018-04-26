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

// FakeAwsLambdaAliases implements AwsLambdaAliasInterface
type FakeAwsLambdaAliases struct {
	Fake *FakeTrussleV1
	ns   string
}

var awslambdaaliasesResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awslambdaaliases"}

var awslambdaaliasesKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsLambdaAlias"}

// Get takes name of the awsLambdaAlias, and returns the corresponding awsLambdaAlias object, and an error if there is any.
func (c *FakeAwsLambdaAliases) Get(name string, options v1.GetOptions) (result *aws_v1.AwsLambdaAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awslambdaaliasesResource, c.ns, name), &aws_v1.AwsLambdaAlias{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLambdaAlias), err
}

// List takes label and field selectors, and returns the list of AwsLambdaAliases that match those selectors.
func (c *FakeAwsLambdaAliases) List(opts v1.ListOptions) (result *aws_v1.AwsLambdaAliasList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awslambdaaliasesResource, awslambdaaliasesKind, c.ns, opts), &aws_v1.AwsLambdaAliasList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsLambdaAliasList{}
	for _, item := range obj.(*aws_v1.AwsLambdaAliasList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsLambdaAliases.
func (c *FakeAwsLambdaAliases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awslambdaaliasesResource, c.ns, opts))

}

// Create takes the representation of a awsLambdaAlias and creates it.  Returns the server's representation of the awsLambdaAlias, and an error, if there is any.
func (c *FakeAwsLambdaAliases) Create(awsLambdaAlias *aws_v1.AwsLambdaAlias) (result *aws_v1.AwsLambdaAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awslambdaaliasesResource, c.ns, awsLambdaAlias), &aws_v1.AwsLambdaAlias{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLambdaAlias), err
}

// Update takes the representation of a awsLambdaAlias and updates it. Returns the server's representation of the awsLambdaAlias, and an error, if there is any.
func (c *FakeAwsLambdaAliases) Update(awsLambdaAlias *aws_v1.AwsLambdaAlias) (result *aws_v1.AwsLambdaAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awslambdaaliasesResource, c.ns, awsLambdaAlias), &aws_v1.AwsLambdaAlias{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLambdaAlias), err
}

// Delete takes name of the awsLambdaAlias and deletes it. Returns an error if one occurs.
func (c *FakeAwsLambdaAliases) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awslambdaaliasesResource, c.ns, name), &aws_v1.AwsLambdaAlias{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsLambdaAliases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awslambdaaliasesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsLambdaAliasList{})
	return err
}

// Patch applies the patch and returns the patched awsLambdaAlias.
func (c *FakeAwsLambdaAliases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsLambdaAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awslambdaaliasesResource, c.ns, name, data, subresources...), &aws_v1.AwsLambdaAlias{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLambdaAlias), err
}
