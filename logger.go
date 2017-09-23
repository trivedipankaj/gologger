package logger

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/vmihailenco/msgpack"
)

type Logger struct {
	rb      *Ring
	topic   string
	encoder string
	level   int64
	brokers []string
}

func NewLogger(topic string, brokers []string, bufferLength int) *Logger {
	in := make(chan string)
	out := make(chan string, bufferLength)
	return &Logger{topic: topic,
		brokers: brokers,
		level:   INFO,
		encoder: "msgpack",
		rb:      NewRing(in, out)}
}

func (l *Logger) SetTopic(topic string) {
	l.topic = topic
}

func (l *Logger) SetEncoder(encoder string) {
	l.encoder = encoder
}

func (l *Logger) SetCurrentLevel(level int64) {
	currentLevel = level
}

func (l *Logger) KafkaProducer() (sarama.AsyncProducer, error) {
	//logger := log.New(os.Stdout, "logger: ", log.Lshortfile)
	//sarama.Logger = logger
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Compression = sarama.CompressionNone
	//config.Producer.Return.Successes = true
	//config.Producer.Return.Errors = true
	var err error
	producer, err := sarama.NewAsyncProducer(l.brokers, config)
	if err != nil {
		return nil, err
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	go func() {
		<-c
		if err := producer.Close(); err != nil {
			log.Fatal("Error closing async producer", err)
		}
		log.Println("Async Producer closed")
		os.Exit(1)
	}()

	go func() {
		for err := range producer.Errors() {
			log.Println("Failed to write message to topic:", err)
		}
	}()

	return producer, nil
}

func (l *Logger) produce(c chan<- *sarama.ProducerMessage, msg string) {
	message := &sarama.ProducerMessage{
		Topic: l.topic,
		Value: sarama.ByteEncoder(msg),
	}
	c <- message
}

func (l *Logger) Consume(p sarama.AsyncProducer) {
	for res := range l.rb.oc {
		l.produce(p.Input(), res)
	}
}

func (l *Logger) InitLogger() {
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		defer wg.Done()
		l.rb.Run()
	}()
	p, _ := l.KafkaProducer()

	wg.Add(1)
	go func() {
		defer wg.Done()
		l.Consume(p)
	}()
}

func (l *Logger) AsyncLog(kind string, msg map[string]interface{}) (P *Logger) {
	if currentLevel < l.level {
		b, err := Encoded(l.encoder, NewLog(kind, msg).SetLevel(l.level))
		if err != nil {
			panic(err)
		}
		go l.rb.Produce(string(b))
	}
	return
}

func (l *Logger) Info() *Logger {
	l.level = INFO
	return l
}

func (l *Logger) Debug() *Logger {
	l.level = DEBUG
	return l
}

func (l *Logger) Warn() *Logger {
	l.level = WARN
	return l
}

func (l *Logger) Error() *Logger {
	l.level = ERROR
	return l
}

func (l *Logger) Fatal() *Logger {
	l.level = FATAL
	return l
}

func Encoded(encoder string, log *Log) ([]byte, error) {
	var (
		err error
		b   []byte
	)

	switch encoder {
	case "msgpack":
		b, err = msgpack.Marshal(log)
	case "json":
		b, err = json.Marshal(log)
	default:
	}

	return b, err
}
