// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/openshift/api/helm/v1beta1"
	helmv1beta1 "github.com/openshift/client-go/helm/applyconfigurations/helm/v1beta1"
	typedhelmv1beta1 "github.com/openshift/client-go/helm/clientset/versioned/typed/helm/v1beta1"
	gentype "k8s.io/client-go/gentype"
)

// fakeHelmChartRepositories implements HelmChartRepositoryInterface
type fakeHelmChartRepositories struct {
	*gentype.FakeClientWithListAndApply[*v1beta1.HelmChartRepository, *v1beta1.HelmChartRepositoryList, *helmv1beta1.HelmChartRepositoryApplyConfiguration]
	Fake *FakeHelmV1beta1
}

func newFakeHelmChartRepositories(fake *FakeHelmV1beta1) typedhelmv1beta1.HelmChartRepositoryInterface {
	return &fakeHelmChartRepositories{
		gentype.NewFakeClientWithListAndApply[*v1beta1.HelmChartRepository, *v1beta1.HelmChartRepositoryList, *helmv1beta1.HelmChartRepositoryApplyConfiguration](
			fake.Fake,
			"",
			v1beta1.SchemeGroupVersion.WithResource("helmchartrepositories"),
			v1beta1.SchemeGroupVersion.WithKind("HelmChartRepository"),
			func() *v1beta1.HelmChartRepository { return &v1beta1.HelmChartRepository{} },
			func() *v1beta1.HelmChartRepositoryList { return &v1beta1.HelmChartRepositoryList{} },
			func(dst, src *v1beta1.HelmChartRepositoryList) { dst.ListMeta = src.ListMeta },
			func(list *v1beta1.HelmChartRepositoryList) []*v1beta1.HelmChartRepository {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1beta1.HelmChartRepositoryList, items []*v1beta1.HelmChartRepository) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
