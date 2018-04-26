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

// FakeAwsIotCertificates implements AwsIotCertificateInterface
type FakeAwsIotCertificates struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsiotcertificatesResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsiotcertificates"}

var awsiotcertificatesKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsIotCertificate"}

// Get takes name of the awsIotCertificate, and returns the corresponding awsIotCertificate object, and an error if there is any.
func (c *FakeAwsIotCertificates) Get(name string, options v1.GetOptions) (result *aws_v1.AwsIotCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsiotcertificatesResource, c.ns, name), &aws_v1.AwsIotCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsIotCertificate), err
}

// List takes label and field selectors, and returns the list of AwsIotCertificates that match those selectors.
func (c *FakeAwsIotCertificates) List(opts v1.ListOptions) (result *aws_v1.AwsIotCertificateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsiotcertificatesResource, awsiotcertificatesKind, c.ns, opts), &aws_v1.AwsIotCertificateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsIotCertificateList{}
	for _, item := range obj.(*aws_v1.AwsIotCertificateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsIotCertificates.
func (c *FakeAwsIotCertificates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsiotcertificatesResource, c.ns, opts))

}

// Create takes the representation of a awsIotCertificate and creates it.  Returns the server's representation of the awsIotCertificate, and an error, if there is any.
func (c *FakeAwsIotCertificates) Create(awsIotCertificate *aws_v1.AwsIotCertificate) (result *aws_v1.AwsIotCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsiotcertificatesResource, c.ns, awsIotCertificate), &aws_v1.AwsIotCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsIotCertificate), err
}

// Update takes the representation of a awsIotCertificate and updates it. Returns the server's representation of the awsIotCertificate, and an error, if there is any.
func (c *FakeAwsIotCertificates) Update(awsIotCertificate *aws_v1.AwsIotCertificate) (result *aws_v1.AwsIotCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsiotcertificatesResource, c.ns, awsIotCertificate), &aws_v1.AwsIotCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsIotCertificate), err
}

// Delete takes name of the awsIotCertificate and deletes it. Returns an error if one occurs.
func (c *FakeAwsIotCertificates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsiotcertificatesResource, c.ns, name), &aws_v1.AwsIotCertificate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsIotCertificates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsiotcertificatesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsIotCertificateList{})
	return err
}

// Patch applies the patch and returns the patched awsIotCertificate.
func (c *FakeAwsIotCertificates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsIotCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsiotcertificatesResource, c.ns, name, data, subresources...), &aws_v1.AwsIotCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsIotCertificate), err
}
