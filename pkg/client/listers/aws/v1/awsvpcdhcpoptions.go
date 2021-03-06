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

// AwsVpcDhcpOptionsLister helps list AwsVpcDhcpOptionses.
type AwsVpcDhcpOptionsLister interface {
	// List lists all AwsVpcDhcpOptionses in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsVpcDhcpOptions, err error)
	// AwsVpcDhcpOptionses returns an object that can list and get AwsVpcDhcpOptionses.
	AwsVpcDhcpOptionses(namespace string) AwsVpcDhcpOptionsNamespaceLister
	AwsVpcDhcpOptionsListerExpansion
}

// awsVpcDhcpOptionsLister implements the AwsVpcDhcpOptionsLister interface.
type awsVpcDhcpOptionsLister struct {
	indexer cache.Indexer
}

// NewAwsVpcDhcpOptionsLister returns a new AwsVpcDhcpOptionsLister.
func NewAwsVpcDhcpOptionsLister(indexer cache.Indexer) AwsVpcDhcpOptionsLister {
	return &awsVpcDhcpOptionsLister{indexer: indexer}
}

// List lists all AwsVpcDhcpOptionses in the indexer.
func (s *awsVpcDhcpOptionsLister) List(selector labels.Selector) (ret []*v1.AwsVpcDhcpOptions, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsVpcDhcpOptions))
	})
	return ret, err
}

// AwsVpcDhcpOptionses returns an object that can list and get AwsVpcDhcpOptionses.
func (s *awsVpcDhcpOptionsLister) AwsVpcDhcpOptionses(namespace string) AwsVpcDhcpOptionsNamespaceLister {
	return awsVpcDhcpOptionsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsVpcDhcpOptionsNamespaceLister helps list and get AwsVpcDhcpOptionses.
type AwsVpcDhcpOptionsNamespaceLister interface {
	// List lists all AwsVpcDhcpOptionses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsVpcDhcpOptions, err error)
	// Get retrieves the AwsVpcDhcpOptions from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsVpcDhcpOptions, error)
	AwsVpcDhcpOptionsNamespaceListerExpansion
}

// awsVpcDhcpOptionsNamespaceLister implements the AwsVpcDhcpOptionsNamespaceLister
// interface.
type awsVpcDhcpOptionsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsVpcDhcpOptionses in the indexer for a given namespace.
func (s awsVpcDhcpOptionsNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsVpcDhcpOptions, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsVpcDhcpOptions))
	})
	return ret, err
}

// Get retrieves the AwsVpcDhcpOptions from the indexer for a given namespace and name.
func (s awsVpcDhcpOptionsNamespaceLister) Get(name string) (*v1.AwsVpcDhcpOptions, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsvpcdhcpoptions"), name)
	}
	return obj.(*v1.AwsVpcDhcpOptions), nil
}
