package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/satori/go.uuid"
	"logrus-demo/sqsclient"
	"time"
)

var Url = "queurUrl"

func main() {
	sqsclient.InitSqs()
	c := make(chan struct{})
	go func() {
		for {
			t := time.Now().Format("2006-01-02 15:04:05")
			uuid := uuid.NewV4().String()
			result, err := sqsclient.Client.SendMessage(&sqs.SendMessageInput{
				MessageGroupId:         aws.String("gwell"),
				MessageDeduplicationId: aws.String(uuid),
				DelaySeconds:           aws.Int64(0),
				MessageBody:            aws.String(t),
				QueueUrl:               &Url,
			})
			if err != nil {
				ErrLogger.Error(err)
				continue
			}
			msgId := result.MessageId
			ErrLogger.Errorf("发送消息 msgId:%s", *msgId)
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			recResult, err := sqsclient.Client.ReceiveMessage(&sqs.ReceiveMessageInput{
				QueueUrl:            &Url,
				MaxNumberOfMessages: aws.Int64(10),
				WaitTimeSeconds:     aws.Int64(1),
				VisibilityTimeout:   aws.Int64(5), // 如果這個 message 沒刪除，下次再被取出來的時間
			})
			if err != nil {
				ErrLogger.Error("消费消息失败")
				continue
			}
			for _, m := range recResult.Messages {
				ErrLogger.Errorf("消费消息 msgId:%v 消息内容:%s", *m.MessageId, *m.Body)
			}

		}
	}()

	<-c
}
