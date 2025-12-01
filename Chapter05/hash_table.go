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
	bucketIndex := h.hashKey(key)
	newEntry := &entry[K, V]{key: key, value: value, next: h.buckets[bucketIndex]}
	h.buckets[bucketIndex] = newEntry
	h.count++
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
	currentHead := h.buckets[bucketIndex]

	if currentHead != nil && currentHead.key == key {
		h.buckets[bucketIndex] = nil
		h.count--
		return true
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
