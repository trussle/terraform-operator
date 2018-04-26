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

// AwsSsmResourceDataSyncLister helps list AwsSsmResourceDataSyncs.
type AwsSsmResourceDataSyncLister interface {
	// List lists all AwsSsmResourceDataSyncs in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsSsmResourceDataSync, err error)
	// AwsSsmResourceDataSyncs returns an object that can list and get AwsSsmResourceDataSyncs.
	AwsSsmResourceDataSyncs(namespace string) AwsSsmResourceDataSyncNamespaceLister
	AwsSsmResourceDataSyncListerExpansion
}

// awsSsmResourceDataSyncLister implements the AwsSsmResourceDataSyncLister interface.
type awsSsmResourceDataSyncLister struct {
	indexer cache.Indexer
}

// NewAwsSsmResourceDataSyncLister returns a new AwsSsmResourceDataSyncLister.
func NewAwsSsmResourceDataSyncLister(indexer cache.Indexer) AwsSsmResourceDataSyncLister {
	return &awsSsmResourceDataSyncLister{indexer: indexer}
}

// List lists all AwsSsmResourceDataSyncs in the indexer.
func (s *awsSsmResourceDataSyncLister) List(selector labels.Selector) (ret []*v1.AwsSsmResourceDataSync, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsSsmResourceDataSync))
	})
	return ret, err
}

// AwsSsmResourceDataSyncs returns an object that can list and get AwsSsmResourceDataSyncs.
func (s *awsSsmResourceDataSyncLister) AwsSsmResourceDataSyncs(namespace string) AwsSsmResourceDataSyncNamespaceLister {
	return awsSsmResourceDataSyncNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsSsmResourceDataSyncNamespaceLister helps list and get AwsSsmResourceDataSyncs.
type AwsSsmResourceDataSyncNamespaceLister interface {
	// List lists all AwsSsmResourceDataSyncs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsSsmResourceDataSync, err error)
	// Get retrieves the AwsSsmResourceDataSync from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsSsmResourceDataSync, error)
	AwsSsmResourceDataSyncNamespaceListerExpansion
}

// awsSsmResourceDataSyncNamespaceLister implements the AwsSsmResourceDataSyncNamespaceLister
// interface.
type awsSsmResourceDataSyncNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsSsmResourceDataSyncs in the indexer for a given namespace.
func (s awsSsmResourceDataSyncNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsSsmResourceDataSync, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsSsmResourceDataSync))
	})
	return ret, err
}

// Get retrieves the AwsSsmResourceDataSync from the indexer for a given namespace and name.
func (s awsSsmResourceDataSyncNamespaceLister) Get(name string) (*v1.AwsSsmResourceDataSync, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsssmresourcedatasync"), name)
	}
	return obj.(*v1.AwsSsmResourceDataSync), nil
}
