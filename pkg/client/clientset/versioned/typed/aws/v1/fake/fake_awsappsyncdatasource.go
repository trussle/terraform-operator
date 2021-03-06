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

// FakeAwsAppsyncDatasources implements AwsAppsyncDatasourceInterface
type FakeAwsAppsyncDatasources struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsappsyncdatasourcesResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsappsyncdatasources"}

var awsappsyncdatasourcesKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsAppsyncDatasource"}

// Get takes name of the awsAppsyncDatasource, and returns the corresponding awsAppsyncDatasource object, and an error if there is any.
func (c *FakeAwsAppsyncDatasources) Get(name string, options v1.GetOptions) (result *aws_v1.AwsAppsyncDatasource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsappsyncdatasourcesResource, c.ns, name), &aws_v1.AwsAppsyncDatasource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsAppsyncDatasource), err
}

// List takes label and field selectors, and returns the list of AwsAppsyncDatasources that match those selectors.
func (c *FakeAwsAppsyncDatasources) List(opts v1.ListOptions) (result *aws_v1.AwsAppsyncDatasourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsappsyncdatasourcesResource, awsappsyncdatasourcesKind, c.ns, opts), &aws_v1.AwsAppsyncDatasourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsAppsyncDatasourceList{}
	for _, item := range obj.(*aws_v1.AwsAppsyncDatasourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsAppsyncDatasources.
func (c *FakeAwsAppsyncDatasources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsappsyncdatasourcesResource, c.ns, opts))

}

// Create takes the representation of a awsAppsyncDatasource and creates it.  Returns the server's representation of the awsAppsyncDatasource, and an error, if there is any.
func (c *FakeAwsAppsyncDatasources) Create(awsAppsyncDatasource *aws_v1.AwsAppsyncDatasource) (result *aws_v1.AwsAppsyncDatasource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsappsyncdatasourcesResource, c.ns, awsAppsyncDatasource), &aws_v1.AwsAppsyncDatasource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsAppsyncDatasource), err
}

// Update takes the representation of a awsAppsyncDatasource and updates it. Returns the server's representation of the awsAppsyncDatasource, and an error, if there is any.
func (c *FakeAwsAppsyncDatasources) Update(awsAppsyncDatasource *aws_v1.AwsAppsyncDatasource) (result *aws_v1.AwsAppsyncDatasource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsappsyncdatasourcesResource, c.ns, awsAppsyncDatasource), &aws_v1.AwsAppsyncDatasource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsAppsyncDatasource), err
}

// Delete takes name of the awsAppsyncDatasource and deletes it. Returns an error if one occurs.
func (c *FakeAwsAppsyncDatasources) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsappsyncdatasourcesResource, c.ns, name), &aws_v1.AwsAppsyncDatasource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsAppsyncDatasources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsappsyncdatasourcesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsAppsyncDatasourceList{})
	return err
}

// Patch applies the patch and returns the patched awsAppsyncDatasource.
func (c *FakeAwsAppsyncDatasources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsAppsyncDatasource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsappsyncdatasourcesResource, c.ns, name, data, subresources...), &aws_v1.AwsAppsyncDatasource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsAppsyncDatasource), err
}
