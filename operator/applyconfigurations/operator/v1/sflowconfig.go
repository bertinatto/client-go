// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	operatorv1 "github.com/openshift/api/operator/v1"
)

// SFlowConfigApplyConfiguration represents a declarative configuration of the SFlowConfig type for use
// with apply.
type SFlowConfigApplyConfiguration struct {
	Collectors []operatorv1.IPPort `json:"collectors,omitempty"`
}

// SFlowConfigApplyConfiguration constructs a declarative configuration of the SFlowConfig type for use with
// apply.
func SFlowConfig() *SFlowConfigApplyConfiguration {
	return &SFlowConfigApplyConfiguration{}
}

// WithCollectors adds the given value to the Collectors field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Collectors field.
func (b *SFlowConfigApplyConfiguration) WithCollectors(values ...operatorv1.IPPort) *SFlowConfigApplyConfiguration {
	for i := range values {
		b.Collectors = append(b.Collectors, values[i])
	}
	return b
}
