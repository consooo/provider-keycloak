/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ClientRolePolicyInitParameters struct {
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Logic *string `json:"logic,omitempty" tf:"logic,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/openidclient/v1alpha1.Client
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-keycloak/config/common.UUIDExtractor()
	ResourceServerID *string `json:"resourceServerId,omitempty" tf:"resource_server_id,omitempty"`

	// Reference to a Client in openidclient to populate resourceServerId.
	// +kubebuilder:validation:Optional
	ResourceServerIDRef *v1.Reference `json:"resourceServerIdRef,omitempty" tf:"-"`

	// Selector for a Client in openidclient to populate resourceServerId.
	// +kubebuilder:validation:Optional
	ResourceServerIDSelector *v1.Selector `json:"resourceServerIdSelector,omitempty" tf:"-"`

	Role []RoleInitParameters `json:"role,omitempty" tf:"role,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ClientRolePolicyObservation struct {
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Logic *string `json:"logic,omitempty" tf:"logic,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	ResourceServerID *string `json:"resourceServerId,omitempty" tf:"resource_server_id,omitempty"`

	Role []RoleObservation `json:"role,omitempty" tf:"role,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ClientRolePolicyParameters struct {

	// +kubebuilder:validation:Optional
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Logic *string `json:"logic,omitempty" tf:"logic,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	// +kubebuilder:validation:Optional
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/openidclient/v1alpha1.Client
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-keycloak/config/common.UUIDExtractor()
	// +kubebuilder:validation:Optional
	ResourceServerID *string `json:"resourceServerId,omitempty" tf:"resource_server_id,omitempty"`

	// Reference to a Client in openidclient to populate resourceServerId.
	// +kubebuilder:validation:Optional
	ResourceServerIDRef *v1.Reference `json:"resourceServerIdRef,omitempty" tf:"-"`

	// Selector for a Client in openidclient to populate resourceServerId.
	// +kubebuilder:validation:Optional
	ResourceServerIDSelector *v1.Selector `json:"resourceServerIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Role []RoleParameters `json:"role,omitempty" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RoleInitParameters struct {

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/role/v1alpha1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-keycloak/config/common.UUIDExtractor()
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a Role in role to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a Role in role to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`

	Required *bool `json:"required,omitempty" tf:"required,omitempty"`
}

type RoleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Required *bool `json:"required,omitempty" tf:"required,omitempty"`
}

type RoleParameters struct {

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/role/v1alpha1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-keycloak/config/common.UUIDExtractor()
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a Role in role to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a Role in role to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Required *bool `json:"required" tf:"required,omitempty"`
}

// ClientRolePolicySpec defines the desired state of ClientRolePolicy
type ClientRolePolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClientRolePolicyParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ClientRolePolicyInitParameters `json:"initProvider,omitempty"`
}

// ClientRolePolicyStatus defines the observed state of ClientRolePolicy.
type ClientRolePolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClientRolePolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ClientRolePolicy is the Schema for the ClientRolePolicys API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,keycloak}
type ClientRolePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.logic) || (has(self.initProvider) && has(self.initProvider.logic))",message="spec.forProvider.logic is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || (has(self.initProvider) && has(self.initProvider.role))",message="spec.forProvider.role is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   ClientRolePolicySpec   `json:"spec"`
	Status ClientRolePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClientRolePolicyList contains a list of ClientRolePolicys
type ClientRolePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClientRolePolicy `json:"items"`
}

// Repository type metadata.
var (
	ClientRolePolicy_Kind             = "ClientRolePolicy"
	ClientRolePolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClientRolePolicy_Kind}.String()
	ClientRolePolicy_KindAPIVersion   = ClientRolePolicy_Kind + "." + CRDGroupVersion.String()
	ClientRolePolicy_GroupVersionKind = CRDGroupVersion.WithKind(ClientRolePolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ClientRolePolicy{}, &ClientRolePolicyList{})
}
