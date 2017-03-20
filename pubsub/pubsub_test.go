package pubsub

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/CirrusMD/redis-pubsub-test/pubsub"
)

func TestPubSub(t *testing.T) {
	iters := rand.Intn(50)
	fmt.Println("iterations", iters)
	for i := 0; i < iters; i++ {
		numC := rand.Intn(25)
		numP := rand.Intn(10000)

		result := pubsub.PubSub(numC, numP)

		if result != uint64(numC*numP) {
			t.Fatal("consumers:", numC, "publishes:", numP, "expected", numC*numP, "got", result)
		}
	}
}
