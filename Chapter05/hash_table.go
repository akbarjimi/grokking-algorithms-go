package Chapter05

import "hash/fnv"

const loadFactor float64 = 0.75

type entry[K comparable, V any] struct {
	key   K
	value V
	next  *entry[K, V]
}

type HashMap[K comparable, V any] struct {
	buckets    []*entry[K, V]
	count      int
	capacity   int
	loadFactor float64
}

func NewHashMap[K comparable, V any](capacity int) *HashMap[K, V] {
	return &HashMap[K, V]{
		buckets:    make([]*entry[K, V], capacity),
		count:      0,
		loadFactor: loadFactor,
		capacity:   capacity,
	}
}

func (h *HashMap[K, V]) Put(key K, value V) {
	foundEntry, ok := h.findKey(key)
	if ok {
		foundEntry.value = value
		return
	}
	if float64(h.count)/float64(h.capacity) >= loadFactor {
		h.resize()
	}

	bucketIndex := h.hashKey(key)
	newEntry := &entry[K, V]{key: key, value: value, next: h.buckets[bucketIndex]}
	h.buckets[bucketIndex] = newEntry
	h.count++
	return
}

func (h *HashMap[K, V]) Get(key K) (V, bool) {
	entry, ok := h.findKey(key)
	if ok {
		return entry.value, ok
	}
	return *new(V), false
}

func (h *HashMap[K, V]) Delete(key K) bool {
	bucketIndex := h.hashKey(key)
	currentEntry := h.buckets[bucketIndex]
	var previousEntry *entry[K, V]

	for currentEntry != nil {
		if currentEntry.key == key {
			if previousEntry != nil {
				previousEntry.next = currentEntry.next
			} else {
				h.buckets[bucketIndex] = currentEntry.next
			}
			h.count--
			return true
		}
		previousEntry = currentEntry
		currentEntry = currentEntry.next
	}

	return false
}

func (h *HashMap[K, V]) hashKey(key K) int {
	if s, ok := any(key).(string); ok {
		var byteStream = []byte(s)
		hasher := fnv.New64a()
		_, err := hasher.Write(byteStream)
		if err != nil {
			return 0
		}
		return int(hasher.Sum64() % uint64(h.capacity))
	}
	return 0
}

func (h *HashMap[K, V]) findKey(key K) (*entry[K, V], bool) {
	bucketIndex := h.hashKey(key)
	currentEntry := h.buckets[bucketIndex]
	for currentEntry != nil {
		if currentEntry.key == key {
			return currentEntry, true
		}
		currentEntry = currentEntry.next
	}
	return nil, false
}

func (h *HashMap[K, V]) resize() {
	biggerHashMap := NewHashMap[K, V](h.capacity * 2)

	for i := 0; i < h.capacity; i++ {
		currentEntry := h.buckets[i]
		for currentEntry != nil {
			biggerHashMap.Put(currentEntry.key, currentEntry.value)
			currentEntry = currentEntry.next
		}
	}

	h.capacity = biggerHashMap.capacity
	h.buckets = biggerHashMap.buckets
	h.count = biggerHashMap.count
}
