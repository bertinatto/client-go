// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/openshift/api/operator/v1"
	operatorv1 "github.com/openshift/client-go/operator/applyconfigurations/operator/v1"
	typedoperatorv1 "github.com/openshift/client-go/operator/clientset/versioned/typed/operator/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeKubeStorageVersionMigrators implements KubeStorageVersionMigratorInterface
type fakeKubeStorageVersionMigrators struct {
	*gentype.FakeClientWithListAndApply[*v1.KubeStorageVersionMigrator, *v1.KubeStorageVersionMigratorList, *operatorv1.KubeStorageVersionMigratorApplyConfiguration]
	Fake *FakeOperatorV1
}

func newFakeKubeStorageVersionMigrators(fake *FakeOperatorV1) typedoperatorv1.KubeStorageVersionMigratorInterface {
	return &fakeKubeStorageVersionMigrators{
		gentype.NewFakeClientWithListAndApply[*v1.KubeStorageVersionMigrator, *v1.KubeStorageVersionMigratorList, *operatorv1.KubeStorageVersionMigratorApplyConfiguration](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("kubestorageversionmigrators"),
			v1.SchemeGroupVersion.WithKind("KubeStorageVersionMigrator"),
			func() *v1.KubeStorageVersionMigrator { return &v1.KubeStorageVersionMigrator{} },
			func() *v1.KubeStorageVersionMigratorList { return &v1.KubeStorageVersionMigratorList{} },
			func(dst, src *v1.KubeStorageVersionMigratorList) { dst.ListMeta = src.ListMeta },
			func(list *v1.KubeStorageVersionMigratorList) []*v1.KubeStorageVersionMigrator {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1.KubeStorageVersionMigratorList, items []*v1.KubeStorageVersionMigrator) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
