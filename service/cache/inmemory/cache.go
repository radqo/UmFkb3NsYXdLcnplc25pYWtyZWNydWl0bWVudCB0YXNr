package inmemory

import (
	"github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/abstraction"
	"sync"
	"time"
)

type cacheItem struct {
	value     interface{}
	timestamp time.Time
}

type cacheService struct {
	items      map[string]*cacheItem
	timeoutSec float64
	sync.RWMutex
}

// New - creates new cacheService instance
func New(timeout int) abstraction.CacheOperator {
	return &cacheService{items: make(map[string]*cacheItem), timeoutSec: float64(timeout)}
}

func (c *cacheService) Get(key string) (value interface{}, found bool) {
	c.RLock()
	defer c.RUnlock()

	item, ok := c.items[key]
	if !ok || (ok && c.timeoutSec < time.Since(item.timestamp).Seconds()) {
		return nil, false
	}
	return item.value, true
}

func (c *cacheService) Set(key string, value interface{}) {
	c.Lock()
	defer c.Unlock()

	item := &cacheItem{}
	item.timestamp = time.Now()
	item.value = value
	c.items[key] = item
}
