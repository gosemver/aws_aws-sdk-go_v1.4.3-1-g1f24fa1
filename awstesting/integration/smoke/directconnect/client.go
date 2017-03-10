// +build integration

//Package directconnect provides gucumber integration tests support.
package directconnect

import (
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/awstesting/integration/smoke"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/directconnect"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@directconnect", func() {
		gucumber.World["client"] = directconnect.New(smoke.Session)
	})
}
