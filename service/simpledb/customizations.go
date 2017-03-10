package simpledb

import "github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/client"

func init() {
	initClient = func(c *client.Client) {
		// SimpleDB uses custom error unmarshaling logic
		c.Handlers.UnmarshalError.Clear()
		c.Handlers.UnmarshalError.PushBack(unmarshalError)
	}
}
