package cache

import (
	"fmt"
	"time"
)

type Cache struct {
	vars      map[string]string
	deadlines map[string]time.Time
}

func NewCache() Cache {

	vars := make(map[string]string)
	deadlines := make(map[string]time.Time)

	return Cache{vars, deadlines}
}

func (c Cache) Get(key string) (string, bool) {
	if val, ok := c.vars[key]; ok {
		if (&c).CheckExpired(key) {
			return "", false
		}
		return val, ok
	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	c.vars[key] = value
}

func (c *Cache) CheckExpired(key string) bool {
	if _, ok := c.deadlines[key]; c.deadlines[key].Sub(time.Now()) <= time.Second*0 && ok {
		delete(c.deadlines, key)
		return true
	}
	return false
}

func (c Cache) Keys() []string {
	var keys []string
	fmt.Println(c.vars)
	for k, _ := range c.vars {
		if (&c).CheckExpired(k) {
			delete(c.deadlines, k)
			continue
		}
		keys = append(keys, k)
	}
	fmt.Println(keys)
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.vars[key] = value
	c.deadlines[key] = deadline
}
