// +build integration

//Package sns provides gucumber integration tests support.
package sns

import (
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/awstesting/integration/smoke"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/sns"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@sns", func() {
		gucumber.World["client"] = sns.New(smoke.Session)
	})
}
