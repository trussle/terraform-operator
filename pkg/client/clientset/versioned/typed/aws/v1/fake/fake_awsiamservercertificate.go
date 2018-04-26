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

// FakeAwsIamServerCertificates implements AwsIamServerCertificateInterface
type FakeAwsIamServerCertificates struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsiamservercertificatesResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsiamservercertificates"}

var awsiamservercertificatesKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsIamServerCertificate"}

// Get takes name of the awsIamServerCertificate, and returns the corresponding awsIamServerCertificate object, and an error if there is any.
func (c *FakeAwsIamServerCertificates) Get(name string, options v1.GetOptions) (result *aws_v1.AwsIamServerCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsiamservercertificatesResource, c.ns, name), &aws_v1.AwsIamServerCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsIamServerCertificate), err
}

// List takes label and field selectors, and returns the list of AwsIamServerCertificates that match those selectors.
func (c *FakeAwsIamServerCertificates) List(opts v1.ListOptions) (result *aws_v1.AwsIamServerCertificateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsiamservercertificatesResource, awsiamservercertificatesKind, c.ns, opts), &aws_v1.AwsIamServerCertificateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsIamServerCertificateList{}
	for _, item := range obj.(*aws_v1.AwsIamServerCertificateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsIamServerCertificates.
func (c *FakeAwsIamServerCertificates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsiamservercertificatesResource, c.ns, opts))

}

// Create takes the representation of a awsIamServerCertificate and creates it.  Returns the server's representation of the awsIamServerCertificate, and an error, if there is any.
func (c *FakeAwsIamServerCertificates) Create(awsIamServerCertificate *aws_v1.AwsIamServerCertificate) (result *aws_v1.AwsIamServerCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsiamservercertificatesResource, c.ns, awsIamServerCertificate), &aws_v1.AwsIamServerCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsIamServerCertificate), err
}

// Update takes the representation of a awsIamServerCertificate and updates it. Returns the server's representation of the awsIamServerCertificate, and an error, if there is any.
func (c *FakeAwsIamServerCertificates) Update(awsIamServerCertificate *aws_v1.AwsIamServerCertificate) (result *aws_v1.AwsIamServerCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsiamservercertificatesResource, c.ns, awsIamServerCertificate), &aws_v1.AwsIamServerCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsIamServerCertificate), err
}

// Delete takes name of the awsIamServerCertificate and deletes it. Returns an error if one occurs.
func (c *FakeAwsIamServerCertificates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsiamservercertificatesResource, c.ns, name), &aws_v1.AwsIamServerCertificate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsIamServerCertificates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsiamservercertificatesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsIamServerCertificateList{})
	return err
}

// Patch applies the patch and returns the patched awsIamServerCertificate.
func (c *FakeAwsIamServerCertificates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsIamServerCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsiamservercertificatesResource, c.ns, name, data, subresources...), &aws_v1.AwsIamServerCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsIamServerCertificate), err
}
