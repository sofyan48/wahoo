# Configuration

## Configure
return config
```
config.Configure()
```
Object 
```
AwsAccessKeyID string
AwsSecretAccessKey string
APArea = os.Getenv("REGION")
StreamName = os.Getenv("STREAM_NAME")
ShardID = os.Getenv("SHARD_ID")

```

## Credential
return config
```
config.NewConfig().Credential(cfg)
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
