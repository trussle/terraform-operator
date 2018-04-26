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

// FakeAwsAlbListenerCertificates implements AwsAlbListenerCertificateInterface
type FakeAwsAlbListenerCertificates struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsalblistenercertificatesResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsalblistenercertificates"}

var awsalblistenercertificatesKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsAlbListenerCertificate"}

// Get takes name of the awsAlbListenerCertificate, and returns the corresponding awsAlbListenerCertificate object, and an error if there is any.
func (c *FakeAwsAlbListenerCertificates) Get(name string, options v1.GetOptions) (result *aws_v1.AwsAlbListenerCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsalblistenercertificatesResource, c.ns, name), &aws_v1.AwsAlbListenerCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsAlbListenerCertificate), err
}

// List takes label and field selectors, and returns the list of AwsAlbListenerCertificates that match those selectors.
func (c *FakeAwsAlbListenerCertificates) List(opts v1.ListOptions) (result *aws_v1.AwsAlbListenerCertificateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsalblistenercertificatesResource, awsalblistenercertificatesKind, c.ns, opts), &aws_v1.AwsAlbListenerCertificateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsAlbListenerCertificateList{}
	for _, item := range obj.(*aws_v1.AwsAlbListenerCertificateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsAlbListenerCertificates.
func (c *FakeAwsAlbListenerCertificates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsalblistenercertificatesResource, c.ns, opts))

}

// Create takes the representation of a awsAlbListenerCertificate and creates it.  Returns the server's representation of the awsAlbListenerCertificate, and an error, if there is any.
func (c *FakeAwsAlbListenerCertificates) Create(awsAlbListenerCertificate *aws_v1.AwsAlbListenerCertificate) (result *aws_v1.AwsAlbListenerCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsalblistenercertificatesResource, c.ns, awsAlbListenerCertificate), &aws_v1.AwsAlbListenerCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsAlbListenerCertificate), err
}

// Update takes the representation of a awsAlbListenerCertificate and updates it. Returns the server's representation of the awsAlbListenerCertificate, and an error, if there is any.
func (c *FakeAwsAlbListenerCertificates) Update(awsAlbListenerCertificate *aws_v1.AwsAlbListenerCertificate) (result *aws_v1.AwsAlbListenerCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsalblistenercertificatesResource, c.ns, awsAlbListenerCertificate), &aws_v1.AwsAlbListenerCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsAlbListenerCertificate), err
}

// Delete takes name of the awsAlbListenerCertificate and deletes it. Returns an error if one occurs.
func (c *FakeAwsAlbListenerCertificates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsalblistenercertificatesResource, c.ns, name), &aws_v1.AwsAlbListenerCertificate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsAlbListenerCertificates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsalblistenercertificatesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsAlbListenerCertificateList{})
	return err
}

// Patch applies the patch and returns the patched awsAlbListenerCertificate.
func (c *FakeAwsAlbListenerCertificates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsAlbListenerCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsalblistenercertificatesResource, c.ns, name, data, subresources...), &aws_v1.AwsAlbListenerCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsAlbListenerCertificate), err
}
