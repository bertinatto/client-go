// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/openshift/api/oauth/v1"
	oauthv1 "github.com/openshift/client-go/oauth/applyconfigurations/oauth/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOAuthAccessTokens implements OAuthAccessTokenInterface
type FakeOAuthAccessTokens struct {
	Fake *FakeOauthV1
}

var oauthaccesstokensResource = v1.SchemeGroupVersion.WithResource("oauthaccesstokens")

var oauthaccesstokensKind = v1.SchemeGroupVersion.WithKind("OAuthAccessToken")

// Get takes name of the oAuthAccessToken, and returns the corresponding oAuthAccessToken object, and an error if there is any.
func (c *FakeOAuthAccessTokens) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.OAuthAccessToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(oauthaccesstokensResource, name), &v1.OAuthAccessToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.OAuthAccessToken), err
}

// List takes label and field selectors, and returns the list of OAuthAccessTokens that match those selectors.
func (c *FakeOAuthAccessTokens) List(ctx context.Context, opts metav1.ListOptions) (result *v1.OAuthAccessTokenList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(oauthaccesstokensResource, oauthaccesstokensKind, opts), &v1.OAuthAccessTokenList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.OAuthAccessTokenList{ListMeta: obj.(*v1.OAuthAccessTokenList).ListMeta}
	for _, item := range obj.(*v1.OAuthAccessTokenList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested oAuthAccessTokens.
func (c *FakeOAuthAccessTokens) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(oauthaccesstokensResource, opts))
}

// Create takes the representation of a oAuthAccessToken and creates it.  Returns the server's representation of the oAuthAccessToken, and an error, if there is any.
func (c *FakeOAuthAccessTokens) Create(ctx context.Context, oAuthAccessToken *v1.OAuthAccessToken, opts metav1.CreateOptions) (result *v1.OAuthAccessToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(oauthaccesstokensResource, oAuthAccessToken), &v1.OAuthAccessToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.OAuthAccessToken), err
}

// Update takes the representation of a oAuthAccessToken and updates it. Returns the server's representation of the oAuthAccessToken, and an error, if there is any.
func (c *FakeOAuthAccessTokens) Update(ctx context.Context, oAuthAccessToken *v1.OAuthAccessToken, opts metav1.UpdateOptions) (result *v1.OAuthAccessToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(oauthaccesstokensResource, oAuthAccessToken), &v1.OAuthAccessToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.OAuthAccessToken), err
}

// Delete takes name of the oAuthAccessToken and deletes it. Returns an error if one occurs.
func (c *FakeOAuthAccessTokens) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(oauthaccesstokensResource, name, opts), &v1.OAuthAccessToken{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOAuthAccessTokens) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(oauthaccesstokensResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.OAuthAccessTokenList{})
	return err
}

// Patch applies the patch and returns the patched oAuthAccessToken.
func (c *FakeOAuthAccessTokens) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.OAuthAccessToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(oauthaccesstokensResource, name, pt, data, subresources...), &v1.OAuthAccessToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.OAuthAccessToken), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied oAuthAccessToken.
func (c *FakeOAuthAccessTokens) Apply(ctx context.Context, oAuthAccessToken *oauthv1.OAuthAccessTokenApplyConfiguration, opts metav1.ApplyOptions) (result *v1.OAuthAccessToken, err error) {
	if oAuthAccessToken == nil {
		return nil, fmt.Errorf("oAuthAccessToken provided to Apply must not be nil")
	}
	data, err := json.Marshal(oAuthAccessToken)
	if err != nil {
		return nil, err
	}
	name := oAuthAccessToken.Name
	if name == nil {
		return nil, fmt.Errorf("oAuthAccessToken.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(oauthaccesstokensResource, *name, types.ApplyPatchType, data), &v1.OAuthAccessToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.OAuthAccessToken), err
}
