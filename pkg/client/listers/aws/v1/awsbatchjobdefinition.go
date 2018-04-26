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

// AwsBatchJobDefinitionLister helps list AwsBatchJobDefinitions.
type AwsBatchJobDefinitionLister interface {
	// List lists all AwsBatchJobDefinitions in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsBatchJobDefinition, err error)
	// AwsBatchJobDefinitions returns an object that can list and get AwsBatchJobDefinitions.
	AwsBatchJobDefinitions(namespace string) AwsBatchJobDefinitionNamespaceLister
	AwsBatchJobDefinitionListerExpansion
}

// awsBatchJobDefinitionLister implements the AwsBatchJobDefinitionLister interface.
type awsBatchJobDefinitionLister struct {
	indexer cache.Indexer
}

// NewAwsBatchJobDefinitionLister returns a new AwsBatchJobDefinitionLister.
func NewAwsBatchJobDefinitionLister(indexer cache.Indexer) AwsBatchJobDefinitionLister {
	return &awsBatchJobDefinitionLister{indexer: indexer}
}

// List lists all AwsBatchJobDefinitions in the indexer.
func (s *awsBatchJobDefinitionLister) List(selector labels.Selector) (ret []*v1.AwsBatchJobDefinition, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsBatchJobDefinition))
	})
	return ret, err
}

// AwsBatchJobDefinitions returns an object that can list and get AwsBatchJobDefinitions.
func (s *awsBatchJobDefinitionLister) AwsBatchJobDefinitions(namespace string) AwsBatchJobDefinitionNamespaceLister {
	return awsBatchJobDefinitionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsBatchJobDefinitionNamespaceLister helps list and get AwsBatchJobDefinitions.
type AwsBatchJobDefinitionNamespaceLister interface {
	// List lists all AwsBatchJobDefinitions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsBatchJobDefinition, err error)
	// Get retrieves the AwsBatchJobDefinition from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsBatchJobDefinition, error)
	AwsBatchJobDefinitionNamespaceListerExpansion
}

// awsBatchJobDefinitionNamespaceLister implements the AwsBatchJobDefinitionNamespaceLister
// interface.
type awsBatchJobDefinitionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsBatchJobDefinitions in the indexer for a given namespace.
func (s awsBatchJobDefinitionNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsBatchJobDefinition, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsBatchJobDefinition))
	})
	return ret, err
}

// Get retrieves the AwsBatchJobDefinition from the indexer for a given namespace and name.
func (s awsBatchJobDefinitionNamespaceLister) Get(name string) (*v1.AwsBatchJobDefinition, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsbatchjobdefinition"), name)
	}
	return obj.(*v1.AwsBatchJobDefinition), nil
}
