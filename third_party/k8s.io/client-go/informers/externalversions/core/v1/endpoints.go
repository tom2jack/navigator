/*
Copyright 2017 Jetstack Ltd.

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

// This file was automatically generated by informer-gen

package v1

import (
	internalinterfaces "github.com/jetstack/navigator/third_party/k8s.io/client-go/informers/externalversions/internalinterfaces"
	core_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	kubernetes "k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/listers/core/v1"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// EndpointsInformer provides access to a shared informer and lister for
// Endpoints.
type EndpointsInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.EndpointsLister
}

type endpointsInformer struct {
	factory internalinterfaces.SharedInformerFactory
	filter  internalinterfaces.FilterFunc
}

// NewEndpointsInformer constructs a new informer for Endpoints type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewEndpointsInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	filter := internalinterfaces.NamespaceFilter(namespace)
	return NewFilteredEndpointsInformer(client, filter, resyncPeriod, indexers)
}

// NewFilteredEndpointsInformer constructs a new informer for Endpoints type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredEndpointsInformer(client kubernetes.Interface, filter internalinterfaces.FilterFunc, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				namespace := filter(&options)
				return client.CoreV1().Endpoints(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				namespace := filter(&options)
				return client.CoreV1().Endpoints(namespace).Watch(options)
			},
		},
		&core_v1.Endpoints{},
		resyncPeriod,
		indexers,
	)
}

func (f *endpointsInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredEndpointsInformer(client, f.filter, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *endpointsInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&core_v1.Endpoints{}, f.defaultInformer)
}

func (f *endpointsInformer) Lister() v1.EndpointsLister {
	return v1.NewEndpointsLister(f.Informer().GetIndexer())
}
