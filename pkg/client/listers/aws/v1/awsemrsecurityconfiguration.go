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

// AwsEmrSecurityConfigurationLister helps list AwsEmrSecurityConfigurations.
type AwsEmrSecurityConfigurationLister interface {
	// List lists all AwsEmrSecurityConfigurations in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsEmrSecurityConfiguration, err error)
	// AwsEmrSecurityConfigurations returns an object that can list and get AwsEmrSecurityConfigurations.
	AwsEmrSecurityConfigurations(namespace string) AwsEmrSecurityConfigurationNamespaceLister
	AwsEmrSecurityConfigurationListerExpansion
}

// awsEmrSecurityConfigurationLister implements the AwsEmrSecurityConfigurationLister interface.
type awsEmrSecurityConfigurationLister struct {
	indexer cache.Indexer
}

// NewAwsEmrSecurityConfigurationLister returns a new AwsEmrSecurityConfigurationLister.
func NewAwsEmrSecurityConfigurationLister(indexer cache.Indexer) AwsEmrSecurityConfigurationLister {
	return &awsEmrSecurityConfigurationLister{indexer: indexer}
}

// List lists all AwsEmrSecurityConfigurations in the indexer.
func (s *awsEmrSecurityConfigurationLister) List(selector labels.Selector) (ret []*v1.AwsEmrSecurityConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsEmrSecurityConfiguration))
	})
	return ret, err
}

// AwsEmrSecurityConfigurations returns an object that can list and get AwsEmrSecurityConfigurations.
func (s *awsEmrSecurityConfigurationLister) AwsEmrSecurityConfigurations(namespace string) AwsEmrSecurityConfigurationNamespaceLister {
	return awsEmrSecurityConfigurationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsEmrSecurityConfigurationNamespaceLister helps list and get AwsEmrSecurityConfigurations.
type AwsEmrSecurityConfigurationNamespaceLister interface {
	// List lists all AwsEmrSecurityConfigurations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsEmrSecurityConfiguration, err error)
	// Get retrieves the AwsEmrSecurityConfiguration from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsEmrSecurityConfiguration, error)
	AwsEmrSecurityConfigurationNamespaceListerExpansion
}

// awsEmrSecurityConfigurationNamespaceLister implements the AwsEmrSecurityConfigurationNamespaceLister
// interface.
type awsEmrSecurityConfigurationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsEmrSecurityConfigurations in the indexer for a given namespace.
func (s awsEmrSecurityConfigurationNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsEmrSecurityConfiguration, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsEmrSecurityConfiguration))
	})
	return ret, err
}

// Get retrieves the AwsEmrSecurityConfiguration from the indexer for a given namespace and name.
func (s awsEmrSecurityConfigurationNamespaceLister) Get(name string) (*v1.AwsEmrSecurityConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsemrsecurityconfiguration"), name)
	}
	return obj.(*v1.AwsEmrSecurityConfiguration), nil
}
