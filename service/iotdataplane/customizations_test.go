package iotdataplane_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/awstesting/unit"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/iotdataplane"
)

func TestRequireEndpointIfRegionProvided(t *testing.T) {
	svc := iotdataplane.New(unit.Session, &aws.Config{
		Region:                 aws.String("mock-region"),
		DisableParamValidation: aws.Bool(true),
	})
	req, _ := svc.GetThingShadowRequest(nil)
	err := req.Build()

	assert.Equal(t, "", svc.Endpoint)
	assert.Error(t, err)
	assert.Equal(t, aws.ErrMissingEndpoint, err)
}

func TestRequireEndpointIfNoRegionProvided(t *testing.T) {
	svc := iotdataplane.New(unit.Session, &aws.Config{
		Region:                 aws.String(""),
		DisableParamValidation: aws.Bool(true),
	})
	fmt.Println(svc.ClientInfo.SigningRegion)

	req, _ := svc.GetThingShadowRequest(nil)
	err := req.Build()

	assert.Equal(t, "", svc.Endpoint)
	assert.Error(t, err)
	assert.Equal(t, aws.ErrMissingEndpoint, err)
}

func TestRequireEndpointUsed(t *testing.T) {
	svc := iotdataplane.New(unit.Session, &aws.Config{
		Region:                 aws.String("mock-region"),
		DisableParamValidation: aws.Bool(true),
		Endpoint:               aws.String("https://endpoint"),
	})
	req, _ := svc.GetThingShadowRequest(nil)
	err := req.Build()

	assert.Equal(t, "https://endpoint", svc.Endpoint)
	assert.NoError(t, err)
}
