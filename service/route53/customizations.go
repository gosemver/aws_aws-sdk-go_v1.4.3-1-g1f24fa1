package route53

import (
	"regexp"

	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/client"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/request"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/private/protocol/restxml"
)

func init() {
	initClient = func(c *client.Client) {
		c.Handlers.Build.PushBack(sanitizeURL)
	}

	initRequest = func(r *request.Request) {
		switch r.Operation.Name {
		case opChangeResourceRecordSets:
			r.Handlers.UnmarshalError.Remove(restxml.UnmarshalErrorHandler)
			r.Handlers.UnmarshalError.PushBack(unmarshalChangeResourceRecordSetsError)
		}
	}
}

var reSanitizeURL = regexp.MustCompile(`\/%2F\w+%2F`)

func sanitizeURL(r *request.Request) {
	r.HTTPRequest.URL.Opaque =
		reSanitizeURL.ReplaceAllString(r.HTTPRequest.URL.Opaque, "/")
}
