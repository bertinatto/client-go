// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	context "context"

	v1 "github.com/openshift/api/authorization/v1"
	authorizationv1 "github.com/openshift/client-go/authorization/clientset/versioned/typed/authorization/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gentype "k8s.io/client-go/gentype"
	testing "k8s.io/client-go/testing"
)

// fakeLocalResourceAccessReviews implements LocalResourceAccessReviewInterface
type fakeLocalResourceAccessReviews struct {
	*gentype.FakeClient[*v1.LocalResourceAccessReview]
	Fake *FakeAuthorizationV1
}

func newFakeLocalResourceAccessReviews(fake *FakeAuthorizationV1, namespace string) authorizationv1.LocalResourceAccessReviewInterface {
	return &fakeLocalResourceAccessReviews{
		gentype.NewFakeClient[*v1.LocalResourceAccessReview](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("localresourceaccessreviews"),
			v1.SchemeGroupVersion.WithKind("LocalResourceAccessReview"),
			func() *v1.LocalResourceAccessReview { return &v1.LocalResourceAccessReview{} },
		),
		fake,
	}
}

// Create takes the representation of a localResourceAccessReview and creates it.  Returns the server's representation of the resourceAccessReviewResponse, and an error, if there is any.
func (c *fakeLocalResourceAccessReviews) Create(ctx context.Context, localResourceAccessReview *v1.LocalResourceAccessReview, opts metav1.CreateOptions) (result *v1.ResourceAccessReviewResponse, err error) {
	emptyResult := &v1.ResourceAccessReviewResponse{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(c.Resource(), c.Namespace(), localResourceAccessReview, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ResourceAccessReviewResponse), err
}
