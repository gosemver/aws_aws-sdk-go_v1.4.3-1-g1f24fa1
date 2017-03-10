// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package codecommit_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/session"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/codecommit"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCodeCommit_BatchGetRepositories() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codecommit.New(sess)

	params := &codecommit.BatchGetRepositoriesInput{
		RepositoryNames: []*string{ // Required
			aws.String("RepositoryName"), // Required
			// More values...
		},
	}
	resp, err := svc.BatchGetRepositories(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_CreateBranch() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codecommit.New(sess)

	params := &codecommit.CreateBranchInput{
		BranchName:     aws.String("BranchName"),     // Required
		CommitId:       aws.String("CommitId"),       // Required
		RepositoryName: aws.String("RepositoryName"), // Required
	}
	resp, err := svc.CreateBranch(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_CreateRepository() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codecommit.New(sess)

	params := &codecommit.CreateRepositoryInput{
		RepositoryName:        aws.String("RepositoryName"), // Required
		RepositoryDescription: aws.String("RepositoryDescription"),
	}
	resp, err := svc.CreateRepository(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_DeleteRepository() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codecommit.New(sess)

	params := &codecommit.DeleteRepositoryInput{
		RepositoryName: aws.String("RepositoryName"), // Required
	}
	resp, err := svc.DeleteRepository(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_GetBranch() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codecommit.New(sess)

	params := &codecommit.GetBranchInput{
		BranchName:     aws.String("BranchName"),
		RepositoryName: aws.String("RepositoryName"),
	}
	resp, err := svc.GetBranch(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_GetCommit() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codecommit.New(sess)

	params := &codecommit.GetCommitInput{
		CommitId:       aws.String("ObjectId"),       // Required
		RepositoryName: aws.String("RepositoryName"), // Required
	}
	resp, err := svc.GetCommit(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_GetRepository() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codecommit.New(sess)

	params := &codecommit.GetRepositoryInput{
		RepositoryName: aws.String("RepositoryName"), // Required
	}
	resp, err := svc.GetRepository(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_GetRepositoryTriggers() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codecommit.New(sess)

	params := &codecommit.GetRepositoryTriggersInput{
		RepositoryName: aws.String("RepositoryName"),
	}
	resp, err := svc.GetRepositoryTriggers(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_ListBranches() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codecommit.New(sess)

	params := &codecommit.ListBranchesInput{
		RepositoryName: aws.String("RepositoryName"), // Required
		NextToken:      aws.String("NextToken"),
	}
	resp, err := svc.ListBranches(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_ListRepositories() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codecommit.New(sess)

	params := &codecommit.ListRepositoriesInput{
		NextToken: aws.String("NextToken"),
		Order:     aws.String("OrderEnum"),
		SortBy:    aws.String("SortByEnum"),
	}
	resp, err := svc.ListRepositories(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_PutRepositoryTriggers() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codecommit.New(sess)

	params := &codecommit.PutRepositoryTriggersInput{
		RepositoryName: aws.String("RepositoryName"),
		Triggers: []*codecommit.RepositoryTrigger{
			{ // Required
				Branches: []*string{
					aws.String("BranchName"), // Required
					// More values...
				},
				CustomData:     aws.String("RepositoryTriggerCustomData"),
				DestinationArn: aws.String("Arn"),
				Events: []*string{
					aws.String("RepositoryTriggerEventEnum"), // Required
					// More values...
				},
				Name: aws.String("RepositoryTriggerName"),
			},
			// More values...
		},
	}
	resp, err := svc.PutRepositoryTriggers(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_TestRepositoryTriggers() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codecommit.New(sess)

	params := &codecommit.TestRepositoryTriggersInput{
		RepositoryName: aws.String("RepositoryName"),
		Triggers: []*codecommit.RepositoryTrigger{
			{ // Required
				Branches: []*string{
					aws.String("BranchName"), // Required
					// More values...
				},
				CustomData:     aws.String("RepositoryTriggerCustomData"),
				DestinationArn: aws.String("Arn"),
				Events: []*string{
					aws.String("RepositoryTriggerEventEnum"), // Required
					// More values...
				},
				Name: aws.String("RepositoryTriggerName"),
			},
			// More values...
		},
	}
	resp, err := svc.TestRepositoryTriggers(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_UpdateDefaultBranch() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codecommit.New(sess)

	params := &codecommit.UpdateDefaultBranchInput{
		DefaultBranchName: aws.String("BranchName"),     // Required
		RepositoryName:    aws.String("RepositoryName"), // Required
	}
	resp, err := svc.UpdateDefaultBranch(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_UpdateRepositoryDescription() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codecommit.New(sess)

	params := &codecommit.UpdateRepositoryDescriptionInput{
		RepositoryName:        aws.String("RepositoryName"), // Required
		RepositoryDescription: aws.String("RepositoryDescription"),
	}
	resp, err := svc.UpdateRepositoryDescription(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCodeCommit_UpdateRepositoryName() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := codecommit.New(sess)

	params := &codecommit.UpdateRepositoryNameInput{
		NewName: aws.String("RepositoryName"), // Required
		OldName: aws.String("RepositoryName"), // Required
	}
	resp, err := svc.UpdateRepositoryName(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
