package cache

import "time"

type Cache struct {
	vars      map[string]string
	deadlines map[string]time.Time
}

func NewCache() Cache {
	return Cache{}
}

func (c Cache) Get(key string) (string, bool) {
	if val, ok := c.vars[key]; ok {
		if c.deadlines[key].Sub(time.Now()) <= time.Second*0 {
			delete(c.deadlines, key)
			return "", false
		}
		return val, ok
	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	c.vars[key] = value
}

func (c Cache) Keys() []string {
	var keys []string
	for k, _ := range c.vars {
		keys = append(keys, k)
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.vars[key] = value
	c.deadlines[key] = deadline
}
