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

// AwsSesDomainDkimLister helps list AwsSesDomainDkims.
type AwsSesDomainDkimLister interface {
	// List lists all AwsSesDomainDkims in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsSesDomainDkim, err error)
	// AwsSesDomainDkims returns an object that can list and get AwsSesDomainDkims.
	AwsSesDomainDkims(namespace string) AwsSesDomainDkimNamespaceLister
	AwsSesDomainDkimListerExpansion
}

// awsSesDomainDkimLister implements the AwsSesDomainDkimLister interface.
type awsSesDomainDkimLister struct {
	indexer cache.Indexer
}

// NewAwsSesDomainDkimLister returns a new AwsSesDomainDkimLister.
func NewAwsSesDomainDkimLister(indexer cache.Indexer) AwsSesDomainDkimLister {
	return &awsSesDomainDkimLister{indexer: indexer}
}

// List lists all AwsSesDomainDkims in the indexer.
func (s *awsSesDomainDkimLister) List(selector labels.Selector) (ret []*v1.AwsSesDomainDkim, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsSesDomainDkim))
	})
	return ret, err
}

// AwsSesDomainDkims returns an object that can list and get AwsSesDomainDkims.
func (s *awsSesDomainDkimLister) AwsSesDomainDkims(namespace string) AwsSesDomainDkimNamespaceLister {
	return awsSesDomainDkimNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsSesDomainDkimNamespaceLister helps list and get AwsSesDomainDkims.
type AwsSesDomainDkimNamespaceLister interface {
	// List lists all AwsSesDomainDkims in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsSesDomainDkim, err error)
	// Get retrieves the AwsSesDomainDkim from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsSesDomainDkim, error)
	AwsSesDomainDkimNamespaceListerExpansion
}

// awsSesDomainDkimNamespaceLister implements the AwsSesDomainDkimNamespaceLister
// interface.
type awsSesDomainDkimNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsSesDomainDkims in the indexer for a given namespace.
func (s awsSesDomainDkimNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsSesDomainDkim, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsSesDomainDkim))
	})
	return ret, err
}

// Get retrieves the AwsSesDomainDkim from the indexer for a given namespace and name.
func (s awsSesDomainDkimNamespaceLister) Get(name string) (*v1.AwsSesDomainDkim, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awssesdomaindkim"), name)
	}
	return obj.(*v1.AwsSesDomainDkim), nil
}