// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v3beta2

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Profile is a specification for a Profile resource
type Profile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProfileSpec   `json:"spec"`
	Status ProfileStatus `json:"status"`
}

// ProfileSpec is the spec for a Profile resource
type ProfileSpec struct {
	Proxy    ProxySpec `json:"proxy"`
	Bindings []Binding `json:"bindings"`
}

type ProxySpec struct {
	// Image is the proxy image
	Image string `json:"image,omitempty"`

	// ImagePullPolicy is the pull policy to apply
	ImagePullPolicy corev1.PullPolicy `json:"imagePullPolicy,omitempty"`

	// SecurityContext is a pod security context
	SecurityContext *corev1.SecurityContext `json:"securityContext,omitempty"`

	// Logging is the proxy logging configuration
	Logging LoggingConfig `json:"logging,omitempty"`
}

type Binding struct {
	Store    corev1.ObjectReference `json:"store"`
	Priority *uint32                `json:"priority"`
	Selector map[string]string      `json:"selector"`
	Services []ServiceBinding       `json:"services"`
}

type ServiceBinding struct {
	Name   string               `json:"name"`
	Config runtime.RawExtension `json:"config"`
}

// LoggingConfig logging configuration
type LoggingConfig struct {
	Loggers map[string]LoggerConfig `json:"loggers" yaml:"loggers"`
	Sinks   map[string]SinkConfig   `json:"sinks" yaml:"sinks"`
}

// LoggerConfig is the configuration for a logger
type LoggerConfig struct {
	Level  *string                 `json:"level,omitempty" yaml:"level,omitempty"`
	Output map[string]OutputConfig `json:"output" yaml:"output"`
}

// OutputConfig is the configuration for a sink output
type OutputConfig struct {
	Sink  *string `json:"sink,omitempty" yaml:"sink,omitempty"`
	Level *string `json:"level,omitempty" yaml:"level,omitempty"`
}

// SinkConfig is the configuration for a sink
type SinkConfig struct {
	Encoding *string           `json:"encoding,omitempty" yaml:"encoding,omitempty"`
	Stdout   *StdoutSinkConfig `json:"stdout" yaml:"stdout,omitempty"`
	Stderr   *StderrSinkConfig `json:"stderr" yaml:"stderr,omitempty"`
	File     *FileSinkConfig   `json:"file" yaml:"file,omitempty"`
}

// StdoutSinkConfig is the configuration for an stdout sink
type StdoutSinkConfig struct {
}

// StderrSinkConfig is the configuration for an stderr sink
type StderrSinkConfig struct {
}

// FileSinkConfig is the configuration for a file sink
type FileSinkConfig struct {
	Path string `json:"path" yaml:"path"`
}

type ProfileStatus struct {
	PodStatuses []PodStatus `json:"podStatuses,omitempty"`
}

type PodStatus struct {
	Name  string      `json:"name"`
	Proxy ProxyStatus `json:"proxy"`
}

type ProxyStatus struct {
	Routes []RouteStatus `json:"routes"`
}

type RouteState string

const (
	RoutePending       RouteState = "Pending"
	RouteConnecting    RouteState = "Connecting"
	RouteConnected     RouteState = "Connected"
	RouteConfiguring   RouteState = "Configuring"
	RouteDisconnecting RouteState = "Disconnecting"
	RouteDisconnected  RouteState = "Disconnected"
)

type RouteStatus struct {
	Store   corev1.ObjectReference `json:"store"`
	State   RouteState             `json:"state"`
	Version string                 `json:"version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ProfileList is a list of Profile resources
type ProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Profile `json:"items"`
}