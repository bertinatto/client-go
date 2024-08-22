// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/console/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// ConsoleNotificationLister helps list ConsoleNotifications.
// All objects returned here must be treated as read-only.
type ConsoleNotificationLister interface {
	// List lists all ConsoleNotifications in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ConsoleNotification, err error)
	// Get retrieves the ConsoleNotification from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ConsoleNotification, error)
	ConsoleNotificationListerExpansion
}

// consoleNotificationLister implements the ConsoleNotificationLister interface.
type consoleNotificationLister struct {
	listers.ResourceIndexer[*v1.ConsoleNotification]
}

// NewConsoleNotificationLister returns a new ConsoleNotificationLister.
func NewConsoleNotificationLister(indexer cache.Indexer) ConsoleNotificationLister {
	return &consoleNotificationLister{listers.New[*v1.ConsoleNotification](indexer, v1.Resource("consolenotification"))}
}
