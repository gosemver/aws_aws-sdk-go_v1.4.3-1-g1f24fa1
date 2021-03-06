// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package workspacesiface provides an interface for the Amazon WorkSpaces.
package workspacesiface

import (
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/request"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/workspaces"
)

// WorkSpacesAPI is the interface type for workspaces.WorkSpaces.
type WorkSpacesAPI interface {
	CreateTagsRequest(*workspaces.CreateTagsInput) (*request.Request, *workspaces.CreateTagsOutput)

	CreateTags(*workspaces.CreateTagsInput) (*workspaces.CreateTagsOutput, error)

	CreateWorkspacesRequest(*workspaces.CreateWorkspacesInput) (*request.Request, *workspaces.CreateWorkspacesOutput)

	CreateWorkspaces(*workspaces.CreateWorkspacesInput) (*workspaces.CreateWorkspacesOutput, error)

	DeleteTagsRequest(*workspaces.DeleteTagsInput) (*request.Request, *workspaces.DeleteTagsOutput)

	DeleteTags(*workspaces.DeleteTagsInput) (*workspaces.DeleteTagsOutput, error)

	DescribeTagsRequest(*workspaces.DescribeTagsInput) (*request.Request, *workspaces.DescribeTagsOutput)

	DescribeTags(*workspaces.DescribeTagsInput) (*workspaces.DescribeTagsOutput, error)

	DescribeWorkspaceBundlesRequest(*workspaces.DescribeWorkspaceBundlesInput) (*request.Request, *workspaces.DescribeWorkspaceBundlesOutput)

	DescribeWorkspaceBundles(*workspaces.DescribeWorkspaceBundlesInput) (*workspaces.DescribeWorkspaceBundlesOutput, error)

	DescribeWorkspaceBundlesPages(*workspaces.DescribeWorkspaceBundlesInput, func(*workspaces.DescribeWorkspaceBundlesOutput, bool) bool) error

	DescribeWorkspaceDirectoriesRequest(*workspaces.DescribeWorkspaceDirectoriesInput) (*request.Request, *workspaces.DescribeWorkspaceDirectoriesOutput)

	DescribeWorkspaceDirectories(*workspaces.DescribeWorkspaceDirectoriesInput) (*workspaces.DescribeWorkspaceDirectoriesOutput, error)

	DescribeWorkspaceDirectoriesPages(*workspaces.DescribeWorkspaceDirectoriesInput, func(*workspaces.DescribeWorkspaceDirectoriesOutput, bool) bool) error

	DescribeWorkspacesRequest(*workspaces.DescribeWorkspacesInput) (*request.Request, *workspaces.DescribeWorkspacesOutput)

	DescribeWorkspaces(*workspaces.DescribeWorkspacesInput) (*workspaces.DescribeWorkspacesOutput, error)

	DescribeWorkspacesPages(*workspaces.DescribeWorkspacesInput, func(*workspaces.DescribeWorkspacesOutput, bool) bool) error

	DescribeWorkspacesConnectionStatusRequest(*workspaces.DescribeWorkspacesConnectionStatusInput) (*request.Request, *workspaces.DescribeWorkspacesConnectionStatusOutput)

	DescribeWorkspacesConnectionStatus(*workspaces.DescribeWorkspacesConnectionStatusInput) (*workspaces.DescribeWorkspacesConnectionStatusOutput, error)

	ModifyWorkspacePropertiesRequest(*workspaces.ModifyWorkspacePropertiesInput) (*request.Request, *workspaces.ModifyWorkspacePropertiesOutput)

	ModifyWorkspaceProperties(*workspaces.ModifyWorkspacePropertiesInput) (*workspaces.ModifyWorkspacePropertiesOutput, error)

	RebootWorkspacesRequest(*workspaces.RebootWorkspacesInput) (*request.Request, *workspaces.RebootWorkspacesOutput)

	RebootWorkspaces(*workspaces.RebootWorkspacesInput) (*workspaces.RebootWorkspacesOutput, error)

	RebuildWorkspacesRequest(*workspaces.RebuildWorkspacesInput) (*request.Request, *workspaces.RebuildWorkspacesOutput)

	RebuildWorkspaces(*workspaces.RebuildWorkspacesInput) (*workspaces.RebuildWorkspacesOutput, error)

	StartWorkspacesRequest(*workspaces.StartWorkspacesInput) (*request.Request, *workspaces.StartWorkspacesOutput)

	StartWorkspaces(*workspaces.StartWorkspacesInput) (*workspaces.StartWorkspacesOutput, error)

	StopWorkspacesRequest(*workspaces.StopWorkspacesInput) (*request.Request, *workspaces.StopWorkspacesOutput)

	StopWorkspaces(*workspaces.StopWorkspacesInput) (*workspaces.StopWorkspacesOutput, error)

	TerminateWorkspacesRequest(*workspaces.TerminateWorkspacesInput) (*request.Request, *workspaces.TerminateWorkspacesOutput)

	TerminateWorkspaces(*workspaces.TerminateWorkspacesInput) (*workspaces.TerminateWorkspacesOutput, error)
}

var _ WorkSpacesAPI = (*workspaces.WorkSpaces)(nil)
