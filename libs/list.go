package libs

import (
	"github.com/aws/aws-sdk-go/service/kinesis"
)

// GetListSharInput ...
func (aw *Aws) GetListSharInput() *kinesis.ListShardsInput {
	return &kinesis.ListShardsInput{}
}

// GetListShard ...
func (aw *Aws) GetListShard(svc *kinesis.Kinesis, data *kinesis.ListShardsInput) (*kinesis.ListShardsOutput, error) {
	return svc.ListShards(data)
}
