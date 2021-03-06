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

// AwsKeyPairLister helps list AwsKeyPairs.
type AwsKeyPairLister interface {
	// List lists all AwsKeyPairs in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsKeyPair, err error)
	// AwsKeyPairs returns an object that can list and get AwsKeyPairs.
	AwsKeyPairs(namespace string) AwsKeyPairNamespaceLister
	AwsKeyPairListerExpansion
}

// awsKeyPairLister implements the AwsKeyPairLister interface.
type awsKeyPairLister struct {
	indexer cache.Indexer
}

// NewAwsKeyPairLister returns a new AwsKeyPairLister.
func NewAwsKeyPairLister(indexer cache.Indexer) AwsKeyPairLister {
	return &awsKeyPairLister{indexer: indexer}
}

// List lists all AwsKeyPairs in the indexer.
func (s *awsKeyPairLister) List(selector labels.Selector) (ret []*v1.AwsKeyPair, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsKeyPair))
	})
	return ret, err
}

// AwsKeyPairs returns an object that can list and get AwsKeyPairs.
func (s *awsKeyPairLister) AwsKeyPairs(namespace string) AwsKeyPairNamespaceLister {
	return awsKeyPairNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsKeyPairNamespaceLister helps list and get AwsKeyPairs.
type AwsKeyPairNamespaceLister interface {
	// List lists all AwsKeyPairs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsKeyPair, err error)
	// Get retrieves the AwsKeyPair from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsKeyPair, error)
	AwsKeyPairNamespaceListerExpansion
}

// awsKeyPairNamespaceLister implements the AwsKeyPairNamespaceLister
// interface.
type awsKeyPairNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsKeyPairs in the indexer for a given namespace.
func (s awsKeyPairNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsKeyPair, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsKeyPair))
	})
	return ret, err
}

// Get retrieves the AwsKeyPair from the indexer for a given namespace and name.
func (s awsKeyPairNamespaceLister) Get(name string) (*v1.AwsKeyPair, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awskeypair"), name)
	}
	return obj.(*v1.AwsKeyPair), nil
}
