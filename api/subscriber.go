package api

import (
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/sofyan48/wahoo/entity"
	"github.com/sofyan48/wahoo/libs"
)

// Subscriber ...
type Subscriber struct {
	session   *kinesis.Kinesis
	config    *entity.AwsConfig
	awsLibs   *libs.Aws
	shardIter string
}

// NewSubscriber subs Data
// @client: *entity.NewClient
// @shardID: string
// return *Consumer
func NewSubscriber(client *entity.NewClient, shardIterator string) *Subscriber {
	subs := &Subscriber{}
	subs.config = client.Configs
	subs.session = client.Sessions
	subs.awsLibs = &libs.Aws{}
	subs.shardIter = shardIterator
	return subs
}

// GetRecord ...
func (subs *Subscriber) GetRecord() (*kinesis.GetRecordsOutput, error) {
	msgInput := subs.awsLibs.GetRecordInput()
	msgInput.SetShardIterator(subs.shardIter)
	return subs.awsLibs.GetRecord(subs.session, msgInput)
}

// GetByShard ...
func (subs *Subscriber) GetByShard(arn string) (*kinesis.SubscribeToShardOutput, error) {
	data := &kinesis.SubscribeToShardInput{}
	data.SetShardId(subs.config.ShardID)
	data.SetConsumerARN(arn)
	return subs.session.SubscribeToShard(data)
}

// Register ...
func (subs *Subscriber) Register(arn string) (*kinesis.RegisterStreamConsumerOutput, error) {
	data := &kinesis.RegisterStreamConsumerInput{}
	data.SetConsumerName(subs.config.StreamName)
	data.SetStreamARN(arn)
	return subs.session.RegisterStreamConsumer(data)
}
