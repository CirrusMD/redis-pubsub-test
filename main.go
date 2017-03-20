package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/CirrusMD/redis-pubsub-test/pubsub"
)

func main() {
	numConsumers, _ := strconv.Atoi(os.Args[1])
	numPublishes, _ := strconv.Atoi(os.Args[2])
	count := pubsub.PubSub(numConsumers, numPublishes)
	fmt.Println("expected", numConsumers*numPublishes)
	fmt.Println("actual", count)
	fmt.Println("expected == actual is", count == uint64(numConsumers*numPublishes))
}
