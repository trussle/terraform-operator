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

// AwsDmsReplicationTaskLister helps list AwsDmsReplicationTasks.
type AwsDmsReplicationTaskLister interface {
	// List lists all AwsDmsReplicationTasks in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsDmsReplicationTask, err error)
	// AwsDmsReplicationTasks returns an object that can list and get AwsDmsReplicationTasks.
	AwsDmsReplicationTasks(namespace string) AwsDmsReplicationTaskNamespaceLister
	AwsDmsReplicationTaskListerExpansion
}

// awsDmsReplicationTaskLister implements the AwsDmsReplicationTaskLister interface.
type awsDmsReplicationTaskLister struct {
	indexer cache.Indexer
}

// NewAwsDmsReplicationTaskLister returns a new AwsDmsReplicationTaskLister.
func NewAwsDmsReplicationTaskLister(indexer cache.Indexer) AwsDmsReplicationTaskLister {
	return &awsDmsReplicationTaskLister{indexer: indexer}
}

// List lists all AwsDmsReplicationTasks in the indexer.
func (s *awsDmsReplicationTaskLister) List(selector labels.Selector) (ret []*v1.AwsDmsReplicationTask, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsDmsReplicationTask))
	})
	return ret, err
}

// AwsDmsReplicationTasks returns an object that can list and get AwsDmsReplicationTasks.
func (s *awsDmsReplicationTaskLister) AwsDmsReplicationTasks(namespace string) AwsDmsReplicationTaskNamespaceLister {
	return awsDmsReplicationTaskNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsDmsReplicationTaskNamespaceLister helps list and get AwsDmsReplicationTasks.
type AwsDmsReplicationTaskNamespaceLister interface {
	// List lists all AwsDmsReplicationTasks in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsDmsReplicationTask, err error)
	// Get retrieves the AwsDmsReplicationTask from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsDmsReplicationTask, error)
	AwsDmsReplicationTaskNamespaceListerExpansion
}

// awsDmsReplicationTaskNamespaceLister implements the AwsDmsReplicationTaskNamespaceLister
// interface.
type awsDmsReplicationTaskNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsDmsReplicationTasks in the indexer for a given namespace.
func (s awsDmsReplicationTaskNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsDmsReplicationTask, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsDmsReplicationTask))
	})
	return ret, err
}

// Get retrieves the AwsDmsReplicationTask from the indexer for a given namespace and name.
func (s awsDmsReplicationTaskNamespaceLister) Get(name string) (*v1.AwsDmsReplicationTask, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsdmsreplicationtask"), name)
	}
	return obj.(*v1.AwsDmsReplicationTask), nil
}
