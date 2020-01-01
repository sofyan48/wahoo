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

// NewDelete dlt Data
// @client: *entity.NewClient
// @shardID: string
// return *Consumer
func NewDelete(client *entity.NewClient) *Delete {
	dlt := &Delete{}
	// dlt.config = client.Configs
	dlt.session = client.Sessions
	dlt.awsLibs = &libs.Aws{}
	return dlt
}

// DeleteStream ...
func (dlt *Delete) DeleteStream(streamName string) (*kinesis.DeleteStreamOutput, error) {
	data := &kinesis.DeleteStreamInput{}
	data.SetStreamName(streamName)
	dlt.streamName = streamName
	return dlt.session.DeleteStream(data)
}

// DeleteStreamConsumer ...
func (dlt *Delete) DeleteStreamConsumer(arn, csmArn, name string) (*kinesis.DeregisterStreamConsumerOutput, error) {
	data := &kinesis.DeregisterStreamConsumerInput{}
	data.SetConsumerARN(csmArn)
	data.SetStreamARN(arn)
	data.SetConsumerName(name)
	return dlt.session.DeregisterStreamConsumer(data)
}
