package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	main()

	assert.Equal(t, concurrent*2, deposit)
}
