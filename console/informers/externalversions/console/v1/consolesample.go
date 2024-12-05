// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	context "context"
	time "time"

	apiconsolev1 "github.com/openshift/api/console/v1"
	versioned "github.com/openshift/client-go/console/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/console/informers/externalversions/internalinterfaces"
	consolev1 "github.com/openshift/client-go/console/listers/console/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ConsoleSampleInformer provides access to a shared informer and lister for
// ConsoleSamples.
type ConsoleSampleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() consolev1.ConsoleSampleLister
}

type consoleSampleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewConsoleSampleInformer constructs a new informer for ConsoleSample type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewConsoleSampleInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredConsoleSampleInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredConsoleSampleInformer constructs a new informer for ConsoleSample type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredConsoleSampleInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConsoleV1().ConsoleSamples().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConsoleV1().ConsoleSamples().Watch(context.TODO(), options)
			},
		},
		&apiconsolev1.ConsoleSample{},
		resyncPeriod,
		indexers,
	)
}

func (f *consoleSampleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredConsoleSampleInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *consoleSampleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiconsolev1.ConsoleSample{}, f.defaultInformer)
}

func (f *consoleSampleInformer) Lister() consolev1.ConsoleSampleLister {
	return consolev1.NewConsoleSampleLister(f.Informer().GetIndexer())
}
