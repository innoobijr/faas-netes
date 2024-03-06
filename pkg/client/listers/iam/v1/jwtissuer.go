/*
Copyright 2019-2021 OpenFaaS Authors

Licensed under the MIT license. See LICENSE file in the project root for full license information.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/innoobijr/faas-netes/pkg/apis/iam/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// JwtIssuerLister helps list JwtIssuers.
// All objects returned here must be treated as read-only.
type JwtIssuerLister interface {
	// List lists all JwtIssuers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.JwtIssuer, err error)
	// JwtIssuers returns an object that can list and get JwtIssuers.
	JwtIssuers(namespace string) JwtIssuerNamespaceLister
	JwtIssuerListerExpansion
}

// jwtIssuerLister implements the JwtIssuerLister interface.
type jwtIssuerLister struct {
	indexer cache.Indexer
}

// NewJwtIssuerLister returns a new JwtIssuerLister.
func NewJwtIssuerLister(indexer cache.Indexer) JwtIssuerLister {
	return &jwtIssuerLister{indexer: indexer}
}

// List lists all JwtIssuers in the indexer.
func (s *jwtIssuerLister) List(selector labels.Selector) (ret []*v1.JwtIssuer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.JwtIssuer))
	})
	return ret, err
}

// JwtIssuers returns an object that can list and get JwtIssuers.
func (s *jwtIssuerLister) JwtIssuers(namespace string) JwtIssuerNamespaceLister {
	return jwtIssuerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// JwtIssuerNamespaceLister helps list and get JwtIssuers.
// All objects returned here must be treated as read-only.
type JwtIssuerNamespaceLister interface {
	// List lists all JwtIssuers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.JwtIssuer, err error)
	// Get retrieves the JwtIssuer from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.JwtIssuer, error)
	JwtIssuerNamespaceListerExpansion
}

// jwtIssuerNamespaceLister implements the JwtIssuerNamespaceLister
// interface.
type jwtIssuerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all JwtIssuers in the indexer for a given namespace.
func (s jwtIssuerNamespaceLister) List(selector labels.Selector) (ret []*v1.JwtIssuer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.JwtIssuer))
	})
	return ret, err
}

// Get retrieves the JwtIssuer from the indexer for a given namespace and name.
func (s jwtIssuerNamespaceLister) Get(name string) (*v1.JwtIssuer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("jwtissuer"), name)
	}
	return obj.(*v1.JwtIssuer), nil
}
