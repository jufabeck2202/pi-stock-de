package messaging

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adjust/rmq/v4"

	"github.com/jufabeck2202/piScraper/messaging/consumer"
	"github.com/jufabeck2202/piScraper/messaging/types"
	"github.com/jufabeck2202/piScraper/storage"
)

const (
	prefetchLimit = 1000
	pollDuration  = 100 * time.Millisecond
	numConsumers  = 5
)

var queue rmq.Queue

func Init() {
	errChan := make(chan error, 10)
	go logErrors(errChan)
	client, err := storage.GetRedisConnection()
	if err != nil {
		log.Fatal(err)
	}
	connection, err := rmq.OpenConnectionWithRedisClient("service", client, errChan)
	if err != nil {
		panic(err)
	}
	queue, err = connection.OpenQueue("alerts")
	if err != nil {
		panic(err)
	}

	if err := queue.StartConsuming(prefetchLimit, pollDuration); err != nil {
		panic(err)
	}

	for i := 0; i < numConsumers; i++ {
		name := fmt.Sprintf("consumer %d", i)
		if _, err := queue.AddConsumer(name, consumer.NewConsumer(i)); err != nil {
			panic(err)
		}
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)
	defer signal.Stop(signals)

	<-signals // wait for signal
	go func() {
		<-signals // hard exit on second signal (in case shutdown gets stuck)
		os.Exit(1)
	}()

}

type Messanger struct {
	ConcurrentJobs int
}

//TODO remove rmq dependency

func AddToQueue(task []types.AlertTask) {
	for _, v := range task {
		taskBytes, _ := json.Marshal(v)
		if err := queue.Publish(string(taskBytes)); err != nil {
			log.Printf("failed to publish: %s", err)
		}
	}
}

func logErrors(errChan <-chan error) {
	for err := range errChan {
		switch err := err.(type) {
		case *rmq.HeartbeatError:
			if err.Count == rmq.HeartbeatErrorLimit {
				log.Print("heartbeat error (limit): ", err)
			} else {
				log.Print("heartbeat error: ", err)
			}
		case *rmq.ConsumeError:
			log.Print("consume error: ", err)
		case *rmq.DeliveryError:
			log.Print("delivery error: ", err.Delivery, err)
		default:
			log.Print("other error: ", err)
		}
	}
}
