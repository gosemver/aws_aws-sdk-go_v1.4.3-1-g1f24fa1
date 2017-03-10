// +build integration

//Package route53domains provides gucumber integration tests support.
package route53domains

import (
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/awstesting/integration/smoke"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/route53domains"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@route53domains", func() {
		gucumber.World["client"] = route53domains.New(smoke.Session)
	})
}
