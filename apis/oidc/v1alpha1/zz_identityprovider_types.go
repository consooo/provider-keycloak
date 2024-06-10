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

type IdentityProviderInitParameters struct {

	// When true, the IDP will accept forwarded authentication requests that contain the prompt=none query parameter. Defaults to false.
	// This is just used together with Identity Provider Authenticator or when kc_idp_hint points to this identity provider. In case that client sends a request with prompt=none and user is not yet authenticated, the error will not be directly returned to client, but the request with prompt=none will be forwarded to this identity provider.
	AcceptsPromptNoneForwardFromClient *bool `json:"acceptsPromptNoneForwardFromClient,omitempty" tf:"accepts_prompt_none_forward_from_client,omitempty"`

	// When true, new users will be able to read stored tokens. This will automatically assign the broker.read-token role. Defaults to false.
	// Enable/disable if new users can read any stored tokens. This assigns the broker.read-token role.
	AddReadTokenRoleOnCreate *bool `json:"addReadTokenRoleOnCreate,omitempty" tf:"add_read_token_role_on_create,omitempty"`

	// The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
	// The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// Enable/disable authenticate users by default.
	AuthenticateByDefault *bool `json:"authenticateByDefault,omitempty" tf:"authenticate_by_default,omitempty"`

	// The Authorization Url.
	// OIDC authorization URL.
	AuthorizationURL *string `json:"authorizationUrl,omitempty" tf:"authorization_url,omitempty"`

	// Does the external IDP support backchannel logout? Defaults to true.
	// Does the external IDP support backchannel logout?
	BackchannelSupported *bool `json:"backchannelSupported,omitempty" tf:"backchannel_supported,omitempty"`

	// The client or client identifier registered within the identity provider.
	// Client ID.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/openidclient/v1alpha1.Client
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-keycloak/config/common.UUIDExtractor()
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Reference to a Client in openidclient to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDRef *v1.Reference `json:"clientIdRef,omitempty" tf:"-"`

	// Selector for a Client in openidclient to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDSelector *v1.Selector `json:"clientIdSelector,omitempty" tf:"-"`

	// The client or client secret registered within the identity provider. This field is able to obtain its value from vault, use $${vault.ID} format.
	// Client Secret.
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to openid.
	// The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to 'openid'.
	DefaultScopes *string `json:"defaultScopes,omitempty" tf:"default_scopes,omitempty"`

	// When true, disables the usage of the user info service to obtain additional user information. Defaults to false.
	// Disable usage of User Info service to obtain additional user information?  Default is to use this OIDC service.
	DisableUserInfo *bool `json:"disableUserInfo,omitempty" tf:"disable_user_info,omitempty"`

	// Display name for the identity provider in the GUI.
	// Friendly name for Identity Providers.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// When true, users will be able to log in to this realm using this identity provider. Defaults to true.
	// Enable/disable this identity provider.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A map of key/value pairs to add extra configuration to this identity provider. Use this attribute at your own risk, as custom attributes may conflict with top-level configuration attributes in future provider updates.
	// +mapType=granular
	ExtraConfig map[string]*string `json:"extraConfig,omitempty" tf:"extra_config,omitempty"`

	// The authentication flow to use when users log in for the first time through this identity provider. Defaults to first broker login.
	// Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login' means that there is not yet existing Keycloak account linked with the authenticated identity provider account.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/authenticationflow/v1alpha1.Flow
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-keycloak/config/common.AuthenticationFlowAliasExtractor()
	FirstBrokerLoginFlowAlias *string `json:"firstBrokerLoginFlowAlias,omitempty" tf:"first_broker_login_flow_alias,omitempty"`

	// Reference to a Flow in authenticationflow to populate firstBrokerLoginFlowAlias.
	// +kubebuilder:validation:Optional
	FirstBrokerLoginFlowAliasRef *v1.Reference `json:"firstBrokerLoginFlowAliasRef,omitempty" tf:"-"`

	// Selector for a Flow in authenticationflow to populate firstBrokerLoginFlowAlias.
	// +kubebuilder:validation:Optional
	FirstBrokerLoginFlowAliasSelector *v1.Selector `json:"firstBrokerLoginFlowAliasSelector,omitempty" tf:"-"`

	// A number defining the order of this identity provider in the GUI.
	// GUI Order
	GuiOrder *string `json:"guiOrder,omitempty" tf:"gui_order,omitempty"`

	// When true, this provider will be hidden on the login page, and is only accessible when requested explicitly. Defaults to false.
	// Hide On Login Page.
	HideOnLoginPage *bool `json:"hideOnLoginPage,omitempty" tf:"hide_on_login_page,omitempty"`

	// The issuer identifier for the issuer of the response. If not provided, no validation will be performed.
	// The issuer identifier for the issuer of the response. If not provided, no validation will be performed.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// JSON Web Key Set URL.
	// JSON Web Key Set URL
	JwksURL *string `json:"jwksUrl,omitempty" tf:"jwks_url,omitempty"`

	// When true, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to false.
	// If true, users cannot log in through this provider.  They can only link to this provider.  This is useful if you don't want to allow login from the provider, but want to integrate with a provider
	LinkOnly *bool `json:"linkOnly,omitempty" tf:"link_only,omitempty"`

	// Pass login hint to identity provider.
	// Login Hint.
	LoginHint *string `json:"loginHint,omitempty" tf:"login_hint,omitempty"`

	// The Logout URL is the end session endpoint to use to logout user from external identity provider.
	// Logout URL
	LogoutURL *string `json:"logoutUrl,omitempty" tf:"logout_url,omitempty"`

	// The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
	// Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that authenticator implementations must assume that user is already set in ClientSession as identity provider already set it.
	PostBrokerLoginFlowAlias *string `json:"postBrokerLoginFlowAlias,omitempty" tf:"post_broker_login_flow_alias,omitempty"`

	// The ID of the identity provider to use. Defaults to oidc, which should be used unless you have extended Keycloak and provided your own implementation.
	// provider id, is always oidc, unless you have a custom implementation
	ProviderID *string `json:"providerId,omitempty" tf:"provider_id,omitempty"`

	// The name of the realm. This is unique across Keycloak.
	// Realm Name
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	Realm *string `json:"realm,omitempty" tf:"realm,omitempty"`

	// Reference to a Realm in realm to populate realm.
	// +kubebuilder:validation:Optional
	RealmRef *v1.Reference `json:"realmRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realm.
	// +kubebuilder:validation:Optional
	RealmSelector *v1.Selector `json:"realmSelector,omitempty" tf:"-"`

	// When true, tokens will be stored after authenticating users. Defaults to true.
	// Enable/disable if tokens must be stored after authenticating users.
	StoreToken *bool `json:"storeToken,omitempty" tf:"store_token,omitempty"`

	// The default sync mode to use for all mappers attached to this identity provider. Can be once of IMPORT, FORCE, or LEGACY.
	// Sync Mode
	SyncMode *string `json:"syncMode,omitempty" tf:"sync_mode,omitempty"`

	// The Token URL.
	// Token URL.
	TokenURL *string `json:"tokenUrl,omitempty" tf:"token_url,omitempty"`

	// When true, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to false.
	// If enabled then email provided by this provider is not verified even if verification is enabled for the realm.
	TrustEmail *bool `json:"trustEmail,omitempty" tf:"trust_email,omitempty"`

	// Pass current locale to identity provider. Defaults to false.
	// Pass current locale to identity provider
	UILocales *bool `json:"uiLocales,omitempty" tf:"ui_locales,omitempty"`

	// User Info URL.
	// User Info URL
	UserInfoURL *string `json:"userInfoUrl,omitempty" tf:"user_info_url,omitempty"`

	// Enable/disable signature validation of external IDP signatures. Defaults to false.
	// Enable/disable signature validation of external IDP signatures.
	ValidateSignature *bool `json:"validateSignature,omitempty" tf:"validate_signature,omitempty"`
}

type IdentityProviderObservation struct {

	// When true, the IDP will accept forwarded authentication requests that contain the prompt=none query parameter. Defaults to false.
	// This is just used together with Identity Provider Authenticator or when kc_idp_hint points to this identity provider. In case that client sends a request with prompt=none and user is not yet authenticated, the error will not be directly returned to client, but the request with prompt=none will be forwarded to this identity provider.
	AcceptsPromptNoneForwardFromClient *bool `json:"acceptsPromptNoneForwardFromClient,omitempty" tf:"accepts_prompt_none_forward_from_client,omitempty"`

	// When true, new users will be able to read stored tokens. This will automatically assign the broker.read-token role. Defaults to false.
	// Enable/disable if new users can read any stored tokens. This assigns the broker.read-token role.
	AddReadTokenRoleOnCreate *bool `json:"addReadTokenRoleOnCreate,omitempty" tf:"add_read_token_role_on_create,omitempty"`

	// The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
	// The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// Enable/disable authenticate users by default.
	AuthenticateByDefault *bool `json:"authenticateByDefault,omitempty" tf:"authenticate_by_default,omitempty"`

	// The Authorization Url.
	// OIDC authorization URL.
	AuthorizationURL *string `json:"authorizationUrl,omitempty" tf:"authorization_url,omitempty"`

	// Does the external IDP support backchannel logout? Defaults to true.
	// Does the external IDP support backchannel logout?
	BackchannelSupported *bool `json:"backchannelSupported,omitempty" tf:"backchannel_supported,omitempty"`

	// The client or client identifier registered within the identity provider.
	// Client ID.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to openid.
	// The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to 'openid'.
	DefaultScopes *string `json:"defaultScopes,omitempty" tf:"default_scopes,omitempty"`

	// When true, disables the usage of the user info service to obtain additional user information. Defaults to false.
	// Disable usage of User Info service to obtain additional user information?  Default is to use this OIDC service.
	DisableUserInfo *bool `json:"disableUserInfo,omitempty" tf:"disable_user_info,omitempty"`

	// Display name for the identity provider in the GUI.
	// Friendly name for Identity Providers.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// When true, users will be able to log in to this realm using this identity provider. Defaults to true.
	// Enable/disable this identity provider.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A map of key/value pairs to add extra configuration to this identity provider. Use this attribute at your own risk, as custom attributes may conflict with top-level configuration attributes in future provider updates.
	// +mapType=granular
	ExtraConfig map[string]*string `json:"extraConfig,omitempty" tf:"extra_config,omitempty"`

	// The authentication flow to use when users log in for the first time through this identity provider. Defaults to first broker login.
	// Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login' means that there is not yet existing Keycloak account linked with the authenticated identity provider account.
	FirstBrokerLoginFlowAlias *string `json:"firstBrokerLoginFlowAlias,omitempty" tf:"first_broker_login_flow_alias,omitempty"`

	// A number defining the order of this identity provider in the GUI.
	// GUI Order
	GuiOrder *string `json:"guiOrder,omitempty" tf:"gui_order,omitempty"`

	// When true, this provider will be hidden on the login page, and is only accessible when requested explicitly. Defaults to false.
	// Hide On Login Page.
	HideOnLoginPage *bool `json:"hideOnLoginPage,omitempty" tf:"hide_on_login_page,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Computed) The unique ID that Keycloak assigns to the identity provider upon creation.
	// Internal Identity Provider Id
	InternalID *string `json:"internalId,omitempty" tf:"internal_id,omitempty"`

	// The issuer identifier for the issuer of the response. If not provided, no validation will be performed.
	// The issuer identifier for the issuer of the response. If not provided, no validation will be performed.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// JSON Web Key Set URL.
	// JSON Web Key Set URL
	JwksURL *string `json:"jwksUrl,omitempty" tf:"jwks_url,omitempty"`

	// When true, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to false.
	// If true, users cannot log in through this provider.  They can only link to this provider.  This is useful if you don't want to allow login from the provider, but want to integrate with a provider
	LinkOnly *bool `json:"linkOnly,omitempty" tf:"link_only,omitempty"`

	// Pass login hint to identity provider.
	// Login Hint.
	LoginHint *string `json:"loginHint,omitempty" tf:"login_hint,omitempty"`

	// The Logout URL is the end session endpoint to use to logout user from external identity provider.
	// Logout URL
	LogoutURL *string `json:"logoutUrl,omitempty" tf:"logout_url,omitempty"`

	// The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
	// Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that authenticator implementations must assume that user is already set in ClientSession as identity provider already set it.
	PostBrokerLoginFlowAlias *string `json:"postBrokerLoginFlowAlias,omitempty" tf:"post_broker_login_flow_alias,omitempty"`

	// The ID of the identity provider to use. Defaults to oidc, which should be used unless you have extended Keycloak and provided your own implementation.
	// provider id, is always oidc, unless you have a custom implementation
	ProviderID *string `json:"providerId,omitempty" tf:"provider_id,omitempty"`

	// The name of the realm. This is unique across Keycloak.
	// Realm Name
	Realm *string `json:"realm,omitempty" tf:"realm,omitempty"`

	// When true, tokens will be stored after authenticating users. Defaults to true.
	// Enable/disable if tokens must be stored after authenticating users.
	StoreToken *bool `json:"storeToken,omitempty" tf:"store_token,omitempty"`

	// The default sync mode to use for all mappers attached to this identity provider. Can be once of IMPORT, FORCE, or LEGACY.
	// Sync Mode
	SyncMode *string `json:"syncMode,omitempty" tf:"sync_mode,omitempty"`

	// The Token URL.
	// Token URL.
	TokenURL *string `json:"tokenUrl,omitempty" tf:"token_url,omitempty"`

	// When true, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to false.
	// If enabled then email provided by this provider is not verified even if verification is enabled for the realm.
	TrustEmail *bool `json:"trustEmail,omitempty" tf:"trust_email,omitempty"`

	// Pass current locale to identity provider. Defaults to false.
	// Pass current locale to identity provider
	UILocales *bool `json:"uiLocales,omitempty" tf:"ui_locales,omitempty"`

	// User Info URL.
	// User Info URL
	UserInfoURL *string `json:"userInfoUrl,omitempty" tf:"user_info_url,omitempty"`

	// Enable/disable signature validation of external IDP signatures. Defaults to false.
	// Enable/disable signature validation of external IDP signatures.
	ValidateSignature *bool `json:"validateSignature,omitempty" tf:"validate_signature,omitempty"`
}

