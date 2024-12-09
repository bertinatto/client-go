// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	buildv1 "github.com/openshift/api/build/v1"
	applyconfigurationsbuildv1 "github.com/openshift/client-go/build/applyconfigurations/build/v1"
	scheme "github.com/openshift/client-go/build/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// BuildsGetter has a method to return a BuildInterface.
// A group's client should implement this interface.
type BuildsGetter interface {
	Builds(namespace string) BuildInterface
}

// BuildInterface has methods to work with Build resources.
type BuildInterface interface {
	Create(ctx context.Context, build *buildv1.Build, opts metav1.CreateOptions) (*buildv1.Build, error)
	Update(ctx context.Context, build *buildv1.Build, opts metav1.UpdateOptions) (*buildv1.Build, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, build *buildv1.Build, opts metav1.UpdateOptions) (*buildv1.Build, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*buildv1.Build, error)
	List(ctx context.Context, opts metav1.ListOptions) (*buildv1.BuildList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *buildv1.Build, err error)
	Apply(ctx context.Context, build *applyconfigurationsbuildv1.BuildApplyConfiguration, opts metav1.ApplyOptions) (result *buildv1.Build, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, build *applyconfigurationsbuildv1.BuildApplyConfiguration, opts metav1.ApplyOptions) (result *buildv1.Build, err error)
	UpdateDetails(ctx context.Context, buildName string, build *buildv1.Build, opts metav1.UpdateOptions) (*buildv1.Build, error)
	Clone(ctx context.Context, buildName string, buildRequest *buildv1.BuildRequest, opts metav1.CreateOptions) (*buildv1.Build, error)

	BuildExpansion
}

// builds implements BuildInterface
type builds struct {
	*gentype.ClientWithListAndApply[*buildv1.Build, *buildv1.BuildList, *applyconfigurationsbuildv1.BuildApplyConfiguration]
}

// newBuilds returns a Builds
func newBuilds(c *BuildV1Client, namespace string) *builds {
	return &builds{
		gentype.NewClientWithListAndApply[*buildv1.Build, *buildv1.BuildList, *applyconfigurationsbuildv1.BuildApplyConfiguration](
			"builds",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *buildv1.Build { return &buildv1.Build{} },
			func() *buildv1.BuildList { return &buildv1.BuildList{} },
		),
	}
}

// UpdateDetails takes the top resource name and the representation of a build and updates it. Returns the server's representation of the build, and an error, if there is any.
func (c *builds) UpdateDetails(ctx context.Context, buildName string, build *buildv1.Build, opts metav1.UpdateOptions) (result *buildv1.Build, err error) {
	result = &buildv1.Build{}
	err = c.GetClient().Put().
		Namespace(c.GetNamespace()).
		Resource("builds").
		Name(buildName).
		SubResource("details").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(build).
		Do(ctx).
		Into(result)
	return
}

// Clone takes the representation of a buildRequest and creates it.  Returns the server's representation of the build, and an error, if there is any.
func (c *builds) Clone(ctx context.Context, buildName string, buildRequest *buildv1.BuildRequest, opts metav1.CreateOptions) (result *buildv1.Build, err error) {
	result = &buildv1.Build{}
	err = c.GetClient().Post().
		Namespace(c.GetNamespace()).
		Resource("builds").
		Name(buildName).
		SubResource("clone").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(buildRequest).
		Do(ctx).
		Into(result)
	return
}
