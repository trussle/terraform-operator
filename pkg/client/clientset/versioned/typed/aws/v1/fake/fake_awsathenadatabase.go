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

// FakeAwsAthenaDatabases implements AwsAthenaDatabaseInterface
type FakeAwsAthenaDatabases struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsathenadatabasesResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsathenadatabases"}

var awsathenadatabasesKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsAthenaDatabase"}

// Get takes name of the awsAthenaDatabase, and returns the corresponding awsAthenaDatabase object, and an error if there is any.
func (c *FakeAwsAthenaDatabases) Get(name string, options v1.GetOptions) (result *aws_v1.AwsAthenaDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsathenadatabasesResource, c.ns, name), &aws_v1.AwsAthenaDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsAthenaDatabase), err
}

// List takes label and field selectors, and returns the list of AwsAthenaDatabases that match those selectors.
func (c *FakeAwsAthenaDatabases) List(opts v1.ListOptions) (result *aws_v1.AwsAthenaDatabaseList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsathenadatabasesResource, awsathenadatabasesKind, c.ns, opts), &aws_v1.AwsAthenaDatabaseList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsAthenaDatabaseList{}
	for _, item := range obj.(*aws_v1.AwsAthenaDatabaseList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsAthenaDatabases.
func (c *FakeAwsAthenaDatabases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsathenadatabasesResource, c.ns, opts))

}

// Create takes the representation of a awsAthenaDatabase and creates it.  Returns the server's representation of the awsAthenaDatabase, and an error, if there is any.
func (c *FakeAwsAthenaDatabases) Create(awsAthenaDatabase *aws_v1.AwsAthenaDatabase) (result *aws_v1.AwsAthenaDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsathenadatabasesResource, c.ns, awsAthenaDatabase), &aws_v1.AwsAthenaDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsAthenaDatabase), err
}

// Update takes the representation of a awsAthenaDatabase and updates it. Returns the server's representation of the awsAthenaDatabase, and an error, if there is any.
func (c *FakeAwsAthenaDatabases) Update(awsAthenaDatabase *aws_v1.AwsAthenaDatabase) (result *aws_v1.AwsAthenaDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsathenadatabasesResource, c.ns, awsAthenaDatabase), &aws_v1.AwsAthenaDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsAthenaDatabase), err
}

// Delete takes name of the awsAthenaDatabase and deletes it. Returns an error if one occurs.
func (c *FakeAwsAthenaDatabases) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsathenadatabasesResource, c.ns, name), &aws_v1.AwsAthenaDatabase{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsAthenaDatabases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsathenadatabasesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsAthenaDatabaseList{})
	return err
}

// Patch applies the patch and returns the patched awsAthenaDatabase.
func (c *FakeAwsAthenaDatabases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsAthenaDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsathenadatabasesResource, c.ns, name, data, subresources...), &aws_v1.AwsAthenaDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsAthenaDatabase), err
}
