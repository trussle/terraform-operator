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

// AwsCloud9EnvironmentEc2Lister helps list AwsCloud9EnvironmentEc2s.
type AwsCloud9EnvironmentEc2Lister interface {
	// List lists all AwsCloud9EnvironmentEc2s in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsCloud9EnvironmentEc2, err error)
	// AwsCloud9EnvironmentEc2s returns an object that can list and get AwsCloud9EnvironmentEc2s.
	AwsCloud9EnvironmentEc2s(namespace string) AwsCloud9EnvironmentEc2NamespaceLister
	AwsCloud9EnvironmentEc2ListerExpansion
}

// awsCloud9EnvironmentEc2Lister implements the AwsCloud9EnvironmentEc2Lister interface.
type awsCloud9EnvironmentEc2Lister struct {
	indexer cache.Indexer
}

// NewAwsCloud9EnvironmentEc2Lister returns a new AwsCloud9EnvironmentEc2Lister.
func NewAwsCloud9EnvironmentEc2Lister(indexer cache.Indexer) AwsCloud9EnvironmentEc2Lister {
	return &awsCloud9EnvironmentEc2Lister{indexer: indexer}
}

// List lists all AwsCloud9EnvironmentEc2s in the indexer.
func (s *awsCloud9EnvironmentEc2Lister) List(selector labels.Selector) (ret []*v1.AwsCloud9EnvironmentEc2, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsCloud9EnvironmentEc2))
	})
	return ret, err
}

// AwsCloud9EnvironmentEc2s returns an object that can list and get AwsCloud9EnvironmentEc2s.
func (s *awsCloud9EnvironmentEc2Lister) AwsCloud9EnvironmentEc2s(namespace string) AwsCloud9EnvironmentEc2NamespaceLister {
	return awsCloud9EnvironmentEc2NamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsCloud9EnvironmentEc2NamespaceLister helps list and get AwsCloud9EnvironmentEc2s.
type AwsCloud9EnvironmentEc2NamespaceLister interface {
	// List lists all AwsCloud9EnvironmentEc2s in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsCloud9EnvironmentEc2, err error)
	// Get retrieves the AwsCloud9EnvironmentEc2 from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsCloud9EnvironmentEc2, error)
	AwsCloud9EnvironmentEc2NamespaceListerExpansion
}

// awsCloud9EnvironmentEc2NamespaceLister implements the AwsCloud9EnvironmentEc2NamespaceLister
// interface.
type awsCloud9EnvironmentEc2NamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsCloud9EnvironmentEc2s in the indexer for a given namespace.
func (s awsCloud9EnvironmentEc2NamespaceLister) List(selector labels.Selector) (ret []*v1.AwsCloud9EnvironmentEc2, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsCloud9EnvironmentEc2))
	})
	return ret, err
}

// Get retrieves the AwsCloud9EnvironmentEc2 from the indexer for a given namespace and name.
func (s awsCloud9EnvironmentEc2NamespaceLister) Get(name string) (*v1.AwsCloud9EnvironmentEc2, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awscloud9environmentec2"), name)
	}
	return obj.(*v1.AwsCloud9EnvironmentEc2), nil
}