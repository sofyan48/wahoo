package api

import (
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/sofyan48/wahoo/entity"
	"github.com/sofyan48/wahoo/libs"
)

// List ...
type List struct {
	session   *kinesis.Kinesis
	config    *entity.AwsConfig
	awsLibs   *libs.Aws
	shardIter string
}

// NewList pubs Data
// @client: *entity.NewClient
// @shardID: string
// return *Consumer
func NewList(client *entity.NewClient, shardIter string) *List {
	desc := &List{}
	desc.config = client.Configs
	desc.session = client.Sessions
	desc.awsLibs = &libs.Aws{}
	desc.shardIter = shardIter
	return desc
}

// GetListShard ...
func (lst *List) GetListShard() (*kinesis.ListShardsOutput, error) {
	data := lst.awsLibs.GetListSharInput()
	data.SetStreamName(lst.config.StreamName)
	return lst.awsLibs.GetListShard(lst.session, data)
}

// GetListStreamConsumer ...
func (lst *List) GetListStreamConsumer(arn string, limit int64) (*kinesis.ListStreamConsumersOutput, error) {
	data := &kinesis.ListStreamConsumersInput{}
	data.SetMaxResults(limit)
	data.SetStreamARN(arn)
	return lst.session.ListStreamConsumers(data)
}

// GetListStream ...
func (lst *List) GetListStream(limit int64) (*kinesis.ListStreamsOutput, error) {
	data := &kinesis.ListStreamsInput{}
	data.SetLimit(limit)
	return lst.session.ListStreams(data)
}

// GetListTagsForStream ...
func (lst *List) GetListTagsForStream(limit int64) (*kinesis.ListTagsForStreamOutput, error) {
	data := &kinesis.ListTagsForStreamInput{}
	data.SetStreamName(lst.config.StreamName)
	data.SetLimit(limit)
	return lst.session.ListTagsForStream(data)
}
