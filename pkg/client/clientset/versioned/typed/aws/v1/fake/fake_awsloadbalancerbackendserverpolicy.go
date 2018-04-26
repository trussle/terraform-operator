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

// FakeAwsLoadBalancerBackendServerPolicies implements AwsLoadBalancerBackendServerPolicyInterface
type FakeAwsLoadBalancerBackendServerPolicies struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsloadbalancerbackendserverpoliciesResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsloadbalancerbackendserverpolicies"}

var awsloadbalancerbackendserverpoliciesKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsLoadBalancerBackendServerPolicy"}

// Get takes name of the awsLoadBalancerBackendServerPolicy, and returns the corresponding awsLoadBalancerBackendServerPolicy object, and an error if there is any.
func (c *FakeAwsLoadBalancerBackendServerPolicies) Get(name string, options v1.GetOptions) (result *aws_v1.AwsLoadBalancerBackendServerPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsloadbalancerbackendserverpoliciesResource, c.ns, name), &aws_v1.AwsLoadBalancerBackendServerPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLoadBalancerBackendServerPolicy), err
}

// List takes label and field selectors, and returns the list of AwsLoadBalancerBackendServerPolicies that match those selectors.
func (c *FakeAwsLoadBalancerBackendServerPolicies) List(opts v1.ListOptions) (result *aws_v1.AwsLoadBalancerBackendServerPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsloadbalancerbackendserverpoliciesResource, awsloadbalancerbackendserverpoliciesKind, c.ns, opts), &aws_v1.AwsLoadBalancerBackendServerPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsLoadBalancerBackendServerPolicyList{}
	for _, item := range obj.(*aws_v1.AwsLoadBalancerBackendServerPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsLoadBalancerBackendServerPolicies.
func (c *FakeAwsLoadBalancerBackendServerPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsloadbalancerbackendserverpoliciesResource, c.ns, opts))

}

// Create takes the representation of a awsLoadBalancerBackendServerPolicy and creates it.  Returns the server's representation of the awsLoadBalancerBackendServerPolicy, and an error, if there is any.
func (c *FakeAwsLoadBalancerBackendServerPolicies) Create(awsLoadBalancerBackendServerPolicy *aws_v1.AwsLoadBalancerBackendServerPolicy) (result *aws_v1.AwsLoadBalancerBackendServerPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsloadbalancerbackendserverpoliciesResource, c.ns, awsLoadBalancerBackendServerPolicy), &aws_v1.AwsLoadBalancerBackendServerPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLoadBalancerBackendServerPolicy), err
}

// Update takes the representation of a awsLoadBalancerBackendServerPolicy and updates it. Returns the server's representation of the awsLoadBalancerBackendServerPolicy, and an error, if there is any.
func (c *FakeAwsLoadBalancerBackendServerPolicies) Update(awsLoadBalancerBackendServerPolicy *aws_v1.AwsLoadBalancerBackendServerPolicy) (result *aws_v1.AwsLoadBalancerBackendServerPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsloadbalancerbackendserverpoliciesResource, c.ns, awsLoadBalancerBackendServerPolicy), &aws_v1.AwsLoadBalancerBackendServerPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLoadBalancerBackendServerPolicy), err
}

// Delete takes name of the awsLoadBalancerBackendServerPolicy and deletes it. Returns an error if one occurs.
func (c *FakeAwsLoadBalancerBackendServerPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsloadbalancerbackendserverpoliciesResource, c.ns, name), &aws_v1.AwsLoadBalancerBackendServerPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsLoadBalancerBackendServerPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsloadbalancerbackendserverpoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsLoadBalancerBackendServerPolicyList{})
	return err
}

// Patch applies the patch and returns the patched awsLoadBalancerBackendServerPolicy.
func (c *FakeAwsLoadBalancerBackendServerPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsLoadBalancerBackendServerPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsloadbalancerbackendserverpoliciesResource, c.ns, name, data, subresources...), &aws_v1.AwsLoadBalancerBackendServerPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLoadBalancerBackendServerPolicy), err
}
