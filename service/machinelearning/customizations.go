package machinelearning

import (
	"net/url"

	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/request"
)

func init() {
	initRequest = func(r *request.Request) {
		switch r.Operation.Name {
		case opPredict:
			r.Handlers.Build.PushBack(updatePredictEndpoint)
		}
	}
}

// updatePredictEndpoint rewrites the request endpoint to use the
// "PredictEndpoint" parameter of the Predict operation.
func updatePredictEndpoint(r *request.Request) {
	if !r.ParamsFilled() {
		return
	}

	r.ClientInfo.Endpoint = *r.Params.(*PredictInput).PredictEndpoint

	uri, err := url.Parse(r.ClientInfo.Endpoint)
	if err != nil {
		r.Error = err
		return
	}
	r.HTTPRequest.URL = uri
}
