package main

import (
	"sync"

	l "github.com/trivedipankaj/gologger"
)

func main() {
	var brokers = []string{"172.16.10.156:9094", "172.16.10.157:9094", "172.16.10.158:9094"}
	var topic = "example1"
	var kind = "merchant"
	var bufferLength = 500

	logger := l.NewLogger(topic, brokers, bufferLength)
	logger.InitLogger()
	// encoder can be "json" and "msgpack"
	logger.SetEncoder("json")
	logger.SetCurrentLevel(7)

	log := map[string]interface{}{
		"mid":  129954,
		"gw":   8,
		"prob": 0.91,
		"mean": 0.78,
	}
	//sending 10 messages to kafka
	for i := 0; i < 10; i++ {
		logger.AsyncLog(kind, log)
	}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	wg.Wait()
}
