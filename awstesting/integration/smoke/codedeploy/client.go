// +build integration

//Package codedeploy provides gucumber integration tests support.
package codedeploy

import (
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/awstesting/integration/smoke"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/codedeploy"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@codedeploy", func() {
		gucumber.World["client"] = codedeploy.New(smoke.Session)
	})
}
