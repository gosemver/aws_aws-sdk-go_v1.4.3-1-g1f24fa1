package request

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/awserr"
)

func TestRequestThrottling(t *testing.T) {
	req := Request{}

	req.Error = awserr.New("Throttling", "", nil)
	assert.True(t, req.IsErrorThrottle())
}
