package api

import (
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/sofyan48/wahoo/entity"
	"github.com/sofyan48/wahoo/libs"
)

// Delete ...
type Delete struct {
	session    *kinesis.Kinesis
	awsLibs    *libs.Aws
	streamName string
}

// NewDelete cr Data
// @client: *entity.NewClient
// @shardID: string
// return *Consumer
func NewDelete(client *entity.NewClient) *Delete {
	cr := &Delete{}
	// cr.config = client.Configs
	cr.session = client.Sessions
	cr.awsLibs = &libs.Aws{}
	return cr
}

// DeleteStream ...
func (cr *Delete) DeleteStream(streamName string) (*kinesis.DeleteStreamOutput, error) {
	data := &kinesis.DeleteStreamInput{}
	data.SetStreamName(streamName)
	cr.streamName = streamName
	return cr.session.DeleteStream(data)
}
