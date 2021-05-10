package auth

import configv1alpha1 "github.com/liqotech/liqo/apis/config/v1alpha1"

// ConfigProvider is a provider for the Authentication Configuration.
type ConfigProvider interface {
	// GetConfig retrieves the AuthConfiguration, such as the peering permission and the token authentication settings.
	GetConfig() *configv1alpha1.AuthConfig
	// GetAPIServerConfig retrieves the ApiServerConfiguration (i.e. Address, Port and TrustedCA).
	GetAPIServerConfig() *configv1alpha1.APIServerConfig
}
