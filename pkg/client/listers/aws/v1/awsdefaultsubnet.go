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

// AwsDefaultSubnetLister helps list AwsDefaultSubnets.
type AwsDefaultSubnetLister interface {
	// List lists all AwsDefaultSubnets in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsDefaultSubnet, err error)
	// AwsDefaultSubnets returns an object that can list and get AwsDefaultSubnets.
	AwsDefaultSubnets(namespace string) AwsDefaultSubnetNamespaceLister
	AwsDefaultSubnetListerExpansion
}

// awsDefaultSubnetLister implements the AwsDefaultSubnetLister interface.
type awsDefaultSubnetLister struct {
	indexer cache.Indexer
}

// NewAwsDefaultSubnetLister returns a new AwsDefaultSubnetLister.
func NewAwsDefaultSubnetLister(indexer cache.Indexer) AwsDefaultSubnetLister {
	return &awsDefaultSubnetLister{indexer: indexer}
}

// List lists all AwsDefaultSubnets in the indexer.
func (s *awsDefaultSubnetLister) List(selector labels.Selector) (ret []*v1.AwsDefaultSubnet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsDefaultSubnet))
	})
	return ret, err
}

// AwsDefaultSubnets returns an object that can list and get AwsDefaultSubnets.
func (s *awsDefaultSubnetLister) AwsDefaultSubnets(namespace string) AwsDefaultSubnetNamespaceLister {
	return awsDefaultSubnetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsDefaultSubnetNamespaceLister helps list and get AwsDefaultSubnets.
type AwsDefaultSubnetNamespaceLister interface {
	// List lists all AwsDefaultSubnets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsDefaultSubnet, err error)
	// Get retrieves the AwsDefaultSubnet from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsDefaultSubnet, error)
	AwsDefaultSubnetNamespaceListerExpansion
}

// awsDefaultSubnetNamespaceLister implements the AwsDefaultSubnetNamespaceLister
// interface.
type awsDefaultSubnetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsDefaultSubnets in the indexer for a given namespace.
func (s awsDefaultSubnetNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsDefaultSubnet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsDefaultSubnet))
	})
	return ret, err
}

// Get retrieves the AwsDefaultSubnet from the indexer for a given namespace and name.
func (s awsDefaultSubnetNamespaceLister) Get(name string) (*v1.AwsDefaultSubnet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsdefaultsubnet"), name)
	}
	return obj.(*v1.AwsDefaultSubnet), nil
}