type IdentityProviderParameters struct {

	// When true, the IDP will accept forwarded authentication requests that contain the prompt=none query parameter. Defaults to false.
	// This is just used together with Identity Provider Authenticator or when kc_idp_hint points to this identity provider. In case that client sends a request with prompt=none and user is not yet authenticated, the error will not be directly returned to client, but the request with prompt=none will be forwarded to this identity provider.
	// +kubebuilder:validation:Optional
	AcceptsPromptNoneForwardFromClient *bool `json:"acceptsPromptNoneForwardFromClient,omitempty" tf:"accepts_prompt_none_forward_from_client,omitempty"`

	// When true, new users will be able to read stored tokens. This will automatically assign the broker.read-token role. Defaults to false.
	// Enable/disable if new users can read any stored tokens. This assigns the broker.read-token role.
	// +kubebuilder:validation:Optional
	AddReadTokenRoleOnCreate *bool `json:"addReadTokenRoleOnCreate,omitempty" tf:"add_read_token_role_on_create,omitempty"`

	// The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
	// The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
	// +kubebuilder:validation:Optional
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// Enable/disable authenticate users by default.
	// +kubebuilder:validation:Optional
	AuthenticateByDefault *bool `json:"authenticateByDefault,omitempty" tf:"authenticate_by_default,omitempty"`

	// The Authorization Url.
	// OIDC authorization URL.
	// +kubebuilder:validation:Optional
	AuthorizationURL *string `json:"authorizationUrl,omitempty" tf:"authorization_url,omitempty"`

	// Does the external IDP support backchannel logout? Defaults to true.
	// Does the external IDP support backchannel logout?
	// +kubebuilder:validation:Optional
	BackchannelSupported *bool `json:"backchannelSupported,omitempty" tf:"backchannel_supported,omitempty"`

	// The client or client identifier registered within the identity provider.
	// Client ID.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/openidclient/v1alpha1.Client
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-keycloak/config/common.UUIDExtractor()
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Reference to a Client in openidclient to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDRef *v1.Reference `json:"clientIdRef,omitempty" tf:"-"`

	// Selector for a Client in openidclient to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDSelector *v1.Selector `json:"clientIdSelector,omitempty" tf:"-"`

	// The client or client secret registered within the identity provider. This field is able to obtain its value from vault, use $${vault.ID} format.
	// Client Secret.
	// +kubebuilder:validation:Optional
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to openid.
	// The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to 'openid'.
	// +kubebuilder:validation:Optional
	DefaultScopes *string `json:"defaultScopes,omitempty" tf:"default_scopes,omitempty"`

	// When true, disables the usage of the user info service to obtain additional user information. Defaults to false.
	// Disable usage of User Info service to obtain additional user information?  Default is to use this OIDC service.
	// +kubebuilder:validation:Optional
	DisableUserInfo *bool `json:"disableUserInfo,omitempty" tf:"disable_user_info,omitempty"`

	// Display name for the identity provider in the GUI.
	// Friendly name for Identity Providers.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// When true, users will be able to log in to this realm using this identity provider. Defaults to true.
	// Enable/disable this identity provider.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A map of key/value pairs to add extra configuration to this identity provider. Use this attribute at your own risk, as custom attributes may conflict with top-level configuration attributes in future provider updates.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	ExtraConfig map[string]*string `json:"extraConfig,omitempty" tf:"extra_config,omitempty"`

	// The authentication flow to use when users log in for the first time through this identity provider. Defaults to first broker login.
	// Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login' means that there is not yet existing Keycloak account linked with the authenticated identity provider account.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/authenticationflow/v1alpha1.Flow
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-keycloak/config/common.AuthenticationFlowAliasExtractor()
	// +kubebuilder:validation:Optional
	FirstBrokerLoginFlowAlias *string `json:"firstBrokerLoginFlowAlias,omitempty" tf:"first_broker_login_flow_alias,omitempty"`

	// Reference to a Flow in authenticationflow to populate firstBrokerLoginFlowAlias.
	// +kubebuilder:validation:Optional
	FirstBrokerLoginFlowAliasRef *v1.Reference `json:"firstBrokerLoginFlowAliasRef,omitempty" tf:"-"`

	// Selector for a Flow in authenticationflow to populate firstBrokerLoginFlowAlias.
	// +kubebuilder:validation:Optional
	FirstBrokerLoginFlowAliasSelector *v1.Selector `json:"firstBrokerLoginFlowAliasSelector,omitempty" tf:"-"`

	// A number defining the order of this identity provider in the GUI.
	// GUI Order
	// +kubebuilder:validation:Optional
	GuiOrder *string `json:"guiOrder,omitempty" tf:"gui_order,omitempty"`

	// When true, this provider will be hidden on the login page, and is only accessible when requested explicitly. Defaults to false.
	// Hide On Login Page.
	// +kubebuilder:validation:Optional
	HideOnLoginPage *bool `json:"hideOnLoginPage,omitempty" tf:"hide_on_login_page,omitempty"`

	// The issuer identifier for the issuer of the response. If not provided, no validation will be performed.
	// The issuer identifier for the issuer of the response. If not provided, no validation will be performed.
	// +kubebuilder:validation:Optional
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// JSON Web Key Set URL.
	// JSON Web Key Set URL
	// +kubebuilder:validation:Optional
	JwksURL *string `json:"jwksUrl,omitempty" tf:"jwks_url,omitempty"`

	// When true, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to false.
	// If true, users cannot log in through this provider.  They can only link to this provider.  This is useful if you don't want to allow login from the provider, but want to integrate with a provider
	// +kubebuilder:validation:Optional
	LinkOnly *bool `json:"linkOnly,omitempty" tf:"link_only,omitempty"`

	// Pass login hint to identity provider.
	// Login Hint.
	// +kubebuilder:validation:Optional
	LoginHint *string `json:"loginHint,omitempty" tf:"login_hint,omitempty"`

	// The Logout URL is the end session endpoint to use to logout user from external identity provider.
	// Logout URL
	// +kubebuilder:validation:Optional
	LogoutURL *string `json:"logoutUrl,omitempty" tf:"logout_url,omitempty"`

	// The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
	// Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that authenticator implementations must assume that user is already set in ClientSession as identity provider already set it.
	// +kubebuilder:validation:Optional
	PostBrokerLoginFlowAlias *string `json:"postBrokerLoginFlowAlias,omitempty" tf:"post_broker_login_flow_alias,omitempty"`

	// The ID of the identity provider to use. Defaults to oidc, which should be used unless you have extended Keycloak and provided your own implementation.
	// provider id, is always oidc, unless you have a custom implementation
	// +kubebuilder:validation:Optional
	ProviderID *string `json:"providerId,omitempty" tf:"provider_id,omitempty"`

	// The name of the realm. This is unique across Keycloak.
	// Realm Name
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	// +kubebuilder:validation:Optional
	Realm *string `json:"realm,omitempty" tf:"realm,omitempty"`

	// Reference to a Realm in realm to populate realm.
	// +kubebuilder:validation:Optional
	RealmRef *v1.Reference `json:"realmRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realm.
	// +kubebuilder:validation:Optional
	RealmSelector *v1.Selector `json:"realmSelector,omitempty" tf:"-"`

	// When true, tokens will be stored after authenticating users. Defaults to true.
	// Enable/disable if tokens must be stored after authenticating users.
	// +kubebuilder:validation:Optional
	StoreToken *bool `json:"storeToken,omitempty" tf:"store_token,omitempty"`

	// The default sync mode to use for all mappers attached to this identity provider. Can be once of IMPORT, FORCE, or LEGACY.
	// Sync Mode
	// +kubebuilder:validation:Optional
	SyncMode *string `json:"syncMode,omitempty" tf:"sync_mode,omitempty"`

	// The Token URL.
	// Token URL.
	// +kubebuilder:validation:Optional
	TokenURL *string `json:"tokenUrl,omitempty" tf:"token_url,omitempty"`

	// When true, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to false.
	// If enabled then email provided by this provider is not verified even if verification is enabled for the realm.
	// +kubebuilder:validation:Optional
	TrustEmail *bool `json:"trustEmail,omitempty" tf:"trust_email,omitempty"`

	// Pass current locale to identity provider. Defaults to false.
	// Pass current locale to identity provider
	// +kubebuilder:validation:Optional
	UILocales *bool `json:"uiLocales,omitempty" tf:"ui_locales,omitempty"`

	// User Info URL.
	// User Info URL
	// +kubebuilder:validation:Optional
	UserInfoURL *string `json:"userInfoUrl,omitempty" tf:"user_info_url,omitempty"`

	// Enable/disable signature validation of external IDP signatures. Defaults to false.
	// Enable/disable signature validation of external IDP signatures.
	// +kubebuilder:validation:Optional
	ValidateSignature *bool `json:"validateSignature,omitempty" tf:"validate_signature,omitempty"`
}

// IdentityProviderSpec defines the desired state of IdentityProvider
type IdentityProviderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IdentityProviderParameters `json:"forProvider"`
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
	InitProvider IdentityProviderInitParameters `json:"initProvider,omitempty"`
}

// IdentityProviderStatus defines the observed state of IdentityProvider.
type IdentityProviderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IdentityProviderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// IdentityProvider is the Schema for the IdentityProviders API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,keycloak}
type IdentityProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.alias) || (has(self.initProvider) && has(self.initProvider.alias))",message="spec.forProvider.alias is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.authorizationUrl) || (has(self.initProvider) && has(self.initProvider.authorizationUrl))",message="spec.forProvider.authorizationUrl is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientSecretSecretRef)",message="spec.forProvider.clientSecretSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.tokenUrl) || (has(self.initProvider) && has(self.initProvider.tokenUrl))",message="spec.forProvider.tokenUrl is a required parameter"
	Spec   IdentityProviderSpec   `json:"spec"`
	Status IdentityProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityProviderList contains a list of IdentityProviders
type IdentityProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IdentityProvider `json:"items"`
}

// Repository type metadata.
var (
	IdentityProvider_Kind             = "IdentityProvider"
	IdentityProvider_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IdentityProvider_Kind}.String()
	IdentityProvider_KindAPIVersion   = IdentityProvider_Kind + "." + CRDGroupVersion.String()
	IdentityProvider_GroupVersionKind = CRDGroupVersion.WithKind(IdentityProvider_Kind)
)

func init() {
	SchemeBuilder.Register(&IdentityProvider{}, &IdentityProviderList{})
}
