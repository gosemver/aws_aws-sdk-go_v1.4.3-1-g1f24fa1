// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package applicationdiscoveryservice_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/session"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/applicationdiscoveryservice"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleApplicationDiscoveryService_CreateTags() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := applicationdiscoveryservice.New(sess)

	params := &applicationdiscoveryservice.CreateTagsInput{
		ConfigurationIds: []*string{ // Required
			aws.String("ConfigurationId"), // Required
			// More values...
		},
		Tags: []*applicationdiscoveryservice.Tag{ // Required
			{ // Required
				Key:   aws.String("TagKey"),   // Required
				Value: aws.String("TagValue"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.CreateTags(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleApplicationDiscoveryService_DeleteTags() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := applicationdiscoveryservice.New(sess)

	params := &applicationdiscoveryservice.DeleteTagsInput{
		ConfigurationIds: []*string{ // Required
			aws.String("ConfigurationId"), // Required
			// More values...
		},
		Tags: []*applicationdiscoveryservice.Tag{
			{ // Required
				Key:   aws.String("TagKey"),   // Required
				Value: aws.String("TagValue"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.DeleteTags(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleApplicationDiscoveryService_DescribeAgents() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := applicationdiscoveryservice.New(sess)

	params := &applicationdiscoveryservice.DescribeAgentsInput{
		AgentIds: []*string{
			aws.String("AgentId"), // Required
			// More values...
		},
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("NextToken"),
	}
	resp, err := svc.DescribeAgents(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleApplicationDiscoveryService_DescribeConfigurations() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := applicationdiscoveryservice.New(sess)

	params := &applicationdiscoveryservice.DescribeConfigurationsInput{
		ConfigurationIds: []*string{ // Required
			aws.String("ConfigurationId"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeConfigurations(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleApplicationDiscoveryService_DescribeExportConfigurations() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := applicationdiscoveryservice.New(sess)

	params := &applicationdiscoveryservice.DescribeExportConfigurationsInput{
		ExportIds: []*string{
			aws.String("ConfigurationsExportId"), // Required
			// More values...
		},
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("NextToken"),
	}
	resp, err := svc.DescribeExportConfigurations(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleApplicationDiscoveryService_DescribeTags() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := applicationdiscoveryservice.New(sess)

	params := &applicationdiscoveryservice.DescribeTagsInput{
		Filters: []*applicationdiscoveryservice.TagFilter{
			{ // Required
				Name: aws.String("FilterName"), // Required
				Values: []*string{ // Required
					aws.String("FilterValue"), // Required
					// More values...
				},
			},
			// More values...
		},
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("NextToken"),
	}
	resp, err := svc.DescribeTags(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleApplicationDiscoveryService_ExportConfigurations() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := applicationdiscoveryservice.New(sess)

	var params *applicationdiscoveryservice.ExportConfigurationsInput
	resp, err := svc.ExportConfigurations(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleApplicationDiscoveryService_ListConfigurations() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := applicationdiscoveryservice.New(sess)

	params := &applicationdiscoveryservice.ListConfigurationsInput{
		ConfigurationType: aws.String("ConfigurationItemType"), // Required
		Filters: []*applicationdiscoveryservice.Filter{
			{ // Required
				Condition: aws.String("Condition"), // Required
				Name:      aws.String("String"),    // Required
				Values: []*string{ // Required
					aws.String("FilterValue"), // Required
					// More values...
				},
			},
			// More values...
		},
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("NextToken"),
	}
	resp, err := svc.ListConfigurations(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleApplicationDiscoveryService_StartDataCollectionByAgentIds() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := applicationdiscoveryservice.New(sess)

	params := &applicationdiscoveryservice.StartDataCollectionByAgentIdsInput{
		AgentIds: []*string{ // Required
			aws.String("AgentId"), // Required
			// More values...
		},
	}
	resp, err := svc.StartDataCollectionByAgentIds(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleApplicationDiscoveryService_StopDataCollectionByAgentIds() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := applicationdiscoveryservice.New(sess)

	params := &applicationdiscoveryservice.StopDataCollectionByAgentIdsInput{
		AgentIds: []*string{ // Required
			aws.String("AgentId"), // Required
			// More values...
		},
	}
	resp, err := svc.StopDataCollectionByAgentIds(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
