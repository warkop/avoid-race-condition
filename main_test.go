package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUseChannel(t *testing.T) {
	assert.Equal(t, concurrent*totalExecution, UseChannel())
}

func TestUseMutex(t *testing.T) {
	assert.Equal(t, concurrent*totalExecution, UseMutex())
}

func TestSynchronous(t *testing.T) {
	assert.Equal(t, concurrent*totalExecution, Synchronous())
}

func TestUseArray(t *testing.T) {
	assert.Equal(t, concurrent*totalExecution, UseArray())
}

func BenchmarkGoroutine(b *testing.B) {
	b.Run("synchronous", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Synchronous()
		}
	})

	b.Run("use mutex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			UseMutex()
		}
	})

	b.Run("use channel", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			UseChannel()
		}
	})

	b.Run("use array", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			UseArray()
		}
	})
}
