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

// AwsRouteLister helps list AwsRoutes.
type AwsRouteLister interface {
	// List lists all AwsRoutes in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsRoute, err error)
	// AwsRoutes returns an object that can list and get AwsRoutes.
	AwsRoutes(namespace string) AwsRouteNamespaceLister
	AwsRouteListerExpansion
}

// awsRouteLister implements the AwsRouteLister interface.
type awsRouteLister struct {
	indexer cache.Indexer
}

// NewAwsRouteLister returns a new AwsRouteLister.
func NewAwsRouteLister(indexer cache.Indexer) AwsRouteLister {
	return &awsRouteLister{indexer: indexer}
}

// List lists all AwsRoutes in the indexer.
func (s *awsRouteLister) List(selector labels.Selector) (ret []*v1.AwsRoute, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsRoute))
	})
	return ret, err
}

// AwsRoutes returns an object that can list and get AwsRoutes.
func (s *awsRouteLister) AwsRoutes(namespace string) AwsRouteNamespaceLister {
	return awsRouteNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsRouteNamespaceLister helps list and get AwsRoutes.
type AwsRouteNamespaceLister interface {
	// List lists all AwsRoutes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsRoute, err error)
	// Get retrieves the AwsRoute from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsRoute, error)
	AwsRouteNamespaceListerExpansion
}

// awsRouteNamespaceLister implements the AwsRouteNamespaceLister
// interface.
type awsRouteNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsRoutes in the indexer for a given namespace.
func (s awsRouteNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsRoute, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsRoute))
	})
	return ret, err
}

// Get retrieves the AwsRoute from the indexer for a given namespace and name.
func (s awsRouteNamespaceLister) Get(name string) (*v1.AwsRoute, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsroute"), name)
	}
	return obj.(*v1.AwsRoute), nil
}
