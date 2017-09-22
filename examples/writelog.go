package main

import (
	"sync"

	l "github.com/trivedipankaj/gologger"
)

func main() {
	var brokers = []string{"10.50.21.117:9094", "10.50.21.118:9094", "10.50.21.119:9094"}
	var topic = "test"
	var kind = "merchant"
	var bufferLength = 500

	logger := l.NewLogger(topic, brokers, bufferLength)
	logger.InitLogger()
	log := map[string]interface{}{
		"mid":  129954,
		"gw":   8,
		"prob": 0.91,
		"mean": 0.78,
	}
	for i := 0; i < 10; i++ {
		logger.AsyncLog(kind, log)
	}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	wg.Wait()
}
