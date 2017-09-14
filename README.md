# gologger
 An asynchronous kafka logger with channel-based ring buffer in golang

## Instructions

```
go get -u github.com/trivedipankaj/gologger
```

### Usage


```
import "github.com/trivedipankaj/gologger"

var brokers   = []string{"172.16.5.10:9092", "172.16.5.11:9092", "172.16.5.12:9092"}
var topic     = "example"
var type      = "merchant"

logger    = gologger.NewLogger(topic, brokers)
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
* `log` - The log data in map format.

