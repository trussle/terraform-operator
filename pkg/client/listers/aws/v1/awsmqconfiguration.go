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

// AwsMqConfigurationLister helps list AwsMqConfigurations.
type AwsMqConfigurationLister interface {
	// List lists all AwsMqConfigurations in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsMqConfiguration, err error)
	// AwsMqConfigurations returns an object that can list and get AwsMqConfigurations.
	AwsMqConfigurations(namespace string) AwsMqConfigurationNamespaceLister
	AwsMqConfigurationListerExpansion
}

// awsMqConfigurationLister implements the AwsMqConfigurationLister interface.
type awsMqConfigurationLister struct {
	indexer cache.Indexer
}

// NewAwsMqConfigurationLister returns a new AwsMqConfigurationLister.
func NewAwsMqConfigurationLister(indexer cache.Indexer) AwsMqConfigurationLister {
	return &awsMqConfigurationLister{indexer: indexer}
}

// List lists all AwsMqConfigurations in the indexer.
func (s *awsMqConfigurationLister) List(selector labels.Selector) (ret []*v1.AwsMqConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsMqConfiguration))
	})
	return ret, err
}

// AwsMqConfigurations returns an object that can list and get AwsMqConfigurations.
func (s *awsMqConfigurationLister) AwsMqConfigurations(namespace string) AwsMqConfigurationNamespaceLister {
	return awsMqConfigurationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsMqConfigurationNamespaceLister helps list and get AwsMqConfigurations.
type AwsMqConfigurationNamespaceLister interface {
	// List lists all AwsMqConfigurations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsMqConfiguration, err error)
	// Get retrieves the AwsMqConfiguration from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsMqConfiguration, error)
	AwsMqConfigurationNamespaceListerExpansion
}

// awsMqConfigurationNamespaceLister implements the AwsMqConfigurationNamespaceLister
// interface.
type awsMqConfigurationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsMqConfigurations in the indexer for a given namespace.
func (s awsMqConfigurationNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsMqConfiguration, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsMqConfiguration))
	})
	return ret, err
}

// Get retrieves the AwsMqConfiguration from the indexer for a given namespace and name.
func (s awsMqConfigurationNamespaceLister) Get(name string) (*v1.AwsMqConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsmqconfiguration"), name)
	}
	return obj.(*v1.AwsMqConfiguration), nil
}