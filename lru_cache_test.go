package cache

import "testing"

func TestPut(t *testing.T) {
	c := newCache(2)
	c.put("key", 4)

	if c.store["key"].value != 4 {
		t.Errorf("Expected 4, got %v", c.store["key"])
	}
}

func TestGetWhenItemFound(t *testing.T) {
	c := newCache(2)
	c.put("key", 4)

	c.assertKey("key", 4, t)
}
func TestGetWhenItemNotFound(t *testing.T) {
	c := newCache(2)

	c.assertKey("keyOne", -1, t)
}

func TestGetItemWithMultipleSaved(t *testing.T) {
	c := newCache(2)
	c.put("keyOne", 1)
	c.put("keyTwo", 2)

	c.assertKey("keyOne", 1, t)
	c.assertKey("keyTwo", 2, t)
}

func TestCacheEvictsPassedCapacity(t *testing.T) {
	c := newCache(1)

	c.put("keyOne", 1)
	c.put("keyTwo", 2)

	c.assertKey("keyOne", -1, t)
	c.assertKey("keyTwo", 2, t)
}

func TestLengthNeverGreaterThanCapacity(t *testing.T) {
	c := newCache(2)

	c.put("keyOne", 1)
	c.put("keyTwo", 2)
	c.put("keyThree", 3)

	if len(c.store) != 2 {
		t.Errorf("Expected 3, got %v", len(c.store))
	}
}
func TestEvictsOldest(t *testing.T) {
	c := newCache(2)

	c.put("keyOne", 1)
	c.put("keyTwo", 2)
	c.put("keyThree", 3)

	c.assertKey("keyOne", -1, t)
	c.assertKey("keyTwo", 2, t)
	c.assertKey("keyThree", 3, t)
}
func TestCacheMostRecentlyReadMovesTailToHead(t *testing.T) {
	c := newCache(3)

	c.put("keyOne", 1)
	c.put("keyTwo", 2)
	c.put("keyThree", 3)
	c.get("keyOne")
	c.put("keyFour", 4)

	c.assertKey("keyOne", 1, t)
	c.assertKey("keyTwo", -1, t)
	c.assertKey("keyThree", 3, t)
	c.assertKey("keyFour", 4, t)
}
func TestCacheMostRecentlyReadMovesDoesNotMoveHead(t *testing.T) {
	c := newCache(2)

	c.put("keyOne", 1)
	c.put("keyTwo", 2)
	c.get("keyTwo")
	c.put("keyThree", 3)

	c.assertKey("keyOne", -1, t)
	c.assertKey("keyTwo", 2, t)
	c.assertKey("keyThree", 3, t)
}
func TestCacheMostRecentlyReadMovesMiddleToHead(t *testing.T) {
	c := newCache(3)

	c.put("keyOne", 1)
	c.put("keyTwo", 2)
	c.put("keyThree", 3)
	c.get("keyTwo")
	c.put("keyFour", 4)
	c.put("keyFive", 5)

	c.assertKey("keyOne", -1, t)
	c.assertKey("keyTwo", 2, t)
	c.assertKey("keyThree", -1, t)
	c.assertKey("keyFour", 4, t)
	c.assertKey("keyFive", 5, t)
}

func TestCacheUpdatesValueWithoutChangingLength(t *testing.T) {
	c := newCache(3)

	c.put("keyOne", 1)
	c.put("keyOne", 10)

	c.assertKey("keyOne", 10, t)

	if len(c.store) != 1 {
		t.Errorf("Expected 1, got %v", len(c.store))
	}
}

func TestCacheMostRecentlyUpdatedMovesToHead(t *testing.T) {
	c := newCache(2)

	c.put("keyOne", 1)
	c.put("keyTwo", 2)
	c.put("keyOne", 10)
	c.put("keyThree", 3)

	c.assertKey("keyOne", 10, t)
	c.assertKey("keyTwo", -1, t)
	c.assertKey("keyThree", 3, t)
}

func (c lruCache) assertKey(k string, expected int, t *testing.T) {
	value := c.get(k)

	if value != expected {
		t.Errorf("Expected %v, got %v", expected, value)
	}
}
