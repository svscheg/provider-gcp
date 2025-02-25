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

type EnvironmentObservation struct {

	// an identifier for the resource with format {{org_id}}/environments/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// NodeConfig for setting the min/max number of nodes associated with the environment.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	NodeConfig []NodeConfigObservation `json:"nodeConfig,omitempty" tf:"node_config,omitempty"`
}

type EnvironmentParameters struct {

	// Optional. API Proxy type supported by the environment. The type can be set when creating
	// the Environment and cannot be changed.
	// Possible values are API_PROXY_TYPE_UNSPECIFIED, PROGRAMMABLE, and CONFIGURABLE.
	// +kubebuilder:validation:Optional
	APIProxyType *string `json:"apiProxyType,omitempty" tf:"api_proxy_type,omitempty"`

	// Optional. Deployment type supported by the environment. The deployment type can be
	// set when creating the environment and cannot be changed. When you enable archive
	// deployment, you will be prevented from performing a subset of actions within the
	// environment, including:
	// Managing the deployment of API proxy or shared flow revisions;
	// Creating, updating, or deleting resource files;
	// Creating, updating, or deleting target servers.
	// Possible values are DEPLOYMENT_TYPE_UNSPECIFIED, PROXY, and ARCHIVE.
	// +kubebuilder:validation:Optional
	DeploymentType *string `json:"deploymentType,omitempty" tf:"deployment_type,omitempty"`

	// Description of the environment.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Display name of the environment.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// NodeConfig for setting the min/max number of nodes associated with the environment.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	NodeConfig []NodeConfigParameters `json:"nodeConfig,omitempty" tf:"node_config,omitempty"`

	// The Apigee Organization associated with the Apigee environment,
	// in the format organizations/{{org_name}}.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/apigee/v1beta1.Organization
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Reference to a Organization in apigee to populate orgId.
	// +kubebuilder:validation:Optional
	OrgIDRef *v1.Reference `json:"orgIdRef,omitempty" tf:"-"`

	// Selector for a Organization in apigee to populate orgId.
	// +kubebuilder:validation:Optional
	OrgIDSelector *v1.Selector `json:"orgIdSelector,omitempty" tf:"-"`
}

type NodeConfigObservation struct {

	// The current total number of gateway nodes that each environment currently has across
	// all instances.
	CurrentAggregateNodeCount *string `json:"currentAggregateNodeCount,omitempty" tf:"current_aggregate_node_count,omitempty"`
}

type NodeConfigParameters struct {

	// The maximum total number of gateway nodes that the is reserved for all instances that
	// has the specified environment. If not specified, the default is determined by the
	// recommended maximum number of nodes for that gateway.
	// +kubebuilder:validation:Optional
	MaxNodeCount *string `json:"maxNodeCount,omitempty" tf:"max_node_count,omitempty"`

	// The minimum total number of gateway nodes that the is reserved for all instances that
	// has the specified environment. If not specified, the default is determined by the
	// recommended minimum number of nodes for that gateway.
	// +kubebuilder:validation:Optional
	MinNodeCount *string `json:"minNodeCount,omitempty" tf:"min_node_count,omitempty"`
}

// EnvironmentSpec defines the desired state of Environment
type EnvironmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EnvironmentParameters `json:"forProvider"`
}

// EnvironmentStatus defines the observed state of Environment.
type EnvironmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EnvironmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Environment is the Schema for the Environments API. An
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Environment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EnvironmentSpec   `json:"spec"`
	Status            EnvironmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EnvironmentList contains a list of Environments
type EnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Environment `json:"items"`
}

// Repository type metadata.
var (
	Environment_Kind             = "Environment"
	Environment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Environment_Kind}.String()
	Environment_KindAPIVersion   = Environment_Kind + "." + CRDGroupVersion.String()
	Environment_GroupVersionKind = CRDGroupVersion.WithKind(Environment_Kind)
)

func init() {
	SchemeBuilder.Register(&Environment{}, &EnvironmentList{})
}
