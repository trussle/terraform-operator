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

// AwsKmsKeyLister helps list AwsKmsKeys.
type AwsKmsKeyLister interface {
	// List lists all AwsKmsKeys in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsKmsKey, err error)
	// AwsKmsKeys returns an object that can list and get AwsKmsKeys.
	AwsKmsKeys(namespace string) AwsKmsKeyNamespaceLister
	AwsKmsKeyListerExpansion
}

// awsKmsKeyLister implements the AwsKmsKeyLister interface.
type awsKmsKeyLister struct {
	indexer cache.Indexer
}

// NewAwsKmsKeyLister returns a new AwsKmsKeyLister.
func NewAwsKmsKeyLister(indexer cache.Indexer) AwsKmsKeyLister {
	return &awsKmsKeyLister{indexer: indexer}
}

// List lists all AwsKmsKeys in the indexer.
func (s *awsKmsKeyLister) List(selector labels.Selector) (ret []*v1.AwsKmsKey, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsKmsKey))
	})
	return ret, err
}

// AwsKmsKeys returns an object that can list and get AwsKmsKeys.
func (s *awsKmsKeyLister) AwsKmsKeys(namespace string) AwsKmsKeyNamespaceLister {
	return awsKmsKeyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsKmsKeyNamespaceLister helps list and get AwsKmsKeys.
type AwsKmsKeyNamespaceLister interface {
	// List lists all AwsKmsKeys in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsKmsKey, err error)
	// Get retrieves the AwsKmsKey from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsKmsKey, error)
	AwsKmsKeyNamespaceListerExpansion
}

// awsKmsKeyNamespaceLister implements the AwsKmsKeyNamespaceLister
// interface.
type awsKmsKeyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsKmsKeys in the indexer for a given namespace.
func (s awsKmsKeyNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsKmsKey, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsKmsKey))
	})
	return ret, err
}

// Get retrieves the AwsKmsKey from the indexer for a given namespace and name.
func (s awsKmsKeyNamespaceLister) Get(name string) (*v1.AwsKmsKey, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awskmskey"), name)
	}
	return obj.(*v1.AwsKmsKey), nil
}
