// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package mobileanalytics provides a client for Amazon Mobile Analytics.
package mobileanalytics

import (
	"fmt"

	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/awsutil"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/request"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/private/protocol"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/private/protocol/restjson"
)

const opPutEvents = "PutEvents"

// PutEventsRequest generates a "aws/request.Request" representing the
// client's request for the PutEvents operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the PutEvents method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the PutEventsRequest method.
//    req, resp := client.PutEventsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
func (c *MobileAnalytics) PutEventsRequest(input *PutEventsInput) (req *request.Request, output *PutEventsOutput) {
	op := &request.Operation{
		Name:       opPutEvents,
		HTTPMethod: "POST",
		HTTPPath:   "/2014-06-05/events",
	}

	if input == nil {
		input = &PutEventsInput{}
	}

	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	output = &PutEventsOutput{}
	req.Data = output
	return
}

// The PutEvents operation records one or more events. You can have up to 1,500
// unique custom events per app, any combination of up to 40 attributes and
// metrics per custom event, and any number of attribute or metric values.
func (c *MobileAnalytics) PutEvents(input *PutEventsInput) (*PutEventsOutput, error) {
	req, out := c.PutEventsRequest(input)
	err := req.Send()
	return out, err
}

// A JSON object representing a batch of unique event occurrences in your app.
type Event struct {
	_ struct{} `type:"structure"`

	// A collection of key-value pairs that give additional context to the event.
	// The key-value pairs are specified by the developer.
	//
	// This collection can be empty or the attribute object can be omitted.
	Attributes map[string]*string `locationName:"attributes" type:"map"`

	// A name signifying an event that occurred in your app. This is used for grouping
	// and aggregating like events together for reporting purposes.
	EventType *string `locationName:"eventType" min:"1" type:"string" required:"true"`

	// A collection of key-value pairs that gives additional, measurable context
	// to the event. The key-value pairs are specified by the developer.
	//
	// This collection can be empty or the attribute object can be omitted.
	Metrics map[string]*float64 `locationName:"metrics" type:"map"`

	// The session the event occured within.
	Session *Session `locationName:"session" type:"structure"`

	// The time the event occurred in ISO 8601 standard date time format. For example,
	// 2014-06-30T19:07:47.885Z
	Timestamp *string `locationName:"timestamp" type:"string" required:"true"`

	// The version of the event.
	Version *string `locationName:"version" min:"1" type:"string"`
}

// String returns the string representation
func (s Event) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Event) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Event) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "Event"}
	if s.EventType == nil {
		invalidParams.Add(request.NewErrParamRequired("EventType"))
	}
	if s.EventType != nil && len(*s.EventType) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("EventType", 1))
	}
	if s.Timestamp == nil {
		invalidParams.Add(request.NewErrParamRequired("Timestamp"))
	}
	if s.Version != nil && len(*s.Version) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Version", 1))
	}
	if s.Session != nil {
		if err := s.Session.Validate(); err != nil {
			invalidParams.AddNested("Session", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A container for the data needed for a PutEvent operation
type PutEventsInput struct {
	_ struct{} `type:"structure"`

	// The client context including the client ID, app title, app version and package
	// name.
	ClientContext *string `location:"header" locationName:"x-amz-Client-Context" type:"string" required:"true"`

	// The encoding used for the client context.
	ClientContextEncoding *string `location:"header" locationName:"x-amz-Client-Context-Encoding" type:"string"`

	// An array of Event JSON objects
	Events []*Event `locationName:"events" type:"list" required:"true"`
}

// String returns the string representation
func (s PutEventsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PutEventsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutEventsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PutEventsInput"}
	if s.ClientContext == nil {
		invalidParams.Add(request.NewErrParamRequired("ClientContext"))
	}
	if s.Events == nil {
		invalidParams.Add(request.NewErrParamRequired("Events"))
	}
	if s.Events != nil {
		for i, v := range s.Events {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Events", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutEventsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutEventsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PutEventsOutput) GoString() string {
	return s.String()
}

// Describes the session. Session information is required on ALL events.
type Session struct {
	_ struct{} `type:"structure"`

	// The duration of the session.
	Duration *int64 `locationName:"duration" type:"long"`

	// A unique identifier for the session
	Id *string `locationName:"id" min:"1" type:"string"`

	// The time the event started in ISO 8601 standard date time format. For example,
	// 2014-06-30T19:07:47.885Z
	StartTimestamp *string `locationName:"startTimestamp" type:"string"`

	// The time the event terminated in ISO 8601 standard date time format. For
	// example, 2014-06-30T19:07:47.885Z
	StopTimestamp *string `locationName:"stopTimestamp" type:"string"`
}

// String returns the string representation
func (s Session) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Session) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Session) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "Session"}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Id", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
