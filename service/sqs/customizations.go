package sqs

import "github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/request"

func init() {
	initRequest = func(r *request.Request) {
		setupChecksumValidation(r)
	}
}
