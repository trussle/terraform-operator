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

// FakeAwsSsmDocuments implements AwsSsmDocumentInterface
type FakeAwsSsmDocuments struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsssmdocumentsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsssmdocuments"}

var awsssmdocumentsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsSsmDocument"}

// Get takes name of the awsSsmDocument, and returns the corresponding awsSsmDocument object, and an error if there is any.
func (c *FakeAwsSsmDocuments) Get(name string, options v1.GetOptions) (result *aws_v1.AwsSsmDocument, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsssmdocumentsResource, c.ns, name), &aws_v1.AwsSsmDocument{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSsmDocument), err
}

// List takes label and field selectors, and returns the list of AwsSsmDocuments that match those selectors.
func (c *FakeAwsSsmDocuments) List(opts v1.ListOptions) (result *aws_v1.AwsSsmDocumentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsssmdocumentsResource, awsssmdocumentsKind, c.ns, opts), &aws_v1.AwsSsmDocumentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsSsmDocumentList{}
	for _, item := range obj.(*aws_v1.AwsSsmDocumentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSsmDocuments.
func (c *FakeAwsSsmDocuments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsssmdocumentsResource, c.ns, opts))

}

// Create takes the representation of a awsSsmDocument and creates it.  Returns the server's representation of the awsSsmDocument, and an error, if there is any.
func (c *FakeAwsSsmDocuments) Create(awsSsmDocument *aws_v1.AwsSsmDocument) (result *aws_v1.AwsSsmDocument, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsssmdocumentsResource, c.ns, awsSsmDocument), &aws_v1.AwsSsmDocument{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSsmDocument), err
}

// Update takes the representation of a awsSsmDocument and updates it. Returns the server's representation of the awsSsmDocument, and an error, if there is any.
func (c *FakeAwsSsmDocuments) Update(awsSsmDocument *aws_v1.AwsSsmDocument) (result *aws_v1.AwsSsmDocument, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsssmdocumentsResource, c.ns, awsSsmDocument), &aws_v1.AwsSsmDocument{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSsmDocument), err
}

// Delete takes name of the awsSsmDocument and deletes it. Returns an error if one occurs.
func (c *FakeAwsSsmDocuments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsssmdocumentsResource, c.ns, name), &aws_v1.AwsSsmDocument{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSsmDocuments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsssmdocumentsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsSsmDocumentList{})
	return err
}

// Patch applies the patch and returns the patched awsSsmDocument.
func (c *FakeAwsSsmDocuments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsSsmDocument, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsssmdocumentsResource, c.ns, name, data, subresources...), &aws_v1.AwsSsmDocument{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSsmDocument), err
}
