# gologger
 An asynchronous kafka logger with channel-based ring buffer

## Instructions

```
go get -u github.com/trivedipankaj/gologger
```

### Usage


```
import "github.com/trivedipankaj/gologger"

var brokers   = []string{"172.16.12.146:9092", "172.16.12.147:9092", "172.16.12.148:9092"}
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

