// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package gamelift_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/session"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/gamelift"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleGameLift_CreateAlias() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.CreateAliasInput{
		Name: aws.String("NonZeroAndMaxString"), // Required
		RoutingStrategy: &gamelift.RoutingStrategy{ // Required
			FleetId: aws.String("FleetId"),
			Message: aws.String("FreeText"),
			Type:    aws.String("RoutingStrategyType"),
		},
		Description: aws.String("NonZeroAndMaxString"),
	}
	resp, err := svc.CreateAlias(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_CreateBuild() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.CreateBuildInput{
		Name: aws.String("NonZeroAndMaxString"),
		StorageLocation: &gamelift.S3Location{
			Bucket:  aws.String("NonEmptyString"),
			Key:     aws.String("NonEmptyString"),
			RoleArn: aws.String("NonEmptyString"),
		},
		Version: aws.String("NonZeroAndMaxString"),
	}
	resp, err := svc.CreateBuild(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_CreateFleet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.CreateFleetInput{
		BuildId:         aws.String("BuildId"),             // Required
		EC2InstanceType: aws.String("EC2InstanceType"),     // Required
		Name:            aws.String("NonZeroAndMaxString"), // Required
		Description:     aws.String("NonZeroAndMaxString"),
		EC2InboundPermissions: []*gamelift.IpPermission{
			{ // Required
				FromPort: aws.Int64(1),                 // Required
				IpRange:  aws.String("NonBlankString"), // Required
				Protocol: aws.String("IpProtocol"),     // Required
				ToPort:   aws.Int64(1),                 // Required
			},
			// More values...
		},
		LogPaths: []*string{
			aws.String("NonZeroAndMaxString"), // Required
			// More values...
		},
		NewGameSessionProtectionPolicy: aws.String("ProtectionPolicy"),
		RuntimeConfiguration: &gamelift.RuntimeConfiguration{
			ServerProcesses: []*gamelift.ServerProcess{
				{ // Required
					ConcurrentExecutions: aws.Int64(1),                      // Required
					LaunchPath:           aws.String("NonZeroAndMaxString"), // Required
					Parameters:           aws.String("NonZeroAndMaxString"),
				},
				// More values...
			},
		},
		ServerLaunchParameters: aws.String("NonZeroAndMaxString"),
		ServerLaunchPath:       aws.String("NonZeroAndMaxString"),
	}
	resp, err := svc.CreateFleet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_CreateGameSession() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.CreateGameSessionInput{
		MaximumPlayerSessionCount: aws.Int64(1), // Required
		AliasId:                   aws.String("AliasId"),
		FleetId:                   aws.String("FleetId"),
		GameProperties: []*gamelift.GameProperty{
			{ // Required
				Key:   aws.String("GamePropertyKey"),   // Required
				Value: aws.String("GamePropertyValue"), // Required
			},
			// More values...
		},
		Name: aws.String("NonZeroAndMaxString"),
	}
	resp, err := svc.CreateGameSession(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_CreatePlayerSession() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.CreatePlayerSessionInput{
		GameSessionId: aws.String("GameSessionId"),       // Required
		PlayerId:      aws.String("NonZeroAndMaxString"), // Required
	}
	resp, err := svc.CreatePlayerSession(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_CreatePlayerSessions() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.CreatePlayerSessionsInput{
		GameSessionId: aws.String("GameSessionId"), // Required
		PlayerIds: []*string{ // Required
			aws.String("NonZeroAndMaxString"), // Required
			// More values...
		},
	}
	resp, err := svc.CreatePlayerSessions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DeleteAlias() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.DeleteAliasInput{
		AliasId: aws.String("AliasId"), // Required
	}
	resp, err := svc.DeleteAlias(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DeleteBuild() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.DeleteBuildInput{
		BuildId: aws.String("BuildId"), // Required
	}
	resp, err := svc.DeleteBuild(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DeleteFleet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.DeleteFleetInput{
		FleetId: aws.String("FleetId"), // Required
	}
	resp, err := svc.DeleteFleet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DeleteScalingPolicy() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.DeleteScalingPolicyInput{
		FleetId: aws.String("FleetId"),             // Required
		Name:    aws.String("NonZeroAndMaxString"), // Required
	}
	resp, err := svc.DeleteScalingPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DescribeAlias() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.DescribeAliasInput{
		AliasId: aws.String("AliasId"), // Required
	}
	resp, err := svc.DescribeAlias(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DescribeBuild() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.DescribeBuildInput{
		BuildId: aws.String("BuildId"), // Required
	}
	resp, err := svc.DescribeBuild(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DescribeEC2InstanceLimits() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.DescribeEC2InstanceLimitsInput{
		EC2InstanceType: aws.String("EC2InstanceType"),
	}
	resp, err := svc.DescribeEC2InstanceLimits(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DescribeFleetAttributes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.DescribeFleetAttributesInput{
		FleetIds: []*string{
			aws.String("FleetId"), // Required
			// More values...
		},
		Limit:     aws.Int64(1),
		NextToken: aws.String("NonZeroAndMaxString"),
	}
	resp, err := svc.DescribeFleetAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DescribeFleetCapacity() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.DescribeFleetCapacityInput{
		FleetIds: []*string{
			aws.String("FleetId"), // Required
			// More values...
		},
		Limit:     aws.Int64(1),
		NextToken: aws.String("NonZeroAndMaxString"),
	}
	resp, err := svc.DescribeFleetCapacity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DescribeFleetEvents() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.DescribeFleetEventsInput{
		FleetId:   aws.String("FleetId"), // Required
		EndTime:   aws.Time(time.Now()),
		Limit:     aws.Int64(1),
		NextToken: aws.String("NonZeroAndMaxString"),
		StartTime: aws.Time(time.Now()),
	}
	resp, err := svc.DescribeFleetEvents(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DescribeFleetPortSettings() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.DescribeFleetPortSettingsInput{
		FleetId: aws.String("FleetId"), // Required
	}
	resp, err := svc.DescribeFleetPortSettings(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DescribeFleetUtilization() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.DescribeFleetUtilizationInput{
		FleetIds: []*string{
			aws.String("FleetId"), // Required
			// More values...
		},
		Limit:     aws.Int64(1),
		NextToken: aws.String("NonZeroAndMaxString"),
	}
	resp, err := svc.DescribeFleetUtilization(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DescribeGameSessionDetails() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.DescribeGameSessionDetailsInput{
		AliasId:       aws.String("AliasId"),
		FleetId:       aws.String("FleetId"),
		GameSessionId: aws.String("GameSessionId"),
		Limit:         aws.Int64(1),
		NextToken:     aws.String("NonZeroAndMaxString"),
		StatusFilter:  aws.String("NonZeroAndMaxString"),
	}
	resp, err := svc.DescribeGameSessionDetails(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DescribeGameSessions() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.DescribeGameSessionsInput{
		AliasId:       aws.String("AliasId"),
		FleetId:       aws.String("FleetId"),
		GameSessionId: aws.String("GameSessionId"),
		Limit:         aws.Int64(1),
		NextToken:     aws.String("NonZeroAndMaxString"),
		StatusFilter:  aws.String("NonZeroAndMaxString"),
	}
	resp, err := svc.DescribeGameSessions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DescribePlayerSessions() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.DescribePlayerSessionsInput{
		GameSessionId:             aws.String("GameSessionId"),
		Limit:                     aws.Int64(1),
		NextToken:                 aws.String("NonZeroAndMaxString"),
		PlayerId:                  aws.String("NonZeroAndMaxString"),
		PlayerSessionId:           aws.String("PlayerSessionId"),
		PlayerSessionStatusFilter: aws.String("NonZeroAndMaxString"),
	}
	resp, err := svc.DescribePlayerSessions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DescribeRuntimeConfiguration() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.DescribeRuntimeConfigurationInput{
		FleetId: aws.String("FleetId"), // Required
	}
	resp, err := svc.DescribeRuntimeConfiguration(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DescribeScalingPolicies() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.DescribeScalingPoliciesInput{
		FleetId:      aws.String("FleetId"), // Required
		Limit:        aws.Int64(1),
		NextToken:    aws.String("NonZeroAndMaxString"),
		StatusFilter: aws.String("ScalingStatusType"),
	}
	resp, err := svc.DescribeScalingPolicies(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_GetGameSessionLogUrl() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.GetGameSessionLogUrlInput{
		GameSessionId: aws.String("GameSessionId"), // Required
	}
	resp, err := svc.GetGameSessionLogUrl(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_ListAliases() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.ListAliasesInput{
		Limit:               aws.Int64(1),
		Name:                aws.String("NonEmptyString"),
		NextToken:           aws.String("NonEmptyString"),
		RoutingStrategyType: aws.String("RoutingStrategyType"),
	}
	resp, err := svc.ListAliases(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_ListBuilds() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.ListBuildsInput{
		Limit:     aws.Int64(1),
		NextToken: aws.String("NonEmptyString"),
		Status:    aws.String("BuildStatus"),
	}
	resp, err := svc.ListBuilds(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_ListFleets() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.ListFleetsInput{
		BuildId:   aws.String("BuildId"),
		Limit:     aws.Int64(1),
		NextToken: aws.String("NonZeroAndMaxString"),
	}
	resp, err := svc.ListFleets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_PutScalingPolicy() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.PutScalingPolicyInput{
		ComparisonOperator:    aws.String("ComparisonOperatorType"), // Required
		EvaluationPeriods:     aws.Int64(1),                         // Required
		FleetId:               aws.String("FleetId"),                // Required
		MetricName:            aws.String("MetricName"),             // Required
		Name:                  aws.String("NonZeroAndMaxString"),    // Required
		ScalingAdjustment:     aws.Int64(1),                         // Required
		ScalingAdjustmentType: aws.String("ScalingAdjustmentType"),  // Required
		Threshold:             aws.Float64(1.0),                     // Required
	}
	resp, err := svc.PutScalingPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_RequestUploadCredentials() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.RequestUploadCredentialsInput{
		BuildId: aws.String("BuildId"), // Required
	}
	resp, err := svc.RequestUploadCredentials(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_ResolveAlias() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.ResolveAliasInput{
		AliasId: aws.String("AliasId"), // Required
	}
	resp, err := svc.ResolveAlias(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_SearchGameSessions() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.SearchGameSessionsInput{
		AliasId:          aws.String("AliasId"),
		FilterExpression: aws.String("NonZeroAndMaxString"),
		FleetId:          aws.String("FleetId"),
		Limit:            aws.Int64(1),
		NextToken:        aws.String("NonZeroAndMaxString"),
		SortExpression:   aws.String("NonZeroAndMaxString"),
	}
	resp, err := svc.SearchGameSessions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_UpdateAlias() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.UpdateAliasInput{
		AliasId:     aws.String("AliasId"), // Required
		Description: aws.String("NonZeroAndMaxString"),
		Name:        aws.String("NonZeroAndMaxString"),
		RoutingStrategy: &gamelift.RoutingStrategy{
			FleetId: aws.String("FleetId"),
			Message: aws.String("FreeText"),
			Type:    aws.String("RoutingStrategyType"),
		},
	}
	resp, err := svc.UpdateAlias(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_UpdateBuild() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.UpdateBuildInput{
		BuildId: aws.String("BuildId"), // Required
		Name:    aws.String("NonZeroAndMaxString"),
		Version: aws.String("NonZeroAndMaxString"),
	}
	resp, err := svc.UpdateBuild(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_UpdateFleetAttributes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.UpdateFleetAttributesInput{
		FleetId:     aws.String("FleetId"), // Required
		Description: aws.String("NonZeroAndMaxString"),
		Name:        aws.String("NonZeroAndMaxString"),
		NewGameSessionProtectionPolicy: aws.String("ProtectionPolicy"),
	}
	resp, err := svc.UpdateFleetAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_UpdateFleetCapacity() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.UpdateFleetCapacityInput{
		FleetId:          aws.String("FleetId"), // Required
		DesiredInstances: aws.Int64(1),
		MaxSize:          aws.Int64(1),
		MinSize:          aws.Int64(1),
	}
	resp, err := svc.UpdateFleetCapacity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_UpdateFleetPortSettings() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.UpdateFleetPortSettingsInput{
		FleetId: aws.String("FleetId"), // Required
		InboundPermissionAuthorizations: []*gamelift.IpPermission{
			{ // Required
				FromPort: aws.Int64(1),                 // Required
				IpRange:  aws.String("NonBlankString"), // Required
				Protocol: aws.String("IpProtocol"),     // Required
				ToPort:   aws.Int64(1),                 // Required
			},
			// More values...
		},
		InboundPermissionRevocations: []*gamelift.IpPermission{
			{ // Required
				FromPort: aws.Int64(1),                 // Required
				IpRange:  aws.String("NonBlankString"), // Required
				Protocol: aws.String("IpProtocol"),     // Required
				ToPort:   aws.Int64(1),                 // Required
			},
			// More values...
		},
	}
	resp, err := svc.UpdateFleetPortSettings(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_UpdateGameSession() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.UpdateGameSessionInput{
		GameSessionId:             aws.String("GameSessionId"), // Required
		MaximumPlayerSessionCount: aws.Int64(1),
		Name: aws.String("NonZeroAndMaxString"),
		PlayerSessionCreationPolicy: aws.String("PlayerSessionCreationPolicy"),
		ProtectionPolicy:            aws.String("ProtectionPolicy"),
	}
	resp, err := svc.UpdateGameSession(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_UpdateRuntimeConfiguration() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := gamelift.New(sess)

	params := &gamelift.UpdateRuntimeConfigurationInput{
		FleetId: aws.String("FleetId"), // Required
		RuntimeConfiguration: &gamelift.RuntimeConfiguration{ // Required
			ServerProcesses: []*gamelift.ServerProcess{
				{ // Required
					ConcurrentExecutions: aws.Int64(1),                      // Required
					LaunchPath:           aws.String("NonZeroAndMaxString"), // Required
					Parameters:           aws.String("NonZeroAndMaxString"),
				},
				// More values...
			},
		},
	}
	resp, err := svc.UpdateRuntimeConfiguration(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
