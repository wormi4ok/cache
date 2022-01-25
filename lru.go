package cache

import (
	"container/list"
)

type LRU struct {
	maxSize int
	keys    *list.List

	cache map[string]item
}

type item struct {
	key   *list.Element
	value interface{}
}

func NewLRU(maxSize int) Cache {
	cache := make(map[string]item, maxSize)

	return &LRU{
		maxSize: maxSize,
		keys:    list.New(),
		cache:   cache,
	}
}

func (c *LRU) Write(key string, value interface{}) {
	if v, ok := c.cache[key]; ok {
		c.keys.MoveToFront(v.key)
		c.cache[key] = item{
			key:   v.key,
			value: value,
		}

		return
	}
	if len(c.cache) >= c.maxSize {
		oldestKey := c.keys.Back()

		delete(c.cache, oldestKey.Value.(string))
		c.keys.Remove(oldestKey)
	}

	c.cache[key] = item{
		key:   c.keys.PushFront(key),
		value: value,
	}
}

func (c *LRU) Read(key string) (interface{}, error) {
	v, ok := c.cache[key]
	if !ok {
		return "", ErrorNoValue
	}

	c.keys.MoveToFront(v.key)
	return v.value, nil
}
