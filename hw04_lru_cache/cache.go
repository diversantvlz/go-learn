package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mx       sync.RWMutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	cache.mx.Lock()
	defer cache.mx.Unlock()
	if item, exist := cache.items[key]; exist {
		item.Value = cacheItem{
			key:   key,
			value: value,
		}

		cache.queue.MoveToFront(item)

		return true
	}

	cache.items[key] = cache.queue.PushFront(cacheItem{
		key:   key,
		value: value,
	})

	if cache.capacity < cache.queue.Len() {
		back := cache.queue.Back()
		removedItem, _ := back.Value.(cacheItem)
		delete(cache.items, removedItem.key)
		cache.queue.Remove(back)
	}

	return false
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	cache.mx.RLock()
	defer cache.mx.RUnlock()
	if item, exist := cache.items[key]; exist {
		cache.queue.MoveToFront(item)
		itemValie, _ := item.Value.(cacheItem)
		return itemValie.value, true
	}

	return nil, false
}

func (cache *lruCache) Clear() {
	cache.mx.Lock()
	defer cache.mx.Unlock()
	cache.items = make(map[Key]*ListItem, cache.capacity)
	for item := cache.queue.Front(); item != nil; item = item.Next {
		cache.queue.Remove(item)
	}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
