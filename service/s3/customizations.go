package s3

import (
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/client"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/request"
)

func init() {
	initClient = defaultInitClientFn
	initRequest = defaultInitRequestFn
}

func defaultInitClientFn(c *client.Client) {
	// Support building custom endpoints based on config
	c.Handlers.Build.PushFront(updateEndpointForS3Config)

	// Require SSL when using SSE keys
	c.Handlers.Validate.PushBack(validateSSERequiresSSL)
	c.Handlers.Build.PushBack(computeSSEKeys)

	// S3 uses custom error unmarshaling logic
	c.Handlers.UnmarshalError.Clear()
	c.Handlers.UnmarshalError.PushBack(unmarshalError)
}

func defaultInitRequestFn(r *request.Request) {
	// Add reuest handlers for specific platforms.
	// e.g. 100-continue support for PUT requests using Go 1.6
	platformRequestHandlers(r)

	switch r.Operation.Name {
	case opPutBucketCors, opPutBucketLifecycle, opPutBucketPolicy,
		opPutBucketTagging, opDeleteObjects, opPutBucketLifecycleConfiguration,
		opPutBucketReplication:
		// These S3 operations require Content-MD5 to be set
		r.Handlers.Build.PushBack(contentMD5)
	case opGetBucketLocation:
		// GetBucketLocation has custom parsing logic
		r.Handlers.Unmarshal.PushFront(buildGetBucketLocation)
	case opCreateBucket:
		// Auto-populate LocationConstraint with current region
		r.Handlers.Validate.PushFront(populateLocationConstraint)
	case opCopyObject, opUploadPartCopy, opCompleteMultipartUpload:
		r.Handlers.Unmarshal.PushFront(copyMultipartStatusOKUnmarhsalError)
	}
}
