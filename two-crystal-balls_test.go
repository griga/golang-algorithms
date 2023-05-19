package algorithms_test

import (
	"cybe/algorithms"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoCrystalBalls(t *testing.T) {
	case00 := make([]bool, 10000)
	idx00 := rand.Intn(10000)

	for i := idx00; i < 10000; i++ {
		case00[i] = true
	}
	assert.Equal(t, idx00, algorithms.TwoCrystalBalls(&case00))

	case01 := make([]bool, 821)
	assert.Equal(t, -1, algorithms.TwoCrystalBalls(&case01))
}
