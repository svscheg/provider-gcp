/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type IdsEndpointObservation struct {

	// Creation timestamp in RFC 3339 text format.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// URL of the endpoint's network address to which traffic is to be sent by Packet Mirroring.
	EndpointForwardingRule *string `json:"endpointForwardingRule,omitempty" tf:"endpoint_forwarding_rule,omitempty"`

	// Internal IP address of the endpoint's network entry point.
	EndpointIP *string `json:"endpointIp,omitempty" tf:"endpoint_ip,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/endpoints/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Last update timestamp in RFC 3339 text format.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type IdsEndpointParameters struct {

	// An optional description of the endpoint.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The location for the endpoint.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Name of the endpoint in the format projects/{project_id}/locations/{locationId}/endpoints/{endpointId}.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Name of the VPC network that is connected to the IDS endpoint. This can either contain the VPC network name itself (like "src-net") or the full URL to the network (like "projects/{project_id}/global/networks/src-net").
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Network
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Reference to a Network in compute to populate network.
	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// Selector for a Network in compute to populate network.
	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The minimum alert severity level that is reported by the endpoint.
	// Possible values are INFORMATIONAL, LOW, MEDIUM, HIGH, and CRITICAL.
	// +kubebuilder:validation:Required
	Severity *string `json:"severity" tf:"severity,omitempty"`

	// Configuration for threat IDs excluded from generating alerts. Limit: 99 IDs.
	// +kubebuilder:validation:Optional
	ThreatExceptions []*string `json:"threatExceptions,omitempty" tf:"threat_exceptions,omitempty"`
}

// IdsEndpointSpec defines the desired state of IdsEndpoint
type IdsEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IdsEndpointParameters `json:"forProvider"`
}

// IdsEndpointStatus defines the observed state of IdsEndpoint.
type IdsEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IdsEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IdsEndpoint is the Schema for the IdsEndpoints API. Cloud IDS is an intrusion detection service that provides threat detection for intrusions, malware, spyware, and command-and-control attacks on your network.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type IdsEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IdsEndpointSpec   `json:"spec"`
	Status            IdsEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IdsEndpointList contains a list of IdsEndpoints
type IdsEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IdsEndpoint `json:"items"`
}

// Repository type metadata.
var (
	IdsEndpoint_Kind             = "IdsEndpoint"
	IdsEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IdsEndpoint_Kind}.String()
	IdsEndpoint_KindAPIVersion   = IdsEndpoint_Kind + "." + CRDGroupVersion.String()
	IdsEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(IdsEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&IdsEndpoint{}, &IdsEndpointList{})
}
