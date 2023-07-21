// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/openshift/api/imageregistry/v1"
	imageregistryv1 "github.com/openshift/client-go/imageregistry/applyconfigurations/imageregistry/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeImagePruners implements ImagePrunerInterface
type FakeImagePruners struct {
	Fake *FakeImageregistryV1
}

var imageprunersResource = v1.SchemeGroupVersion.WithResource("imagepruners")

var imageprunersKind = v1.SchemeGroupVersion.WithKind("ImagePruner")

// Get takes name of the imagePruner, and returns the corresponding imagePruner object, and an error if there is any.
func (c *FakeImagePruners) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ImagePruner, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(imageprunersResource, name), &v1.ImagePruner{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ImagePruner), err
}

// List takes label and field selectors, and returns the list of ImagePruners that match those selectors.
func (c *FakeImagePruners) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ImagePrunerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(imageprunersResource, imageprunersKind, opts), &v1.ImagePrunerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ImagePrunerList{ListMeta: obj.(*v1.ImagePrunerList).ListMeta}
	for _, item := range obj.(*v1.ImagePrunerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested imagePruners.
func (c *FakeImagePruners) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(imageprunersResource, opts))
}

// Create takes the representation of a imagePruner and creates it.  Returns the server's representation of the imagePruner, and an error, if there is any.
func (c *FakeImagePruners) Create(ctx context.Context, imagePruner *v1.ImagePruner, opts metav1.CreateOptions) (result *v1.ImagePruner, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(imageprunersResource, imagePruner), &v1.ImagePruner{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ImagePruner), err
}

// Update takes the representation of a imagePruner and updates it. Returns the server's representation of the imagePruner, and an error, if there is any.
func (c *FakeImagePruners) Update(ctx context.Context, imagePruner *v1.ImagePruner, opts metav1.UpdateOptions) (result *v1.ImagePruner, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(imageprunersResource, imagePruner), &v1.ImagePruner{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ImagePruner), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeImagePruners) UpdateStatus(ctx context.Context, imagePruner *v1.ImagePruner, opts metav1.UpdateOptions) (*v1.ImagePruner, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(imageprunersResource, "status", imagePruner), &v1.ImagePruner{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ImagePruner), err
}

// Delete takes name of the imagePruner and deletes it. Returns an error if one occurs.
func (c *FakeImagePruners) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(imageprunersResource, name, opts), &v1.ImagePruner{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeImagePruners) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(imageprunersResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ImagePrunerList{})
	return err
}

// Patch applies the patch and returns the patched imagePruner.
func (c *FakeImagePruners) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ImagePruner, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(imageprunersResource, name, pt, data, subresources...), &v1.ImagePruner{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ImagePruner), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied imagePruner.
func (c *FakeImagePruners) Apply(ctx context.Context, imagePruner *imageregistryv1.ImagePrunerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ImagePruner, err error) {
	if imagePruner == nil {
		return nil, fmt.Errorf("imagePruner provided to Apply must not be nil")
	}
	data, err := json.Marshal(imagePruner)
	if err != nil {
		return nil, err
	}
	name := imagePruner.Name
	if name == nil {
		return nil, fmt.Errorf("imagePruner.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(imageprunersResource, *name, types.ApplyPatchType, data), &v1.ImagePruner{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ImagePruner), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeImagePruners) ApplyStatus(ctx context.Context, imagePruner *imageregistryv1.ImagePrunerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ImagePruner, err error) {
	if imagePruner == nil {
		return nil, fmt.Errorf("imagePruner provided to Apply must not be nil")
	}
	data, err := json.Marshal(imagePruner)
	if err != nil {
		return nil, err
	}
	name := imagePruner.Name
	if name == nil {
		return nil, fmt.Errorf("imagePruner.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(imageprunersResource, *name, types.ApplyPatchType, data, "status"), &v1.ImagePruner{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ImagePruner), err
}
