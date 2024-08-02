/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1alpha1

import (
	"context"

	common "github.com/crossplane-contrib/provider-keycloak/config/common"
	apisresolver "github.com/crossplane-contrib/provider-keycloak/internal/apis"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *ProtocolMapper) ResolveReferences( // ResolveReferences of this ProtocolMapper.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("openidclient.keycloak.crossplane.io", "v1alpha1", "Client", "ClientList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClientID),
			Extract:      common.UUIDExtractor(),
			Reference:    mg.Spec.ForProvider.ClientIDRef,
			Selector:     mg.Spec.ForProvider.ClientIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClientID")
	}
	mg.Spec.ForProvider.ClientID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClientIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("openidclient.keycloak.crossplane.io", "v1alpha1", "ClientScope", "ClientScopeList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClientScopeID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ClientScopeIDRef,
			Selector:     mg.Spec.ForProvider.ClientScopeIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClientScopeID")
	}
	mg.Spec.ForProvider.ClientScopeID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClientScopeIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("realm.keycloak.crossplane.io", "v1alpha1", "Realm", "RealmList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RealmIDRef,
			Selector:     mg.Spec.ForProvider.RealmIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("openidclient.keycloak.crossplane.io", "v1alpha1", "Client", "ClientList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClientID),
			Extract:      common.UUIDExtractor(),
			Reference:    mg.Spec.InitProvider.ClientIDRef,
			Selector:     mg.Spec.InitProvider.ClientIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClientID")
	}
	mg.Spec.InitProvider.ClientID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClientIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("openidclient.keycloak.crossplane.io", "v1alpha1", "ClientScope", "ClientScopeList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClientScopeID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ClientScopeIDRef,
			Selector:     mg.Spec.InitProvider.ClientScopeIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClientScopeID")
	}
	mg.Spec.InitProvider.ClientScopeID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClientScopeIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("realm.keycloak.crossplane.io", "v1alpha1", "Realm", "RealmList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.RealmIDRef,
			Selector:     mg.Spec.InitProvider.RealmIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RoleMapper.
func (mg *RoleMapper) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("openidclient.keycloak.crossplane.io", "v1alpha1", "Client", "ClientList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClientID),
			Extract:      common.UUIDExtractor(),
			Reference:    mg.Spec.ForProvider.ClientIDRef,
			Selector:     mg.Spec.ForProvider.ClientIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClientID")
	}
	mg.Spec.ForProvider.ClientID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClientIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("openidclient.keycloak.crossplane.io", "v1alpha1", "ClientScope", "ClientScopeList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClientScopeID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ClientScopeIDRef,
			Selector:     mg.Spec.ForProvider.ClientScopeIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClientScopeID")
	}
	mg.Spec.ForProvider.ClientScopeID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClientScopeIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("realm.keycloak.crossplane.io", "v1alpha1", "Realm", "RealmList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RealmIDRef,
			Selector:     mg.Spec.ForProvider.RealmIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("role.keycloak.crossplane.io", "v1alpha1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RoleIDRef,
			Selector:     mg.Spec.ForProvider.RoleIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleID")
	}
	mg.Spec.ForProvider.RoleID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("openidclient.keycloak.crossplane.io", "v1alpha1", "Client", "ClientList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClientID),
			Extract:      common.UUIDExtractor(),
			Reference:    mg.Spec.InitProvider.ClientIDRef,
			Selector:     mg.Spec.InitProvider.ClientIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClientID")
	}
	mg.Spec.InitProvider.ClientID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClientIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("openidclient.keycloak.crossplane.io", "v1alpha1", "ClientScope", "ClientScopeList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClientScopeID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ClientScopeIDRef,
			Selector:     mg.Spec.InitProvider.ClientScopeIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClientScopeID")
	}
	mg.Spec.InitProvider.ClientScopeID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClientScopeIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("realm.keycloak.crossplane.io", "v1alpha1", "Realm", "RealmList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.RealmIDRef,
			Selector:     mg.Spec.InitProvider.RealmIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("role.keycloak.crossplane.io", "v1alpha1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RoleID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.RoleIDRef,
			Selector:     mg.Spec.InitProvider.RoleIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RoleID")
	}
	mg.Spec.InitProvider.RoleID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleIDRef = rsp.ResolvedReference

	return nil
}
