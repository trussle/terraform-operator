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

// AwsWafIpsetLister helps list AwsWafIpsets.
type AwsWafIpsetLister interface {
	// List lists all AwsWafIpsets in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsWafIpset, err error)
	// AwsWafIpsets returns an object that can list and get AwsWafIpsets.
	AwsWafIpsets(namespace string) AwsWafIpsetNamespaceLister
	AwsWafIpsetListerExpansion
}

// awsWafIpsetLister implements the AwsWafIpsetLister interface.
type awsWafIpsetLister struct {
	indexer cache.Indexer
}

// NewAwsWafIpsetLister returns a new AwsWafIpsetLister.
func NewAwsWafIpsetLister(indexer cache.Indexer) AwsWafIpsetLister {
	return &awsWafIpsetLister{indexer: indexer}
}

// List lists all AwsWafIpsets in the indexer.
func (s *awsWafIpsetLister) List(selector labels.Selector) (ret []*v1.AwsWafIpset, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsWafIpset))
	})
	return ret, err
}

// AwsWafIpsets returns an object that can list and get AwsWafIpsets.
func (s *awsWafIpsetLister) AwsWafIpsets(namespace string) AwsWafIpsetNamespaceLister {
	return awsWafIpsetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsWafIpsetNamespaceLister helps list and get AwsWafIpsets.
type AwsWafIpsetNamespaceLister interface {
	// List lists all AwsWafIpsets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsWafIpset, err error)
	// Get retrieves the AwsWafIpset from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsWafIpset, error)
	AwsWafIpsetNamespaceListerExpansion
}

// awsWafIpsetNamespaceLister implements the AwsWafIpsetNamespaceLister
// interface.
type awsWafIpsetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsWafIpsets in the indexer for a given namespace.
func (s awsWafIpsetNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsWafIpset, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsWafIpset))
	})
	return ret, err
}

// Get retrieves the AwsWafIpset from the indexer for a given namespace and name.
func (s awsWafIpsetNamespaceLister) Get(name string) (*v1.AwsWafIpset, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awswafipset"), name)
	}
	return obj.(*v1.AwsWafIpset), nil
}
