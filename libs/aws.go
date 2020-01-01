package libs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/sofyan48/wahoo/entity"
)

// Sessions get session kinesis
// @cfg: *entity.AwsConfig
// return *kinesis.Kinesis
func (aw *Aws) Sessions(cfg *entity.AwsConfig) *kinesis.Kinesis {
	creds := credentials.NewStaticCredentials(cfg.AwsAccessKeyID, cfg.AwsSecretAccessKey, "")
	creds.Get()
	cfgAws := aws.NewConfig().WithRegion(cfg.APArea).WithCredentials(creds)
	svc := kinesis.New(session.New(), cfgAws)
	return svc
}

// GetShardID global Shard ID
// @cfg: *entity.AwsConfig
// return *kinesis.GetShardIteratorOutput
func (aw *Aws) GetShardIterator(svc *kinesis.Kinesis, cfg *entity.AwsConfig, shardType string) (*kinesis.GetShardIteratorOutput, error) {
	dsIter := &kinesis.GetShardIteratorInput{}
	dsIter.SetStreamName(cfg.StreamName)
	dsIter.SetShardId(cfg.ShardID)
	dsIter.SetShardIteratorType(shardType)
	shardIter, err := svc.GetShardIterator(dsIter)
	return shardIter, err
}

// GetMessagesInput ...
func (aw *Aws) GetMessagesInput() *kinesis.PutRecordInput {
	return &kinesis.PutRecordInput{}
}

// Send ...
func (aw *Aws) Send(svc *kinesis.Kinesis, data *kinesis.PutRecordInput) (*kinesis.PutRecordOutput, error) {
	req, err := svc.PutRecord(data)
	return req, err
}

// GetRecordInput ...
func (aw *Aws) GetRecordInput() *kinesis.GetRecordsInput {
	return &kinesis.GetRecordsInput{}
}

// GetRecord ...
func (aw *Aws) GetRecord(svc *kinesis.Kinesis, data *kinesis.GetRecordsInput) (*kinesis.GetRecordsOutput, error) {
	records, err := svc.GetRecords(data)
	return records, err
}

// GetDescribeInput ...
func (aw *Aws) GetDescribeInput() *kinesis.DescribeStreamInput {
	return &kinesis.DescribeStreamInput{}
}

// Describe ...
func (aw *Aws) Describe(svc *kinesis.Kinesis, data *kinesis.DescribeStreamInput) (*kinesis.DescribeStreamOutput, error) {
	descData, err := svc.DescribeStream(data)
	return descData, err
}
