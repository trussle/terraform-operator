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

// AwsLbSslNegotiationPolicyLister helps list AwsLbSslNegotiationPolicies.
type AwsLbSslNegotiationPolicyLister interface {
	// List lists all AwsLbSslNegotiationPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsLbSslNegotiationPolicy, err error)
	// AwsLbSslNegotiationPolicies returns an object that can list and get AwsLbSslNegotiationPolicies.
	AwsLbSslNegotiationPolicies(namespace string) AwsLbSslNegotiationPolicyNamespaceLister
	AwsLbSslNegotiationPolicyListerExpansion
}

// awsLbSslNegotiationPolicyLister implements the AwsLbSslNegotiationPolicyLister interface.
type awsLbSslNegotiationPolicyLister struct {
	indexer cache.Indexer
}

// NewAwsLbSslNegotiationPolicyLister returns a new AwsLbSslNegotiationPolicyLister.
func NewAwsLbSslNegotiationPolicyLister(indexer cache.Indexer) AwsLbSslNegotiationPolicyLister {
	return &awsLbSslNegotiationPolicyLister{indexer: indexer}
}

// List lists all AwsLbSslNegotiationPolicies in the indexer.
func (s *awsLbSslNegotiationPolicyLister) List(selector labels.Selector) (ret []*v1.AwsLbSslNegotiationPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsLbSslNegotiationPolicy))
	})
	return ret, err
}

// AwsLbSslNegotiationPolicies returns an object that can list and get AwsLbSslNegotiationPolicies.
func (s *awsLbSslNegotiationPolicyLister) AwsLbSslNegotiationPolicies(namespace string) AwsLbSslNegotiationPolicyNamespaceLister {
	return awsLbSslNegotiationPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsLbSslNegotiationPolicyNamespaceLister helps list and get AwsLbSslNegotiationPolicies.
type AwsLbSslNegotiationPolicyNamespaceLister interface {
	// List lists all AwsLbSslNegotiationPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsLbSslNegotiationPolicy, err error)
	// Get retrieves the AwsLbSslNegotiationPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsLbSslNegotiationPolicy, error)
	AwsLbSslNegotiationPolicyNamespaceListerExpansion
}

// awsLbSslNegotiationPolicyNamespaceLister implements the AwsLbSslNegotiationPolicyNamespaceLister
// interface.
type awsLbSslNegotiationPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsLbSslNegotiationPolicies in the indexer for a given namespace.
func (s awsLbSslNegotiationPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsLbSslNegotiationPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsLbSslNegotiationPolicy))
	})
	return ret, err
}

// Get retrieves the AwsLbSslNegotiationPolicy from the indexer for a given namespace and name.
func (s awsLbSslNegotiationPolicyNamespaceLister) Get(name string) (*v1.AwsLbSslNegotiationPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awslbsslnegotiationpolicy"), name)
	}
	return obj.(*v1.AwsLbSslNegotiationPolicy), nil
}
