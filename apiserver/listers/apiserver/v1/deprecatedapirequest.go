// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/apiserver/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DeprecatedAPIRequestLister helps list DeprecatedAPIRequests.
// All objects returned here must be treated as read-only.
type DeprecatedAPIRequestLister interface {
	// List lists all DeprecatedAPIRequests in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.DeprecatedAPIRequest, err error)
	// DeprecatedAPIRequests returns an object that can list and get DeprecatedAPIRequests.
	DeprecatedAPIRequests(namespace string) DeprecatedAPIRequestNamespaceLister
	DeprecatedAPIRequestListerExpansion
}

// deprecatedAPIRequestLister implements the DeprecatedAPIRequestLister interface.
type deprecatedAPIRequestLister struct {
	indexer cache.Indexer
}

// NewDeprecatedAPIRequestLister returns a new DeprecatedAPIRequestLister.
func NewDeprecatedAPIRequestLister(indexer cache.Indexer) DeprecatedAPIRequestLister {
	return &deprecatedAPIRequestLister{indexer: indexer}
}

// List lists all DeprecatedAPIRequests in the indexer.
func (s *deprecatedAPIRequestLister) List(selector labels.Selector) (ret []*v1.DeprecatedAPIRequest, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.DeprecatedAPIRequest))
	})
	return ret, err
}

// DeprecatedAPIRequests returns an object that can list and get DeprecatedAPIRequests.
func (s *deprecatedAPIRequestLister) DeprecatedAPIRequests(namespace string) DeprecatedAPIRequestNamespaceLister {
	return deprecatedAPIRequestNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DeprecatedAPIRequestNamespaceLister helps list and get DeprecatedAPIRequests.
// All objects returned here must be treated as read-only.
type DeprecatedAPIRequestNamespaceLister interface {
	// List lists all DeprecatedAPIRequests in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.DeprecatedAPIRequest, err error)
	// Get retrieves the DeprecatedAPIRequest from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.DeprecatedAPIRequest, error)
	DeprecatedAPIRequestNamespaceListerExpansion
}

// deprecatedAPIRequestNamespaceLister implements the DeprecatedAPIRequestNamespaceLister
// interface.
type deprecatedAPIRequestNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DeprecatedAPIRequests in the indexer for a given namespace.
func (s deprecatedAPIRequestNamespaceLister) List(selector labels.Selector) (ret []*v1.DeprecatedAPIRequest, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.DeprecatedAPIRequest))
	})
	return ret, err
}

// Get retrieves the DeprecatedAPIRequest from the indexer for a given namespace and name.
func (s deprecatedAPIRequestNamespaceLister) Get(name string) (*v1.DeprecatedAPIRequest, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("deprecatedapirequest"), name)
	}
	return obj.(*v1.DeprecatedAPIRequest), nil
}