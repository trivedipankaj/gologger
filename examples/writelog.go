package main

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
