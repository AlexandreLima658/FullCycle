package main

import (
	"encoding/json"
	"fmt"
	"fullCycle/internal/infra/kafka"
	"fullCycle/internal/market/dto"
	"fullCycle/internal/market/entity"
	"fullCycle/internal/market/transformer"
	"sync"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	ordersIn := make(chan *entity.Order)
	ordersOut := make(chan *entity.Order)
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	kafkaMsgChan := make(chan *ckafka.Message)
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
		"group.id":          "mygroup",
		"auto.offset.reset": "latest",
	}
	producer := kafka.NewKafkaProducer(configMap)
	kafka := kafka.NewConsumer(configMap, []string{"input"})

	go kafka.Consume(kafkaMsgChan)

	book := entity.NewBook(ordersIn, ordersOut, wg)
	go book.Trade()

	go func() {
		for msg := range kafkaMsgChan {
			wg.Add(1)
			tradeInput := dto.TradeInput{}
			err := json.Unmarshal(msg.Value, &tradeInput)
			if err != nil {
				panic(err)
			}
			order := transformer.TransformerInput(tradeInput)
			fmt.Println(string(msg.Value))
			ordersIn <- order

		}
	}()

	for order := range ordersOut {
		output := transformer.TransformOutput(order)
		outputJson, err := json.MarshalIndent(output, "", "  ")
		fmt.Println(string(outputJson))
		if err != nil {
			fmt.Println(err)
		}
		producer.Publish(outputJson, []byte("orders"), "output")

	}
}
