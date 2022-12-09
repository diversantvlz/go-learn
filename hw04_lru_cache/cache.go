package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	Cache // Remove me after realization.

	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	if item, exist := cache.items[key]; exist {
		item.Value = value
		cache.queue.MoveToFront(item)

		return true
	} else {
		cache.items[key] = cache.queue.PushFront(value)
		if cache.capacity < cache.queue.Len() {
			removedItem := cache.queue.Back()
			for removeKey, searchItem := range cache.items {
				if searchItem == removedItem {
					delete(cache.items, removeKey)
				}
			}
			cache.queue.Remove(removedItem)
		}

		return false
	}
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	if item, exist := cache.items[key]; exist {
		cache.queue.MoveToFront(item)
		return item.Value, true
	}

	return nil, false
}

func (cache *lruCache) Clear() {
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
