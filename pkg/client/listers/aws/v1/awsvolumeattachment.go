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

// AwsVolumeAttachmentLister helps list AwsVolumeAttachments.
type AwsVolumeAttachmentLister interface {
	// List lists all AwsVolumeAttachments in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsVolumeAttachment, err error)
	// AwsVolumeAttachments returns an object that can list and get AwsVolumeAttachments.
	AwsVolumeAttachments(namespace string) AwsVolumeAttachmentNamespaceLister
	AwsVolumeAttachmentListerExpansion
}

// awsVolumeAttachmentLister implements the AwsVolumeAttachmentLister interface.
type awsVolumeAttachmentLister struct {
	indexer cache.Indexer
}

// NewAwsVolumeAttachmentLister returns a new AwsVolumeAttachmentLister.
func NewAwsVolumeAttachmentLister(indexer cache.Indexer) AwsVolumeAttachmentLister {
	return &awsVolumeAttachmentLister{indexer: indexer}
}

// List lists all AwsVolumeAttachments in the indexer.
func (s *awsVolumeAttachmentLister) List(selector labels.Selector) (ret []*v1.AwsVolumeAttachment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsVolumeAttachment))
	})
	return ret, err
}

// AwsVolumeAttachments returns an object that can list and get AwsVolumeAttachments.
func (s *awsVolumeAttachmentLister) AwsVolumeAttachments(namespace string) AwsVolumeAttachmentNamespaceLister {
	return awsVolumeAttachmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsVolumeAttachmentNamespaceLister helps list and get AwsVolumeAttachments.
type AwsVolumeAttachmentNamespaceLister interface {
	// List lists all AwsVolumeAttachments in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsVolumeAttachment, err error)
	// Get retrieves the AwsVolumeAttachment from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsVolumeAttachment, error)
	AwsVolumeAttachmentNamespaceListerExpansion
}

// awsVolumeAttachmentNamespaceLister implements the AwsVolumeAttachmentNamespaceLister
// interface.
type awsVolumeAttachmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsVolumeAttachments in the indexer for a given namespace.
func (s awsVolumeAttachmentNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsVolumeAttachment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsVolumeAttachment))
	})
	return ret, err
}

// Get retrieves the AwsVolumeAttachment from the indexer for a given namespace and name.
func (s awsVolumeAttachmentNamespaceLister) Get(name string) (*v1.AwsVolumeAttachment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsvolumeattachment"), name)
	}
	return obj.(*v1.AwsVolumeAttachment), nil
}
