/*
Copyright 2019-2021 OpenFaaS Authors

Licensed under the MIT license. See LICENSE file in the project root for full license information.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	iamv1 "github.com/innoobijr/faas-netes/pkg/apis/iam/v1"
	versioned "github.com/innoobijr/faas-netes/pkg/client/clientset/versioned"
	internalinterfaces "github.com/innoobijr/faas-netes/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/innoobijr/faas-netes/pkg/client/listers/iam/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PolicyInformer provides access to a shared informer and lister for
// Policies.
type PolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.PolicyLister
}

type policyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPolicyInformer constructs a new informer for Policy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPolicyInformer constructs a new informer for Policy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IamV1().Policies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IamV1().Policies(namespace).Watch(context.TODO(), options)
			},
		},
		&iamv1.Policy{},
		resyncPeriod,
		indexers,
	)
}

func (f *policyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *policyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&iamv1.Policy{}, f.defaultInformer)
}

func (f *policyInformer) Lister() v1.PolicyLister {
	return v1.NewPolicyLister(f.Informer().GetIndexer())
}
