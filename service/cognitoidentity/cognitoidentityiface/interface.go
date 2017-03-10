// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cognitoidentityiface provides an interface for the Amazon Cognito Identity.
package cognitoidentityiface

import (
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/request"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/cognitoidentity"
)

// CognitoIdentityAPI is the interface type for cognitoidentity.CognitoIdentity.
type CognitoIdentityAPI interface {
	CreateIdentityPoolRequest(*cognitoidentity.CreateIdentityPoolInput) (*request.Request, *cognitoidentity.IdentityPool)

	CreateIdentityPool(*cognitoidentity.CreateIdentityPoolInput) (*cognitoidentity.IdentityPool, error)

	DeleteIdentitiesRequest(*cognitoidentity.DeleteIdentitiesInput) (*request.Request, *cognitoidentity.DeleteIdentitiesOutput)

	DeleteIdentities(*cognitoidentity.DeleteIdentitiesInput) (*cognitoidentity.DeleteIdentitiesOutput, error)

	DeleteIdentityPoolRequest(*cognitoidentity.DeleteIdentityPoolInput) (*request.Request, *cognitoidentity.DeleteIdentityPoolOutput)

	DeleteIdentityPool(*cognitoidentity.DeleteIdentityPoolInput) (*cognitoidentity.DeleteIdentityPoolOutput, error)

	DescribeIdentityRequest(*cognitoidentity.DescribeIdentityInput) (*request.Request, *cognitoidentity.IdentityDescription)

	DescribeIdentity(*cognitoidentity.DescribeIdentityInput) (*cognitoidentity.IdentityDescription, error)

	DescribeIdentityPoolRequest(*cognitoidentity.DescribeIdentityPoolInput) (*request.Request, *cognitoidentity.IdentityPool)

	DescribeIdentityPool(*cognitoidentity.DescribeIdentityPoolInput) (*cognitoidentity.IdentityPool, error)

	GetCredentialsForIdentityRequest(*cognitoidentity.GetCredentialsForIdentityInput) (*request.Request, *cognitoidentity.GetCredentialsForIdentityOutput)

	GetCredentialsForIdentity(*cognitoidentity.GetCredentialsForIdentityInput) (*cognitoidentity.GetCredentialsForIdentityOutput, error)

	GetIdRequest(*cognitoidentity.GetIdInput) (*request.Request, *cognitoidentity.GetIdOutput)

	GetId(*cognitoidentity.GetIdInput) (*cognitoidentity.GetIdOutput, error)

	GetIdentityPoolRolesRequest(*cognitoidentity.GetIdentityPoolRolesInput) (*request.Request, *cognitoidentity.GetIdentityPoolRolesOutput)

	GetIdentityPoolRoles(*cognitoidentity.GetIdentityPoolRolesInput) (*cognitoidentity.GetIdentityPoolRolesOutput, error)

	GetOpenIdTokenRequest(*cognitoidentity.GetOpenIdTokenInput) (*request.Request, *cognitoidentity.GetOpenIdTokenOutput)

	GetOpenIdToken(*cognitoidentity.GetOpenIdTokenInput) (*cognitoidentity.GetOpenIdTokenOutput, error)

	GetOpenIdTokenForDeveloperIdentityRequest(*cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput) (*request.Request, *cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput)

	GetOpenIdTokenForDeveloperIdentity(*cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput) (*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput, error)

	ListIdentitiesRequest(*cognitoidentity.ListIdentitiesInput) (*request.Request, *cognitoidentity.ListIdentitiesOutput)

	ListIdentities(*cognitoidentity.ListIdentitiesInput) (*cognitoidentity.ListIdentitiesOutput, error)

	ListIdentityPoolsRequest(*cognitoidentity.ListIdentityPoolsInput) (*request.Request, *cognitoidentity.ListIdentityPoolsOutput)

	ListIdentityPools(*cognitoidentity.ListIdentityPoolsInput) (*cognitoidentity.ListIdentityPoolsOutput, error)

	LookupDeveloperIdentityRequest(*cognitoidentity.LookupDeveloperIdentityInput) (*request.Request, *cognitoidentity.LookupDeveloperIdentityOutput)

	LookupDeveloperIdentity(*cognitoidentity.LookupDeveloperIdentityInput) (*cognitoidentity.LookupDeveloperIdentityOutput, error)

	MergeDeveloperIdentitiesRequest(*cognitoidentity.MergeDeveloperIdentitiesInput) (*request.Request, *cognitoidentity.MergeDeveloperIdentitiesOutput)

	MergeDeveloperIdentities(*cognitoidentity.MergeDeveloperIdentitiesInput) (*cognitoidentity.MergeDeveloperIdentitiesOutput, error)

	SetIdentityPoolRolesRequest(*cognitoidentity.SetIdentityPoolRolesInput) (*request.Request, *cognitoidentity.SetIdentityPoolRolesOutput)

	SetIdentityPoolRoles(*cognitoidentity.SetIdentityPoolRolesInput) (*cognitoidentity.SetIdentityPoolRolesOutput, error)

	UnlinkDeveloperIdentityRequest(*cognitoidentity.UnlinkDeveloperIdentityInput) (*request.Request, *cognitoidentity.UnlinkDeveloperIdentityOutput)

	UnlinkDeveloperIdentity(*cognitoidentity.UnlinkDeveloperIdentityInput) (*cognitoidentity.UnlinkDeveloperIdentityOutput, error)

	UnlinkIdentityRequest(*cognitoidentity.UnlinkIdentityInput) (*request.Request, *cognitoidentity.UnlinkIdentityOutput)

	UnlinkIdentity(*cognitoidentity.UnlinkIdentityInput) (*cognitoidentity.UnlinkIdentityOutput, error)

	UpdateIdentityPoolRequest(*cognitoidentity.IdentityPool) (*request.Request, *cognitoidentity.IdentityPool)

	UpdateIdentityPool(*cognitoidentity.IdentityPool) (*cognitoidentity.IdentityPool, error)
}

var _ CognitoIdentityAPI = (*cognitoidentity.CognitoIdentity)(nil)
