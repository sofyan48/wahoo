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
}

// NewProducer pubs Data
// @client: *entity.NewClient
// @shardID: string
// return *Consumer
func NewProducer(client *entity.NewClient) *Producer {
	pubs := &Producer{}
	pubs.config = client.Configs
	pubs.session = client.Sessions
	pubs.awsLibs = &libs.Aws{}
	return pubs
}

// Publish ..
func (pubs *Producer) Publish(data []byte, partitionKey string) (*kinesis.PutRecordOutput, error) {
	msgInput := pubs.awsLibs.GetMessagesInput()
	msgInput.SetStreamName(pubs.config.StreamName)
	msgInput.SetPartitionKey(partitionKey)
	msgInput.SetData(data)
	return pubs.awsLibs.Send(pubs.session, msgInput)
}

// ProducerAddTags ...
func (pubs *Producer) ProducerAddTags(tags map[string]*string) (*kinesis.AddTagsToStreamOutput, error) {
	data := &kinesis.AddTagsToStreamInput{}
	data.SetStreamName(pubs.config.StreamName)
	data.SetTags(tags)
	return pubs.session.AddTagsToStream(data)
}
