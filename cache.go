package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type cacheItem struct {
	value     interface{}
	err       error
	timestamp time.Time
	ready     chan struct{}
}

type cache struct {
	items      map[string]*cacheItem
	timeoutSec float64
	sync.Mutex
}

type getFunc func(key string) (interface{}, error)

type getError struct {
	Message string
}

func (e getError) Error() string {
	return e.Message
}

func newCache(timeout int) *cache {
	return &cache{items: make(map[string]*cacheItem), timeoutSec: float64(timeout)}
}

func (c *cache) Get(key string, f getFunc) (value interface{}, err error) {
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
				item.err = getError{Message: "500"}
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
