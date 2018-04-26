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

// FakeAwsBatchJobQueues implements AwsBatchJobQueueInterface
type FakeAwsBatchJobQueues struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsbatchjobqueuesResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsbatchjobqueues"}

var awsbatchjobqueuesKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsBatchJobQueue"}

// Get takes name of the awsBatchJobQueue, and returns the corresponding awsBatchJobQueue object, and an error if there is any.
func (c *FakeAwsBatchJobQueues) Get(name string, options v1.GetOptions) (result *aws_v1.AwsBatchJobQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsbatchjobqueuesResource, c.ns, name), &aws_v1.AwsBatchJobQueue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsBatchJobQueue), err
}

// List takes label and field selectors, and returns the list of AwsBatchJobQueues that match those selectors.
func (c *FakeAwsBatchJobQueues) List(opts v1.ListOptions) (result *aws_v1.AwsBatchJobQueueList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsbatchjobqueuesResource, awsbatchjobqueuesKind, c.ns, opts), &aws_v1.AwsBatchJobQueueList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsBatchJobQueueList{}
	for _, item := range obj.(*aws_v1.AwsBatchJobQueueList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsBatchJobQueues.
func (c *FakeAwsBatchJobQueues) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsbatchjobqueuesResource, c.ns, opts))

}

// Create takes the representation of a awsBatchJobQueue and creates it.  Returns the server's representation of the awsBatchJobQueue, and an error, if there is any.
func (c *FakeAwsBatchJobQueues) Create(awsBatchJobQueue *aws_v1.AwsBatchJobQueue) (result *aws_v1.AwsBatchJobQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsbatchjobqueuesResource, c.ns, awsBatchJobQueue), &aws_v1.AwsBatchJobQueue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsBatchJobQueue), err
}

// Update takes the representation of a awsBatchJobQueue and updates it. Returns the server's representation of the awsBatchJobQueue, and an error, if there is any.
func (c *FakeAwsBatchJobQueues) Update(awsBatchJobQueue *aws_v1.AwsBatchJobQueue) (result *aws_v1.AwsBatchJobQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsbatchjobqueuesResource, c.ns, awsBatchJobQueue), &aws_v1.AwsBatchJobQueue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsBatchJobQueue), err
}

// Delete takes name of the awsBatchJobQueue and deletes it. Returns an error if one occurs.
func (c *FakeAwsBatchJobQueues) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsbatchjobqueuesResource, c.ns, name), &aws_v1.AwsBatchJobQueue{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsBatchJobQueues) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsbatchjobqueuesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsBatchJobQueueList{})
	return err
}

// Patch applies the patch and returns the patched awsBatchJobQueue.
func (c *FakeAwsBatchJobQueues) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsBatchJobQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsbatchjobqueuesResource, c.ns, name, data, subresources...), &aws_v1.AwsBatchJobQueue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsBatchJobQueue), err
}
