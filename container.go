package winter

import (
	"database/sql"
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

	return nil, fmt.Errorf("'%s' does not exist", key)
}

var container = &Container{m: make(map[string]interface{})}

func Register(key string, item interface{}) {
	container.Add(key, item)
}

func DBFromContainer(key string) *sql.DB {

	if item, err := container.Get(key); err == nil {
		return item.(*sql.DB)
	}

	panic("Can not resolve from container")
}
