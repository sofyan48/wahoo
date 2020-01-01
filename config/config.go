package config

import (
	"github.com/sofyan48/wahoo/entity"
	"github.com/sofyan48/wahoo/libs"
)

// Configs ...
type Configs struct {
	AwsAccessKeyID     string
	AwsSecretAccessKey string
	AwsAPArea          string
	ShardID            string
	StreamName         string
}

// Configure call config entity
func Configure() *entity.AwsConfig {
	return &entity.AwsConfig{}
}

// NewConfig ...
func NewConfig() *Configs {
	return &Configs{}
}

// Credential client
func (cfg *Configs) Credential(awsCfg *entity.AwsConfig) *Configs {
	cfg.AwsAccessKeyID = awsCfg.AwsAccessKeyID
	cfg.AwsSecretAccessKey = awsCfg.AwsSecretAccessKey
	cfg.AwsAPArea = awsCfg.APArea
	cfg.ShardID = awsCfg.ShardID
	cfg.StreamName = awsCfg.StreamName
	return cfg
}

// New create new config
func (cfg *Configs) New() *entity.NewClient {
	clients := &entity.NewClient{}
	awsLibs := &libs.Aws{}
	awsCfg := &entity.AwsConfig{}
	awsCfg.AwsAccessKeyID = cfg.AwsAccessKeyID
	awsCfg.AwsSecretAccessKey = cfg.AwsSecretAccessKey
	awsCfg.APArea = cfg.AwsAPArea
	awsCfg.ShardID = cfg.ShardID
	awsCfg.StreamName = cfg.StreamName
	kinesis := awsLibs.Sessions(awsCfg)
	clients.Sessions = kinesis
	clients.Configs = awsCfg
	return clients
}

// GetShardID ..
func (cfg *Configs) GetShardID(client *entity.NewClient, shardType string) (string, error) {
	awsLibs := &libs.Aws{}
	awsCfg := &entity.AwsConfig{}
	awsCfg.AwsAccessKeyID = client.Configs.AwsAccessKeyID
	awsCfg.AwsSecretAccessKey = client.Configs.AwsSecretAccessKey
	awsCfg.APArea = client.Configs.APArea
	awsCfg.ShardID = client.Configs.ShardID
	awsCfg.StreamName = client.Configs.StreamName
	shard, err := awsLibs.GetShardID(client.Sessions, awsCfg, shardType)
	return *shard.ShardIterator, err
}
