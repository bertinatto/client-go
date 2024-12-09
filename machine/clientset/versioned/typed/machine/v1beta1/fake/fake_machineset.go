// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/openshift/api/machine/v1beta1"
	machinev1beta1 "github.com/openshift/client-go/machine/applyconfigurations/machine/v1beta1"
	typedmachinev1beta1 "github.com/openshift/client-go/machine/clientset/versioned/typed/machine/v1beta1"
	gentype "k8s.io/client-go/gentype"
)

// fakeMachineSets implements MachineSetInterface
type fakeMachineSets struct {
	*gentype.FakeClientWithListAndApply[*v1beta1.MachineSet, *v1beta1.MachineSetList, *machinev1beta1.MachineSetApplyConfiguration]
	Fake *FakeMachineV1beta1
}

func newFakeMachineSets(fake *FakeMachineV1beta1, namespace string) typedmachinev1beta1.MachineSetInterface {
	return &fakeMachineSets{
		gentype.NewFakeClientWithListAndApply[*v1beta1.MachineSet, *v1beta1.MachineSetList, *machinev1beta1.MachineSetApplyConfiguration](
			fake.Fake,
			namespace,
			v1beta1.SchemeGroupVersion.WithResource("machinesets"),
			v1beta1.SchemeGroupVersion.WithKind("MachineSet"),
			func() *v1beta1.MachineSet { return &v1beta1.MachineSet{} },
			func() *v1beta1.MachineSetList { return &v1beta1.MachineSetList{} },
			func(dst, src *v1beta1.MachineSetList) { dst.ListMeta = src.ListMeta },
			func(list *v1beta1.MachineSetList) []*v1beta1.MachineSet { return gentype.ToPointerSlice(list.Items) },
			func(list *v1beta1.MachineSetList, items []*v1beta1.MachineSet) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
