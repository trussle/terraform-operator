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

// FakeAwsCloudwatchEventPermissions implements AwsCloudwatchEventPermissionInterface
type FakeAwsCloudwatchEventPermissions struct {
	Fake *FakeTrussleV1
	ns   string
}

var awscloudwatcheventpermissionsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awscloudwatcheventpermissions"}

var awscloudwatcheventpermissionsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsCloudwatchEventPermission"}

// Get takes name of the awsCloudwatchEventPermission, and returns the corresponding awsCloudwatchEventPermission object, and an error if there is any.
func (c *FakeAwsCloudwatchEventPermissions) Get(name string, options v1.GetOptions) (result *aws_v1.AwsCloudwatchEventPermission, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awscloudwatcheventpermissionsResource, c.ns, name), &aws_v1.AwsCloudwatchEventPermission{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCloudwatchEventPermission), err
}

// List takes label and field selectors, and returns the list of AwsCloudwatchEventPermissions that match those selectors.
func (c *FakeAwsCloudwatchEventPermissions) List(opts v1.ListOptions) (result *aws_v1.AwsCloudwatchEventPermissionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awscloudwatcheventpermissionsResource, awscloudwatcheventpermissionsKind, c.ns, opts), &aws_v1.AwsCloudwatchEventPermissionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsCloudwatchEventPermissionList{}
	for _, item := range obj.(*aws_v1.AwsCloudwatchEventPermissionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsCloudwatchEventPermissions.
func (c *FakeAwsCloudwatchEventPermissions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awscloudwatcheventpermissionsResource, c.ns, opts))

}

// Create takes the representation of a awsCloudwatchEventPermission and creates it.  Returns the server's representation of the awsCloudwatchEventPermission, and an error, if there is any.
func (c *FakeAwsCloudwatchEventPermissions) Create(awsCloudwatchEventPermission *aws_v1.AwsCloudwatchEventPermission) (result *aws_v1.AwsCloudwatchEventPermission, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awscloudwatcheventpermissionsResource, c.ns, awsCloudwatchEventPermission), &aws_v1.AwsCloudwatchEventPermission{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCloudwatchEventPermission), err
}

// Update takes the representation of a awsCloudwatchEventPermission and updates it. Returns the server's representation of the awsCloudwatchEventPermission, and an error, if there is any.
func (c *FakeAwsCloudwatchEventPermissions) Update(awsCloudwatchEventPermission *aws_v1.AwsCloudwatchEventPermission) (result *aws_v1.AwsCloudwatchEventPermission, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awscloudwatcheventpermissionsResource, c.ns, awsCloudwatchEventPermission), &aws_v1.AwsCloudwatchEventPermission{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCloudwatchEventPermission), err
}

// Delete takes name of the awsCloudwatchEventPermission and deletes it. Returns an error if one occurs.
func (c *FakeAwsCloudwatchEventPermissions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awscloudwatcheventpermissionsResource, c.ns, name), &aws_v1.AwsCloudwatchEventPermission{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsCloudwatchEventPermissions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awscloudwatcheventpermissionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsCloudwatchEventPermissionList{})
	return err
}

// Patch applies the patch and returns the patched awsCloudwatchEventPermission.
func (c *FakeAwsCloudwatchEventPermissions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsCloudwatchEventPermission, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awscloudwatcheventpermissionsResource, c.ns, name, data, subresources...), &aws_v1.AwsCloudwatchEventPermission{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCloudwatchEventPermission), err
}
