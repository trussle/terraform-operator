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

// FakeAwsCognitoUserPoolDomains implements AwsCognitoUserPoolDomainInterface
type FakeAwsCognitoUserPoolDomains struct {
	Fake *FakeTrussleV1
	ns   string
}

var awscognitouserpooldomainsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awscognitouserpooldomains"}

var awscognitouserpooldomainsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsCognitoUserPoolDomain"}

// Get takes name of the awsCognitoUserPoolDomain, and returns the corresponding awsCognitoUserPoolDomain object, and an error if there is any.
func (c *FakeAwsCognitoUserPoolDomains) Get(name string, options v1.GetOptions) (result *aws_v1.AwsCognitoUserPoolDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awscognitouserpooldomainsResource, c.ns, name), &aws_v1.AwsCognitoUserPoolDomain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCognitoUserPoolDomain), err
}

// List takes label and field selectors, and returns the list of AwsCognitoUserPoolDomains that match those selectors.
func (c *FakeAwsCognitoUserPoolDomains) List(opts v1.ListOptions) (result *aws_v1.AwsCognitoUserPoolDomainList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awscognitouserpooldomainsResource, awscognitouserpooldomainsKind, c.ns, opts), &aws_v1.AwsCognitoUserPoolDomainList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsCognitoUserPoolDomainList{}
	for _, item := range obj.(*aws_v1.AwsCognitoUserPoolDomainList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsCognitoUserPoolDomains.
func (c *FakeAwsCognitoUserPoolDomains) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awscognitouserpooldomainsResource, c.ns, opts))

}

// Create takes the representation of a awsCognitoUserPoolDomain and creates it.  Returns the server's representation of the awsCognitoUserPoolDomain, and an error, if there is any.
func (c *FakeAwsCognitoUserPoolDomains) Create(awsCognitoUserPoolDomain *aws_v1.AwsCognitoUserPoolDomain) (result *aws_v1.AwsCognitoUserPoolDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awscognitouserpooldomainsResource, c.ns, awsCognitoUserPoolDomain), &aws_v1.AwsCognitoUserPoolDomain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCognitoUserPoolDomain), err
}

// Update takes the representation of a awsCognitoUserPoolDomain and updates it. Returns the server's representation of the awsCognitoUserPoolDomain, and an error, if there is any.
func (c *FakeAwsCognitoUserPoolDomains) Update(awsCognitoUserPoolDomain *aws_v1.AwsCognitoUserPoolDomain) (result *aws_v1.AwsCognitoUserPoolDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awscognitouserpooldomainsResource, c.ns, awsCognitoUserPoolDomain), &aws_v1.AwsCognitoUserPoolDomain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCognitoUserPoolDomain), err
}

// Delete takes name of the awsCognitoUserPoolDomain and deletes it. Returns an error if one occurs.
func (c *FakeAwsCognitoUserPoolDomains) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awscognitouserpooldomainsResource, c.ns, name), &aws_v1.AwsCognitoUserPoolDomain{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsCognitoUserPoolDomains) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awscognitouserpooldomainsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsCognitoUserPoolDomainList{})
	return err
}

// Patch applies the patch and returns the patched awsCognitoUserPoolDomain.
func (c *FakeAwsCognitoUserPoolDomains) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsCognitoUserPoolDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awscognitouserpooldomainsResource, c.ns, name, data, subresources...), &aws_v1.AwsCognitoUserPoolDomain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCognitoUserPoolDomain), err
}
