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

// FakeAwsEcsTaskDefinitions implements AwsEcsTaskDefinitionInterface
type FakeAwsEcsTaskDefinitions struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsecstaskdefinitionsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsecstaskdefinitions"}

var awsecstaskdefinitionsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsEcsTaskDefinition"}

// Get takes name of the awsEcsTaskDefinition, and returns the corresponding awsEcsTaskDefinition object, and an error if there is any.
func (c *FakeAwsEcsTaskDefinitions) Get(name string, options v1.GetOptions) (result *aws_v1.AwsEcsTaskDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsecstaskdefinitionsResource, c.ns, name), &aws_v1.AwsEcsTaskDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsEcsTaskDefinition), err
}

// List takes label and field selectors, and returns the list of AwsEcsTaskDefinitions that match those selectors.
func (c *FakeAwsEcsTaskDefinitions) List(opts v1.ListOptions) (result *aws_v1.AwsEcsTaskDefinitionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsecstaskdefinitionsResource, awsecstaskdefinitionsKind, c.ns, opts), &aws_v1.AwsEcsTaskDefinitionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsEcsTaskDefinitionList{}
	for _, item := range obj.(*aws_v1.AwsEcsTaskDefinitionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsEcsTaskDefinitions.
func (c *FakeAwsEcsTaskDefinitions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsecstaskdefinitionsResource, c.ns, opts))

}

// Create takes the representation of a awsEcsTaskDefinition and creates it.  Returns the server's representation of the awsEcsTaskDefinition, and an error, if there is any.
func (c *FakeAwsEcsTaskDefinitions) Create(awsEcsTaskDefinition *aws_v1.AwsEcsTaskDefinition) (result *aws_v1.AwsEcsTaskDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsecstaskdefinitionsResource, c.ns, awsEcsTaskDefinition), &aws_v1.AwsEcsTaskDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsEcsTaskDefinition), err
}

// Update takes the representation of a awsEcsTaskDefinition and updates it. Returns the server's representation of the awsEcsTaskDefinition, and an error, if there is any.
func (c *FakeAwsEcsTaskDefinitions) Update(awsEcsTaskDefinition *aws_v1.AwsEcsTaskDefinition) (result *aws_v1.AwsEcsTaskDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsecstaskdefinitionsResource, c.ns, awsEcsTaskDefinition), &aws_v1.AwsEcsTaskDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsEcsTaskDefinition), err
}

// Delete takes name of the awsEcsTaskDefinition and deletes it. Returns an error if one occurs.
func (c *FakeAwsEcsTaskDefinitions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsecstaskdefinitionsResource, c.ns, name), &aws_v1.AwsEcsTaskDefinition{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsEcsTaskDefinitions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsecstaskdefinitionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsEcsTaskDefinitionList{})
	return err
}

// Patch applies the patch and returns the patched awsEcsTaskDefinition.
func (c *FakeAwsEcsTaskDefinitions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsEcsTaskDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsecstaskdefinitionsResource, c.ns, name, data, subresources...), &aws_v1.AwsEcsTaskDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsEcsTaskDefinition), err
}
