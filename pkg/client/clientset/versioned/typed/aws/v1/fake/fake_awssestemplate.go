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

// FakeAwsSesTemplates implements AwsSesTemplateInterface
type FakeAwsSesTemplates struct {
	Fake *FakeTrussleV1
	ns   string
}

var awssestemplatesResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awssestemplates"}

var awssestemplatesKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsSesTemplate"}

// Get takes name of the awsSesTemplate, and returns the corresponding awsSesTemplate object, and an error if there is any.
func (c *FakeAwsSesTemplates) Get(name string, options v1.GetOptions) (result *aws_v1.AwsSesTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awssestemplatesResource, c.ns, name), &aws_v1.AwsSesTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSesTemplate), err
}

// List takes label and field selectors, and returns the list of AwsSesTemplates that match those selectors.
func (c *FakeAwsSesTemplates) List(opts v1.ListOptions) (result *aws_v1.AwsSesTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awssestemplatesResource, awssestemplatesKind, c.ns, opts), &aws_v1.AwsSesTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsSesTemplateList{}
	for _, item := range obj.(*aws_v1.AwsSesTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSesTemplates.
func (c *FakeAwsSesTemplates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awssestemplatesResource, c.ns, opts))

}

// Create takes the representation of a awsSesTemplate and creates it.  Returns the server's representation of the awsSesTemplate, and an error, if there is any.
func (c *FakeAwsSesTemplates) Create(awsSesTemplate *aws_v1.AwsSesTemplate) (result *aws_v1.AwsSesTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awssestemplatesResource, c.ns, awsSesTemplate), &aws_v1.AwsSesTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSesTemplate), err
}

// Update takes the representation of a awsSesTemplate and updates it. Returns the server's representation of the awsSesTemplate, and an error, if there is any.
func (c *FakeAwsSesTemplates) Update(awsSesTemplate *aws_v1.AwsSesTemplate) (result *aws_v1.AwsSesTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awssestemplatesResource, c.ns, awsSesTemplate), &aws_v1.AwsSesTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSesTemplate), err
}

// Delete takes name of the awsSesTemplate and deletes it. Returns an error if one occurs.
func (c *FakeAwsSesTemplates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awssestemplatesResource, c.ns, name), &aws_v1.AwsSesTemplate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSesTemplates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awssestemplatesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsSesTemplateList{})
	return err
}

// Patch applies the patch and returns the patched awsSesTemplate.
func (c *FakeAwsSesTemplates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsSesTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awssestemplatesResource, c.ns, name, data, subresources...), &aws_v1.AwsSesTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSesTemplate), err
}
