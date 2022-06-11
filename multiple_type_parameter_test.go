package golang_generic

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func MultipleParameter[T1 any, T2 any](param1 T1, param2 T2) (T1, T2) {
	fmt.Println(param1)
	fmt.Println(param2)

	return param1, param2
}

//memanggil function MultipleParameter
func TestMultipleParam(t *testing.T) {
	result1, result2 := MultipleParameter[string, int]("Renaldi", 24)
	assert.Equal(t, "Renaldi", result1)
	assert.Equal(t, 24, result2)

	result1, result3 := MultipleParameter[string, string]("Aku", "Kamu")
	assert.Equal(t, "Aku", result1)
	assert.Equal(t, "Kamu", result3)
}
