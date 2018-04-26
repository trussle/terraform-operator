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

// AwsIotTopicRuleLister helps list AwsIotTopicRules.
type AwsIotTopicRuleLister interface {
	// List lists all AwsIotTopicRules in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsIotTopicRule, err error)
	// AwsIotTopicRules returns an object that can list and get AwsIotTopicRules.
	AwsIotTopicRules(namespace string) AwsIotTopicRuleNamespaceLister
	AwsIotTopicRuleListerExpansion
}

// awsIotTopicRuleLister implements the AwsIotTopicRuleLister interface.
type awsIotTopicRuleLister struct {
	indexer cache.Indexer
}

// NewAwsIotTopicRuleLister returns a new AwsIotTopicRuleLister.
func NewAwsIotTopicRuleLister(indexer cache.Indexer) AwsIotTopicRuleLister {
	return &awsIotTopicRuleLister{indexer: indexer}
}

// List lists all AwsIotTopicRules in the indexer.
func (s *awsIotTopicRuleLister) List(selector labels.Selector) (ret []*v1.AwsIotTopicRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsIotTopicRule))
	})
	return ret, err
}

// AwsIotTopicRules returns an object that can list and get AwsIotTopicRules.
func (s *awsIotTopicRuleLister) AwsIotTopicRules(namespace string) AwsIotTopicRuleNamespaceLister {
	return awsIotTopicRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsIotTopicRuleNamespaceLister helps list and get AwsIotTopicRules.
type AwsIotTopicRuleNamespaceLister interface {
	// List lists all AwsIotTopicRules in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsIotTopicRule, err error)
	// Get retrieves the AwsIotTopicRule from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsIotTopicRule, error)
	AwsIotTopicRuleNamespaceListerExpansion
}

// awsIotTopicRuleNamespaceLister implements the AwsIotTopicRuleNamespaceLister
// interface.
type awsIotTopicRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsIotTopicRules in the indexer for a given namespace.
func (s awsIotTopicRuleNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsIotTopicRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsIotTopicRule))
	})
	return ret, err
}

// Get retrieves the AwsIotTopicRule from the indexer for a given namespace and name.
func (s awsIotTopicRuleNamespaceLister) Get(name string) (*v1.AwsIotTopicRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsiottopicrule"), name)
	}
	return obj.(*v1.AwsIotTopicRule), nil
}
