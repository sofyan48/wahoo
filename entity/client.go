package entity

import (
	"github.com/aws/aws-sdk-go/service/kinesis"
)

// NewClient ...
type NewClient struct {
	Sessions *kinesis.Kinesis
	Configs  *AwsConfig
	ShardID  *kinesis.GetShardIteratorOutput
}
