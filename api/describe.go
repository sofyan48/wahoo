package api

import (
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/sofyan48/wahoo/entity"
	"github.com/sofyan48/wahoo/libs"
)

// Describe ...
type Describe struct {
	session   *kinesis.Kinesis
	config    *entity.AwsConfig
	awsLibs   *libs.Aws
	shardIter string
}

// NewDescribe pubs Data
// @client: *entity.NewClient
// @shardID: string
// return *Consumer
func NewDescribe(client *entity.NewClient, shardIter string) *Describe {
	desc := &Describe{}
	desc.config = client.Configs
	desc.session = client.Sessions
	desc.awsLibs = &libs.Aws{}
	desc.shardIter = shardIter
	return desc
}

// Describe ...
func (dsc *Describe) Describe() (*kinesis.DescribeStreamOutput, error) {
	dscInput := dsc.awsLibs.GetDescribeInput()
	dscInput.SetStreamName(dsc.config.StreamName)
	return dsc.awsLibs.Describe(dsc.session, dscInput)
}
