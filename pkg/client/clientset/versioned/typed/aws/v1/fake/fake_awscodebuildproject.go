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

// FakeAwsCodebuildProjects implements AwsCodebuildProjectInterface
type FakeAwsCodebuildProjects struct {
	Fake *FakeTrussleV1
	ns   string
}

var awscodebuildprojectsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awscodebuildprojects"}

var awscodebuildprojectsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsCodebuildProject"}

// Get takes name of the awsCodebuildProject, and returns the corresponding awsCodebuildProject object, and an error if there is any.
func (c *FakeAwsCodebuildProjects) Get(name string, options v1.GetOptions) (result *aws_v1.AwsCodebuildProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awscodebuildprojectsResource, c.ns, name), &aws_v1.AwsCodebuildProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCodebuildProject), err
}

// List takes label and field selectors, and returns the list of AwsCodebuildProjects that match those selectors.
func (c *FakeAwsCodebuildProjects) List(opts v1.ListOptions) (result *aws_v1.AwsCodebuildProjectList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awscodebuildprojectsResource, awscodebuildprojectsKind, c.ns, opts), &aws_v1.AwsCodebuildProjectList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsCodebuildProjectList{}
	for _, item := range obj.(*aws_v1.AwsCodebuildProjectList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsCodebuildProjects.
func (c *FakeAwsCodebuildProjects) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awscodebuildprojectsResource, c.ns, opts))

}

// Create takes the representation of a awsCodebuildProject and creates it.  Returns the server's representation of the awsCodebuildProject, and an error, if there is any.
func (c *FakeAwsCodebuildProjects) Create(awsCodebuildProject *aws_v1.AwsCodebuildProject) (result *aws_v1.AwsCodebuildProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awscodebuildprojectsResource, c.ns, awsCodebuildProject), &aws_v1.AwsCodebuildProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCodebuildProject), err
}

// Update takes the representation of a awsCodebuildProject and updates it. Returns the server's representation of the awsCodebuildProject, and an error, if there is any.
func (c *FakeAwsCodebuildProjects) Update(awsCodebuildProject *aws_v1.AwsCodebuildProject) (result *aws_v1.AwsCodebuildProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awscodebuildprojectsResource, c.ns, awsCodebuildProject), &aws_v1.AwsCodebuildProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCodebuildProject), err
}

// Delete takes name of the awsCodebuildProject and deletes it. Returns an error if one occurs.
func (c *FakeAwsCodebuildProjects) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awscodebuildprojectsResource, c.ns, name), &aws_v1.AwsCodebuildProject{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsCodebuildProjects) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awscodebuildprojectsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsCodebuildProjectList{})
	return err
}

// Patch applies the patch and returns the patched awsCodebuildProject.
func (c *FakeAwsCodebuildProjects) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsCodebuildProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awscodebuildprojectsResource, c.ns, name, data, subresources...), &aws_v1.AwsCodebuildProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCodebuildProject), err
}
