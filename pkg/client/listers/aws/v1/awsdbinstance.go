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

// AwsDbInstanceLister helps list AwsDbInstances.
type AwsDbInstanceLister interface {
	// List lists all AwsDbInstances in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsDbInstance, err error)
	// AwsDbInstances returns an object that can list and get AwsDbInstances.
	AwsDbInstances(namespace string) AwsDbInstanceNamespaceLister
	AwsDbInstanceListerExpansion
}

// awsDbInstanceLister implements the AwsDbInstanceLister interface.
type awsDbInstanceLister struct {
	indexer cache.Indexer
}

// NewAwsDbInstanceLister returns a new AwsDbInstanceLister.
func NewAwsDbInstanceLister(indexer cache.Indexer) AwsDbInstanceLister {
	return &awsDbInstanceLister{indexer: indexer}
}

// List lists all AwsDbInstances in the indexer.
func (s *awsDbInstanceLister) List(selector labels.Selector) (ret []*v1.AwsDbInstance, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsDbInstance))
	})
	return ret, err
}

// AwsDbInstances returns an object that can list and get AwsDbInstances.
func (s *awsDbInstanceLister) AwsDbInstances(namespace string) AwsDbInstanceNamespaceLister {
	return awsDbInstanceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsDbInstanceNamespaceLister helps list and get AwsDbInstances.
type AwsDbInstanceNamespaceLister interface {
	// List lists all AwsDbInstances in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsDbInstance, err error)
	// Get retrieves the AwsDbInstance from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsDbInstance, error)
	AwsDbInstanceNamespaceListerExpansion
}

// awsDbInstanceNamespaceLister implements the AwsDbInstanceNamespaceLister
// interface.
type awsDbInstanceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsDbInstances in the indexer for a given namespace.
func (s awsDbInstanceNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsDbInstance, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsDbInstance))
	})
	return ret, err
}

// Get retrieves the AwsDbInstance from the indexer for a given namespace and name.
func (s awsDbInstanceNamespaceLister) Get(name string) (*v1.AwsDbInstance, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsdbinstance"), name)
	}
	return obj.(*v1.AwsDbInstance), nil
}
