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

// FakeAwsCloudwatchDashboards implements AwsCloudwatchDashboardInterface
type FakeAwsCloudwatchDashboards struct {
	Fake *FakeTrussleV1
	ns   string
}

var awscloudwatchdashboardsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awscloudwatchdashboards"}

var awscloudwatchdashboardsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsCloudwatchDashboard"}

// Get takes name of the awsCloudwatchDashboard, and returns the corresponding awsCloudwatchDashboard object, and an error if there is any.
func (c *FakeAwsCloudwatchDashboards) Get(name string, options v1.GetOptions) (result *aws_v1.AwsCloudwatchDashboard, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awscloudwatchdashboardsResource, c.ns, name), &aws_v1.AwsCloudwatchDashboard{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCloudwatchDashboard), err
}

// List takes label and field selectors, and returns the list of AwsCloudwatchDashboards that match those selectors.
func (c *FakeAwsCloudwatchDashboards) List(opts v1.ListOptions) (result *aws_v1.AwsCloudwatchDashboardList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awscloudwatchdashboardsResource, awscloudwatchdashboardsKind, c.ns, opts), &aws_v1.AwsCloudwatchDashboardList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsCloudwatchDashboardList{}
	for _, item := range obj.(*aws_v1.AwsCloudwatchDashboardList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsCloudwatchDashboards.
func (c *FakeAwsCloudwatchDashboards) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awscloudwatchdashboardsResource, c.ns, opts))

}

// Create takes the representation of a awsCloudwatchDashboard and creates it.  Returns the server's representation of the awsCloudwatchDashboard, and an error, if there is any.
func (c *FakeAwsCloudwatchDashboards) Create(awsCloudwatchDashboard *aws_v1.AwsCloudwatchDashboard) (result *aws_v1.AwsCloudwatchDashboard, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awscloudwatchdashboardsResource, c.ns, awsCloudwatchDashboard), &aws_v1.AwsCloudwatchDashboard{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCloudwatchDashboard), err
}

// Update takes the representation of a awsCloudwatchDashboard and updates it. Returns the server's representation of the awsCloudwatchDashboard, and an error, if there is any.
func (c *FakeAwsCloudwatchDashboards) Update(awsCloudwatchDashboard *aws_v1.AwsCloudwatchDashboard) (result *aws_v1.AwsCloudwatchDashboard, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awscloudwatchdashboardsResource, c.ns, awsCloudwatchDashboard), &aws_v1.AwsCloudwatchDashboard{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCloudwatchDashboard), err
}

// Delete takes name of the awsCloudwatchDashboard and deletes it. Returns an error if one occurs.
func (c *FakeAwsCloudwatchDashboards) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awscloudwatchdashboardsResource, c.ns, name), &aws_v1.AwsCloudwatchDashboard{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsCloudwatchDashboards) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awscloudwatchdashboardsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsCloudwatchDashboardList{})
	return err
}

// Patch applies the patch and returns the patched awsCloudwatchDashboard.
func (c *FakeAwsCloudwatchDashboards) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsCloudwatchDashboard, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awscloudwatchdashboardsResource, c.ns, name, data, subresources...), &aws_v1.AwsCloudwatchDashboard{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCloudwatchDashboard), err
}
