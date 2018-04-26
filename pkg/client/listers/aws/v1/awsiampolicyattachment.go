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

// AwsIamPolicyAttachmentLister helps list AwsIamPolicyAttachments.
type AwsIamPolicyAttachmentLister interface {
	// List lists all AwsIamPolicyAttachments in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsIamPolicyAttachment, err error)
	// AwsIamPolicyAttachments returns an object that can list and get AwsIamPolicyAttachments.
	AwsIamPolicyAttachments(namespace string) AwsIamPolicyAttachmentNamespaceLister
	AwsIamPolicyAttachmentListerExpansion
}

// awsIamPolicyAttachmentLister implements the AwsIamPolicyAttachmentLister interface.
type awsIamPolicyAttachmentLister struct {
	indexer cache.Indexer
}

// NewAwsIamPolicyAttachmentLister returns a new AwsIamPolicyAttachmentLister.
func NewAwsIamPolicyAttachmentLister(indexer cache.Indexer) AwsIamPolicyAttachmentLister {
	return &awsIamPolicyAttachmentLister{indexer: indexer}
}

// List lists all AwsIamPolicyAttachments in the indexer.
func (s *awsIamPolicyAttachmentLister) List(selector labels.Selector) (ret []*v1.AwsIamPolicyAttachment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsIamPolicyAttachment))
	})
	return ret, err
}

// AwsIamPolicyAttachments returns an object that can list and get AwsIamPolicyAttachments.
func (s *awsIamPolicyAttachmentLister) AwsIamPolicyAttachments(namespace string) AwsIamPolicyAttachmentNamespaceLister {
	return awsIamPolicyAttachmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsIamPolicyAttachmentNamespaceLister helps list and get AwsIamPolicyAttachments.
type AwsIamPolicyAttachmentNamespaceLister interface {
	// List lists all AwsIamPolicyAttachments in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsIamPolicyAttachment, err error)
	// Get retrieves the AwsIamPolicyAttachment from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsIamPolicyAttachment, error)
	AwsIamPolicyAttachmentNamespaceListerExpansion
}

// awsIamPolicyAttachmentNamespaceLister implements the AwsIamPolicyAttachmentNamespaceLister
// interface.
type awsIamPolicyAttachmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsIamPolicyAttachments in the indexer for a given namespace.
func (s awsIamPolicyAttachmentNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsIamPolicyAttachment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsIamPolicyAttachment))
	})
	return ret, err
}

// Get retrieves the AwsIamPolicyAttachment from the indexer for a given namespace and name.
func (s awsIamPolicyAttachmentNamespaceLister) Get(name string) (*v1.AwsIamPolicyAttachment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsiampolicyattachment"), name)
	}
	return obj.(*v1.AwsIamPolicyAttachment), nil
}
