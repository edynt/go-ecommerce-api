package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

var (
	kafkaProducer *kafka.Writer
)

const (
	kafkaURL   = "localhost:9092"
	kafkaTopic = "user_topic_001ww"
)

// for producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// for consumer
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers, // []string{"localhost:9092"},
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3,              // 10KB
		MaxBytes:       10e6,              // 10MB
		CommitInterval: time.Second,       // time commit after 1s
		StartOffset:    kafka.FirstOffset, //co nghia la dat gia tri offset ban dau khi consumer bat dau lang nghe
	})
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

// mua ban chung khoan
func newStock(msg, typeMsg string) *StockInfo {
	s := StockInfo{}
	s.Message = msg
	s.Type = typeMsg
	return &s
}

//

func actionStock(c *gin.Context) {
	s := newStock(c.Query("msg"), c.Query("type"))
	body := make(map[string]interface{})
	body["action"] = "action"
	body["info"] = s

	jsonBody, _ := json.Marshal(body)

	// tao message trong kafka
	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonBody),
	}

	// viet message by  producer
	err := kafkaProducer.WriteMessages(context.Background(), msg)
	if err != nil {
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"err":     "",
		"message": "success",
	})
}

// consumer hong mua ATC
func RegisterConsumerATC(id int) {
	kafkaGroupId := "consumer-group"
	reader := getKafkaReader(kafkaURL, kafkaTopic, kafkaGroupId)

	// giai phong ket noi
	defer reader.Close()

	fmt.Printf("Consumer(%d) Hong Phien ATC::", id)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Consumer(%d) error: %v\n", id, err)
			continue
		}

		fmt.Printf("Consumer(%d), hong topic: %v, partition: %v, offset: %v, time:%d %s = %s\n", id, m.Topic, m.Partition, m.Offset, m.Time.Unix(), m.Key, string(m.Value))

	}
}

func main() {
	r := gin.Default()
	kafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic)
	defer kafkaProducer.Close()

	r.POST("/action/stock", actionStock)

	// dang ky 2 user de mua stock trong 2 atc
	go RegisterConsumerATC(1)
	go RegisterConsumerATC(2)

	r.Run()
}
