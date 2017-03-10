package s3crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws"
)

func TestEncodeMaterialDescription(t *testing.T) {
	md := MaterialDescription{}
	md["foo"] = aws.String("bar")
	b, err := md.encodeDescription()
	expected := `{"foo":"bar"}`
	assert.NoError(t, err)
	assert.Equal(t, expected, string(b))
}
func TestDecodeMaterialDescription(t *testing.T) {
	md := MaterialDescription{}
	json := `{"foo":"bar"}`
	err := md.decodeDescription([]byte(json))
	expected := MaterialDescription{
		"foo": aws.String("bar"),
	}
	assert.NoError(t, err)
	assert.Equal(t, expected, md)
}
