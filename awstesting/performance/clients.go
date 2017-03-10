// +build integration

package performance

import (
	"reflect"

	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/acm"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/apigateway"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/autoscaling"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/cloudformation"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/cloudfront"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/cloudhsm"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/cloudsearch"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/cloudsearchdomain"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/cloudtrail"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/cloudwatch"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/cloudwatchevents"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/cloudwatchlogs"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/codecommit"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/codedeploy"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/codepipeline"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/cognitoidentity"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/cognitosync"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/configservice"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/datapipeline"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/devicefarm"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/directconnect"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/directoryservice"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/dynamodb"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/dynamodbstreams"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/ec2"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/ecr"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/ecs"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/efs"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/elasticache"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/elasticbeanstalk"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/elasticsearchservice"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/elastictranscoder"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/elb"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/emr"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/firehose"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/glacier"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/iam"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/inspector"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/iot"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/iotdataplane"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/kinesis"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/kms"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/lambda"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/machinelearning"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/marketplacecommerceanalytics"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/mobileanalytics"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/opsworks"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/rds"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/redshift"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/route53"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/route53domains"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/s3"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/ses"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/simpledb"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/sns"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/sqs"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/ssm"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/storagegateway"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/sts"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/support"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/swf"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/waf"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/workspaces"
)

var clients = []reflect.Value{
	reflect.ValueOf(acm.New),
	reflect.ValueOf(apigateway.New),
	reflect.ValueOf(autoscaling.New),
	reflect.ValueOf(cloudformation.New),
	reflect.ValueOf(cloudfront.New),
	reflect.ValueOf(cloudhsm.New),
	reflect.ValueOf(cloudsearch.New),
	reflect.ValueOf(cloudsearchdomain.New),
	reflect.ValueOf(cloudtrail.New),
	reflect.ValueOf(cloudwatch.New),
	reflect.ValueOf(cloudwatchevents.New),
	reflect.ValueOf(cloudwatchlogs.New),
	reflect.ValueOf(codecommit.New),
	reflect.ValueOf(codedeploy.New),
	reflect.ValueOf(codepipeline.New),
	reflect.ValueOf(cognitoidentity.New),
	reflect.ValueOf(cognitosync.New),
	reflect.ValueOf(configservice.New),
	reflect.ValueOf(datapipeline.New),
	reflect.ValueOf(devicefarm.New),
	reflect.ValueOf(directconnect.New),
	reflect.ValueOf(directoryservice.New),
	reflect.ValueOf(dynamodb.New),
	reflect.ValueOf(dynamodbstreams.New),
	reflect.ValueOf(ec2.New),
	reflect.ValueOf(ecr.New),
	reflect.ValueOf(ecs.New),
	reflect.ValueOf(efs.New),
	reflect.ValueOf(elasticache.New),
	reflect.ValueOf(elasticbeanstalk.New),
	reflect.ValueOf(elasticsearchservice.New),
	reflect.ValueOf(elastictranscoder.New),
	reflect.ValueOf(elb.New),
	reflect.ValueOf(emr.New),
	reflect.ValueOf(firehose.New),
	reflect.ValueOf(glacier.New),
	reflect.ValueOf(iam.New),
	reflect.ValueOf(inspector.New),
	reflect.ValueOf(iot.New),
	reflect.ValueOf(iotdataplane.New),
	reflect.ValueOf(kinesis.New),
	reflect.ValueOf(kms.New),
	reflect.ValueOf(lambda.New),
	reflect.ValueOf(machinelearning.New),
	reflect.ValueOf(marketplacecommerceanalytics.New),
	reflect.ValueOf(mobileanalytics.New),
	reflect.ValueOf(opsworks.New),
	reflect.ValueOf(rds.New),
	reflect.ValueOf(redshift.New),
	reflect.ValueOf(route53.New),
	reflect.ValueOf(route53domains.New),
	reflect.ValueOf(s3.New),
	reflect.ValueOf(ses.New),
	reflect.ValueOf(simpledb.New),
	reflect.ValueOf(sns.New),
	reflect.ValueOf(sqs.New),
	reflect.ValueOf(ssm.New),
	reflect.ValueOf(storagegateway.New),
	reflect.ValueOf(sts.New),
	reflect.ValueOf(support.New),
	reflect.ValueOf(swf.New),
	reflect.ValueOf(waf.New),
	reflect.ValueOf(workspaces.New),
}
