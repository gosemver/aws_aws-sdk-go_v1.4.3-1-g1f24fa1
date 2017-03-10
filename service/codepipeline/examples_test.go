// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package codepipeline_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/session"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/codepipeline"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCodePipeline_AcknowledgeJob() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codepipeline.New(sess)

	params := &codepipeline.AcknowledgeJobInput{
		JobId: aws.String("JobId"), // Required
		Nonce: aws.String("Nonce"), // Required
	}
	resp, err := svc.AcknowledgeJob(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodePipeline_AcknowledgeThirdPartyJob() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codepipeline.New(sess)

	params := &codepipeline.AcknowledgeThirdPartyJobInput{
		ClientToken: aws.String("ClientToken"),     // Required
		JobId:       aws.String("ThirdPartyJobId"), // Required
		Nonce:       aws.String("Nonce"),           // Required
	}
	resp, err := svc.AcknowledgeThirdPartyJob(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodePipeline_CreateCustomActionType() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codepipeline.New(sess)

	params := &codepipeline.CreateCustomActionTypeInput{
		Category: aws.String("ActionCategory"), // Required
		InputArtifactDetails: &codepipeline.ArtifactDetails{ // Required
			MaximumCount: aws.Int64(1), // Required
			MinimumCount: aws.Int64(1), // Required
		},
		OutputArtifactDetails: &codepipeline.ArtifactDetails{ // Required
			MaximumCount: aws.Int64(1), // Required
			MinimumCount: aws.Int64(1), // Required
		},
		Provider: aws.String("ActionProvider"), // Required
		Version:  aws.String("Version"),        // Required
		ConfigurationProperties: []*codepipeline.ActionConfigurationProperty{
			{ // Required
				Key:         aws.Bool(true),                       // Required
				Name:        aws.String("ActionConfigurationKey"), // Required
				Required:    aws.Bool(true),                       // Required
				Secret:      aws.Bool(true),                       // Required
				Description: aws.String("Description"),
				Queryable:   aws.Bool(true),
				Type:        aws.String("ActionConfigurationPropertyType"),
			},
			// More values...
		},
		Settings: &codepipeline.ActionTypeSettings{
			EntityUrlTemplate:          aws.String("UrlTemplate"),
			ExecutionUrlTemplate:       aws.String("UrlTemplate"),
			RevisionUrlTemplate:        aws.String("UrlTemplate"),
			ThirdPartyConfigurationUrl: aws.String("Url"),
		},
	}
	resp, err := svc.CreateCustomActionType(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodePipeline_CreatePipeline() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codepipeline.New(sess)

	params := &codepipeline.CreatePipelineInput{
		Pipeline: &codepipeline.PipelineDeclaration{ // Required
			ArtifactStore: &codepipeline.ArtifactStore{ // Required
				Location: aws.String("ArtifactStoreLocation"), // Required
				Type:     aws.String("ArtifactStoreType"),     // Required
				EncryptionKey: &codepipeline.EncryptionKey{
					Id:   aws.String("EncryptionKeyId"),   // Required
					Type: aws.String("EncryptionKeyType"), // Required
				},
			},
			Name:    aws.String("PipelineName"), // Required
			RoleArn: aws.String("RoleArn"),      // Required
			Stages: []*codepipeline.StageDeclaration{ // Required
				{ // Required
					Actions: []*codepipeline.ActionDeclaration{ // Required
						{ // Required
							ActionTypeId: &codepipeline.ActionTypeId{ // Required
								Category: aws.String("ActionCategory"), // Required
								Owner:    aws.String("ActionOwner"),    // Required
								Provider: aws.String("ActionProvider"), // Required
								Version:  aws.String("Version"),        // Required
							},
							Name: aws.String("ActionName"), // Required
							Configuration: map[string]*string{
								"Key": aws.String("ActionConfigurationValue"), // Required
								// More values...
							},
							InputArtifacts: []*codepipeline.InputArtifact{
								{ // Required
									Name: aws.String("ArtifactName"), // Required
								},
								// More values...
							},
							OutputArtifacts: []*codepipeline.OutputArtifact{
								{ // Required
									Name: aws.String("ArtifactName"), // Required
								},
								// More values...
							},
							RoleArn:  aws.String("RoleArn"),
							RunOrder: aws.Int64(1),
						},
						// More values...
					},
					Name: aws.String("StageName"), // Required
					Blockers: []*codepipeline.BlockerDeclaration{
						{ // Required
							Name: aws.String("BlockerName"), // Required
							Type: aws.String("BlockerType"), // Required
						},
						// More values...
					},
				},
				// More values...
			},
			Version: aws.Int64(1),
		},
	}
	resp, err := svc.CreatePipeline(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodePipeline_DeleteCustomActionType() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codepipeline.New(sess)

	params := &codepipeline.DeleteCustomActionTypeInput{
		Category: aws.String("ActionCategory"), // Required
		Provider: aws.String("ActionProvider"), // Required
		Version:  aws.String("Version"),        // Required
	}
	resp, err := svc.DeleteCustomActionType(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodePipeline_DeletePipeline() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codepipeline.New(sess)

	params := &codepipeline.DeletePipelineInput{
		Name: aws.String("PipelineName"), // Required
	}
	resp, err := svc.DeletePipeline(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodePipeline_DisableStageTransition() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codepipeline.New(sess)

	params := &codepipeline.DisableStageTransitionInput{
		PipelineName:   aws.String("PipelineName"),        // Required
		Reason:         aws.String("DisabledReason"),      // Required
		StageName:      aws.String("StageName"),           // Required
		TransitionType: aws.String("StageTransitionType"), // Required
	}
	resp, err := svc.DisableStageTransition(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodePipeline_EnableStageTransition() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codepipeline.New(sess)

	params := &codepipeline.EnableStageTransitionInput{
		PipelineName:   aws.String("PipelineName"),        // Required
		StageName:      aws.String("StageName"),           // Required
		TransitionType: aws.String("StageTransitionType"), // Required
	}
	resp, err := svc.EnableStageTransition(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodePipeline_GetJobDetails() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codepipeline.New(sess)

	params := &codepipeline.GetJobDetailsInput{
		JobId: aws.String("JobId"), // Required
	}
	resp, err := svc.GetJobDetails(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodePipeline_GetPipeline() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codepipeline.New(sess)

	params := &codepipeline.GetPipelineInput{
		Name:    aws.String("PipelineName"), // Required
		Version: aws.Int64(1),
	}
	resp, err := svc.GetPipeline(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodePipeline_GetPipelineState() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codepipeline.New(sess)

	params := &codepipeline.GetPipelineStateInput{
		Name: aws.String("PipelineName"), // Required
	}
	resp, err := svc.GetPipelineState(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodePipeline_GetThirdPartyJobDetails() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codepipeline.New(sess)

	params := &codepipeline.GetThirdPartyJobDetailsInput{
		ClientToken: aws.String("ClientToken"),     // Required
		JobId:       aws.String("ThirdPartyJobId"), // Required
	}
	resp, err := svc.GetThirdPartyJobDetails(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodePipeline_ListActionTypes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codepipeline.New(sess)

	params := &codepipeline.ListActionTypesInput{
		ActionOwnerFilter: aws.String("ActionOwner"),
		NextToken:         aws.String("NextToken"),
	}
	resp, err := svc.ListActionTypes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodePipeline_ListPipelines() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codepipeline.New(sess)

	params := &codepipeline.ListPipelinesInput{
		NextToken: aws.String("NextToken"),
	}
	resp, err := svc.ListPipelines(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodePipeline_PollForJobs() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codepipeline.New(sess)

	params := &codepipeline.PollForJobsInput{
		ActionTypeId: &codepipeline.ActionTypeId{ // Required
			Category: aws.String("ActionCategory"), // Required
			Owner:    aws.String("ActionOwner"),    // Required
			Provider: aws.String("ActionProvider"), // Required
			Version:  aws.String("Version"),        // Required
		},
		MaxBatchSize: aws.Int64(1),
		QueryParam: map[string]*string{
			"Key": aws.String("ActionConfigurationQueryableValue"), // Required
			// More values...
		},
	}
	resp, err := svc.PollForJobs(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodePipeline_PollForThirdPartyJobs() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codepipeline.New(sess)

	params := &codepipeline.PollForThirdPartyJobsInput{
		ActionTypeId: &codepipeline.ActionTypeId{ // Required
			Category: aws.String("ActionCategory"), // Required
			Owner:    aws.String("ActionOwner"),    // Required
			Provider: aws.String("ActionProvider"), // Required
			Version:  aws.String("Version"),        // Required
		},
		MaxBatchSize: aws.Int64(1),
	}
	resp, err := svc.PollForThirdPartyJobs(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodePipeline_PutActionRevision() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codepipeline.New(sess)

	params := &codepipeline.PutActionRevisionInput{
		ActionName: aws.String("ActionName"), // Required
		ActionRevision: &codepipeline.ActionRevision{ // Required
			Created:          aws.Time(time.Now()),                   // Required
			RevisionChangeId: aws.String("RevisionChangeIdentifier"), // Required
			RevisionId:       aws.String("Revision"),                 // Required
		},
		PipelineName: aws.String("PipelineName"), // Required
		StageName:    aws.String("StageName"),    // Required
	}
	resp, err := svc.PutActionRevision(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodePipeline_PutApprovalResult() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codepipeline.New(sess)

	params := &codepipeline.PutApprovalResultInput{
		ActionName:   aws.String("ActionName"),   // Required
		PipelineName: aws.String("PipelineName"), // Required
		Result: &codepipeline.ApprovalResult{ // Required
			Status:  aws.String("ApprovalStatus"),  // Required
			Summary: aws.String("ApprovalSummary"), // Required
		},
		StageName: aws.String("StageName"), // Required
		Token:     aws.String("ApprovalToken"),
	}
	resp, err := svc.PutApprovalResult(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodePipeline_PutJobFailureResult() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codepipeline.New(sess)

	params := &codepipeline.PutJobFailureResultInput{
		FailureDetails: &codepipeline.FailureDetails{ // Required
			Message:             aws.String("Message"),     // Required
			Type:                aws.String("FailureType"), // Required
			ExternalExecutionId: aws.String("ExecutionId"),
		},
		JobId: aws.String("JobId"), // Required
	}
	resp, err := svc.PutJobFailureResult(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodePipeline_PutJobSuccessResult() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codepipeline.New(sess)

	params := &codepipeline.PutJobSuccessResultInput{
		JobId:             aws.String("JobId"), // Required
		ContinuationToken: aws.String("ContinuationToken"),
		CurrentRevision: &codepipeline.CurrentRevision{
			ChangeIdentifier: aws.String("RevisionChangeIdentifier"), // Required
			Revision:         aws.String("Revision"),                 // Required
		},
		ExecutionDetails: &codepipeline.ExecutionDetails{
			ExternalExecutionId: aws.String("ExecutionId"),
			PercentComplete:     aws.Int64(1),
			Summary:             aws.String("ExecutionSummary"),
		},
	}
	resp, err := svc.PutJobSuccessResult(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodePipeline_PutThirdPartyJobFailureResult() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codepipeline.New(sess)

	params := &codepipeline.PutThirdPartyJobFailureResultInput{
		ClientToken: aws.String("ClientToken"), // Required
		FailureDetails: &codepipeline.FailureDetails{ // Required
			Message:             aws.String("Message"),     // Required
			Type:                aws.String("FailureType"), // Required
			ExternalExecutionId: aws.String("ExecutionId"),
		},
		JobId: aws.String("ThirdPartyJobId"), // Required
	}
	resp, err := svc.PutThirdPartyJobFailureResult(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodePipeline_PutThirdPartyJobSuccessResult() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codepipeline.New(sess)

	params := &codepipeline.PutThirdPartyJobSuccessResultInput{
		ClientToken:       aws.String("ClientToken"),     // Required
		JobId:             aws.String("ThirdPartyJobId"), // Required
		ContinuationToken: aws.String("ContinuationToken"),
		CurrentRevision: &codepipeline.CurrentRevision{
			ChangeIdentifier: aws.String("RevisionChangeIdentifier"), // Required
			Revision:         aws.String("Revision"),                 // Required
		},
		ExecutionDetails: &codepipeline.ExecutionDetails{
			ExternalExecutionId: aws.String("ExecutionId"),
			PercentComplete:     aws.Int64(1),
			Summary:             aws.String("ExecutionSummary"),
		},
	}
	resp, err := svc.PutThirdPartyJobSuccessResult(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodePipeline_RetryStageExecution() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codepipeline.New(sess)

	params := &codepipeline.RetryStageExecutionInput{
		PipelineExecutionId: aws.String("PipelineExecutionId"), // Required
		PipelineName:        aws.String("PipelineName"),        // Required
		RetryMode:           aws.String("StageRetryMode"),      // Required
		StageName:           aws.String("StageName"),           // Required
	}
	resp, err := svc.RetryStageExecution(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodePipeline_StartPipelineExecution() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codepipeline.New(sess)

	params := &codepipeline.StartPipelineExecutionInput{
		Name: aws.String("PipelineName"), // Required
	}
	resp, err := svc.StartPipelineExecution(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodePipeline_UpdatePipeline() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codepipeline.New(sess)

	params := &codepipeline.UpdatePipelineInput{
		Pipeline: &codepipeline.PipelineDeclaration{ // Required
			ArtifactStore: &codepipeline.ArtifactStore{ // Required
				Location: aws.String("ArtifactStoreLocation"), // Required
				Type:     aws.String("ArtifactStoreType"),     // Required
				EncryptionKey: &codepipeline.EncryptionKey{
					Id:   aws.String("EncryptionKeyId"),   // Required
					Type: aws.String("EncryptionKeyType"), // Required
				},
			},
			Name:    aws.String("PipelineName"), // Required
			RoleArn: aws.String("RoleArn"),      // Required
			Stages: []*codepipeline.StageDeclaration{ // Required
				{ // Required
					Actions: []*codepipeline.ActionDeclaration{ // Required
						{ // Required
							ActionTypeId: &codepipeline.ActionTypeId{ // Required
								Category: aws.String("ActionCategory"), // Required
								Owner:    aws.String("ActionOwner"),    // Required
								Provider: aws.String("ActionProvider"), // Required
								Version:  aws.String("Version"),        // Required
							},
							Name: aws.String("ActionName"), // Required
							Configuration: map[string]*string{
								"Key": aws.String("ActionConfigurationValue"), // Required
								// More values...
							},
							InputArtifacts: []*codepipeline.InputArtifact{
								{ // Required
									Name: aws.String("ArtifactName"), // Required
								},
								// More values...
							},
							OutputArtifacts: []*codepipeline.OutputArtifact{
								{ // Required
									Name: aws.String("ArtifactName"), // Required
								},
								// More values...
							},
							RoleArn:  aws.String("RoleArn"),
							RunOrder: aws.Int64(1),
						},
						// More values...
					},
					Name: aws.String("StageName"), // Required
					Blockers: []*codepipeline.BlockerDeclaration{
						{ // Required
							Name: aws.String("BlockerName"), // Required
							Type: aws.String("BlockerType"), // Required
						},
						// More values...
					},
				},
				// More values...
			},
			Version: aws.Int64(1),
		},
	}
	resp, err := svc.UpdatePipeline(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
