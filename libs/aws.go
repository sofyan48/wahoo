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
func (aw *Aws) GetShardID(svc *kinesis.Kinesis, cfg *entity.AwsConfig) (*kinesis.GetShardIteratorOutput, error) {
	dsIter := &kinesis.GetShardIteratorInput{}
	dsIter.SetStreamName(cfg.StreamName)
	dsIter.SetShardId(cfg.ShardID)
	dsIter.SetShardIteratorType("TRIM_HORIZON")
	shardIter, err := svc.GetShardIterator(dsIter)
	return shardIter, err
}
