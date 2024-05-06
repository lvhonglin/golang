package cov

import "testing"

func TestAdd(t *testing.T) {
	println(123)
	t.Run("test1", func(t *testing.T) {
		Add(2, 2)
	})
}
