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

// AwsIamSamlProviderLister helps list AwsIamSamlProviders.
type AwsIamSamlProviderLister interface {
	// List lists all AwsIamSamlProviders in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsIamSamlProvider, err error)
	// AwsIamSamlProviders returns an object that can list and get AwsIamSamlProviders.
	AwsIamSamlProviders(namespace string) AwsIamSamlProviderNamespaceLister
	AwsIamSamlProviderListerExpansion
}

// awsIamSamlProviderLister implements the AwsIamSamlProviderLister interface.
type awsIamSamlProviderLister struct {
	indexer cache.Indexer
}

// NewAwsIamSamlProviderLister returns a new AwsIamSamlProviderLister.
func NewAwsIamSamlProviderLister(indexer cache.Indexer) AwsIamSamlProviderLister {
	return &awsIamSamlProviderLister{indexer: indexer}
}

// List lists all AwsIamSamlProviders in the indexer.
func (s *awsIamSamlProviderLister) List(selector labels.Selector) (ret []*v1.AwsIamSamlProvider, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsIamSamlProvider))
	})
	return ret, err
}

// AwsIamSamlProviders returns an object that can list and get AwsIamSamlProviders.
func (s *awsIamSamlProviderLister) AwsIamSamlProviders(namespace string) AwsIamSamlProviderNamespaceLister {
	return awsIamSamlProviderNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsIamSamlProviderNamespaceLister helps list and get AwsIamSamlProviders.
type AwsIamSamlProviderNamespaceLister interface {
	// List lists all AwsIamSamlProviders in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsIamSamlProvider, err error)
	// Get retrieves the AwsIamSamlProvider from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsIamSamlProvider, error)
	AwsIamSamlProviderNamespaceListerExpansion
}

// awsIamSamlProviderNamespaceLister implements the AwsIamSamlProviderNamespaceLister
// interface.
type awsIamSamlProviderNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsIamSamlProviders in the indexer for a given namespace.
func (s awsIamSamlProviderNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsIamSamlProvider, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsIamSamlProvider))
	})
	return ret, err
}

// Get retrieves the AwsIamSamlProvider from the indexer for a given namespace and name.
func (s awsIamSamlProviderNamespaceLister) Get(name string) (*v1.AwsIamSamlProvider, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsiamsamlprovider"), name)
	}
	return obj.(*v1.AwsIamSamlProvider), nil
}
