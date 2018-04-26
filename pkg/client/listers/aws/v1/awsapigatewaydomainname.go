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

// AwsApiGatewayDomainNameLister helps list AwsApiGatewayDomainNames.
type AwsApiGatewayDomainNameLister interface {
	// List lists all AwsApiGatewayDomainNames in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsApiGatewayDomainName, err error)
	// AwsApiGatewayDomainNames returns an object that can list and get AwsApiGatewayDomainNames.
	AwsApiGatewayDomainNames(namespace string) AwsApiGatewayDomainNameNamespaceLister
	AwsApiGatewayDomainNameListerExpansion
}

// awsApiGatewayDomainNameLister implements the AwsApiGatewayDomainNameLister interface.
type awsApiGatewayDomainNameLister struct {
	indexer cache.Indexer
}

// NewAwsApiGatewayDomainNameLister returns a new AwsApiGatewayDomainNameLister.
func NewAwsApiGatewayDomainNameLister(indexer cache.Indexer) AwsApiGatewayDomainNameLister {
	return &awsApiGatewayDomainNameLister{indexer: indexer}
}

// List lists all AwsApiGatewayDomainNames in the indexer.
func (s *awsApiGatewayDomainNameLister) List(selector labels.Selector) (ret []*v1.AwsApiGatewayDomainName, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsApiGatewayDomainName))
	})
	return ret, err
}

// AwsApiGatewayDomainNames returns an object that can list and get AwsApiGatewayDomainNames.
func (s *awsApiGatewayDomainNameLister) AwsApiGatewayDomainNames(namespace string) AwsApiGatewayDomainNameNamespaceLister {
	return awsApiGatewayDomainNameNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsApiGatewayDomainNameNamespaceLister helps list and get AwsApiGatewayDomainNames.
type AwsApiGatewayDomainNameNamespaceLister interface {
	// List lists all AwsApiGatewayDomainNames in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsApiGatewayDomainName, err error)
	// Get retrieves the AwsApiGatewayDomainName from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsApiGatewayDomainName, error)
	AwsApiGatewayDomainNameNamespaceListerExpansion
}

// awsApiGatewayDomainNameNamespaceLister implements the AwsApiGatewayDomainNameNamespaceLister
// interface.
type awsApiGatewayDomainNameNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsApiGatewayDomainNames in the indexer for a given namespace.
func (s awsApiGatewayDomainNameNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsApiGatewayDomainName, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsApiGatewayDomainName))
	})
	return ret, err
}

// Get retrieves the AwsApiGatewayDomainName from the indexer for a given namespace and name.
func (s awsApiGatewayDomainNameNamespaceLister) Get(name string) (*v1.AwsApiGatewayDomainName, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsapigatewaydomainname"), name)
	}
	return obj.(*v1.AwsApiGatewayDomainName), nil
}