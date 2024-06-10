/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha11 "github.com/crossplane-contrib/provider-keycloak/apis/authenticationflow/v1alpha1"
	v1alpha1 "github.com/crossplane-contrib/provider-keycloak/apis/openidclient/v1alpha1"
	v1alpha12 "github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1"
	common "github.com/crossplane-contrib/provider-keycloak/config/common"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this IdentityProvider.
func (mg *IdentityProvider) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClientID),
		Extract:      common.UUIDExtractor(),
		Reference:    mg.Spec.ForProvider.ClientIDRef,
		Selector:     mg.Spec.ForProvider.ClientIDSelector,
		To: reference.To{
			List:    &v1alpha1.ClientList{},
			Managed: &v1alpha1.Client{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClientID")
	}
	mg.Spec.ForProvider.ClientID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClientIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FirstBrokerLoginFlowAlias),
		Extract:      common.AuthenticationFlowAliasExtractor(),
		Reference:    mg.Spec.ForProvider.FirstBrokerLoginFlowAliasRef,
		Selector:     mg.Spec.ForProvider.FirstBrokerLoginFlowAliasSelector,
		To: reference.To{
			List:    &v1alpha11.FlowList{},
			Managed: &v1alpha11.Flow{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FirstBrokerLoginFlowAlias")
	}
	mg.Spec.ForProvider.FirstBrokerLoginFlowAlias = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FirstBrokerLoginFlowAliasRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Realm),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmRef,
		Selector:     mg.Spec.ForProvider.RealmSelector,
		To: reference.To{
			List:    &v1alpha12.RealmList{},
			Managed: &v1alpha12.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Realm")
	}
	mg.Spec.ForProvider.Realm = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClientID),
		Extract:      common.UUIDExtractor(),
		Reference:    mg.Spec.InitProvider.ClientIDRef,
		Selector:     mg.Spec.InitProvider.ClientIDSelector,
		To: reference.To{
			List:    &v1alpha1.ClientList{},
			Managed: &v1alpha1.Client{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClientID")
	}
	mg.Spec.InitProvider.ClientID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClientIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FirstBrokerLoginFlowAlias),
		Extract:      common.AuthenticationFlowAliasExtractor(),
		Reference:    mg.Spec.InitProvider.FirstBrokerLoginFlowAliasRef,
		Selector:     mg.Spec.InitProvider.FirstBrokerLoginFlowAliasSelector,
		To: reference.To{
			List:    &v1alpha11.FlowList{},
			Managed: &v1alpha11.Flow{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FirstBrokerLoginFlowAlias")
	}
	mg.Spec.InitProvider.FirstBrokerLoginFlowAlias = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FirstBrokerLoginFlowAliasRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Realm),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RealmRef,
		Selector:     mg.Spec.InitProvider.RealmSelector,
		To: reference.To{
			List:    &v1alpha12.RealmList{},
			Managed: &v1alpha12.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Realm")
	}
	mg.Spec.InitProvider.Realm = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmRef = rsp.ResolvedReference

	return nil
}
