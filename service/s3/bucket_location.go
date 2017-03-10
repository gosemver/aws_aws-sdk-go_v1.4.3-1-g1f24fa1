package s3

import (
	"io/ioutil"
	"regexp"

	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/awserr"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/awsutil"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/request"
)

var reBucketLocation = regexp.MustCompile(`>([^<>]+)<\/Location`)

func buildGetBucketLocation(r *request.Request) {
	if r.DataFilled() {
		out := r.Data.(*GetBucketLocationOutput)
		b, err := ioutil.ReadAll(r.HTTPResponse.Body)
		if err != nil {
			r.Error = awserr.New("SerializationError", "failed reading response body", err)
			return
		}

		match := reBucketLocation.FindSubmatch(b)
		if len(match) > 1 {
			loc := string(match[1])
			out.LocationConstraint = &loc
		}
	}
}

func populateLocationConstraint(r *request.Request) {
	if r.ParamsFilled() && aws.StringValue(r.Config.Region) != "us-east-1" {
		in := r.Params.(*CreateBucketInput)
		if in.CreateBucketConfiguration == nil {
			r.Params = awsutil.CopyOf(r.Params)
			in = r.Params.(*CreateBucketInput)
			in.CreateBucketConfiguration = &CreateBucketConfiguration{
				LocationConstraint: r.Config.Region,
			}
		}
	}
}
