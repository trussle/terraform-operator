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

// FakeAwsLightsailDomains implements AwsLightsailDomainInterface
type FakeAwsLightsailDomains struct {
	Fake *FakeTrussleV1
	ns   string
}

var awslightsaildomainsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awslightsaildomains"}

var awslightsaildomainsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsLightsailDomain"}

// Get takes name of the awsLightsailDomain, and returns the corresponding awsLightsailDomain object, and an error if there is any.
func (c *FakeAwsLightsailDomains) Get(name string, options v1.GetOptions) (result *aws_v1.AwsLightsailDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awslightsaildomainsResource, c.ns, name), &aws_v1.AwsLightsailDomain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLightsailDomain), err
}

// List takes label and field selectors, and returns the list of AwsLightsailDomains that match those selectors.
func (c *FakeAwsLightsailDomains) List(opts v1.ListOptions) (result *aws_v1.AwsLightsailDomainList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awslightsaildomainsResource, awslightsaildomainsKind, c.ns, opts), &aws_v1.AwsLightsailDomainList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsLightsailDomainList{}
	for _, item := range obj.(*aws_v1.AwsLightsailDomainList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsLightsailDomains.
func (c *FakeAwsLightsailDomains) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awslightsaildomainsResource, c.ns, opts))

}

// Create takes the representation of a awsLightsailDomain and creates it.  Returns the server's representation of the awsLightsailDomain, and an error, if there is any.
func (c *FakeAwsLightsailDomains) Create(awsLightsailDomain *aws_v1.AwsLightsailDomain) (result *aws_v1.AwsLightsailDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awslightsaildomainsResource, c.ns, awsLightsailDomain), &aws_v1.AwsLightsailDomain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLightsailDomain), err
}

// Update takes the representation of a awsLightsailDomain and updates it. Returns the server's representation of the awsLightsailDomain, and an error, if there is any.
func (c *FakeAwsLightsailDomains) Update(awsLightsailDomain *aws_v1.AwsLightsailDomain) (result *aws_v1.AwsLightsailDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awslightsaildomainsResource, c.ns, awsLightsailDomain), &aws_v1.AwsLightsailDomain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLightsailDomain), err
}

// Delete takes name of the awsLightsailDomain and deletes it. Returns an error if one occurs.
func (c *FakeAwsLightsailDomains) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awslightsaildomainsResource, c.ns, name), &aws_v1.AwsLightsailDomain{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsLightsailDomains) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awslightsaildomainsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsLightsailDomainList{})
	return err
}

// Patch applies the patch and returns the patched awsLightsailDomain.
func (c *FakeAwsLightsailDomains) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsLightsailDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awslightsaildomainsResource, c.ns, name, data, subresources...), &aws_v1.AwsLightsailDomain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLightsailDomain), err
}