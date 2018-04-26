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

// AwsNetworkAclLister helps list AwsNetworkAcls.
type AwsNetworkAclLister interface {
	// List lists all AwsNetworkAcls in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsNetworkAcl, err error)
	// AwsNetworkAcls returns an object that can list and get AwsNetworkAcls.
	AwsNetworkAcls(namespace string) AwsNetworkAclNamespaceLister
	AwsNetworkAclListerExpansion
}

// awsNetworkAclLister implements the AwsNetworkAclLister interface.
type awsNetworkAclLister struct {
	indexer cache.Indexer
}

// NewAwsNetworkAclLister returns a new AwsNetworkAclLister.
func NewAwsNetworkAclLister(indexer cache.Indexer) AwsNetworkAclLister {
	return &awsNetworkAclLister{indexer: indexer}
}

// List lists all AwsNetworkAcls in the indexer.
func (s *awsNetworkAclLister) List(selector labels.Selector) (ret []*v1.AwsNetworkAcl, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsNetworkAcl))
	})
	return ret, err
}

// AwsNetworkAcls returns an object that can list and get AwsNetworkAcls.
func (s *awsNetworkAclLister) AwsNetworkAcls(namespace string) AwsNetworkAclNamespaceLister {
	return awsNetworkAclNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsNetworkAclNamespaceLister helps list and get AwsNetworkAcls.
type AwsNetworkAclNamespaceLister interface {
	// List lists all AwsNetworkAcls in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsNetworkAcl, err error)
	// Get retrieves the AwsNetworkAcl from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsNetworkAcl, error)
	AwsNetworkAclNamespaceListerExpansion
}

// awsNetworkAclNamespaceLister implements the AwsNetworkAclNamespaceLister
// interface.
type awsNetworkAclNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsNetworkAcls in the indexer for a given namespace.
func (s awsNetworkAclNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsNetworkAcl, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsNetworkAcl))
	})
	return ret, err
}

// Get retrieves the AwsNetworkAcl from the indexer for a given namespace and name.
func (s awsNetworkAclNamespaceLister) Get(name string) (*v1.AwsNetworkAcl, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsnetworkacl"), name)
	}
	return obj.(*v1.AwsNetworkAcl), nil
}