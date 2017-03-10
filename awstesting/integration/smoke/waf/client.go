// +build integration

//Package waf provides gucumber integration tests support.
package waf

import (
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/awstesting/integration/smoke"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/waf"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@waf", func() {
		gucumber.World["client"] = waf.New(smoke.Session)
	})
}
