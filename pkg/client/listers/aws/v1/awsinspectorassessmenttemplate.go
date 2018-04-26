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

// AwsInspectorAssessmentTemplateLister helps list AwsInspectorAssessmentTemplates.
type AwsInspectorAssessmentTemplateLister interface {
	// List lists all AwsInspectorAssessmentTemplates in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsInspectorAssessmentTemplate, err error)
	// AwsInspectorAssessmentTemplates returns an object that can list and get AwsInspectorAssessmentTemplates.
	AwsInspectorAssessmentTemplates(namespace string) AwsInspectorAssessmentTemplateNamespaceLister
	AwsInspectorAssessmentTemplateListerExpansion
}

// awsInspectorAssessmentTemplateLister implements the AwsInspectorAssessmentTemplateLister interface.
type awsInspectorAssessmentTemplateLister struct {
	indexer cache.Indexer
}

// NewAwsInspectorAssessmentTemplateLister returns a new AwsInspectorAssessmentTemplateLister.
func NewAwsInspectorAssessmentTemplateLister(indexer cache.Indexer) AwsInspectorAssessmentTemplateLister {
	return &awsInspectorAssessmentTemplateLister{indexer: indexer}
}

// List lists all AwsInspectorAssessmentTemplates in the indexer.
func (s *awsInspectorAssessmentTemplateLister) List(selector labels.Selector) (ret []*v1.AwsInspectorAssessmentTemplate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsInspectorAssessmentTemplate))
	})
	return ret, err
}

// AwsInspectorAssessmentTemplates returns an object that can list and get AwsInspectorAssessmentTemplates.
func (s *awsInspectorAssessmentTemplateLister) AwsInspectorAssessmentTemplates(namespace string) AwsInspectorAssessmentTemplateNamespaceLister {
	return awsInspectorAssessmentTemplateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsInspectorAssessmentTemplateNamespaceLister helps list and get AwsInspectorAssessmentTemplates.
type AwsInspectorAssessmentTemplateNamespaceLister interface {
	// List lists all AwsInspectorAssessmentTemplates in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsInspectorAssessmentTemplate, err error)
	// Get retrieves the AwsInspectorAssessmentTemplate from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsInspectorAssessmentTemplate, error)
	AwsInspectorAssessmentTemplateNamespaceListerExpansion
}

// awsInspectorAssessmentTemplateNamespaceLister implements the AwsInspectorAssessmentTemplateNamespaceLister
// interface.
type awsInspectorAssessmentTemplateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsInspectorAssessmentTemplates in the indexer for a given namespace.
func (s awsInspectorAssessmentTemplateNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsInspectorAssessmentTemplate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsInspectorAssessmentTemplate))
	})
	return ret, err
}

// Get retrieves the AwsInspectorAssessmentTemplate from the indexer for a given namespace and name.
func (s awsInspectorAssessmentTemplateNamespaceLister) Get(name string) (*v1.AwsInspectorAssessmentTemplate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsinspectorassessmenttemplate"), name)
	}
	return obj.(*v1.AwsInspectorAssessmentTemplate), nil
}
