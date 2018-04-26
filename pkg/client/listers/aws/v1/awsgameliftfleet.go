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

// AwsGameliftFleetLister helps list AwsGameliftFleets.
type AwsGameliftFleetLister interface {
	// List lists all AwsGameliftFleets in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsGameliftFleet, err error)
	// AwsGameliftFleets returns an object that can list and get AwsGameliftFleets.
	AwsGameliftFleets(namespace string) AwsGameliftFleetNamespaceLister
	AwsGameliftFleetListerExpansion
}

// awsGameliftFleetLister implements the AwsGameliftFleetLister interface.
type awsGameliftFleetLister struct {
	indexer cache.Indexer
}

// NewAwsGameliftFleetLister returns a new AwsGameliftFleetLister.
func NewAwsGameliftFleetLister(indexer cache.Indexer) AwsGameliftFleetLister {
	return &awsGameliftFleetLister{indexer: indexer}
}

// List lists all AwsGameliftFleets in the indexer.
func (s *awsGameliftFleetLister) List(selector labels.Selector) (ret []*v1.AwsGameliftFleet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsGameliftFleet))
	})
	return ret, err
}

// AwsGameliftFleets returns an object that can list and get AwsGameliftFleets.
func (s *awsGameliftFleetLister) AwsGameliftFleets(namespace string) AwsGameliftFleetNamespaceLister {
	return awsGameliftFleetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsGameliftFleetNamespaceLister helps list and get AwsGameliftFleets.
type AwsGameliftFleetNamespaceLister interface {
	// List lists all AwsGameliftFleets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsGameliftFleet, err error)
	// Get retrieves the AwsGameliftFleet from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsGameliftFleet, error)
	AwsGameliftFleetNamespaceListerExpansion
}

// awsGameliftFleetNamespaceLister implements the AwsGameliftFleetNamespaceLister
// interface.
type awsGameliftFleetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsGameliftFleets in the indexer for a given namespace.
func (s awsGameliftFleetNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsGameliftFleet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsGameliftFleet))
	})
	return ret, err
}

// Get retrieves the AwsGameliftFleet from the indexer for a given namespace and name.
func (s awsGameliftFleetNamespaceLister) Get(name string) (*v1.AwsGameliftFleet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsgameliftfleet"), name)
	}
	return obj.(*v1.AwsGameliftFleet), nil
}
