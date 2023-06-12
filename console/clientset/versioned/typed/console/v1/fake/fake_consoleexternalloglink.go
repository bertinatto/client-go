// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/openshift/api/console/v1"
	consolev1 "github.com/openshift/client-go/console/applyconfigurations/console/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeConsoleExternalLogLinks implements ConsoleExternalLogLinkInterface
type FakeConsoleExternalLogLinks struct {
	Fake *FakeConsoleV1
}

var consoleexternalloglinksResource = v1.SchemeGroupVersion.WithResource("consoleexternalloglinks")

var consoleexternalloglinksKind = v1.SchemeGroupVersion.WithKind("ConsoleExternalLogLink")

// Get takes name of the consoleExternalLogLink, and returns the corresponding consoleExternalLogLink object, and an error if there is any.
func (c *FakeConsoleExternalLogLinks) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ConsoleExternalLogLink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(consoleexternalloglinksResource, name), &v1.ConsoleExternalLogLink{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ConsoleExternalLogLink), err
}

// List takes label and field selectors, and returns the list of ConsoleExternalLogLinks that match those selectors.
func (c *FakeConsoleExternalLogLinks) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ConsoleExternalLogLinkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(consoleexternalloglinksResource, consoleexternalloglinksKind, opts), &v1.ConsoleExternalLogLinkList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ConsoleExternalLogLinkList{ListMeta: obj.(*v1.ConsoleExternalLogLinkList).ListMeta}
	for _, item := range obj.(*v1.ConsoleExternalLogLinkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested consoleExternalLogLinks.
func (c *FakeConsoleExternalLogLinks) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(consoleexternalloglinksResource, opts))
}

// Create takes the representation of a consoleExternalLogLink and creates it.  Returns the server's representation of the consoleExternalLogLink, and an error, if there is any.
func (c *FakeConsoleExternalLogLinks) Create(ctx context.Context, consoleExternalLogLink *v1.ConsoleExternalLogLink, opts metav1.CreateOptions) (result *v1.ConsoleExternalLogLink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(consoleexternalloglinksResource, consoleExternalLogLink), &v1.ConsoleExternalLogLink{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ConsoleExternalLogLink), err
}

// Update takes the representation of a consoleExternalLogLink and updates it. Returns the server's representation of the consoleExternalLogLink, and an error, if there is any.
func (c *FakeConsoleExternalLogLinks) Update(ctx context.Context, consoleExternalLogLink *v1.ConsoleExternalLogLink, opts metav1.UpdateOptions) (result *v1.ConsoleExternalLogLink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(consoleexternalloglinksResource, consoleExternalLogLink), &v1.ConsoleExternalLogLink{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ConsoleExternalLogLink), err
}

// Delete takes name of the consoleExternalLogLink and deletes it. Returns an error if one occurs.
func (c *FakeConsoleExternalLogLinks) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(consoleexternalloglinksResource, name, opts), &v1.ConsoleExternalLogLink{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConsoleExternalLogLinks) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(consoleexternalloglinksResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ConsoleExternalLogLinkList{})
	return err
}

// Patch applies the patch and returns the patched consoleExternalLogLink.
func (c *FakeConsoleExternalLogLinks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ConsoleExternalLogLink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(consoleexternalloglinksResource, name, pt, data, subresources...), &v1.ConsoleExternalLogLink{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ConsoleExternalLogLink), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied consoleExternalLogLink.
func (c *FakeConsoleExternalLogLinks) Apply(ctx context.Context, consoleExternalLogLink *consolev1.ConsoleExternalLogLinkApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ConsoleExternalLogLink, err error) {
	if consoleExternalLogLink == nil {
		return nil, fmt.Errorf("consoleExternalLogLink provided to Apply must not be nil")
	}
	data, err := json.Marshal(consoleExternalLogLink)
	if err != nil {
		return nil, err
	}
	name := consoleExternalLogLink.Name
	if name == nil {
		return nil, fmt.Errorf("consoleExternalLogLink.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(consoleexternalloglinksResource, *name, types.ApplyPatchType, data), &v1.ConsoleExternalLogLink{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ConsoleExternalLogLink), err
}
