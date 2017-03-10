// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package machinelearningiface provides an interface for the Amazon Machine Learning.
package machinelearningiface

import (
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/aws/request"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/machinelearning"
)

// MachineLearningAPI is the interface type for machinelearning.MachineLearning.
type MachineLearningAPI interface {
	AddTagsRequest(*machinelearning.AddTagsInput) (*request.Request, *machinelearning.AddTagsOutput)

	AddTags(*machinelearning.AddTagsInput) (*machinelearning.AddTagsOutput, error)

	CreateBatchPredictionRequest(*machinelearning.CreateBatchPredictionInput) (*request.Request, *machinelearning.CreateBatchPredictionOutput)

	CreateBatchPrediction(*machinelearning.CreateBatchPredictionInput) (*machinelearning.CreateBatchPredictionOutput, error)

	CreateDataSourceFromRDSRequest(*machinelearning.CreateDataSourceFromRDSInput) (*request.Request, *machinelearning.CreateDataSourceFromRDSOutput)

	CreateDataSourceFromRDS(*machinelearning.CreateDataSourceFromRDSInput) (*machinelearning.CreateDataSourceFromRDSOutput, error)

	CreateDataSourceFromRedshiftRequest(*machinelearning.CreateDataSourceFromRedshiftInput) (*request.Request, *machinelearning.CreateDataSourceFromRedshiftOutput)

	CreateDataSourceFromRedshift(*machinelearning.CreateDataSourceFromRedshiftInput) (*machinelearning.CreateDataSourceFromRedshiftOutput, error)

	CreateDataSourceFromS3Request(*machinelearning.CreateDataSourceFromS3Input) (*request.Request, *machinelearning.CreateDataSourceFromS3Output)

	CreateDataSourceFromS3(*machinelearning.CreateDataSourceFromS3Input) (*machinelearning.CreateDataSourceFromS3Output, error)

	CreateEvaluationRequest(*machinelearning.CreateEvaluationInput) (*request.Request, *machinelearning.CreateEvaluationOutput)

	CreateEvaluation(*machinelearning.CreateEvaluationInput) (*machinelearning.CreateEvaluationOutput, error)

	CreateMLModelRequest(*machinelearning.CreateMLModelInput) (*request.Request, *machinelearning.CreateMLModelOutput)

	CreateMLModel(*machinelearning.CreateMLModelInput) (*machinelearning.CreateMLModelOutput, error)

	CreateRealtimeEndpointRequest(*machinelearning.CreateRealtimeEndpointInput) (*request.Request, *machinelearning.CreateRealtimeEndpointOutput)

	CreateRealtimeEndpoint(*machinelearning.CreateRealtimeEndpointInput) (*machinelearning.CreateRealtimeEndpointOutput, error)

	DeleteBatchPredictionRequest(*machinelearning.DeleteBatchPredictionInput) (*request.Request, *machinelearning.DeleteBatchPredictionOutput)

	DeleteBatchPrediction(*machinelearning.DeleteBatchPredictionInput) (*machinelearning.DeleteBatchPredictionOutput, error)

	DeleteDataSourceRequest(*machinelearning.DeleteDataSourceInput) (*request.Request, *machinelearning.DeleteDataSourceOutput)

	DeleteDataSource(*machinelearning.DeleteDataSourceInput) (*machinelearning.DeleteDataSourceOutput, error)

	DeleteEvaluationRequest(*machinelearning.DeleteEvaluationInput) (*request.Request, *machinelearning.DeleteEvaluationOutput)

	DeleteEvaluation(*machinelearning.DeleteEvaluationInput) (*machinelearning.DeleteEvaluationOutput, error)

	DeleteMLModelRequest(*machinelearning.DeleteMLModelInput) (*request.Request, *machinelearning.DeleteMLModelOutput)

	DeleteMLModel(*machinelearning.DeleteMLModelInput) (*machinelearning.DeleteMLModelOutput, error)

	DeleteRealtimeEndpointRequest(*machinelearning.DeleteRealtimeEndpointInput) (*request.Request, *machinelearning.DeleteRealtimeEndpointOutput)

	DeleteRealtimeEndpoint(*machinelearning.DeleteRealtimeEndpointInput) (*machinelearning.DeleteRealtimeEndpointOutput, error)

	DeleteTagsRequest(*machinelearning.DeleteTagsInput) (*request.Request, *machinelearning.DeleteTagsOutput)

	DeleteTags(*machinelearning.DeleteTagsInput) (*machinelearning.DeleteTagsOutput, error)

	DescribeBatchPredictionsRequest(*machinelearning.DescribeBatchPredictionsInput) (*request.Request, *machinelearning.DescribeBatchPredictionsOutput)

	DescribeBatchPredictions(*machinelearning.DescribeBatchPredictionsInput) (*machinelearning.DescribeBatchPredictionsOutput, error)

	DescribeBatchPredictionsPages(*machinelearning.DescribeBatchPredictionsInput, func(*machinelearning.DescribeBatchPredictionsOutput, bool) bool) error

	DescribeDataSourcesRequest(*machinelearning.DescribeDataSourcesInput) (*request.Request, *machinelearning.DescribeDataSourcesOutput)

	DescribeDataSources(*machinelearning.DescribeDataSourcesInput) (*machinelearning.DescribeDataSourcesOutput, error)

	DescribeDataSourcesPages(*machinelearning.DescribeDataSourcesInput, func(*machinelearning.DescribeDataSourcesOutput, bool) bool) error

	DescribeEvaluationsRequest(*machinelearning.DescribeEvaluationsInput) (*request.Request, *machinelearning.DescribeEvaluationsOutput)

	DescribeEvaluations(*machinelearning.DescribeEvaluationsInput) (*machinelearning.DescribeEvaluationsOutput, error)

	DescribeEvaluationsPages(*machinelearning.DescribeEvaluationsInput, func(*machinelearning.DescribeEvaluationsOutput, bool) bool) error

	DescribeMLModelsRequest(*machinelearning.DescribeMLModelsInput) (*request.Request, *machinelearning.DescribeMLModelsOutput)

	DescribeMLModels(*machinelearning.DescribeMLModelsInput) (*machinelearning.DescribeMLModelsOutput, error)

	DescribeMLModelsPages(*machinelearning.DescribeMLModelsInput, func(*machinelearning.DescribeMLModelsOutput, bool) bool) error

	DescribeTagsRequest(*machinelearning.DescribeTagsInput) (*request.Request, *machinelearning.DescribeTagsOutput)

	DescribeTags(*machinelearning.DescribeTagsInput) (*machinelearning.DescribeTagsOutput, error)

	GetBatchPredictionRequest(*machinelearning.GetBatchPredictionInput) (*request.Request, *machinelearning.GetBatchPredictionOutput)

	GetBatchPrediction(*machinelearning.GetBatchPredictionInput) (*machinelearning.GetBatchPredictionOutput, error)

	GetDataSourceRequest(*machinelearning.GetDataSourceInput) (*request.Request, *machinelearning.GetDataSourceOutput)

	GetDataSource(*machinelearning.GetDataSourceInput) (*machinelearning.GetDataSourceOutput, error)

	GetEvaluationRequest(*machinelearning.GetEvaluationInput) (*request.Request, *machinelearning.GetEvaluationOutput)

	GetEvaluation(*machinelearning.GetEvaluationInput) (*machinelearning.GetEvaluationOutput, error)

	GetMLModelRequest(*machinelearning.GetMLModelInput) (*request.Request, *machinelearning.GetMLModelOutput)

	GetMLModel(*machinelearning.GetMLModelInput) (*machinelearning.GetMLModelOutput, error)

	PredictRequest(*machinelearning.PredictInput) (*request.Request, *machinelearning.PredictOutput)

	Predict(*machinelearning.PredictInput) (*machinelearning.PredictOutput, error)

	UpdateBatchPredictionRequest(*machinelearning.UpdateBatchPredictionInput) (*request.Request, *machinelearning.UpdateBatchPredictionOutput)

	UpdateBatchPrediction(*machinelearning.UpdateBatchPredictionInput) (*machinelearning.UpdateBatchPredictionOutput, error)

	UpdateDataSourceRequest(*machinelearning.UpdateDataSourceInput) (*request.Request, *machinelearning.UpdateDataSourceOutput)

	UpdateDataSource(*machinelearning.UpdateDataSourceInput) (*machinelearning.UpdateDataSourceOutput, error)

	UpdateEvaluationRequest(*machinelearning.UpdateEvaluationInput) (*request.Request, *machinelearning.UpdateEvaluationOutput)

	UpdateEvaluation(*machinelearning.UpdateEvaluationInput) (*machinelearning.UpdateEvaluationOutput, error)

	UpdateMLModelRequest(*machinelearning.UpdateMLModelInput) (*request.Request, *machinelearning.UpdateMLModelOutput)

	UpdateMLModel(*machinelearning.UpdateMLModelInput) (*machinelearning.UpdateMLModelOutput, error)
}

var _ MachineLearningAPI = (*machinelearning.MachineLearning)(nil)
