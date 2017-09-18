# gologger
 An asynchronous kafka logger with channel-based ring buffer in golang
 * uses https://github.com/Shopify/sarama for interacting with Kafka

## Installation

```
go get -u github.com/trivedipankaj/gologger
```

### Usage


```
import "github.com/trivedipankaj/gologger"

var brokers   = []string{"172.16.5.10:9092", "172.16.5.11:9092", "172.16.5.12:9092"}
var topic     = "example"
var type      = "merchant"
var bufferLength = 500

logger    = gologger.NewLogger(topic, brokers, bufferLength)
log 	  := map[string]interface{}{
			"mid":  129954,
			"gw":   8,
			"prob": 0.91,
			"mean": 0.78,
		}
logger.AsyncLog(type, log)

```
* `brokers` - The Kafka broker list.
* `topic` - The Kafka topic to publish to.
* `type` - The type of the log stream.
* `bufferLength` - Buffer length of the ring.
* `log` - The log data in map format.

