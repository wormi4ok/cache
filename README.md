# Cache

A simple cache implementation

## LRU Cache

An in memory cache implementation that expires the least recently used items, and limits cache maxSize by a maximum
number of items.

Example usage:

```
import (
	"errors"
	"github.com/wormi4ok/cache"
)

func example() {
	c := cache.NewLRU(100)
	c.Write("key", "value")
	val, err := c.Read("key")
	if !errors.Is(err,cache.ErrorNoValue){
	    // cache hit
    }
    // cache miss
}
```
