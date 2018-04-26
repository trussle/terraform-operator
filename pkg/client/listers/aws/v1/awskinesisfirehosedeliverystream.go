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

// AwsKinesisFirehoseDeliveryStreamLister helps list AwsKinesisFirehoseDeliveryStreams.
type AwsKinesisFirehoseDeliveryStreamLister interface {
	// List lists all AwsKinesisFirehoseDeliveryStreams in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsKinesisFirehoseDeliveryStream, err error)
	// AwsKinesisFirehoseDeliveryStreams returns an object that can list and get AwsKinesisFirehoseDeliveryStreams.
	AwsKinesisFirehoseDeliveryStreams(namespace string) AwsKinesisFirehoseDeliveryStreamNamespaceLister
	AwsKinesisFirehoseDeliveryStreamListerExpansion
}

// awsKinesisFirehoseDeliveryStreamLister implements the AwsKinesisFirehoseDeliveryStreamLister interface.
type awsKinesisFirehoseDeliveryStreamLister struct {
	indexer cache.Indexer
}

// NewAwsKinesisFirehoseDeliveryStreamLister returns a new AwsKinesisFirehoseDeliveryStreamLister.
func NewAwsKinesisFirehoseDeliveryStreamLister(indexer cache.Indexer) AwsKinesisFirehoseDeliveryStreamLister {
	return &awsKinesisFirehoseDeliveryStreamLister{indexer: indexer}
}

// List lists all AwsKinesisFirehoseDeliveryStreams in the indexer.
func (s *awsKinesisFirehoseDeliveryStreamLister) List(selector labels.Selector) (ret []*v1.AwsKinesisFirehoseDeliveryStream, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsKinesisFirehoseDeliveryStream))
	})
	return ret, err
}

// AwsKinesisFirehoseDeliveryStreams returns an object that can list and get AwsKinesisFirehoseDeliveryStreams.
func (s *awsKinesisFirehoseDeliveryStreamLister) AwsKinesisFirehoseDeliveryStreams(namespace string) AwsKinesisFirehoseDeliveryStreamNamespaceLister {
	return awsKinesisFirehoseDeliveryStreamNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsKinesisFirehoseDeliveryStreamNamespaceLister helps list and get AwsKinesisFirehoseDeliveryStreams.
type AwsKinesisFirehoseDeliveryStreamNamespaceLister interface {
	// List lists all AwsKinesisFirehoseDeliveryStreams in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsKinesisFirehoseDeliveryStream, err error)
	// Get retrieves the AwsKinesisFirehoseDeliveryStream from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsKinesisFirehoseDeliveryStream, error)
	AwsKinesisFirehoseDeliveryStreamNamespaceListerExpansion
}

// awsKinesisFirehoseDeliveryStreamNamespaceLister implements the AwsKinesisFirehoseDeliveryStreamNamespaceLister
// interface.
type awsKinesisFirehoseDeliveryStreamNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsKinesisFirehoseDeliveryStreams in the indexer for a given namespace.
func (s awsKinesisFirehoseDeliveryStreamNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsKinesisFirehoseDeliveryStream, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsKinesisFirehoseDeliveryStream))
	})
	return ret, err
}

// Get retrieves the AwsKinesisFirehoseDeliveryStream from the indexer for a given namespace and name.
func (s awsKinesisFirehoseDeliveryStreamNamespaceLister) Get(name string) (*v1.AwsKinesisFirehoseDeliveryStream, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awskinesisfirehosedeliverystream"), name)
	}
	return obj.(*v1.AwsKinesisFirehoseDeliveryStream), nil
}
