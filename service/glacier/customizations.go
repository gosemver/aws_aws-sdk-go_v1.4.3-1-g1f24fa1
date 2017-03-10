package glacier

import (
	"encoding/hex"
	"reflect"

	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/awsutil"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/request"
)

var (
	defaultAccountID = "-"
)

func init() {
	initRequest = func(r *request.Request) {
		r.Handlers.Validate.PushFront(addAccountID)
		r.Handlers.Validate.PushFront(copyParams) // this happens first
		r.Handlers.Build.PushBack(addChecksum)
		r.Handlers.Build.PushBack(addAPIVersion)
	}
}

func copyParams(r *request.Request) {
	r.Params = awsutil.CopyOf(r.Params)
}

func addAccountID(r *request.Request) {
	if !r.ParamsFilled() {
		return
	}

	v := reflect.Indirect(reflect.ValueOf(r.Params))
	if f := v.FieldByName("AccountId"); f.IsNil() {
		f.Set(reflect.ValueOf(&defaultAccountID))
	}
}

func addChecksum(r *request.Request) {
	if r.Body == nil || r.HTTPRequest.Header.Get("X-Amz-Sha256-Tree-Hash") != "" {
		return
	}

	h := ComputeHashes(r.Body)
	hstr := hex.EncodeToString(h.TreeHash)
	r.HTTPRequest.Header.Set("X-Amz-Sha256-Tree-Hash", hstr)

	hLstr := hex.EncodeToString(h.LinearHash)
	r.HTTPRequest.Header.Set("X-Amz-Content-Sha256", hLstr)
}

func addAPIVersion(r *request.Request) {
	r.HTTPRequest.Header.Set("X-Amz-Glacier-Version", r.ClientInfo.APIVersion)
}
