package fizz_buzz

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	testCases := []struct{
		input uint
		output string
	} {
		{0, "FizzBuzz"},
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{6, "Fizz"},
		{7, "7"},
		{8, "8"},
		{9, "Fizz"},
		{10, "Buzz"},
		{11, "11"},
		{12, "Fizz"},
		{13, "13"},
		{14, "14"},
		{15, "FizzBuzz"},
		{16, "16"},
		{17, "17"},
		{18, "Fizz"},
	}

	for _, test := range testCases {
		output := FizzBuzz(test.input)
		outputSlow := fizzBuzzSlow(test.input)
		assert.Equal(t, test.output, output)
		assert.Equal(t, test.output, outputSlow)
	}
}

func TestFizzBuzz_DifferentialTesting(t *testing.T) {
	for i := 0; i < 10000; i++ {
		input := uint(rand.Uint64())
		outputSlow := fizzBuzzSlow(input)
		output := FizzBuzz(input)
		assert.Equal(t, outputSlow, output)
	}
}

func BenchmarkFizzBuzzSlow(b *testing.B) {
	input := make([]uint, b.N)
	for i, _ := range input {
		input[i] = uint(rand.Uint64())
	}
	b.ResetTimer()
	for _, v := range input {
		fizzBuzzSlow(v);
	}
}
func BenchmarkFizzBuzz(b *testing.B) {
	input := make([]uint, b.N)
	for i, _ := range input {
		input[i] = uint(rand.Uint64())
	}
	b.ResetTimer()
	for _, v := range input {
		FizzBuzz(v);
	}
}

func BenchmarkFizzBuzzFakeSlow(b *testing.B) {
	input := make([]uint, b.N)
	for i, _ := range input {
		input[i] = uint(rand.Uint64())
	}
	b.ResetTimer()
	for _, v := range input {
		fizzBuzzFakeSlow(v);
	}
}
func BenchmarkFizzBuzzFake(b *testing.B) {
	input := make([]uint, b.N)
	for i, _ := range input {
		input[i] = uint(rand.Uint64())
	}
	b.ResetTimer()
	for _, v := range input {
		fizzBuzzFake(v);
	}
}
