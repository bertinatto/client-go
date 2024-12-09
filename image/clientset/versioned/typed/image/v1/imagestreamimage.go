// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	imagev1 "github.com/openshift/api/image/v1"
	scheme "github.com/openshift/client-go/image/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gentype "k8s.io/client-go/gentype"
)

// ImageStreamImagesGetter has a method to return a ImageStreamImageInterface.
// A group's client should implement this interface.
type ImageStreamImagesGetter interface {
	ImageStreamImages(namespace string) ImageStreamImageInterface
}

// ImageStreamImageInterface has methods to work with ImageStreamImage resources.
type ImageStreamImageInterface interface {
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*imagev1.ImageStreamImage, error)
	ImageStreamImageExpansion
}

// imageStreamImages implements ImageStreamImageInterface
type imageStreamImages struct {
	*gentype.Client[*imagev1.ImageStreamImage]
}

// newImageStreamImages returns a ImageStreamImages
func newImageStreamImages(c *ImageV1Client, namespace string) *imageStreamImages {
	return &imageStreamImages{
		gentype.NewClient[*imagev1.ImageStreamImage](
			"imagestreamimages",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *imagev1.ImageStreamImage { return &imagev1.ImageStreamImage{} },
		),
	}
}
