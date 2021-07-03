package sqsclient

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)


var (
	Client *sqs.SQS
)


func InitSqs()  {
	creds := credentials.NewStaticCredentials("key",
		"secret", "")
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: creds,
	}))
	sqsClient := sqs.New(sess)
	Client = sqsClient
}