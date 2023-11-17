package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMyFunction1(t *testing.T) {
	t.Parallel()
	t.Run("TestCase1", func(t *testing.T) {
		t.Parallel()
		time.Sleep(60 * time.Second)
	})

	t.Run("TestCase2", func(t *testing.T) {
		t.Parallel()
		time.Sleep(5 * time.Second)
		assert.Equal(t, 1, 1)
	})

	t.Run("TestCase3", func(t *testing.T) {
		t.Parallel()
		time.Sleep(5 * time.Second)
		assert.Equal(t, 1, 1)
	})
}
