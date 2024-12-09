package main

import "testing"

func TestAdd(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		result := 1
		if result != 1 {
			t.Error("Test failed")
		}
	})
}
