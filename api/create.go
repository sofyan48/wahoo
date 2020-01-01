package api

import (
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/sofyan48/wahoo/entity"
	"github.com/sofyan48/wahoo/libs"
)

// Create ...
type Create struct {
	session    *kinesis.Kinesis
	awsLibs    *libs.Aws
	streamName string
}

// NewCreate cr Data
// @client: *entity.NewClient
// @shardID: string
// return *Consumer
func NewCreate(client *entity.NewClient) *Create {
	cr := &Create{}
	// cr.config = client.Configs
	cr.session = client.Sessions
	cr.awsLibs = &libs.Aws{}
	return cr
}

// CreateStream ...
func (cr *Create) CreateStream(streamName string, shardCount int64) (*kinesis.CreateStreamOutput, error) {
	data := &kinesis.CreateStreamInput{}
	data.SetStreamName(streamName)
	data.SetShardCount(shardCount)
	cr.streamName = streamName
	return cr.session.CreateStream(data)
}

// AddTags ...
func (cr *Create) AddTags(tags map[string]*string) (*kinesis.AddTagsToStreamOutput, error) {
	data := &kinesis.AddTagsToStreamInput{}
	data.SetStreamName(cr.streamName)
	data.SetTags(tags)
	return cr.session.AddTagsToStream(data)
}
