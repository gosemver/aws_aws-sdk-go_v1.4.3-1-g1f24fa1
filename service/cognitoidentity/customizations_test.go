package cognitoidentity_test

import (
	"testing"

	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/awstesting/unit"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/cognitoidentity"
	"github.com/stretchr/testify/assert"
)

var svc = cognitoidentity.New(unit.Session)

func TestUnsignedRequest_GetID(t *testing.T) {
	req, _ := svc.GetIdRequest(&cognitoidentity.GetIdInput{
		IdentityPoolId: aws.String("IdentityPoolId"),
	})

	err := req.Sign()
	assert.NoError(t, err)
	assert.Equal(t, "", req.HTTPRequest.Header.Get("Authorization"))
}

func TestUnsignedRequest_GetOpenIDToken(t *testing.T) {
	req, _ := svc.GetOpenIdTokenRequest(&cognitoidentity.GetOpenIdTokenInput{
		IdentityId: aws.String("IdentityId"),
	})

	err := req.Sign()
	assert.NoError(t, err)
	assert.Equal(t, "", req.HTTPRequest.Header.Get("Authorization"))
}

func TestUnsignedRequest_GetCredentialsForIdentity(t *testing.T) {
	req, _ := svc.GetCredentialsForIdentityRequest(&cognitoidentity.GetCredentialsForIdentityInput{
		IdentityId: aws.String("IdentityId"),
	})

	err := req.Sign()
	assert.NoError(t, err)
	assert.Equal(t, "", req.HTTPRequest.Header.Get("Authorization"))
}
