# WAHOO
Data Streamer with AWS Kinesis


## GETTING STARTED
### Install
```bash
$go get github.com/sofyan48/wahoo
```

### Publisher

```golang
package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sofyan48/wahoo/api"
	"github.com/sofyan48/wahoo/config"
)

func main() {
	godotenv.Load()
	cfg := config.Configure()
	cfg.AwsAccessKeyID = os.Getenv("ACCESS_KEY")
	cfg.AwsSecretAccessKey = os.Getenv("SECRET_KEY")
	cfg.APArea = os.Getenv("REGION")
	cfg.StreamName = os.Getenv("STREAM_NAME")
	cfg.ShardID = os.Getenv("SHARD_ID")

	client := config.NewConfig().Credential(cfg).New()
	shardID, err := config.NewConfig().GetShardID(client, "TRIM_HORIZON")
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(0)
	}
	producer := api.NewProducer(client, shardID)
	data := []byte("{\"type\": \"json\",\"value\": {\"CompanyCode\": \"12345\",\"CustomerNumber\": \"ABC0012300DEF\",\"RequestID\": \"201507131507262221400000001975\",\"ChannelType\": \"6014\",\"CustomerName\": \"Customer BCA Virtual Account\",\"CurrencyCode\": \"IDR\",\"PaidAmount\": \"150000.00\",\"TotalAmount\": \"150000.00\",\"SubCompany\": \"00000\",\"TransactionDate\": \"15/03/2014 22:07:40\",\"Reference\": \"1234567890\",\"DetailBills\": [],\"FlagAdvide\": \"N\",\"Additionaldata\": \"\"}\n}")

	response, err := producer.Publish(data, shardID)
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(0)
	}
	fmt.Println(response.ShardId)
}

```