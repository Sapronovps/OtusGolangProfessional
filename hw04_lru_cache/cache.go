package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mu       sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

// Установка значения в кеш.
func (cache *lruCache) Set(key Key, value interface{}) bool {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	if cache.capacity == 0 {
		return false
	}
	existValue, ok := cache.items[key]
	if ok {
		cache.queue.Remove(existValue)
		moveElem := cache.queue.PushFront(value)
		cache.items[key] = moveElem
	} else {
		if cache.capacity > cache.queue.Len() {
			lastElem := cache.queue.Back()
			cache.queue.Remove(lastElem)
		}
		newElem := cache.queue.PushFront(value)
		cache.items[key] = newElem
	}

	return ok
}

// Получение значения из кеша.
func (cache *lruCache) Get(key Key) (interface{}, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	val, ok := cache.items[key]
	if !ok {
		return nil, false
	}

	cache.queue.MoveToFront(val)

	return val.Value, true
}

// Очистка кеша.
func (cache *lruCache) Clear() {
	cache.items = make(map[Key]*ListItem, cache.capacity)
	cache.queue = NewList()
}
