package pblocal

import (
	"Delivery_Food/common"
	"Delivery_Food/pubsub"
	"context"
	"log"
	"sync"
)

type localPubSub struct {
	messageQueue chan *pubsub.Message
	mapChannel   map[pubsub.Topic][]chan *pubsub.Message
	locker       *sync.RWMutex
}

func NewLocalPubSub() *localPubSub {
	pb := &localPubSub{
		messageQueue: make(chan *pubsub.Message, 10000),
		mapChannel:   make(map[pubsub.Topic][]chan *pubsub.Message),
		locker:       &sync.RWMutex{},
	}

	pb.Start()

	return pb
}

func (pb *localPubSub) Publish(ctx context.Context, topic pubsub.Topic,
	data *pubsub.Message) error {
	data.SetChannel(topic)

	go func() {
		defer common.AppRecover()
		pb.messageQueue <- data
		log.Println("New message in queue", data.String())
	}()

	return nil
}

func (pb *localPubSub) Subscribe(ctx context.Context,
	topic pubsub.Topic) (ch <-chan *pubsub.Message,
	close func()) {
	c := make(chan *pubsub.Message)

	pb.locker.Lock()

	if val, ok := pb.mapChannel[topic]; ok {
		pb.mapChannel[topic] = append(val, c)
	} else {
		pb.mapChannel[topic] = []chan *pubsub.Message{c}
	}

	pb.locker.Unlock()

	return c, func() {
		log.Println("Unsubscribe", topic)
		pb.locker.Lock()
		defer pb.locker.Unlock()

		if val, ok := pb.mapChannel[topic]; ok {
			for i, v := range val {
				if v == c {
					pb.mapChannel[topic] = append(val[:i], val[i+1:]...)
					break
				}
			}
		}
	}
}

func (pb *localPubSub) Start() {
	go func() {
		for {
			select {
			case <-context.Background().Done():
				return
			case msg := <-pb.messageQueue:
				pb.locker.RLock()
				if val, ok := pb.mapChannel[msg.Channel()]; ok {
					for i := range val {
						go func(c chan *pubsub.Message) {
							defer common.AppRecover()
							c <- msg
						}(val[i])
					}
				}
				pb.locker.RUnlock()
			}
		}
	}()
}
