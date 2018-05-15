package v1beta1

import (
	configv1 "github.com/openshift/api/config/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConsoleConfiguration holds the necessary configuration options for serving the web console
type ConsoleConfiguration struct {
	metav1.TypeMeta `json:,inline"`

	// ServingInfo is the HTTP serving information for these assets
	ServingInfo configv1.HTTPServingInfo `json:"servingInfo"`

	// ClusterInfo holds information the web console needs to talk to the cluster such as master public URL
	// and metrics public URL
	ClusterInfo ClusterInfo `json:"clusterInfo"`

	// Customization
	Customization Customization `json:"customization"`
}

// ClusterInfo holds information the web console needs to talk to the cluster such as master public URL and
// metrics public URL
type ClusterInfo struct {
	// ConsolePublicURL is where you can find the web console server
	ConsolePublicURL string `json:"consolePublicURL"`

	// MasterPublicURL is how the web console can access the OpenShift v1 server
	MasterPublicURL string `json:"masterPublicURL"`

	// LogoutPublicURL is an optional, absolute URL to redirect web browsers to after logging out of the web
	// console. If not specified, the built-in logout page is shown.
	LogoutPublicURL string `json:"logoutPublicURL"`
}

type Customization struct {
	// TODO: I plan to update the values to 'origin', 'ocp', and 'online'
	// LogoIconName specifies the icon to use in the masthead. Can be 'os-origin', 'os-platform', or 'os-online'.
	LogoIconName string `json:"logoIconName"`

	// DocumentationURL is the base URL to the documentation.
	DocumentationURL string `json:"documentationURL"`
}
