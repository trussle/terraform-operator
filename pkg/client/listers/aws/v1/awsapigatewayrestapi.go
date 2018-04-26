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

// AwsApiGatewayRestApiLister helps list AwsApiGatewayRestApis.
type AwsApiGatewayRestApiLister interface {
	// List lists all AwsApiGatewayRestApis in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsApiGatewayRestApi, err error)
	// AwsApiGatewayRestApis returns an object that can list and get AwsApiGatewayRestApis.
	AwsApiGatewayRestApis(namespace string) AwsApiGatewayRestApiNamespaceLister
	AwsApiGatewayRestApiListerExpansion
}

// awsApiGatewayRestApiLister implements the AwsApiGatewayRestApiLister interface.
type awsApiGatewayRestApiLister struct {
	indexer cache.Indexer
}

// NewAwsApiGatewayRestApiLister returns a new AwsApiGatewayRestApiLister.
func NewAwsApiGatewayRestApiLister(indexer cache.Indexer) AwsApiGatewayRestApiLister {
	return &awsApiGatewayRestApiLister{indexer: indexer}
}

// List lists all AwsApiGatewayRestApis in the indexer.
func (s *awsApiGatewayRestApiLister) List(selector labels.Selector) (ret []*v1.AwsApiGatewayRestApi, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsApiGatewayRestApi))
	})
	return ret, err
}

// AwsApiGatewayRestApis returns an object that can list and get AwsApiGatewayRestApis.
func (s *awsApiGatewayRestApiLister) AwsApiGatewayRestApis(namespace string) AwsApiGatewayRestApiNamespaceLister {
	return awsApiGatewayRestApiNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsApiGatewayRestApiNamespaceLister helps list and get AwsApiGatewayRestApis.
type AwsApiGatewayRestApiNamespaceLister interface {
	// List lists all AwsApiGatewayRestApis in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsApiGatewayRestApi, err error)
	// Get retrieves the AwsApiGatewayRestApi from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsApiGatewayRestApi, error)
	AwsApiGatewayRestApiNamespaceListerExpansion
}

// awsApiGatewayRestApiNamespaceLister implements the AwsApiGatewayRestApiNamespaceLister
// interface.
type awsApiGatewayRestApiNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsApiGatewayRestApis in the indexer for a given namespace.
func (s awsApiGatewayRestApiNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsApiGatewayRestApi, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsApiGatewayRestApi))
	})
	return ret, err
}

// Get retrieves the AwsApiGatewayRestApi from the indexer for a given namespace and name.
func (s awsApiGatewayRestApiNamespaceLister) Get(name string) (*v1.AwsApiGatewayRestApi, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsapigatewayrestapi"), name)
	}
	return obj.(*v1.AwsApiGatewayRestApi), nil
}
