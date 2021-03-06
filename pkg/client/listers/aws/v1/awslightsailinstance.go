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

// AwsLightsailInstanceLister helps list AwsLightsailInstances.
type AwsLightsailInstanceLister interface {
	// List lists all AwsLightsailInstances in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsLightsailInstance, err error)
	// AwsLightsailInstances returns an object that can list and get AwsLightsailInstances.
	AwsLightsailInstances(namespace string) AwsLightsailInstanceNamespaceLister
	AwsLightsailInstanceListerExpansion
}

// awsLightsailInstanceLister implements the AwsLightsailInstanceLister interface.
type awsLightsailInstanceLister struct {
	indexer cache.Indexer
}

// NewAwsLightsailInstanceLister returns a new AwsLightsailInstanceLister.
func NewAwsLightsailInstanceLister(indexer cache.Indexer) AwsLightsailInstanceLister {
	return &awsLightsailInstanceLister{indexer: indexer}
}

// List lists all AwsLightsailInstances in the indexer.
func (s *awsLightsailInstanceLister) List(selector labels.Selector) (ret []*v1.AwsLightsailInstance, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsLightsailInstance))
	})
	return ret, err
}

// AwsLightsailInstances returns an object that can list and get AwsLightsailInstances.
func (s *awsLightsailInstanceLister) AwsLightsailInstances(namespace string) AwsLightsailInstanceNamespaceLister {
	return awsLightsailInstanceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsLightsailInstanceNamespaceLister helps list and get AwsLightsailInstances.
type AwsLightsailInstanceNamespaceLister interface {
	// List lists all AwsLightsailInstances in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsLightsailInstance, err error)
	// Get retrieves the AwsLightsailInstance from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsLightsailInstance, error)
	AwsLightsailInstanceNamespaceListerExpansion
}

// awsLightsailInstanceNamespaceLister implements the AwsLightsailInstanceNamespaceLister
// interface.
type awsLightsailInstanceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsLightsailInstances in the indexer for a given namespace.
func (s awsLightsailInstanceNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsLightsailInstance, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsLightsailInstance))
	})
	return ret, err
}

// Get retrieves the AwsLightsailInstance from the indexer for a given namespace and name.
func (s awsLightsailInstanceNamespaceLister) Get(name string) (*v1.AwsLightsailInstance, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awslightsailinstance"), name)
	}
	return obj.(*v1.AwsLightsailInstance), nil
}
