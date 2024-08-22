// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// ActionApplyConfiguration represents a declarative configuration of the Action type for use
// with apply.
type ActionApplyConfiguration struct {
	Namespace        *string               `json:"namespace,omitempty"`
	Verb             *string               `json:"verb,omitempty"`
	Group            *string               `json:"resourceAPIGroup,omitempty"`
	Version          *string               `json:"resourceAPIVersion,omitempty"`
	Resource         *string               `json:"resource,omitempty"`
	ResourceName     *string               `json:"resourceName,omitempty"`
	Path             *string               `json:"path,omitempty"`
	IsNonResourceURL *bool                 `json:"isNonResourceURL,omitempty"`
	Content          *runtime.RawExtension `json:"content,omitempty"`
}

// ActionApplyConfiguration constructs a declarative configuration of the Action type for use with
// apply.
func Action() *ActionApplyConfiguration {
	return &ActionApplyConfiguration{}
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *ActionApplyConfiguration) WithNamespace(value string) *ActionApplyConfiguration {
	b.Namespace = &value
	return b
}

// WithVerb sets the Verb field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Verb field is set to the value of the last call.
func (b *ActionApplyConfiguration) WithVerb(value string) *ActionApplyConfiguration {
	b.Verb = &value
	return b
}

// WithGroup sets the Group field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Group field is set to the value of the last call.
func (b *ActionApplyConfiguration) WithGroup(value string) *ActionApplyConfiguration {
	b.Group = &value
	return b
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *ActionApplyConfiguration) WithVersion(value string) *ActionApplyConfiguration {
	b.Version = &value
	return b
}

// WithResource sets the Resource field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resource field is set to the value of the last call.
func (b *ActionApplyConfiguration) WithResource(value string) *ActionApplyConfiguration {
	b.Resource = &value
	return b
}

// WithResourceName sets the ResourceName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceName field is set to the value of the last call.
func (b *ActionApplyConfiguration) WithResourceName(value string) *ActionApplyConfiguration {
	b.ResourceName = &value
	return b
}

// WithPath sets the Path field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Path field is set to the value of the last call.
func (b *ActionApplyConfiguration) WithPath(value string) *ActionApplyConfiguration {
	b.Path = &value
	return b
}

// WithIsNonResourceURL sets the IsNonResourceURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IsNonResourceURL field is set to the value of the last call.
func (b *ActionApplyConfiguration) WithIsNonResourceURL(value bool) *ActionApplyConfiguration {
	b.IsNonResourceURL = &value
	return b
}

// WithContent sets the Content field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Content field is set to the value of the last call.
func (b *ActionApplyConfiguration) WithContent(value runtime.RawExtension) *ActionApplyConfiguration {
	b.Content = &value
	return b
}
