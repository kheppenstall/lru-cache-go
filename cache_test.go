package cache

import "testing"

func testNewCache(t *testing.T) {
	c := newCache(2)

	if c.capacity != 2 {
		t.Errorf("Expected capacity of 2, got %v", c.capacity)
	}
}
