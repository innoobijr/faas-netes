/*
Copyright 2019-2021 OpenFaaS Authors

Licensed under the MIT license. See LICENSE file in the project root for full license information.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	v1 "github.com/openfaas/faas-netes/pkg/apis/iam/v1"
	openfaasv1 "github.com/openfaas/faas-netes/pkg/apis/openfaas/v1"
	iamv1 "github.com/openfaas/faas-netes/pkg/client/applyconfiguration/iam/v1"
	applyconfigurationopenfaasv1 "github.com/openfaas/faas-netes/pkg/client/applyconfiguration/openfaas/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=iam.openfaas.com, Version=v1
	case v1.SchemeGroupVersion.WithKind("JwtIssuer"):
		return &iamv1.JwtIssuerApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("JwtIssuerSpec"):
		return &iamv1.JwtIssuerSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Policy"):
		return &iamv1.PolicyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PolicySpec"):
		return &iamv1.PolicySpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PolicyStatement"):
		return &iamv1.PolicyStatementApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Role"):
		return &iamv1.RoleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RoleSpec"):
		return &iamv1.RoleSpecApplyConfiguration{}

		// Group=openfaas.com, Version=v1
	case openfaasv1.SchemeGroupVersion.WithKind("Function"):
		return &applyconfigurationopenfaasv1.FunctionApplyConfiguration{}
	case openfaasv1.SchemeGroupVersion.WithKind("FunctionResources"):
		return &applyconfigurationopenfaasv1.FunctionResourcesApplyConfiguration{}
	case openfaasv1.SchemeGroupVersion.WithKind("FunctionSpec"):
		return &applyconfigurationopenfaasv1.FunctionSpecApplyConfiguration{}
	case openfaasv1.SchemeGroupVersion.WithKind("Profile"):
		return &applyconfigurationopenfaasv1.ProfileApplyConfiguration{}
	case openfaasv1.SchemeGroupVersion.WithKind("ProfileSpec"):
		return &applyconfigurationopenfaasv1.ProfileSpecApplyConfiguration{}

	}
	return nil
}