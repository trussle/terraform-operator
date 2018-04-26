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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	aws_v1 "github.com/trussle/terraform-operator/pkg/apis/aws/v1"
	versioned "github.com/trussle/terraform-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/trussle/terraform-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/trussle/terraform-operator/pkg/client/listers/aws/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AwsLaunchTemplateInformer provides access to a shared informer and lister for
// AwsLaunchTemplates.
type AwsLaunchTemplateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.AwsLaunchTemplateLister
}

type awsLaunchTemplateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAwsLaunchTemplateInformer constructs a new informer for AwsLaunchTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAwsLaunchTemplateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAwsLaunchTemplateInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAwsLaunchTemplateInformer constructs a new informer for AwsLaunchTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAwsLaunchTemplateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TrussleV1().AwsLaunchTemplates(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TrussleV1().AwsLaunchTemplates(namespace).Watch(options)
			},
		},
		&aws_v1.AwsLaunchTemplate{},
		resyncPeriod,
		indexers,
	)
}

func (f *awsLaunchTemplateInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAwsLaunchTemplateInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *awsLaunchTemplateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&aws_v1.AwsLaunchTemplate{}, f.defaultInformer)
}

func (f *awsLaunchTemplateInformer) Lister() v1.AwsLaunchTemplateLister {
	return v1.NewAwsLaunchTemplateLister(f.Informer().GetIndexer())
}
