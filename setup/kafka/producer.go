package kafka

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)


func Produce(ctx context.Context) {
   writer := kafka.NewWriter(kafka.WriterConfig{
       Brokers: []string{"localhost:9092"},
       Topic: "example-topic",
   })
   defer writer.Close()
   for i := 0; i < 10; i++ {
       err := writer.WriteMessages(ctx, kafka.Message{
           Key: []byte(strconv.Itoa(i)),
           Value: []byte("Message " + strconv.Itoa(i)),
       })
       if err != nil {
           fmt.Println("Error writing message:", err)
       } else {
           fmt.Println("Message written:", i)
       }
       time.Sleep(1 * time.Second)
   }
}
