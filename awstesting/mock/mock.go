package mock

import (
	"net/http"
	"net/http/httptest"

	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/client"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/client/metadata"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/session"
)

// Session is a mock session which is used to hit the mock server
var Session = session.Must(session.NewSession(&aws.Config{
	DisableSSL: aws.Bool(true),
	Endpoint:   aws.String(server.URL[7:]),
}))

// server is the mock server that simply writes a 200 status back to the client
var server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}))

// NewMockClient creates and initializes a client that will connect to the
// mock server
func NewMockClient(cfgs ...*aws.Config) *client.Client {
	c := Session.ClientConfig("Mock", cfgs...)

	svc := client.New(
		*c.Config,
		metadata.ClientInfo{
			ServiceName:   "Mock",
			SigningRegion: c.SigningRegion,
			Endpoint:      c.Endpoint,
			APIVersion:    "2015-12-08",
			JSONVersion:   "1.1",
			TargetPrefix:  "MockServer",
		},
		c.Handlers,
	)

	return svc
}
