// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package types

import (
	"fmt"

	"github.com/hashicorp/consul/internal/catalog"
	"github.com/hashicorp/consul/internal/resource"
	pbmesh "github.com/hashicorp/consul/proto-public/pbmesh/v1alpha1"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

func IsRouteType(typ *pbresource.Type) bool {
	return IsHTTPRouteType(typ) ||
		IsGRPCRouteType(typ) ||
		IsTCPRouteType(typ)
}

func IsHTTPRouteType(typ *pbresource.Type) bool {
	switch {
	case resource.EqualType(typ, HTTPRouteType):
		return true
	}
	return false
}

func IsGRPCRouteType(typ *pbresource.Type) bool {
	switch {
	case resource.EqualType(typ, GRPCRouteType):
		return true
	}
	return false
}

func IsTCPRouteType(typ *pbresource.Type) bool {
	switch {
	case resource.EqualType(typ, TCPRouteType):
		return true
	}
	return false
}

func IsFailoverPolicyType(typ *pbresource.Type) bool {
	switch {
	case resource.EqualType(typ, catalog.FailoverPolicyType):
		return true
	}
	return false
}

func IsDestinationPolicyType(typ *pbresource.Type) bool {
	switch {
	case resource.EqualType(typ, DestinationPolicyType):
		return true
	}
	return false
}

func IsServiceType(typ *pbresource.Type) bool {
	switch {
	case resource.EqualType(typ, catalog.ServiceType):
		return true
	}
	return false
}

func IsComputedRoutesType(typ *pbresource.Type) bool {
	switch {
	case resource.EqualType(typ, ComputedRoutesType):
		return true
	}
	return false
}

// BackendRefToComputedRoutesTarget turns the provided BackendReference into an
// opaque string format suitable for use as a map key and reference in a
// standalone object or reference.
//
// It is opaque in that the caller should not attempt to parse it, and there is
// no implied storage or wire compatibility concern, since the data is treated
// opaquely at use time.
func BackendRefToComputedRoutesTarget(backendRef *pbmesh.BackendReference) string {
	ref := backendRef.Ref

	s := fmt.Sprintf(
		"%s/%s/%s?port=%s",
		resource.TypeToString(ref.Type),
		resource.TenancyToString(ref.Tenancy),
		ref.Name,
		backendRef.Port,
	)

	if backendRef.Datacenter != "" {
		s += "&dc=" + backendRef.Datacenter
	}

	return s
}
