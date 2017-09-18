# gologger
 An asynchronous kafka logger with channel-based ring buffer in golang
 * uses https://github.com/Shopify/sarama for interacting with Kafka
 * uses msgpack to encode logs to send it over kafka
## Installation

```
go get -u github.com/trivedipankaj/gologger
```

### Usage


```
import l "github.com/trivedipankaj/gologger"

func main() {
    var brokers = []string{"172.16.12.146:9092", "172.16.12.147:9092", "172.16.12.148:9092"}
    var topic = "example"
    var kind = "merchant"
    var bufferLength = 500

    logger := l.NewLogger(topic, brokers, bufferLength)
    log := map[string]interface{}{
        "mid":  129954,
        "gw":   8,
        "prob": 0.91,
        "mean": 0.78,
    }
    logger.AsyncLog(kind, log)
}
```
* `brokers` - The Kafka broker list.
* `topic` - The Kafka topic to publish to.
* `type` - The type of the log stream.
* `bufferLength` - Buffer length of the ring.
* `log` - The log data in map format.

