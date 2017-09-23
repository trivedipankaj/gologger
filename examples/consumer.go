package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/vmihailenco/msgpack"
)

type Log struct {
	UID       string
	Type      string
	Level     string
	Host      string
	Timestamp string
	Data      map[string]interface{}
}

func KafkaMessages(servers []string, topic string, offset int64) <-chan interface{} {
	out := make(chan interface{})

	go func() {
		defer close(out)

		consumer := MustSucceed(sarama.NewConsumer(servers, nil)).(sarama.Consumer)
		defer consumer.Close()

		partitions := MustSucceed(consumer.Partitions(topic)).([]int32)
		wg := new(sync.WaitGroup)
		defer wg.Wait()

		for _, p := range partitions {
			wg.Add(1)

			go func() {
				defer wg.Done()

				pc := MustSucceed(consumer.ConsumePartition(topic, p, offset)).(sarama.PartitionConsumer)
				defer pc.Close()

				for msg := range pc.Messages() {
					if true {
						fmt.Println(string(msg.Value))
					}
					out <- string(msg.Value)
				}
			}()
		}
	}()

	return out
}

func MustSucceed(args ...interface{}) interface{} {
	if len(args) == 0 {
		return nil
	}

	if err := args[len(args)-1]; err != nil {
		log.Fatal(err)
	}

	if len(args) > 1 {
		return args[0]
	}

	return nil
}

func psink(in <-chan interface{}) {
	wg := new(sync.WaitGroup)
	defer wg.Wait()

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for msg := range in {
				fmt.Println("here", msg)
				//log := parseLog(msg)
			}
		}()
	}
}

func parseLog(msg interface{}) Log {
	k, _ := msg.([]byte)
	var item Log
	msgpack.Unmarshal(k, &item)
	return item
}

func main() {

	// initialize with empty data or an actual slice of floats

	brokers := []string{"10.50.21.117:9094", "10.50.21.118:9094", "10.50.21.119:9094"}
	topic := "test"

	wg := new(sync.WaitGroup)
	defer wg.Wait()

	wg.Add(1)

	go func() {
		defer wg.Done()
		psink(KafkaMessages(brokers, topic, int64(-1)))
	}()

}
