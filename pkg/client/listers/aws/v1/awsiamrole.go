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

// AwsIamRoleLister helps list AwsIamRoles.
type AwsIamRoleLister interface {
	// List lists all AwsIamRoles in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsIamRole, err error)
	// AwsIamRoles returns an object that can list and get AwsIamRoles.
	AwsIamRoles(namespace string) AwsIamRoleNamespaceLister
	AwsIamRoleListerExpansion
}

// awsIamRoleLister implements the AwsIamRoleLister interface.
type awsIamRoleLister struct {
	indexer cache.Indexer
}

// NewAwsIamRoleLister returns a new AwsIamRoleLister.
func NewAwsIamRoleLister(indexer cache.Indexer) AwsIamRoleLister {
	return &awsIamRoleLister{indexer: indexer}
}

// List lists all AwsIamRoles in the indexer.
func (s *awsIamRoleLister) List(selector labels.Selector) (ret []*v1.AwsIamRole, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsIamRole))
	})
	return ret, err
}

// AwsIamRoles returns an object that can list and get AwsIamRoles.
func (s *awsIamRoleLister) AwsIamRoles(namespace string) AwsIamRoleNamespaceLister {
	return awsIamRoleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsIamRoleNamespaceLister helps list and get AwsIamRoles.
type AwsIamRoleNamespaceLister interface {
	// List lists all AwsIamRoles in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsIamRole, err error)
	// Get retrieves the AwsIamRole from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsIamRole, error)
	AwsIamRoleNamespaceListerExpansion
}

// awsIamRoleNamespaceLister implements the AwsIamRoleNamespaceLister
// interface.
type awsIamRoleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsIamRoles in the indexer for a given namespace.
func (s awsIamRoleNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsIamRole, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsIamRole))
	})
	return ret, err
}

// Get retrieves the AwsIamRole from the indexer for a given namespace and name.
func (s awsIamRoleNamespaceLister) Get(name string) (*v1.AwsIamRole, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsiamrole"), name)
	}
	return obj.(*v1.AwsIamRole), nil
}
