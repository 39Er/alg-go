package cache

import (
	"container/list"
	"errors"
)

type LRUCache struct {
	MaxEntries    int
	cache         map[Key]*list.Element
	ll            *list.List
	OnEntryDelete func(key Key, val interface{})
}

type Key interface{}

type entry struct {
	key   Key
	value interface{}
}

func NewLRUCache(maxentry int) (out *LRUCache, err error) {
	out = &LRUCache{
		MaxEntries: maxentry,
		cache:      make(map[Key]*list.Element),
		ll:         list.New()}

	return out, nil
}

func (l *LRUCache) Add(key Key, val interface{}) {
	if l.cache == nil {
		l.cache = make(map[Key]*list.Element)
		l.ll = list.New()
	}

	if ele, ok := l.cache[key]; ok {
		l.ll.MoveToFront(ele)
		ele.Value.(*entry).value = val
		return
	}

	ele := l.ll.PushFront(&entry{key: key, value: val})
	l.cache[key] = ele

	if l.MaxEntries > 0 && l.ll.Len() > l.MaxEntries {
		l.RemoveOldestEntry()
	}
}

func (l *LRUCache) Get(key Key) (outval interface{}, err error) {
	if l.cache == nil {
		return outval, errors.New("not initialize object")
	}

	if ele, ok := l.cache[key]; ok {
		l.ll.MoveToFront(ele)
		return ele.Value.(*entry).value, nil
	}
	return outval, errors.New("no exist")
}

func (l *LRUCache) RemoveOldestEntry() {
	if l.cache == nil || l.ll == nil || l.ll.Len() == 0 {
		return
	}

	ele := l.ll.Back()
	if ele != nil {
		l.removeEntry(ele)
	}
}

func (l *LRUCache) removeEntry(ele *list.Element) {
	l.ll.Remove(ele)
	delete(l.cache, ele.Value.(*entry).key)
	if l.OnEntryDelete != nil {
		l.OnEntryDelete(ele.Value.(*entry).key, ele.Value.(*entry).value)
	}
}

func (l *LRUCache) Len() int {
	if l.cache == nil {
		return 0
	}

	return l.ll.Len()
}
