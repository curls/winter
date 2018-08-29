package container

import (
	"fmt"
	"sync"
)

type Container struct {
	sync.RWMutex
	m map[string]interface{}
}

func (c *Container) Add(key string, item interface{}) {
	c.Lock()
	c.m[key] = item
	c.Unlock()
}

func (c *Container) Get(key string) (item interface{}, err error) {

	item, exist := c.m[key]
	if exist {
		return item, nil
	}

	return nil, fmt.Errorf("element '%s' does not exist", key)
}

var container = &Container{m: make(map[string]interface{})}

func Register(key string, item interface{}) {
	container.Add(key, item)
}

func Resolve(key string) (interface{}, error) {

	return container.Get(key)
}
