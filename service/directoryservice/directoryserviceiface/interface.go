// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package directoryserviceiface provides an interface for the AWS Directory Service.
package directoryserviceiface

import (
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/request"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/directoryservice"
)

// DirectoryServiceAPI is the interface type for directoryservice.DirectoryService.
type DirectoryServiceAPI interface {
	AddIpRoutesRequest(*directoryservice.AddIpRoutesInput) (*request.Request, *directoryservice.AddIpRoutesOutput)

	AddIpRoutes(*directoryservice.AddIpRoutesInput) (*directoryservice.AddIpRoutesOutput, error)

	AddTagsToResourceRequest(*directoryservice.AddTagsToResourceInput) (*request.Request, *directoryservice.AddTagsToResourceOutput)

	AddTagsToResource(*directoryservice.AddTagsToResourceInput) (*directoryservice.AddTagsToResourceOutput, error)

	ConnectDirectoryRequest(*directoryservice.ConnectDirectoryInput) (*request.Request, *directoryservice.ConnectDirectoryOutput)

	ConnectDirectory(*directoryservice.ConnectDirectoryInput) (*directoryservice.ConnectDirectoryOutput, error)

	CreateAliasRequest(*directoryservice.CreateAliasInput) (*request.Request, *directoryservice.CreateAliasOutput)

	CreateAlias(*directoryservice.CreateAliasInput) (*directoryservice.CreateAliasOutput, error)

	CreateComputerRequest(*directoryservice.CreateComputerInput) (*request.Request, *directoryservice.CreateComputerOutput)

	CreateComputer(*directoryservice.CreateComputerInput) (*directoryservice.CreateComputerOutput, error)

	CreateConditionalForwarderRequest(*directoryservice.CreateConditionalForwarderInput) (*request.Request, *directoryservice.CreateConditionalForwarderOutput)

	CreateConditionalForwarder(*directoryservice.CreateConditionalForwarderInput) (*directoryservice.CreateConditionalForwarderOutput, error)

	CreateDirectoryRequest(*directoryservice.CreateDirectoryInput) (*request.Request, *directoryservice.CreateDirectoryOutput)

	CreateDirectory(*directoryservice.CreateDirectoryInput) (*directoryservice.CreateDirectoryOutput, error)

	CreateMicrosoftADRequest(*directoryservice.CreateMicrosoftADInput) (*request.Request, *directoryservice.CreateMicrosoftADOutput)

	CreateMicrosoftAD(*directoryservice.CreateMicrosoftADInput) (*directoryservice.CreateMicrosoftADOutput, error)

	CreateSnapshotRequest(*directoryservice.CreateSnapshotInput) (*request.Request, *directoryservice.CreateSnapshotOutput)

	CreateSnapshot(*directoryservice.CreateSnapshotInput) (*directoryservice.CreateSnapshotOutput, error)

	CreateTrustRequest(*directoryservice.CreateTrustInput) (*request.Request, *directoryservice.CreateTrustOutput)

	CreateTrust(*directoryservice.CreateTrustInput) (*directoryservice.CreateTrustOutput, error)

	DeleteConditionalForwarderRequest(*directoryservice.DeleteConditionalForwarderInput) (*request.Request, *directoryservice.DeleteConditionalForwarderOutput)

	DeleteConditionalForwarder(*directoryservice.DeleteConditionalForwarderInput) (*directoryservice.DeleteConditionalForwarderOutput, error)

	DeleteDirectoryRequest(*directoryservice.DeleteDirectoryInput) (*request.Request, *directoryservice.DeleteDirectoryOutput)

	DeleteDirectory(*directoryservice.DeleteDirectoryInput) (*directoryservice.DeleteDirectoryOutput, error)

	DeleteSnapshotRequest(*directoryservice.DeleteSnapshotInput) (*request.Request, *directoryservice.DeleteSnapshotOutput)

	DeleteSnapshot(*directoryservice.DeleteSnapshotInput) (*directoryservice.DeleteSnapshotOutput, error)

	DeleteTrustRequest(*directoryservice.DeleteTrustInput) (*request.Request, *directoryservice.DeleteTrustOutput)

	DeleteTrust(*directoryservice.DeleteTrustInput) (*directoryservice.DeleteTrustOutput, error)

	DeregisterEventTopicRequest(*directoryservice.DeregisterEventTopicInput) (*request.Request, *directoryservice.DeregisterEventTopicOutput)

	DeregisterEventTopic(*directoryservice.DeregisterEventTopicInput) (*directoryservice.DeregisterEventTopicOutput, error)

	DescribeConditionalForwardersRequest(*directoryservice.DescribeConditionalForwardersInput) (*request.Request, *directoryservice.DescribeConditionalForwardersOutput)

	DescribeConditionalForwarders(*directoryservice.DescribeConditionalForwardersInput) (*directoryservice.DescribeConditionalForwardersOutput, error)

	DescribeDirectoriesRequest(*directoryservice.DescribeDirectoriesInput) (*request.Request, *directoryservice.DescribeDirectoriesOutput)

	DescribeDirectories(*directoryservice.DescribeDirectoriesInput) (*directoryservice.DescribeDirectoriesOutput, error)

	DescribeEventTopicsRequest(*directoryservice.DescribeEventTopicsInput) (*request.Request, *directoryservice.DescribeEventTopicsOutput)

	DescribeEventTopics(*directoryservice.DescribeEventTopicsInput) (*directoryservice.DescribeEventTopicsOutput, error)

	DescribeSnapshotsRequest(*directoryservice.DescribeSnapshotsInput) (*request.Request, *directoryservice.DescribeSnapshotsOutput)

	DescribeSnapshots(*directoryservice.DescribeSnapshotsInput) (*directoryservice.DescribeSnapshotsOutput, error)

	DescribeTrustsRequest(*directoryservice.DescribeTrustsInput) (*request.Request, *directoryservice.DescribeTrustsOutput)

	DescribeTrusts(*directoryservice.DescribeTrustsInput) (*directoryservice.DescribeTrustsOutput, error)

	DisableRadiusRequest(*directoryservice.DisableRadiusInput) (*request.Request, *directoryservice.DisableRadiusOutput)

	DisableRadius(*directoryservice.DisableRadiusInput) (*directoryservice.DisableRadiusOutput, error)

	DisableSsoRequest(*directoryservice.DisableSsoInput) (*request.Request, *directoryservice.DisableSsoOutput)

	DisableSso(*directoryservice.DisableSsoInput) (*directoryservice.DisableSsoOutput, error)

	EnableRadiusRequest(*directoryservice.EnableRadiusInput) (*request.Request, *directoryservice.EnableRadiusOutput)

	EnableRadius(*directoryservice.EnableRadiusInput) (*directoryservice.EnableRadiusOutput, error)

	EnableSsoRequest(*directoryservice.EnableSsoInput) (*request.Request, *directoryservice.EnableSsoOutput)

	EnableSso(*directoryservice.EnableSsoInput) (*directoryservice.EnableSsoOutput, error)

	GetDirectoryLimitsRequest(*directoryservice.GetDirectoryLimitsInput) (*request.Request, *directoryservice.GetDirectoryLimitsOutput)

	GetDirectoryLimits(*directoryservice.GetDirectoryLimitsInput) (*directoryservice.GetDirectoryLimitsOutput, error)

	GetSnapshotLimitsRequest(*directoryservice.GetSnapshotLimitsInput) (*request.Request, *directoryservice.GetSnapshotLimitsOutput)

	GetSnapshotLimits(*directoryservice.GetSnapshotLimitsInput) (*directoryservice.GetSnapshotLimitsOutput, error)

	ListIpRoutesRequest(*directoryservice.ListIpRoutesInput) (*request.Request, *directoryservice.ListIpRoutesOutput)

	ListIpRoutes(*directoryservice.ListIpRoutesInput) (*directoryservice.ListIpRoutesOutput, error)

	ListTagsForResourceRequest(*directoryservice.ListTagsForResourceInput) (*request.Request, *directoryservice.ListTagsForResourceOutput)

	ListTagsForResource(*directoryservice.ListTagsForResourceInput) (*directoryservice.ListTagsForResourceOutput, error)

	RegisterEventTopicRequest(*directoryservice.RegisterEventTopicInput) (*request.Request, *directoryservice.RegisterEventTopicOutput)

	RegisterEventTopic(*directoryservice.RegisterEventTopicInput) (*directoryservice.RegisterEventTopicOutput, error)

	RemoveIpRoutesRequest(*directoryservice.RemoveIpRoutesInput) (*request.Request, *directoryservice.RemoveIpRoutesOutput)

	RemoveIpRoutes(*directoryservice.RemoveIpRoutesInput) (*directoryservice.RemoveIpRoutesOutput, error)

	RemoveTagsFromResourceRequest(*directoryservice.RemoveTagsFromResourceInput) (*request.Request, *directoryservice.RemoveTagsFromResourceOutput)

	RemoveTagsFromResource(*directoryservice.RemoveTagsFromResourceInput) (*directoryservice.RemoveTagsFromResourceOutput, error)

	RestoreFromSnapshotRequest(*directoryservice.RestoreFromSnapshotInput) (*request.Request, *directoryservice.RestoreFromSnapshotOutput)

	RestoreFromSnapshot(*directoryservice.RestoreFromSnapshotInput) (*directoryservice.RestoreFromSnapshotOutput, error)

	UpdateConditionalForwarderRequest(*directoryservice.UpdateConditionalForwarderInput) (*request.Request, *directoryservice.UpdateConditionalForwarderOutput)

	UpdateConditionalForwarder(*directoryservice.UpdateConditionalForwarderInput) (*directoryservice.UpdateConditionalForwarderOutput, error)

	UpdateRadiusRequest(*directoryservice.UpdateRadiusInput) (*request.Request, *directoryservice.UpdateRadiusOutput)

	UpdateRadius(*directoryservice.UpdateRadiusInput) (*directoryservice.UpdateRadiusOutput, error)

	VerifyTrustRequest(*directoryservice.VerifyTrustInput) (*request.Request, *directoryservice.VerifyTrustOutput)

	VerifyTrust(*directoryservice.VerifyTrustInput) (*directoryservice.VerifyTrustOutput, error)
}

var _ DirectoryServiceAPI = (*directoryservice.DirectoryService)(nil)
