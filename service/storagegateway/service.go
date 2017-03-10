// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package storagegateway

import (
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/client"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/client/metadata"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/request"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/signer/v4"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/private/protocol/jsonrpc"
)

// AWS Storage Gateway is the service that connects an on-premises software
// appliance with cloud-based storage to provide seamless and secure integration
// between an organization's on-premises IT environment and AWS's storage infrastructure.
// The service enables you to securely upload data to the AWS cloud for cost
// effective backup and rapid disaster recovery.
//
// Use the following links to get started using the AWS Storage Gateway Service
// API Reference:
//
//    AWS Storage Gateway Required Request Headers (http://docs.aws.amazon.com/storagegateway/latest/userguide/AWSStorageGatewayHTTPRequestsHeaders.html):
// Describes the required headers that you must send with every POST request
// to AWS Storage Gateway.
//
//    Signing Requests (http://docs.aws.amazon.com/storagegateway/latest/userguide/AWSStorageGatewaySigningRequests.html):
// AWS Storage Gateway requires that you authenticate every request you send;
// this topic describes how sign such a request.
//
//    Error Responses (http://docs.aws.amazon.com/storagegateway/latest/userguide/APIErrorResponses.html):
// Provides reference information about AWS Storage Gateway errors.
//
//    Operations in AWS Storage Gateway (http://docs.aws.amazon.com/storagegateway/latest/userguide/AWSStorageGatewayAPIOperations.html):
// Contains detailed descriptions of all AWS Storage Gateway operations, their
// request parameters, response elements, possible errors, and examples of requests
// and responses.
//
//    AWS Storage Gateway Regions and Endpoints (http://docs.aws.amazon.com/general/latest/gr/index.html?rande.html):
// Provides a list of each of the s and endpoints available for use with AWS
// Storage Gateway.
//
//    AWS Storage Gateway resource IDs are in uppercase. When you use these
// resource IDs with the Amazon EC2 API, EC2 expects resource IDs in lowercase.
// You must change your resource ID to lowercase to use it with the EC2 API.
// For example, in Storage Gateway the ID for a volume might be vol-1122AABB.
// When you use this ID with the EC2 API, you must change it to vol-1122aabb.
// Otherwise, the EC2 API might not behave as expected.
//
//   IDs for Storage Gateway volumes and Amazon EBS snapshots created from
// gateway volumes are changing to a longer format. Starting in December 2016,
// all new volumes and snapshots will be created with a 17-character string.
// Starting in April 2016, you will be able to use these longer IDs so you can
// test your systems with the new format. For more information, see Longer EC2
// and EBS Resource IDs (https://aws.amazon.com/ec2/faqs/#longer-ids).
//
//  For example, a volume ARN with the longer volume ID format will look like
// this:
//
//  arn:aws:storagegateway:us-west-2:111122223333:gateway/sgw-12A3456B/volume/vol-1122AABBCCDDEEFFG.
//
// A snapshot ID with the longer ID format will look like this: snap-78e226633445566ee.
//
// For more information, see Announcement: Heads-up – Longer AWS Storage Gateway
// volume and snapshot IDs coming in 2016 (https://forums.aws.amazon.com/ann.jspa?annID=3557).
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type StorageGateway struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "storagegateway"

// New creates a new instance of the StorageGateway client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a StorageGateway client from just a session.
//     svc := storagegateway.New(mySession)
//
//     // Create a StorageGateway client with additional configuration
//     svc := storagegateway.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *StorageGateway {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *StorageGateway {
	svc := &StorageGateway{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2013-06-30",
				JSONVersion:   "1.1",
				TargetPrefix:  "StorageGateway_20130630",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a StorageGateway operation and runs any
// custom request initialization.
func (c *StorageGateway) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
