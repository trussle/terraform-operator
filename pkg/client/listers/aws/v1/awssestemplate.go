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

// AwsSesTemplateLister helps list AwsSesTemplates.
type AwsSesTemplateLister interface {
	// List lists all AwsSesTemplates in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsSesTemplate, err error)
	// AwsSesTemplates returns an object that can list and get AwsSesTemplates.
	AwsSesTemplates(namespace string) AwsSesTemplateNamespaceLister
	AwsSesTemplateListerExpansion
}

// awsSesTemplateLister implements the AwsSesTemplateLister interface.
type awsSesTemplateLister struct {
	indexer cache.Indexer
}

// NewAwsSesTemplateLister returns a new AwsSesTemplateLister.
func NewAwsSesTemplateLister(indexer cache.Indexer) AwsSesTemplateLister {
	return &awsSesTemplateLister{indexer: indexer}
}

// List lists all AwsSesTemplates in the indexer.
func (s *awsSesTemplateLister) List(selector labels.Selector) (ret []*v1.AwsSesTemplate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsSesTemplate))
	})
	return ret, err
}

// AwsSesTemplates returns an object that can list and get AwsSesTemplates.
func (s *awsSesTemplateLister) AwsSesTemplates(namespace string) AwsSesTemplateNamespaceLister {
	return awsSesTemplateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsSesTemplateNamespaceLister helps list and get AwsSesTemplates.
type AwsSesTemplateNamespaceLister interface {
	// List lists all AwsSesTemplates in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsSesTemplate, err error)
	// Get retrieves the AwsSesTemplate from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsSesTemplate, error)
	AwsSesTemplateNamespaceListerExpansion
}

// awsSesTemplateNamespaceLister implements the AwsSesTemplateNamespaceLister
// interface.
type awsSesTemplateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsSesTemplates in the indexer for a given namespace.
func (s awsSesTemplateNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsSesTemplate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsSesTemplate))
	})
	return ret, err
}

// Get retrieves the AwsSesTemplate from the indexer for a given namespace and name.
func (s awsSesTemplateNamespaceLister) Get(name string) (*v1.AwsSesTemplate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awssestemplate"), name)
	}
	return obj.(*v1.AwsSesTemplate), nil
}
