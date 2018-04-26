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

// AwsSecurityGroupLister helps list AwsSecurityGroups.
type AwsSecurityGroupLister interface {
	// List lists all AwsSecurityGroups in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsSecurityGroup, err error)
	// AwsSecurityGroups returns an object that can list and get AwsSecurityGroups.
	AwsSecurityGroups(namespace string) AwsSecurityGroupNamespaceLister
	AwsSecurityGroupListerExpansion
}

// awsSecurityGroupLister implements the AwsSecurityGroupLister interface.
type awsSecurityGroupLister struct {
	indexer cache.Indexer
}

// NewAwsSecurityGroupLister returns a new AwsSecurityGroupLister.
func NewAwsSecurityGroupLister(indexer cache.Indexer) AwsSecurityGroupLister {
	return &awsSecurityGroupLister{indexer: indexer}
}

// List lists all AwsSecurityGroups in the indexer.
func (s *awsSecurityGroupLister) List(selector labels.Selector) (ret []*v1.AwsSecurityGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsSecurityGroup))
	})
	return ret, err
}

// AwsSecurityGroups returns an object that can list and get AwsSecurityGroups.
func (s *awsSecurityGroupLister) AwsSecurityGroups(namespace string) AwsSecurityGroupNamespaceLister {
	return awsSecurityGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsSecurityGroupNamespaceLister helps list and get AwsSecurityGroups.
type AwsSecurityGroupNamespaceLister interface {
	// List lists all AwsSecurityGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsSecurityGroup, err error)
	// Get retrieves the AwsSecurityGroup from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsSecurityGroup, error)
	AwsSecurityGroupNamespaceListerExpansion
}

// awsSecurityGroupNamespaceLister implements the AwsSecurityGroupNamespaceLister
// interface.
type awsSecurityGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsSecurityGroups in the indexer for a given namespace.
func (s awsSecurityGroupNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsSecurityGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsSecurityGroup))
	})
	return ret, err
}

// Get retrieves the AwsSecurityGroup from the indexer for a given namespace and name.
func (s awsSecurityGroupNamespaceLister) Get(name string) (*v1.AwsSecurityGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awssecuritygroup"), name)
	}
	return obj.(*v1.AwsSecurityGroup), nil
}