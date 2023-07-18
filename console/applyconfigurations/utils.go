// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfigurations

import (
	v1 "github.com/openshift/api/console/v1"
	v1alpha1 "github.com/openshift/api/console/v1alpha1"
	consolev1 "github.com/openshift/client-go/console/applyconfigurations/console/v1"
	consolev1alpha1 "github.com/openshift/client-go/console/applyconfigurations/console/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=console.openshift.io, Version=v1
	case v1.SchemeGroupVersion.WithKind("ApplicationMenuSpec"):
		return &consolev1.ApplicationMenuSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CLIDownloadLink"):
		return &consolev1.CLIDownloadLinkApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleCLIDownload"):
		return &consolev1.ConsoleCLIDownloadApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleCLIDownloadSpec"):
		return &consolev1.ConsoleCLIDownloadSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleExternalLogLink"):
		return &consolev1.ConsoleExternalLogLinkApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleExternalLogLinkSpec"):
		return &consolev1.ConsoleExternalLogLinkSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleLink"):
		return &consolev1.ConsoleLinkApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleLinkSpec"):
		return &consolev1.ConsoleLinkSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleNotification"):
		return &consolev1.ConsoleNotificationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleNotificationSpec"):
		return &consolev1.ConsoleNotificationSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsolePlugin"):
		return &consolev1.ConsolePluginApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsolePluginBackend"):
		return &consolev1.ConsolePluginBackendApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsolePluginI18n"):
		return &consolev1.ConsolePluginI18nApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsolePluginProxy"):
		return &consolev1.ConsolePluginProxyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsolePluginProxyEndpoint"):
		return &consolev1.ConsolePluginProxyEndpointApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsolePluginProxyServiceConfig"):
		return &consolev1.ConsolePluginProxyServiceConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsolePluginService"):
		return &consolev1.ConsolePluginServiceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsolePluginSpec"):
		return &consolev1.ConsolePluginSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleQuickStart"):
		return &consolev1.ConsoleQuickStartApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleQuickStartSpec"):
		return &consolev1.ConsoleQuickStartSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleQuickStartTask"):
		return &consolev1.ConsoleQuickStartTaskApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleQuickStartTaskReview"):
		return &consolev1.ConsoleQuickStartTaskReviewApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleQuickStartTaskSummary"):
		return &consolev1.ConsoleQuickStartTaskSummaryApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleSample"):
		return &consolev1.ConsoleSampleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleSampleContainerImportSource"):
		return &consolev1.ConsoleSampleContainerImportSourceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleSampleContainerImportSourceService"):
		return &consolev1.ConsoleSampleContainerImportSourceServiceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleSampleGitImportSource"):
		return &consolev1.ConsoleSampleGitImportSourceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleSampleGitImportSourceRepository"):
		return &consolev1.ConsoleSampleGitImportSourceRepositoryApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleSampleGitImportSourceService"):
		return &consolev1.ConsoleSampleGitImportSourceServiceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleSampleSource"):
		return &consolev1.ConsoleSampleSourceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleSampleSpec"):
		return &consolev1.ConsoleSampleSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleYAMLSample"):
		return &consolev1.ConsoleYAMLSampleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleYAMLSampleSpec"):
		return &consolev1.ConsoleYAMLSampleSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Link"):
		return &consolev1.LinkApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NamespaceDashboardSpec"):
		return &consolev1.NamespaceDashboardSpecApplyConfiguration{}

		// Group=console.openshift.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("ConsolePlugin"):
		return &consolev1alpha1.ConsolePluginApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ConsolePluginProxy"):
		return &consolev1alpha1.ConsolePluginProxyApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ConsolePluginProxyServiceConfig"):
		return &consolev1alpha1.ConsolePluginProxyServiceConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ConsolePluginService"):
		return &consolev1alpha1.ConsolePluginServiceApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ConsolePluginSpec"):
		return &consolev1alpha1.ConsolePluginSpecApplyConfiguration{}

	}
	return nil
}
