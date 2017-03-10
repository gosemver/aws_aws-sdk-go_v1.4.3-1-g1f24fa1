package ec2

import (
	"time"

	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/awsutil"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/request"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/private/endpoints"
)

func init() {
	initRequest = func(r *request.Request) {
		if r.Operation.Name == opCopySnapshot { // fill the PresignedURL parameter
			r.Handlers.Build.PushFront(fillPresignedURL)
		}
	}
}

func fillPresignedURL(r *request.Request) {
	if !r.ParamsFilled() {
		return
	}

	origParams := r.Params.(*CopySnapshotInput)

	// Stop if PresignedURL/DestinationRegion is set
	if origParams.PresignedUrl != nil || origParams.DestinationRegion != nil {
		return
	}

	origParams.DestinationRegion = r.Config.Region
	newParams := awsutil.CopyOf(r.Params).(*CopySnapshotInput)

	// Create a new request based on the existing request. We will use this to
	// presign the CopySnapshot request against the source region.
	cfg := r.Config.Copy(aws.NewConfig().
		WithEndpoint("").
		WithRegion(aws.StringValue(origParams.SourceRegion)))

	clientInfo := r.ClientInfo
	clientInfo.Endpoint, clientInfo.SigningRegion = endpoints.EndpointForRegion(
		clientInfo.ServiceName,
		aws.StringValue(cfg.Region),
		aws.BoolValue(cfg.DisableSSL),
		aws.BoolValue(cfg.UseDualStack),
	)

	// Presign a CopySnapshot request with modified params
	req := request.New(*cfg, clientInfo, r.Handlers, r.Retryer, r.Operation, newParams, r.Data)
	url, err := req.Presign(5 * time.Minute) // 5 minutes should be enough.
	if err != nil {                          // bubble error back up to original request
		r.Error = err
		return
	}

	// We have our URL, set it on params
	origParams.PresignedUrl = &url
}
