// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package storagegateway_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/session"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/storagegateway"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleStorageGateway_ActivateGateway() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.ActivateGatewayInput{
		ActivationKey:     aws.String("ActivationKey"),   // Required
		GatewayName:       aws.String("GatewayName"),     // Required
		GatewayRegion:     aws.String("RegionId"),        // Required
		GatewayTimezone:   aws.String("GatewayTimezone"), // Required
		GatewayType:       aws.String("GatewayType"),
		MediumChangerType: aws.String("MediumChangerType"),
		TapeDriveType:     aws.String("TapeDriveType"),
	}
	resp, err := svc.ActivateGateway(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_AddCache() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.AddCacheInput{
		DiskIds: []*string{ // Required
			aws.String("DiskId"), // Required
			// More values...
		},
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.AddCache(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_AddTagsToResource() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.AddTagsToResourceInput{
		ResourceARN: aws.String("ResourceARN"), // Required
		Tags: []*storagegateway.Tag{ // Required
			{ // Required
				Key:   aws.String("TagKey"),   // Required
				Value: aws.String("TagValue"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.AddTagsToResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_AddUploadBuffer() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.AddUploadBufferInput{
		DiskIds: []*string{ // Required
			aws.String("DiskId"), // Required
			// More values...
		},
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.AddUploadBuffer(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_AddWorkingStorage() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.AddWorkingStorageInput{
		DiskIds: []*string{ // Required
			aws.String("DiskId"), // Required
			// More values...
		},
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.AddWorkingStorage(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_CancelArchival() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.CancelArchivalInput{
		GatewayARN: aws.String("GatewayARN"), // Required
		TapeARN:    aws.String("TapeARN"),    // Required
	}
	resp, err := svc.CancelArchival(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_CancelRetrieval() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.CancelRetrievalInput{
		GatewayARN: aws.String("GatewayARN"), // Required
		TapeARN:    aws.String("TapeARN"),    // Required
	}
	resp, err := svc.CancelRetrieval(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_CreateCachediSCSIVolume() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.CreateCachediSCSIVolumeInput{
		ClientToken:        aws.String("ClientToken"),        // Required
		GatewayARN:         aws.String("GatewayARN"),         // Required
		NetworkInterfaceId: aws.String("NetworkInterfaceId"), // Required
		TargetName:         aws.String("TargetName"),         // Required
		VolumeSizeInBytes:  aws.Int64(1),                     // Required
		SnapshotId:         aws.String("SnapshotId"),
	}
	resp, err := svc.CreateCachediSCSIVolume(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_CreateSnapshot() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.CreateSnapshotInput{
		SnapshotDescription: aws.String("SnapshotDescription"), // Required
		VolumeARN:           aws.String("VolumeARN"),           // Required
	}
	resp, err := svc.CreateSnapshot(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_CreateSnapshotFromVolumeRecoveryPoint() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.CreateSnapshotFromVolumeRecoveryPointInput{
		SnapshotDescription: aws.String("SnapshotDescription"), // Required
		VolumeARN:           aws.String("VolumeARN"),           // Required
	}
	resp, err := svc.CreateSnapshotFromVolumeRecoveryPoint(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_CreateStorediSCSIVolume() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.CreateStorediSCSIVolumeInput{
		DiskId:               aws.String("DiskId"),             // Required
		GatewayARN:           aws.String("GatewayARN"),         // Required
		NetworkInterfaceId:   aws.String("NetworkInterfaceId"), // Required
		PreserveExistingData: aws.Bool(true),                   // Required
		TargetName:           aws.String("TargetName"),         // Required
		SnapshotId:           aws.String("SnapshotId"),
	}
	resp, err := svc.CreateStorediSCSIVolume(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_CreateTapeWithBarcode() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.CreateTapeWithBarcodeInput{
		GatewayARN:      aws.String("GatewayARN"),  // Required
		TapeBarcode:     aws.String("TapeBarcode"), // Required
		TapeSizeInBytes: aws.Int64(1),              // Required
	}
	resp, err := svc.CreateTapeWithBarcode(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_CreateTapes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.CreateTapesInput{
		ClientToken:       aws.String("ClientToken"),       // Required
		GatewayARN:        aws.String("GatewayARN"),        // Required
		NumTapesToCreate:  aws.Int64(1),                    // Required
		TapeBarcodePrefix: aws.String("TapeBarcodePrefix"), // Required
		TapeSizeInBytes:   aws.Int64(1),                    // Required
	}
	resp, err := svc.CreateTapes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DeleteBandwidthRateLimit() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.DeleteBandwidthRateLimitInput{
		BandwidthType: aws.String("BandwidthType"), // Required
		GatewayARN:    aws.String("GatewayARN"),    // Required
	}
	resp, err := svc.DeleteBandwidthRateLimit(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DeleteChapCredentials() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.DeleteChapCredentialsInput{
		InitiatorName: aws.String("IqnName"),   // Required
		TargetARN:     aws.String("TargetARN"), // Required
	}
	resp, err := svc.DeleteChapCredentials(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DeleteGateway() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.DeleteGatewayInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.DeleteGateway(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DeleteSnapshotSchedule() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.DeleteSnapshotScheduleInput{
		VolumeARN: aws.String("VolumeARN"), // Required
	}
	resp, err := svc.DeleteSnapshotSchedule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DeleteTape() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.DeleteTapeInput{
		GatewayARN: aws.String("GatewayARN"), // Required
		TapeARN:    aws.String("TapeARN"),    // Required
	}
	resp, err := svc.DeleteTape(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DeleteTapeArchive() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.DeleteTapeArchiveInput{
		TapeARN: aws.String("TapeARN"), // Required
	}
	resp, err := svc.DeleteTapeArchive(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DeleteVolume() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.DeleteVolumeInput{
		VolumeARN: aws.String("VolumeARN"), // Required
	}
	resp, err := svc.DeleteVolume(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeBandwidthRateLimit() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeBandwidthRateLimitInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.DescribeBandwidthRateLimit(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeCache() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeCacheInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.DescribeCache(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeCachediSCSIVolumes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeCachediSCSIVolumesInput{
		VolumeARNs: []*string{ // Required
			aws.String("VolumeARN"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeCachediSCSIVolumes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeChapCredentials() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeChapCredentialsInput{
		TargetARN: aws.String("TargetARN"), // Required
	}
	resp, err := svc.DescribeChapCredentials(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeGatewayInformation() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeGatewayInformationInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.DescribeGatewayInformation(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeMaintenanceStartTime() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeMaintenanceStartTimeInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.DescribeMaintenanceStartTime(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeSnapshotSchedule() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeSnapshotScheduleInput{
		VolumeARN: aws.String("VolumeARN"), // Required
	}
	resp, err := svc.DescribeSnapshotSchedule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeStorediSCSIVolumes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeStorediSCSIVolumesInput{
		VolumeARNs: []*string{ // Required
			aws.String("VolumeARN"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeStorediSCSIVolumes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeTapeArchives() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeTapeArchivesInput{
		Limit:  aws.Int64(1),
		Marker: aws.String("Marker"),
		TapeARNs: []*string{
			aws.String("TapeARN"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeTapeArchives(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeTapeRecoveryPoints() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeTapeRecoveryPointsInput{
		GatewayARN: aws.String("GatewayARN"), // Required
		Limit:      aws.Int64(1),
		Marker:     aws.String("Marker"),
	}
	resp, err := svc.DescribeTapeRecoveryPoints(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeTapes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeTapesInput{
		GatewayARN: aws.String("GatewayARN"), // Required
		Limit:      aws.Int64(1),
		Marker:     aws.String("Marker"),
		TapeARNs: []*string{
			aws.String("TapeARN"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeTapes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeUploadBuffer() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeUploadBufferInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.DescribeUploadBuffer(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeVTLDevices() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeVTLDevicesInput{
		GatewayARN: aws.String("GatewayARN"), // Required
		Limit:      aws.Int64(1),
		Marker:     aws.String("Marker"),
		VTLDeviceARNs: []*string{
			aws.String("VTLDeviceARN"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeVTLDevices(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DescribeWorkingStorage() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.DescribeWorkingStorageInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.DescribeWorkingStorage(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_DisableGateway() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.DisableGatewayInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.DisableGateway(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_ListGateways() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.ListGatewaysInput{
		Limit:  aws.Int64(1),
		Marker: aws.String("Marker"),
	}
	resp, err := svc.ListGateways(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_ListLocalDisks() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.ListLocalDisksInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.ListLocalDisks(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_ListTagsForResource() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.ListTagsForResourceInput{
		ResourceARN: aws.String("ResourceARN"), // Required
		Limit:       aws.Int64(1),
		Marker:      aws.String("Marker"),
	}
	resp, err := svc.ListTagsForResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_ListTapes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.ListTapesInput{
		Limit:  aws.Int64(1),
		Marker: aws.String("Marker"),
		TapeARNs: []*string{
			aws.String("TapeARN"), // Required
			// More values...
		},
	}
	resp, err := svc.ListTapes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_ListVolumeInitiators() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.ListVolumeInitiatorsInput{
		VolumeARN: aws.String("VolumeARN"), // Required
	}
	resp, err := svc.ListVolumeInitiators(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_ListVolumeRecoveryPoints() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.ListVolumeRecoveryPointsInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.ListVolumeRecoveryPoints(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_ListVolumes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.ListVolumesInput{
		GatewayARN: aws.String("GatewayARN"),
		Limit:      aws.Int64(1),
		Marker:     aws.String("Marker"),
	}
	resp, err := svc.ListVolumes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_RemoveTagsFromResource() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.RemoveTagsFromResourceInput{
		ResourceARN: aws.String("ResourceARN"), // Required
		TagKeys: []*string{ // Required
			aws.String("TagKey"), // Required
			// More values...
		},
	}
	resp, err := svc.RemoveTagsFromResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_ResetCache() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.ResetCacheInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.ResetCache(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_RetrieveTapeArchive() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.RetrieveTapeArchiveInput{
		GatewayARN: aws.String("GatewayARN"), // Required
		TapeARN:    aws.String("TapeARN"),    // Required
	}
	resp, err := svc.RetrieveTapeArchive(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_RetrieveTapeRecoveryPoint() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.RetrieveTapeRecoveryPointInput{
		GatewayARN: aws.String("GatewayARN"), // Required
		TapeARN:    aws.String("TapeARN"),    // Required
	}
	resp, err := svc.RetrieveTapeRecoveryPoint(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_SetLocalConsolePassword() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.SetLocalConsolePasswordInput{
		GatewayARN:           aws.String("GatewayARN"),           // Required
		LocalConsolePassword: aws.String("LocalConsolePassword"), // Required
	}
	resp, err := svc.SetLocalConsolePassword(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_ShutdownGateway() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.ShutdownGatewayInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.ShutdownGateway(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_StartGateway() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.StartGatewayInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.StartGateway(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_UpdateBandwidthRateLimit() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.UpdateBandwidthRateLimitInput{
		GatewayARN:                           aws.String("GatewayARN"), // Required
		AverageDownloadRateLimitInBitsPerSec: aws.Int64(1),
		AverageUploadRateLimitInBitsPerSec:   aws.Int64(1),
	}
	resp, err := svc.UpdateBandwidthRateLimit(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_UpdateChapCredentials() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.UpdateChapCredentialsInput{
		InitiatorName:                 aws.String("IqnName"),    // Required
		SecretToAuthenticateInitiator: aws.String("ChapSecret"), // Required
		TargetARN:                     aws.String("TargetARN"),  // Required
		SecretToAuthenticateTarget:    aws.String("ChapSecret"),
	}
	resp, err := svc.UpdateChapCredentials(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_UpdateGatewayInformation() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.UpdateGatewayInformationInput{
		GatewayARN:      aws.String("GatewayARN"), // Required
		GatewayName:     aws.String("GatewayName"),
		GatewayTimezone: aws.String("GatewayTimezone"),
	}
	resp, err := svc.UpdateGatewayInformation(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_UpdateGatewaySoftwareNow() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.UpdateGatewaySoftwareNowInput{
		GatewayARN: aws.String("GatewayARN"), // Required
	}
	resp, err := svc.UpdateGatewaySoftwareNow(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_UpdateMaintenanceStartTime() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.UpdateMaintenanceStartTimeInput{
		DayOfWeek:    aws.Int64(1),             // Required
		GatewayARN:   aws.String("GatewayARN"), // Required
		HourOfDay:    aws.Int64(1),             // Required
		MinuteOfHour: aws.Int64(1),             // Required
	}
	resp, err := svc.UpdateMaintenanceStartTime(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_UpdateSnapshotSchedule() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.UpdateSnapshotScheduleInput{
		RecurrenceInHours: aws.Int64(1),            // Required
		StartAt:           aws.Int64(1),            // Required
		VolumeARN:         aws.String("VolumeARN"), // Required
		Description:       aws.String("Description"),
	}
	resp, err := svc.UpdateSnapshotSchedule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleStorageGateway_UpdateVTLDeviceType() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := storagegateway.New(sess)

	params := &storagegateway.UpdateVTLDeviceTypeInput{
		DeviceType:   aws.String("DeviceType"),   // Required
		VTLDeviceARN: aws.String("VTLDeviceARN"), // Required
	}
	resp, err := svc.UpdateVTLDeviceType(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
