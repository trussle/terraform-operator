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

// FakeAwsSesDomainDkims implements AwsSesDomainDkimInterface
type FakeAwsSesDomainDkims struct {
	Fake *FakeTrussleV1
	ns   string
}

var awssesdomaindkimsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awssesdomaindkims"}

var awssesdomaindkimsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsSesDomainDkim"}

// Get takes name of the awsSesDomainDkim, and returns the corresponding awsSesDomainDkim object, and an error if there is any.
func (c *FakeAwsSesDomainDkims) Get(name string, options v1.GetOptions) (result *aws_v1.AwsSesDomainDkim, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awssesdomaindkimsResource, c.ns, name), &aws_v1.AwsSesDomainDkim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSesDomainDkim), err
}

// List takes label and field selectors, and returns the list of AwsSesDomainDkims that match those selectors.
func (c *FakeAwsSesDomainDkims) List(opts v1.ListOptions) (result *aws_v1.AwsSesDomainDkimList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awssesdomaindkimsResource, awssesdomaindkimsKind, c.ns, opts), &aws_v1.AwsSesDomainDkimList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsSesDomainDkimList{}
	for _, item := range obj.(*aws_v1.AwsSesDomainDkimList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSesDomainDkims.
func (c *FakeAwsSesDomainDkims) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awssesdomaindkimsResource, c.ns, opts))

}

// Create takes the representation of a awsSesDomainDkim and creates it.  Returns the server's representation of the awsSesDomainDkim, and an error, if there is any.
func (c *FakeAwsSesDomainDkims) Create(awsSesDomainDkim *aws_v1.AwsSesDomainDkim) (result *aws_v1.AwsSesDomainDkim, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awssesdomaindkimsResource, c.ns, awsSesDomainDkim), &aws_v1.AwsSesDomainDkim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSesDomainDkim), err
}

// Update takes the representation of a awsSesDomainDkim and updates it. Returns the server's representation of the awsSesDomainDkim, and an error, if there is any.
func (c *FakeAwsSesDomainDkims) Update(awsSesDomainDkim *aws_v1.AwsSesDomainDkim) (result *aws_v1.AwsSesDomainDkim, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awssesdomaindkimsResource, c.ns, awsSesDomainDkim), &aws_v1.AwsSesDomainDkim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSesDomainDkim), err
}

// Delete takes name of the awsSesDomainDkim and deletes it. Returns an error if one occurs.
func (c *FakeAwsSesDomainDkims) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awssesdomaindkimsResource, c.ns, name), &aws_v1.AwsSesDomainDkim{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSesDomainDkims) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awssesdomaindkimsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsSesDomainDkimList{})
	return err
}

// Patch applies the patch and returns the patched awsSesDomainDkim.
func (c *FakeAwsSesDomainDkims) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsSesDomainDkim, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awssesdomaindkimsResource, c.ns, name, data, subresources...), &aws_v1.AwsSesDomainDkim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSesDomainDkim), err
}
