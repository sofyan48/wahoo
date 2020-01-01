package config

import (
	"github.com/sofyan48/wahoo/entity"
	"github.com/sofyan48/wahoo/libs"
)

// Configs ...
type Configs struct {
	PathURL            string
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
	cfg.PathURL = awsCfg.PathURL
	cfg.AwsAccessKeyID = awsCfg.AwsAccessKeyID
	cfg.AwsSecretAccessKey = awsCfg.AwsSecretAccessKey
	cfg.AwsAPArea = awsCfg.APArea
	awsCfg.ShardID = cfg.ShardID
	awsCfg.StreamName = cfg.StreamName
	return cfg
}

// New create new config
func (cfg *Configs) New() *entity.NewClient {
	clients := &entity.NewClient{}
	awsLibs := &libs.Aws{}
	awsCfg := &entity.AwsConfig{}
	awsCfg.PathURL = cfg.PathURL
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
