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

// AwsApiGatewayApiKeyLister helps list AwsApiGatewayApiKeys.
type AwsApiGatewayApiKeyLister interface {
	// List lists all AwsApiGatewayApiKeys in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsApiGatewayApiKey, err error)
	// AwsApiGatewayApiKeys returns an object that can list and get AwsApiGatewayApiKeys.
	AwsApiGatewayApiKeys(namespace string) AwsApiGatewayApiKeyNamespaceLister
	AwsApiGatewayApiKeyListerExpansion
}

// awsApiGatewayApiKeyLister implements the AwsApiGatewayApiKeyLister interface.
type awsApiGatewayApiKeyLister struct {
	indexer cache.Indexer
}

// NewAwsApiGatewayApiKeyLister returns a new AwsApiGatewayApiKeyLister.
func NewAwsApiGatewayApiKeyLister(indexer cache.Indexer) AwsApiGatewayApiKeyLister {
	return &awsApiGatewayApiKeyLister{indexer: indexer}
}

// List lists all AwsApiGatewayApiKeys in the indexer.
func (s *awsApiGatewayApiKeyLister) List(selector labels.Selector) (ret []*v1.AwsApiGatewayApiKey, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsApiGatewayApiKey))
	})
	return ret, err
}

// AwsApiGatewayApiKeys returns an object that can list and get AwsApiGatewayApiKeys.
func (s *awsApiGatewayApiKeyLister) AwsApiGatewayApiKeys(namespace string) AwsApiGatewayApiKeyNamespaceLister {
	return awsApiGatewayApiKeyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsApiGatewayApiKeyNamespaceLister helps list and get AwsApiGatewayApiKeys.
type AwsApiGatewayApiKeyNamespaceLister interface {
	// List lists all AwsApiGatewayApiKeys in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsApiGatewayApiKey, err error)
	// Get retrieves the AwsApiGatewayApiKey from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsApiGatewayApiKey, error)
	AwsApiGatewayApiKeyNamespaceListerExpansion
}

// awsApiGatewayApiKeyNamespaceLister implements the AwsApiGatewayApiKeyNamespaceLister
// interface.
type awsApiGatewayApiKeyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsApiGatewayApiKeys in the indexer for a given namespace.
func (s awsApiGatewayApiKeyNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsApiGatewayApiKey, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsApiGatewayApiKey))
	})
	return ret, err
}

// Get retrieves the AwsApiGatewayApiKey from the indexer for a given namespace and name.
func (s awsApiGatewayApiKeyNamespaceLister) Get(name string) (*v1.AwsApiGatewayApiKey, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsapigatewayapikey"), name)
	}
	return obj.(*v1.AwsApiGatewayApiKey), nil
}