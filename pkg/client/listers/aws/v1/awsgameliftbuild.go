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

// AwsGameliftBuildLister helps list AwsGameliftBuilds.
type AwsGameliftBuildLister interface {
	// List lists all AwsGameliftBuilds in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsGameliftBuild, err error)
	// AwsGameliftBuilds returns an object that can list and get AwsGameliftBuilds.
	AwsGameliftBuilds(namespace string) AwsGameliftBuildNamespaceLister
	AwsGameliftBuildListerExpansion
}

// awsGameliftBuildLister implements the AwsGameliftBuildLister interface.
type awsGameliftBuildLister struct {
	indexer cache.Indexer
}

// NewAwsGameliftBuildLister returns a new AwsGameliftBuildLister.
func NewAwsGameliftBuildLister(indexer cache.Indexer) AwsGameliftBuildLister {
	return &awsGameliftBuildLister{indexer: indexer}
}

// List lists all AwsGameliftBuilds in the indexer.
func (s *awsGameliftBuildLister) List(selector labels.Selector) (ret []*v1.AwsGameliftBuild, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsGameliftBuild))
	})
	return ret, err
}

// AwsGameliftBuilds returns an object that can list and get AwsGameliftBuilds.
func (s *awsGameliftBuildLister) AwsGameliftBuilds(namespace string) AwsGameliftBuildNamespaceLister {
	return awsGameliftBuildNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsGameliftBuildNamespaceLister helps list and get AwsGameliftBuilds.
type AwsGameliftBuildNamespaceLister interface {
	// List lists all AwsGameliftBuilds in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsGameliftBuild, err error)
	// Get retrieves the AwsGameliftBuild from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsGameliftBuild, error)
	AwsGameliftBuildNamespaceListerExpansion
}

// awsGameliftBuildNamespaceLister implements the AwsGameliftBuildNamespaceLister
// interface.
type awsGameliftBuildNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsGameliftBuilds in the indexer for a given namespace.
func (s awsGameliftBuildNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsGameliftBuild, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsGameliftBuild))
	})
	return ret, err
}

// Get retrieves the AwsGameliftBuild from the indexer for a given namespace and name.
func (s awsGameliftBuildNamespaceLister) Get(name string) (*v1.AwsGameliftBuild, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsgameliftbuild"), name)
	}
	return obj.(*v1.AwsGameliftBuild), nil
}
