// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/openshift/api/console/v1"
	consolev1 "github.com/openshift/client-go/console/applyconfigurations/console/v1"
	typedconsolev1 "github.com/openshift/client-go/console/clientset/versioned/typed/console/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeConsoleExternalLogLinks implements ConsoleExternalLogLinkInterface
type fakeConsoleExternalLogLinks struct {
	*gentype.FakeClientWithListAndApply[*v1.ConsoleExternalLogLink, *v1.ConsoleExternalLogLinkList, *consolev1.ConsoleExternalLogLinkApplyConfiguration]
	Fake *FakeConsoleV1
}

func newFakeConsoleExternalLogLinks(fake *FakeConsoleV1) typedconsolev1.ConsoleExternalLogLinkInterface {
	return &fakeConsoleExternalLogLinks{
		gentype.NewFakeClientWithListAndApply[*v1.ConsoleExternalLogLink, *v1.ConsoleExternalLogLinkList, *consolev1.ConsoleExternalLogLinkApplyConfiguration](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("consoleexternalloglinks"),
			v1.SchemeGroupVersion.WithKind("ConsoleExternalLogLink"),
			func() *v1.ConsoleExternalLogLink { return &v1.ConsoleExternalLogLink{} },
			func() *v1.ConsoleExternalLogLinkList { return &v1.ConsoleExternalLogLinkList{} },
			func(dst, src *v1.ConsoleExternalLogLinkList) { dst.ListMeta = src.ListMeta },
			func(list *v1.ConsoleExternalLogLinkList) []*v1.ConsoleExternalLogLink {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1.ConsoleExternalLogLinkList, items []*v1.ConsoleExternalLogLink) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
