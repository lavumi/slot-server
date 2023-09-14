package module

import (
	"github.com/goark/mt"
	"github.com/goark/mt/mt19937"
	"time"
)

var prng *mt.PRNG

func init() {
	prng = mt.New(mt19937.New(time.Now().UnixNano()))
}

func GenRandom() float64 {
	return prng.Real(1)
}

func MakeRandomIntArray(max uint, count int) []int {

	var result []int

	for i := 0; i < count; i++ {
		random := prng.Real(1)
		result = append(result, int(random*float64(max)))
	}

	return result
}
