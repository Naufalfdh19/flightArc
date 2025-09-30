package kafka

import (
   "context"
   "fmt"
   "github.com/segmentio/kafka-go"
)
func Consume(ctx context.Context) {
   reader := kafka.NewReader(kafka.ReaderConfig{
       Brokers: []string{"localhost:9092"},
       Topic: "example-topic",
       GroupID: "example-group",
   })
   defer reader.Close()
   for {
       msg, err := reader.ReadMessage(ctx)
       if err != nil {
           fmt.Println("Error reading message:", err)
           break
       }
       fmt.Printf("Received message: %s\n", string(msg.Value))
   }
}
