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

// AwsMediaStoreContainerLister helps list AwsMediaStoreContainers.
type AwsMediaStoreContainerLister interface {
	// List lists all AwsMediaStoreContainers in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsMediaStoreContainer, err error)
	// AwsMediaStoreContainers returns an object that can list and get AwsMediaStoreContainers.
	AwsMediaStoreContainers(namespace string) AwsMediaStoreContainerNamespaceLister
	AwsMediaStoreContainerListerExpansion
}

// awsMediaStoreContainerLister implements the AwsMediaStoreContainerLister interface.
type awsMediaStoreContainerLister struct {
	indexer cache.Indexer
}

// NewAwsMediaStoreContainerLister returns a new AwsMediaStoreContainerLister.
func NewAwsMediaStoreContainerLister(indexer cache.Indexer) AwsMediaStoreContainerLister {
	return &awsMediaStoreContainerLister{indexer: indexer}
}

// List lists all AwsMediaStoreContainers in the indexer.
func (s *awsMediaStoreContainerLister) List(selector labels.Selector) (ret []*v1.AwsMediaStoreContainer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsMediaStoreContainer))
	})
	return ret, err
}

// AwsMediaStoreContainers returns an object that can list and get AwsMediaStoreContainers.
func (s *awsMediaStoreContainerLister) AwsMediaStoreContainers(namespace string) AwsMediaStoreContainerNamespaceLister {
	return awsMediaStoreContainerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsMediaStoreContainerNamespaceLister helps list and get AwsMediaStoreContainers.
type AwsMediaStoreContainerNamespaceLister interface {
	// List lists all AwsMediaStoreContainers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsMediaStoreContainer, err error)
	// Get retrieves the AwsMediaStoreContainer from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsMediaStoreContainer, error)
	AwsMediaStoreContainerNamespaceListerExpansion
}

// awsMediaStoreContainerNamespaceLister implements the AwsMediaStoreContainerNamespaceLister
// interface.
type awsMediaStoreContainerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsMediaStoreContainers in the indexer for a given namespace.
func (s awsMediaStoreContainerNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsMediaStoreContainer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsMediaStoreContainer))
	})
	return ret, err
}

// Get retrieves the AwsMediaStoreContainer from the indexer for a given namespace and name.
func (s awsMediaStoreContainerNamespaceLister) Get(name string) (*v1.AwsMediaStoreContainer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsmediastorecontainer"), name)
	}
	return obj.(*v1.AwsMediaStoreContainer), nil
}
