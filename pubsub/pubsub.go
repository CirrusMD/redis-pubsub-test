package pubsub

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/garyburd/redigo/redis"
)

func PubSub(numConsumers, numPublishes int) uint64 {
	publisher, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	var counter uint64 = 0
	defer publisher.Close()

	fmt.Println("creating", numConsumers, "consumers")
	var wg sync.WaitGroup
	wg.Add(numPublishes * numConsumers)

	consumers := []redis.PubSubConn{}
	for i := 0; i < numConsumers; i++ {
		c, err := redis.Dial("tcp", ":6379")
		if err != nil {
			panic(err)
		}
		sub := redis.PubSubConn{Conn: c}
		sub.Subscribe("test:pubsub")
		go func(c redis.PubSubConn) {
			for {
				switch n := c.Receive().(type) {
				case redis.Message:
					atomic.AddUint64(&counter, 1)
					wg.Done()
				case error:
					fmt.Println("ERROR in receive", n)
					return
				}
			}
		}(sub)
		consumers = append(consumers, sub)
	}

	defer func() {
		for _, c := range consumers {
			consumer := c
			consumer.Conn.Close()
		}
	}()

	fmt.Println("waiting")

	for i := 0; i < numPublishes; i++ {
		publisher.Do("publish", "test:pubsub", "1")
	}

	wg.Wait()

	return atomic.LoadUint64(&counter)
}
