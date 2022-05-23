package cache

import "time"

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
		if _, ok := c.deadlines[key]; c.deadlines[key].Sub(time.Now()) <= time.Second*0 && ok {
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
		if _, ok := c.deadlines[k]; c.deadlines[k].Sub(time.Now()) <= time.Second*0 && ok {
			delete(c.deadlines, k)
			continue
		}
		keys = append(keys, k)
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.vars[key] = value
	c.deadlines[key] = deadline
}
