// +build integration

//Package route53 provides gucumber integration tests support.
package route53

import (
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/awstesting/integration/smoke"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/route53"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@route53", func() {
		gucumber.World["client"] = route53.New(smoke.Session)
	})
}
