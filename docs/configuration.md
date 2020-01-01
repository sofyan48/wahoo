# Configuration

## Configure
return config
```
config.Configure()
```
setting up this configuration :
- AwsAccessKeyID string
- AwsSecretAccessKey string
- APArea string
- StreamName string
- ShardID string

## Credential
return config
```
config.NewConfig().Credential(config)
```

## client 
return client session
```
creds.New()
```


## Shard Iterator
return string and error
```
config.NewConfig().GetShardIterator(client, ITER MODELS)
```
***ITER MODELS*** :
- LATEST
- TRIM_HORIZON
- AT_SEQUENCE_NUMBER
- AFTER_SEQUENCE_NUMBER
- AT_TIMESTAMP
