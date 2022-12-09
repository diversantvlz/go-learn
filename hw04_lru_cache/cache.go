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

func (cache *lruCache) Set(key Key, value interface{}) bool {
	cache.mx.Lock()
	defer cache.mx.Unlock()
	if item, exist := cache.items[key]; exist {
		item.Value = value
		cache.queue.MoveToFront(item)

		return true
	}

	cache.items[key] = cache.queue.PushFront(value)
	if cache.capacity < cache.queue.Len() {
		removedItem := cache.queue.Back()
		for removeKey, searchItem := range cache.items {
			if searchItem == removedItem {
				delete(cache.items, removeKey)
				break
			}
		}
		cache.queue.Remove(removedItem)
	}

	return false
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	cache.mx.RLock()
	defer cache.mx.RUnlock()
	if item, exist := cache.items[key]; exist {
		cache.queue.MoveToFront(item)
		return item.Value, true
	}

	return nil, false
}

func (cache *lruCache) Clear() {
	cache.mx.Lock()
	defer cache.mx.Unlock()
	cache.items = map[Key]*ListItem{}
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
