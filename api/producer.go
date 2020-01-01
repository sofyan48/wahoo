package api

import (
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/sofyan48/wahoo/entity"
	"github.com/sofyan48/wahoo/libs"
)

// Producer ...
type Producer struct {
	session *kinesis.Kinesis
	config  *entity.AwsConfig
	awsLibs *libs.Aws
	shardID string
	// MsgOutput *sqs.ReceiveMessageOutput
	// Worker    int
}

// NewProducer pubs Data
// @client: *entity.NewClient
// @shardID: string
// return *Consumer
func NewProducer(client *entity.NewClient, shardID string) *Producer {
	pubs := &Producer{}
	pubs.config = client.Configs
	pubs.session = client.Sessions
	pubs.awsLibs = &libs.Aws{}
	pubs.shardID = shardID
	return pubs
}

// Publish ..
func (pubs *Producer) Publish(data []byte, partitionKey string) (*kinesis.PutRecordOutput, error) {
	msgInput := pubs.awsLibs.GetMessagesInput()
	msgInput.SetStreamName(pubs.config.StreamName)
	msgInput.SetPartitionKey(partitionKey)
	msgInput.SetData(data)
	resp, err := pubs.awsLibs.Send(pubs.session, msgInput)
	return resp, err
}
