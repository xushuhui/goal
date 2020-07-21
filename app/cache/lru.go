package cache

import (
	"container/list"
)

type Cache struct {
	maxBytes int64
	useBytes int64
	list     *list.List
	cache    map[string]*list.Element
	//回调
	callback func(key string, value Value)
}
type entry struct {
	key   string
	value Value
}
type Value interface {
	Len() int
}

func New(maxBytes int64, callback func(key string, value Value)) *Cache {
	return &Cache{
		maxBytes: maxBytes,
		list:     list.New(),
		cache:    make(map[string]*list.Element),
		callback: callback,
	}
}
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.list.MoveToBack(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}
func (c *Cache) RemoveOldest() {
	ele := c.list.Front()
	if ele != nil {
		c.list.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.useBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.callback != nil {
			c.callback(kv.key, kv.value)
		}
	}
}
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok { //值已经存在，更新value
		c.list.MoveToBack(ele)
		kv := ele.Value.(*entry)
		c.useBytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.list.PushBack(&entry{key: key, value: value})
		c.cache[key] = ele
		c.useBytes += int64(len(key)) + int64(value.Len())
	}
	for c.maxBytes != 0 && c.maxBytes < c.useBytes {
		c.RemoveOldest()
	}
}
func (c *Cache) Len() int {
	return c.list.Len()
}
