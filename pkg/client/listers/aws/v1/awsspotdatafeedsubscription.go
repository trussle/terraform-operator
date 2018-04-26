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

// AwsSpotDatafeedSubscriptionLister helps list AwsSpotDatafeedSubscriptions.
type AwsSpotDatafeedSubscriptionLister interface {
	// List lists all AwsSpotDatafeedSubscriptions in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsSpotDatafeedSubscription, err error)
	// AwsSpotDatafeedSubscriptions returns an object that can list and get AwsSpotDatafeedSubscriptions.
	AwsSpotDatafeedSubscriptions(namespace string) AwsSpotDatafeedSubscriptionNamespaceLister
	AwsSpotDatafeedSubscriptionListerExpansion
}

// awsSpotDatafeedSubscriptionLister implements the AwsSpotDatafeedSubscriptionLister interface.
type awsSpotDatafeedSubscriptionLister struct {
	indexer cache.Indexer
}

// NewAwsSpotDatafeedSubscriptionLister returns a new AwsSpotDatafeedSubscriptionLister.
func NewAwsSpotDatafeedSubscriptionLister(indexer cache.Indexer) AwsSpotDatafeedSubscriptionLister {
	return &awsSpotDatafeedSubscriptionLister{indexer: indexer}
}

// List lists all AwsSpotDatafeedSubscriptions in the indexer.
func (s *awsSpotDatafeedSubscriptionLister) List(selector labels.Selector) (ret []*v1.AwsSpotDatafeedSubscription, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsSpotDatafeedSubscription))
	})
	return ret, err
}

// AwsSpotDatafeedSubscriptions returns an object that can list and get AwsSpotDatafeedSubscriptions.
func (s *awsSpotDatafeedSubscriptionLister) AwsSpotDatafeedSubscriptions(namespace string) AwsSpotDatafeedSubscriptionNamespaceLister {
	return awsSpotDatafeedSubscriptionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsSpotDatafeedSubscriptionNamespaceLister helps list and get AwsSpotDatafeedSubscriptions.
type AwsSpotDatafeedSubscriptionNamespaceLister interface {
	// List lists all AwsSpotDatafeedSubscriptions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsSpotDatafeedSubscription, err error)
	// Get retrieves the AwsSpotDatafeedSubscription from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsSpotDatafeedSubscription, error)
	AwsSpotDatafeedSubscriptionNamespaceListerExpansion
}

// awsSpotDatafeedSubscriptionNamespaceLister implements the AwsSpotDatafeedSubscriptionNamespaceLister
// interface.
type awsSpotDatafeedSubscriptionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsSpotDatafeedSubscriptions in the indexer for a given namespace.
func (s awsSpotDatafeedSubscriptionNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsSpotDatafeedSubscription, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsSpotDatafeedSubscription))
	})
	return ret, err
}

// Get retrieves the AwsSpotDatafeedSubscription from the indexer for a given namespace and name.
func (s awsSpotDatafeedSubscriptionNamespaceLister) Get(name string) (*v1.AwsSpotDatafeedSubscription, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsspotdatafeedsubscription"), name)
	}
	return obj.(*v1.AwsSpotDatafeedSubscription), nil
}
