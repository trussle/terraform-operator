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

// AwsCustomerGatewayLister helps list AwsCustomerGateways.
type AwsCustomerGatewayLister interface {
	// List lists all AwsCustomerGateways in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsCustomerGateway, err error)
	// AwsCustomerGateways returns an object that can list and get AwsCustomerGateways.
	AwsCustomerGateways(namespace string) AwsCustomerGatewayNamespaceLister
	AwsCustomerGatewayListerExpansion
}

// awsCustomerGatewayLister implements the AwsCustomerGatewayLister interface.
type awsCustomerGatewayLister struct {
	indexer cache.Indexer
}

// NewAwsCustomerGatewayLister returns a new AwsCustomerGatewayLister.
func NewAwsCustomerGatewayLister(indexer cache.Indexer) AwsCustomerGatewayLister {
	return &awsCustomerGatewayLister{indexer: indexer}
}

// List lists all AwsCustomerGateways in the indexer.
func (s *awsCustomerGatewayLister) List(selector labels.Selector) (ret []*v1.AwsCustomerGateway, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsCustomerGateway))
	})
	return ret, err
}

// AwsCustomerGateways returns an object that can list and get AwsCustomerGateways.
func (s *awsCustomerGatewayLister) AwsCustomerGateways(namespace string) AwsCustomerGatewayNamespaceLister {
	return awsCustomerGatewayNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsCustomerGatewayNamespaceLister helps list and get AwsCustomerGateways.
type AwsCustomerGatewayNamespaceLister interface {
	// List lists all AwsCustomerGateways in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsCustomerGateway, err error)
	// Get retrieves the AwsCustomerGateway from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsCustomerGateway, error)
	AwsCustomerGatewayNamespaceListerExpansion
}

// awsCustomerGatewayNamespaceLister implements the AwsCustomerGatewayNamespaceLister
// interface.
type awsCustomerGatewayNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsCustomerGateways in the indexer for a given namespace.
func (s awsCustomerGatewayNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsCustomerGateway, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsCustomerGateway))
	})
	return ret, err
}

// Get retrieves the AwsCustomerGateway from the indexer for a given namespace and name.
func (s awsCustomerGatewayNamespaceLister) Get(name string) (*v1.AwsCustomerGateway, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awscustomergateway"), name)
	}
	return obj.(*v1.AwsCustomerGateway), nil
}
