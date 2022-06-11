package golang_generic

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
func Length[T]() {

}

*/

/*
func Length[T interface{}]() {

}
*/

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func kali(a int, b int, c int) int {
	hasil := a * b * c
	return hasil
}

func bagi(a int, b int, c int) int {
	hasil := a / b / c
	return hasil
}

func jumlah(a int, b int, c int) int {
	hasil := a + b + c
	return hasil
}

func kurang(a int, b int, c int) int {
	hasil := a - b - c
	return hasil
}

func TestSample(t *testing.T) {
	// memanggil function kali
	hasil := kali(2, 3, 2)
	assert.Equal(t, 12, hasil)

	//memanggil function bagi
	hasil = bagi(20, 2, 5)
	assert.Equal(t, 2, hasil)

	//memanggil function jumlah
	hasil = jumlah(2, 5, 5)
	assert.Equal(t, 12, hasil)

	//memanggil function kurang
	hasil = kurang(10, 5, 4)
	assert.Equal(t, 1, hasil)
	assert.True(t, true)

	//memanggil function yang type parameter
	var result1 = Length[string]("Renaldi")
	assert.Equal(t, "Renaldi", result1)

	var resultNumber = Length[int](100)
	assert.Equal(t, 100, resultNumber)

	var resultBool = Length[bool](true)
	assert.True(t, resultBool)

	var resultFloat = Length[float64](1.23)
	assert.Equal(t, 1.23, resultFloat)
}
