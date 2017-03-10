// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package inspector_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/session"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/inspector"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleInspector_AddAttributesToFindings() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.AddAttributesToFindingsInput{
		Attributes: []*inspector.Attribute{ // Required
			{ // Required
				Key:   aws.String("AttributeKey"), // Required
				Value: aws.String("AttributeValue"),
			},
			// More values...
		},
		FindingArns: []*string{ // Required
			aws.String("Arn"), // Required
			// More values...
		},
	}
	resp, err := svc.AddAttributesToFindings(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_CreateAssessmentTarget() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.CreateAssessmentTargetInput{
		AssessmentTargetName: aws.String("AssessmentTargetName"), // Required
		ResourceGroupArn:     aws.String("Arn"),                  // Required
	}
	resp, err := svc.CreateAssessmentTarget(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_CreateAssessmentTemplate() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.CreateAssessmentTemplateInput{
		AssessmentTargetArn:    aws.String("Arn"),                    // Required
		AssessmentTemplateName: aws.String("AssessmentTemplateName"), // Required
		DurationInSeconds:      aws.Int64(1),                         // Required
		RulesPackageArns: []*string{ // Required
			aws.String("Arn"), // Required
			// More values...
		},
		UserAttributesForFindings: []*inspector.Attribute{
			{ // Required
				Key:   aws.String("AttributeKey"), // Required
				Value: aws.String("AttributeValue"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateAssessmentTemplate(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_CreateResourceGroup() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.CreateResourceGroupInput{
		ResourceGroupTags: []*inspector.ResourceGroupTag{ // Required
			{ // Required
				Key:   aws.String("TagKey"), // Required
				Value: aws.String("TagValue"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateResourceGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_DeleteAssessmentRun() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.DeleteAssessmentRunInput{
		AssessmentRunArn: aws.String("Arn"), // Required
	}
	resp, err := svc.DeleteAssessmentRun(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_DeleteAssessmentTarget() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.DeleteAssessmentTargetInput{
		AssessmentTargetArn: aws.String("Arn"), // Required
	}
	resp, err := svc.DeleteAssessmentTarget(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_DeleteAssessmentTemplate() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.DeleteAssessmentTemplateInput{
		AssessmentTemplateArn: aws.String("Arn"), // Required
	}
	resp, err := svc.DeleteAssessmentTemplate(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_DescribeAssessmentRuns() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.DescribeAssessmentRunsInput{
		AssessmentRunArns: []*string{ // Required
			aws.String("Arn"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeAssessmentRuns(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_DescribeAssessmentTargets() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.DescribeAssessmentTargetsInput{
		AssessmentTargetArns: []*string{ // Required
			aws.String("Arn"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeAssessmentTargets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_DescribeAssessmentTemplates() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.DescribeAssessmentTemplatesInput{
		AssessmentTemplateArns: []*string{ // Required
			aws.String("Arn"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeAssessmentTemplates(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_DescribeCrossAccountAccessRole() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	var params *inspector.DescribeCrossAccountAccessRoleInput
	resp, err := svc.DescribeCrossAccountAccessRole(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_DescribeFindings() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.DescribeFindingsInput{
		FindingArns: []*string{ // Required
			aws.String("Arn"), // Required
			// More values...
		},
		Locale: aws.String("Locale"),
	}
	resp, err := svc.DescribeFindings(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_DescribeResourceGroups() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.DescribeResourceGroupsInput{
		ResourceGroupArns: []*string{ // Required
			aws.String("Arn"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeResourceGroups(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_DescribeRulesPackages() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.DescribeRulesPackagesInput{
		RulesPackageArns: []*string{ // Required
			aws.String("Arn"), // Required
			// More values...
		},
		Locale: aws.String("Locale"),
	}
	resp, err := svc.DescribeRulesPackages(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_GetTelemetryMetadata() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.GetTelemetryMetadataInput{
		AssessmentRunArn: aws.String("Arn"), // Required
	}
	resp, err := svc.GetTelemetryMetadata(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_ListAssessmentRunAgents() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.ListAssessmentRunAgentsInput{
		AssessmentRunArn: aws.String("Arn"), // Required
		Filter: &inspector.AgentFilter{
			AgentHealthCodes: []*string{ // Required
				aws.String("AgentHealthCode"), // Required
				// More values...
			},
			AgentHealths: []*string{ // Required
				aws.String("AgentHealth"), // Required
				// More values...
			},
		},
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("PaginationToken"),
	}
	resp, err := svc.ListAssessmentRunAgents(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_ListAssessmentRuns() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.ListAssessmentRunsInput{
		AssessmentTemplateArns: []*string{
			aws.String("Arn"), // Required
			// More values...
		},
		Filter: &inspector.AssessmentRunFilter{
			CompletionTimeRange: &inspector.TimestampRange{
				BeginDate: aws.Time(time.Now()),
				EndDate:   aws.Time(time.Now()),
			},
			DurationRange: &inspector.DurationRange{
				MaxSeconds: aws.Int64(1),
				MinSeconds: aws.Int64(1),
			},
			NamePattern: aws.String("NamePattern"),
			RulesPackageArns: []*string{
				aws.String("Arn"), // Required
				// More values...
			},
			StartTimeRange: &inspector.TimestampRange{
				BeginDate: aws.Time(time.Now()),
				EndDate:   aws.Time(time.Now()),
			},
			StateChangeTimeRange: &inspector.TimestampRange{
				BeginDate: aws.Time(time.Now()),
				EndDate:   aws.Time(time.Now()),
			},
			States: []*string{
				aws.String("AssessmentRunState"), // Required
				// More values...
			},
		},
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("PaginationToken"),
	}
	resp, err := svc.ListAssessmentRuns(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_ListAssessmentTargets() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.ListAssessmentTargetsInput{
		Filter: &inspector.AssessmentTargetFilter{
			AssessmentTargetNamePattern: aws.String("NamePattern"),
		},
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("PaginationToken"),
	}
	resp, err := svc.ListAssessmentTargets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_ListAssessmentTemplates() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.ListAssessmentTemplatesInput{
		AssessmentTargetArns: []*string{
			aws.String("Arn"), // Required
			// More values...
		},
		Filter: &inspector.AssessmentTemplateFilter{
			DurationRange: &inspector.DurationRange{
				MaxSeconds: aws.Int64(1),
				MinSeconds: aws.Int64(1),
			},
			NamePattern: aws.String("NamePattern"),
			RulesPackageArns: []*string{
				aws.String("Arn"), // Required
				// More values...
			},
		},
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("PaginationToken"),
	}
	resp, err := svc.ListAssessmentTemplates(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_ListEventSubscriptions() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.ListEventSubscriptionsInput{
		MaxResults:  aws.Int64(1),
		NextToken:   aws.String("PaginationToken"),
		ResourceArn: aws.String("Arn"),
	}
	resp, err := svc.ListEventSubscriptions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_ListFindings() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.ListFindingsInput{
		AssessmentRunArns: []*string{
			aws.String("Arn"), // Required
			// More values...
		},
		Filter: &inspector.FindingFilter{
			AgentIds: []*string{
				aws.String("AgentId"), // Required
				// More values...
			},
			Attributes: []*inspector.Attribute{
				{ // Required
					Key:   aws.String("AttributeKey"), // Required
					Value: aws.String("AttributeValue"),
				},
				// More values...
			},
			AutoScalingGroups: []*string{
				aws.String("AutoScalingGroup"), // Required
				// More values...
			},
			CreationTimeRange: &inspector.TimestampRange{
				BeginDate: aws.Time(time.Now()),
				EndDate:   aws.Time(time.Now()),
			},
			RuleNames: []*string{
				aws.String("RuleName"), // Required
				// More values...
			},
			RulesPackageArns: []*string{
				aws.String("Arn"), // Required
				// More values...
			},
			Severities: []*string{
				aws.String("Severity"), // Required
				// More values...
			},
			UserAttributes: []*inspector.Attribute{
				{ // Required
					Key:   aws.String("AttributeKey"), // Required
					Value: aws.String("AttributeValue"),
				},
				// More values...
			},
		},
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("PaginationToken"),
	}
	resp, err := svc.ListFindings(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_ListRulesPackages() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.ListRulesPackagesInput{
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("PaginationToken"),
	}
	resp, err := svc.ListRulesPackages(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_ListTagsForResource() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.ListTagsForResourceInput{
		ResourceArn: aws.String("Arn"), // Required
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

func ExampleInspector_PreviewAgents() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.PreviewAgentsInput{
		PreviewAgentsArn: aws.String("Arn"), // Required
		MaxResults:       aws.Int64(1),
		NextToken:        aws.String("PaginationToken"),
	}
	resp, err := svc.PreviewAgents(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_RegisterCrossAccountAccessRole() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.RegisterCrossAccountAccessRoleInput{
		RoleArn: aws.String("Arn"), // Required
	}
	resp, err := svc.RegisterCrossAccountAccessRole(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_RemoveAttributesFromFindings() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.RemoveAttributesFromFindingsInput{
		AttributeKeys: []*string{ // Required
			aws.String("AttributeKey"), // Required
			// More values...
		},
		FindingArns: []*string{ // Required
			aws.String("Arn"), // Required
			// More values...
		},
	}
	resp, err := svc.RemoveAttributesFromFindings(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_SetTagsForResource() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.SetTagsForResourceInput{
		ResourceArn: aws.String("Arn"), // Required
		Tags: []*inspector.Tag{
			{ // Required
				Key:   aws.String("TagKey"), // Required
				Value: aws.String("TagValue"),
			},
			// More values...
		},
	}
	resp, err := svc.SetTagsForResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_StartAssessmentRun() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.StartAssessmentRunInput{
		AssessmentTemplateArn: aws.String("Arn"), // Required
		AssessmentRunName:     aws.String("AssessmentRunName"),
	}
	resp, err := svc.StartAssessmentRun(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_StopAssessmentRun() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.StopAssessmentRunInput{
		AssessmentRunArn: aws.String("Arn"), // Required
	}
	resp, err := svc.StopAssessmentRun(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_SubscribeToEvent() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.SubscribeToEventInput{
		Event:       aws.String("Event"), // Required
		ResourceArn: aws.String("Arn"),   // Required
		TopicArn:    aws.String("Arn"),   // Required
	}
	resp, err := svc.SubscribeToEvent(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_UnsubscribeFromEvent() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.UnsubscribeFromEventInput{
		Event:       aws.String("Event"), // Required
		ResourceArn: aws.String("Arn"),   // Required
		TopicArn:    aws.String("Arn"),   // Required
	}
	resp, err := svc.UnsubscribeFromEvent(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_UpdateAssessmentTarget() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := inspector.New(sess)

	params := &inspector.UpdateAssessmentTargetInput{
		AssessmentTargetArn:  aws.String("Arn"),                  // Required
		AssessmentTargetName: aws.String("AssessmentTargetName"), // Required
		ResourceGroupArn:     aws.String("Arn"),                  // Required
	}
	resp, err := svc.UpdateAssessmentTarget(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
