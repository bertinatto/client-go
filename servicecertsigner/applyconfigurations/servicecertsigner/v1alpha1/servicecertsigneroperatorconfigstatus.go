// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/openshift/client-go/operator/applyconfigurations/operator/v1"
)

// ServiceCertSignerOperatorConfigStatusApplyConfiguration represents a declarative configuration of the ServiceCertSignerOperatorConfigStatus type for use
// with apply.
type ServiceCertSignerOperatorConfigStatusApplyConfiguration struct {
	v1.OperatorStatusApplyConfiguration `json:",inline"`
}

// ServiceCertSignerOperatorConfigStatusApplyConfiguration constructs a declarative configuration of the ServiceCertSignerOperatorConfigStatus type for use with
// apply.
func ServiceCertSignerOperatorConfigStatus() *ServiceCertSignerOperatorConfigStatusApplyConfiguration {
	return &ServiceCertSignerOperatorConfigStatusApplyConfiguration{}
}

// WithObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedGeneration field is set to the value of the last call.
func (b *ServiceCertSignerOperatorConfigStatusApplyConfiguration) WithObservedGeneration(value int64) *ServiceCertSignerOperatorConfigStatusApplyConfiguration {
	b.OperatorStatusApplyConfiguration.ObservedGeneration = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *ServiceCertSignerOperatorConfigStatusApplyConfiguration) WithConditions(values ...*v1.OperatorConditionApplyConfiguration) *ServiceCertSignerOperatorConfigStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.OperatorStatusApplyConfiguration.Conditions = append(b.OperatorStatusApplyConfiguration.Conditions, *values[i])
	}
	return b
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *ServiceCertSignerOperatorConfigStatusApplyConfiguration) WithVersion(value string) *ServiceCertSignerOperatorConfigStatusApplyConfiguration {
	b.OperatorStatusApplyConfiguration.Version = &value
	return b
}

// WithReadyReplicas sets the ReadyReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadyReplicas field is set to the value of the last call.
func (b *ServiceCertSignerOperatorConfigStatusApplyConfiguration) WithReadyReplicas(value int32) *ServiceCertSignerOperatorConfigStatusApplyConfiguration {
	b.OperatorStatusApplyConfiguration.ReadyReplicas = &value
	return b
}

// WithLatestAvailableRevision sets the LatestAvailableRevision field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LatestAvailableRevision field is set to the value of the last call.
func (b *ServiceCertSignerOperatorConfigStatusApplyConfiguration) WithLatestAvailableRevision(value int32) *ServiceCertSignerOperatorConfigStatusApplyConfiguration {
	b.OperatorStatusApplyConfiguration.LatestAvailableRevision = &value
	return b
}

// WithGenerations adds the given value to the Generations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Generations field.
func (b *ServiceCertSignerOperatorConfigStatusApplyConfiguration) WithGenerations(values ...*v1.GenerationStatusApplyConfiguration) *ServiceCertSignerOperatorConfigStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithGenerations")
		}
		b.OperatorStatusApplyConfiguration.Generations = append(b.OperatorStatusApplyConfiguration.Generations, *values[i])
	}
	return b
}
