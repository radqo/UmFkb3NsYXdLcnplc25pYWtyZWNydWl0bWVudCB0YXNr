package inmemory

import (
	"fmt"
	"log"
	"sync"
	"time"
	"errors"
	"github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/abstraction"
)

type cacheItem struct {
	value     interface{}
	err       error
	timestamp time.Time
	ready     chan struct{}
}

type cacheService struct {
	items      map[string]*cacheItem
	timeoutSec float64
	sync.Mutex
}

func New(timeout int) abstraction.CacheGetter {
	return &cacheService{items: make(map[string]*cacheItem), timeoutSec: float64(timeout)}
}

func (c *cacheService) Get(key string, f abstraction.GetFunc) (value interface{}, err error) {
	c.Lock()
	item, ok := c.items[key]
	if !ok || (ok && c.timeoutSec < time.Since(item.timestamp).Seconds()) {
		item = &cacheItem{ready: make(chan struct{})}
		item.timestamp = time.Now()
		c.items[key] = item
		c.Unlock()

		defer func() {
			if r := recover(); r != nil {
				log.Println("panic in getFunc")
				log.Println(fmt.Sprintf("Recovered in Get from cache for key %s (%s)", key, r))
				item.value = nil
				item.err = errors.New("Internal server error")
				value = item.value
				err = item.err
				close(item.ready)
			}
		}()

		item.value, item.err = f(key)
		close(item.ready)
	} else {
		c.Unlock()
		<-item.ready
	}
	return item.value, item.err
}