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

// AwsDxLagLister helps list AwsDxLags.
type AwsDxLagLister interface {
	// List lists all AwsDxLags in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsDxLag, err error)
	// AwsDxLags returns an object that can list and get AwsDxLags.
	AwsDxLags(namespace string) AwsDxLagNamespaceLister
	AwsDxLagListerExpansion
}

// awsDxLagLister implements the AwsDxLagLister interface.
type awsDxLagLister struct {
	indexer cache.Indexer
}

// NewAwsDxLagLister returns a new AwsDxLagLister.
func NewAwsDxLagLister(indexer cache.Indexer) AwsDxLagLister {
	return &awsDxLagLister{indexer: indexer}
}

// List lists all AwsDxLags in the indexer.
func (s *awsDxLagLister) List(selector labels.Selector) (ret []*v1.AwsDxLag, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsDxLag))
	})
	return ret, err
}

// AwsDxLags returns an object that can list and get AwsDxLags.
func (s *awsDxLagLister) AwsDxLags(namespace string) AwsDxLagNamespaceLister {
	return awsDxLagNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsDxLagNamespaceLister helps list and get AwsDxLags.
type AwsDxLagNamespaceLister interface {
	// List lists all AwsDxLags in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsDxLag, err error)
	// Get retrieves the AwsDxLag from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsDxLag, error)
	AwsDxLagNamespaceListerExpansion
}

// awsDxLagNamespaceLister implements the AwsDxLagNamespaceLister
// interface.
type awsDxLagNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsDxLags in the indexer for a given namespace.
func (s awsDxLagNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsDxLag, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsDxLag))
	})
	return ret, err
}

// Get retrieves the AwsDxLag from the indexer for a given namespace and name.
func (s awsDxLagNamespaceLister) Get(name string) (*v1.AwsDxLag, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsdxlag"), name)
	}
	return obj.(*v1.AwsDxLag), nil
}
