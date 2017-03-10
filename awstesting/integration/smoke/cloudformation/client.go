// +build integration

//Package cloudformation provides gucumber integration tests support.
package cloudformation

import (
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/awstesting/integration/smoke"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/cloudformation"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@cloudformation", func() {
		gucumber.World["client"] = cloudformation.New(smoke.Session)
	})
}
