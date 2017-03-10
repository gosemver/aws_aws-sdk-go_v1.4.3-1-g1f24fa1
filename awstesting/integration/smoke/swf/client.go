// +build integration

//Package swf provides gucumber integration tests support.
package swf

import (
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/awstesting/integration/smoke"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/swf"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@swf", func() {
		gucumber.World["client"] = swf.New(smoke.Session)
	})
}
