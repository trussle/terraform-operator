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

// AwsDirectoryServiceConditionalForwarderLister helps list AwsDirectoryServiceConditionalForwarders.
type AwsDirectoryServiceConditionalForwarderLister interface {
	// List lists all AwsDirectoryServiceConditionalForwarders in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsDirectoryServiceConditionalForwarder, err error)
	// AwsDirectoryServiceConditionalForwarders returns an object that can list and get AwsDirectoryServiceConditionalForwarders.
	AwsDirectoryServiceConditionalForwarders(namespace string) AwsDirectoryServiceConditionalForwarderNamespaceLister
	AwsDirectoryServiceConditionalForwarderListerExpansion
}

// awsDirectoryServiceConditionalForwarderLister implements the AwsDirectoryServiceConditionalForwarderLister interface.
type awsDirectoryServiceConditionalForwarderLister struct {
	indexer cache.Indexer
}

// NewAwsDirectoryServiceConditionalForwarderLister returns a new AwsDirectoryServiceConditionalForwarderLister.
func NewAwsDirectoryServiceConditionalForwarderLister(indexer cache.Indexer) AwsDirectoryServiceConditionalForwarderLister {
	return &awsDirectoryServiceConditionalForwarderLister{indexer: indexer}
}

// List lists all AwsDirectoryServiceConditionalForwarders in the indexer.
func (s *awsDirectoryServiceConditionalForwarderLister) List(selector labels.Selector) (ret []*v1.AwsDirectoryServiceConditionalForwarder, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsDirectoryServiceConditionalForwarder))
	})
	return ret, err
}

// AwsDirectoryServiceConditionalForwarders returns an object that can list and get AwsDirectoryServiceConditionalForwarders.
func (s *awsDirectoryServiceConditionalForwarderLister) AwsDirectoryServiceConditionalForwarders(namespace string) AwsDirectoryServiceConditionalForwarderNamespaceLister {
	return awsDirectoryServiceConditionalForwarderNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsDirectoryServiceConditionalForwarderNamespaceLister helps list and get AwsDirectoryServiceConditionalForwarders.
type AwsDirectoryServiceConditionalForwarderNamespaceLister interface {
	// List lists all AwsDirectoryServiceConditionalForwarders in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsDirectoryServiceConditionalForwarder, err error)
	// Get retrieves the AwsDirectoryServiceConditionalForwarder from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsDirectoryServiceConditionalForwarder, error)
	AwsDirectoryServiceConditionalForwarderNamespaceListerExpansion
}

// awsDirectoryServiceConditionalForwarderNamespaceLister implements the AwsDirectoryServiceConditionalForwarderNamespaceLister
// interface.
type awsDirectoryServiceConditionalForwarderNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsDirectoryServiceConditionalForwarders in the indexer for a given namespace.
func (s awsDirectoryServiceConditionalForwarderNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsDirectoryServiceConditionalForwarder, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsDirectoryServiceConditionalForwarder))
	})
	return ret, err
}

// Get retrieves the AwsDirectoryServiceConditionalForwarder from the indexer for a given namespace and name.
func (s awsDirectoryServiceConditionalForwarderNamespaceLister) Get(name string) (*v1.AwsDirectoryServiceConditionalForwarder, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsdirectoryserviceconditionalforwarder"), name)
	}
	return obj.(*v1.AwsDirectoryServiceConditionalForwarder), nil
}