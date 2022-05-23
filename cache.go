package cache

import (
	"fmt"
	"time"
)

type Cache struct {
	vars         map[string]string
	deadlines    map[string]time.Time
	varsInitTime map[string]time.Time
}

func NewCache() Cache {

	vars := make(map[string]string)
	varsInitTime := make(map[string]time.Time)
	deadlines := make(map[string]time.Time)

	return Cache{vars, varsInitTime, deadlines}
}

func (c Cache) Get(key string) (string, bool) {
	if val, ok := c.vars[key]; ok {
		if c.CheckExpired(key) {
			return "", false
		}
		return val, ok
	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	c.vars[key] = value
	c.varsInitTime[key] = time.Now()
}

func (c *Cache) CheckExpired(key string) bool {
	if _, ok := (*c).deadlines[key]; ok {
		fmt.Println((*c).deadlines[key].Sub(time.Now()))
		exp := (*c).deadlines[key].Sub(time.Now())
		if exp <= time.Second*0 {
			delete((*c).deadlines, key)
			return true
		}
	}
	return false
}

func (c Cache) Keys() []string {
	var keys []string
	fmt.Println(c.vars)
	for k, _ := range c.vars {
		if !c.CheckExpired(k) {
			keys = append(keys, k)
		}
	}
	fmt.Println(keys)
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.vars[key] = value
	c.deadlines[key] = deadline
	c.varsInitTime[key] = time.Now()
}
