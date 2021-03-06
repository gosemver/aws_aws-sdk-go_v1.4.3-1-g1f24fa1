// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package applicationdiscoveryserviceiface provides an interface for the AWS Application Discovery Service.
package applicationdiscoveryserviceiface

import (
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/request"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/applicationdiscoveryservice"
)

// ApplicationDiscoveryServiceAPI is the interface type for applicationdiscoveryservice.ApplicationDiscoveryService.
type ApplicationDiscoveryServiceAPI interface {
	CreateTagsRequest(*applicationdiscoveryservice.CreateTagsInput) (*request.Request, *applicationdiscoveryservice.CreateTagsOutput)

	CreateTags(*applicationdiscoveryservice.CreateTagsInput) (*applicationdiscoveryservice.CreateTagsOutput, error)

	DeleteTagsRequest(*applicationdiscoveryservice.DeleteTagsInput) (*request.Request, *applicationdiscoveryservice.DeleteTagsOutput)

	DeleteTags(*applicationdiscoveryservice.DeleteTagsInput) (*applicationdiscoveryservice.DeleteTagsOutput, error)

	DescribeAgentsRequest(*applicationdiscoveryservice.DescribeAgentsInput) (*request.Request, *applicationdiscoveryservice.DescribeAgentsOutput)

	DescribeAgents(*applicationdiscoveryservice.DescribeAgentsInput) (*applicationdiscoveryservice.DescribeAgentsOutput, error)

	DescribeConfigurationsRequest(*applicationdiscoveryservice.DescribeConfigurationsInput) (*request.Request, *applicationdiscoveryservice.DescribeConfigurationsOutput)

	DescribeConfigurations(*applicationdiscoveryservice.DescribeConfigurationsInput) (*applicationdiscoveryservice.DescribeConfigurationsOutput, error)

	DescribeExportConfigurationsRequest(*applicationdiscoveryservice.DescribeExportConfigurationsInput) (*request.Request, *applicationdiscoveryservice.DescribeExportConfigurationsOutput)

	DescribeExportConfigurations(*applicationdiscoveryservice.DescribeExportConfigurationsInput) (*applicationdiscoveryservice.DescribeExportConfigurationsOutput, error)

	DescribeTagsRequest(*applicationdiscoveryservice.DescribeTagsInput) (*request.Request, *applicationdiscoveryservice.DescribeTagsOutput)

	DescribeTags(*applicationdiscoveryservice.DescribeTagsInput) (*applicationdiscoveryservice.DescribeTagsOutput, error)

	ExportConfigurationsRequest(*applicationdiscoveryservice.ExportConfigurationsInput) (*request.Request, *applicationdiscoveryservice.ExportConfigurationsOutput)

	ExportConfigurations(*applicationdiscoveryservice.ExportConfigurationsInput) (*applicationdiscoveryservice.ExportConfigurationsOutput, error)

	ListConfigurationsRequest(*applicationdiscoveryservice.ListConfigurationsInput) (*request.Request, *applicationdiscoveryservice.ListConfigurationsOutput)

	ListConfigurations(*applicationdiscoveryservice.ListConfigurationsInput) (*applicationdiscoveryservice.ListConfigurationsOutput, error)

	StartDataCollectionByAgentIdsRequest(*applicationdiscoveryservice.StartDataCollectionByAgentIdsInput) (*request.Request, *applicationdiscoveryservice.StartDataCollectionByAgentIdsOutput)

	StartDataCollectionByAgentIds(*applicationdiscoveryservice.StartDataCollectionByAgentIdsInput) (*applicationdiscoveryservice.StartDataCollectionByAgentIdsOutput, error)

	StopDataCollectionByAgentIdsRequest(*applicationdiscoveryservice.StopDataCollectionByAgentIdsInput) (*request.Request, *applicationdiscoveryservice.StopDataCollectionByAgentIdsOutput)

	StopDataCollectionByAgentIds(*applicationdiscoveryservice.StopDataCollectionByAgentIdsInput) (*applicationdiscoveryservice.StopDataCollectionByAgentIdsOutput, error)
}

var _ ApplicationDiscoveryServiceAPI = (*applicationdiscoveryservice.ApplicationDiscoveryService)(nil)
