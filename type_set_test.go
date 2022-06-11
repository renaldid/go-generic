package golang_generic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Age int

type Number interface {
	~int | int8 | int16 | int32 | int64 | float32 | float64
}

func Min[T Number](number1 T, number2 T) T {
	if number1 < number2 {
		return number1
	} else {
		return number2
	}
}

func Max[T Number](number1 T, number2 T) T {
	if number1 < number2 {
		return number2
	} else {
		return number1
	}
}

func TestMinMax(t *testing.T) {
	assert.Equal(t, int(100), Min[int](100, 200))
	assert.Equal(t, int(200), Max[int](100, 200))
	assert.Equal(t, Age(20), Max[Age](10, 20))
	assert.Equal(t, int64(100), Min[int64](100, 200))
	assert.Equal(t, float64(100), Min[float64](100.0, 200.0))
}

func TestInference(t *testing.T) {
	assert.Equal(t, 100, Min(100, 200))
	assert.Equal(t, int64(200), Max(int64(100), int64(200)))
	assert.Equal(t, Age(20), Max(Age(10), Age(20)))
	assert.Equal(t, float64(20000), Max(1000.0, 20000.0))
}
