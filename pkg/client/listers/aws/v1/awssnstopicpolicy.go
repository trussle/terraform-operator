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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/trussle/terraform-operator/pkg/apis/aws/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AwsSnsTopicPolicyLister helps list AwsSnsTopicPolicies.
type AwsSnsTopicPolicyLister interface {
	// List lists all AwsSnsTopicPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsSnsTopicPolicy, err error)
	// AwsSnsTopicPolicies returns an object that can list and get AwsSnsTopicPolicies.
	AwsSnsTopicPolicies(namespace string) AwsSnsTopicPolicyNamespaceLister
	AwsSnsTopicPolicyListerExpansion
}

// awsSnsTopicPolicyLister implements the AwsSnsTopicPolicyLister interface.
type awsSnsTopicPolicyLister struct {
	indexer cache.Indexer
}

// NewAwsSnsTopicPolicyLister returns a new AwsSnsTopicPolicyLister.
func NewAwsSnsTopicPolicyLister(indexer cache.Indexer) AwsSnsTopicPolicyLister {
	return &awsSnsTopicPolicyLister{indexer: indexer}
}

// List lists all AwsSnsTopicPolicies in the indexer.
func (s *awsSnsTopicPolicyLister) List(selector labels.Selector) (ret []*v1.AwsSnsTopicPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsSnsTopicPolicy))
	})
	return ret, err
}

// AwsSnsTopicPolicies returns an object that can list and get AwsSnsTopicPolicies.
func (s *awsSnsTopicPolicyLister) AwsSnsTopicPolicies(namespace string) AwsSnsTopicPolicyNamespaceLister {
	return awsSnsTopicPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsSnsTopicPolicyNamespaceLister helps list and get AwsSnsTopicPolicies.
type AwsSnsTopicPolicyNamespaceLister interface {
	// List lists all AwsSnsTopicPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsSnsTopicPolicy, err error)
	// Get retrieves the AwsSnsTopicPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsSnsTopicPolicy, error)
	AwsSnsTopicPolicyNamespaceListerExpansion
}

// awsSnsTopicPolicyNamespaceLister implements the AwsSnsTopicPolicyNamespaceLister
// interface.
type awsSnsTopicPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsSnsTopicPolicies in the indexer for a given namespace.
func (s awsSnsTopicPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsSnsTopicPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsSnsTopicPolicy))
	})
	return ret, err
}

// Get retrieves the AwsSnsTopicPolicy from the indexer for a given namespace and name.
func (s awsSnsTopicPolicyNamespaceLister) Get(name string) (*v1.AwsSnsTopicPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awssnstopicpolicy"), name)
	}
	return obj.(*v1.AwsSnsTopicPolicy), nil
}
