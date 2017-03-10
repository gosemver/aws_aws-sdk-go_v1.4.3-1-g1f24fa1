// +build integration

//Package sqs provides gucumber integration tests support.
package sqs

import (
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/awstesting/integration/smoke"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/sqs"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@sqs", func() {
		gucumber.World["client"] = sqs.New(smoke.Session)
	})
}
