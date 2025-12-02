package Chapter05

import (
	"strconv"
	"testing"
)

func setupMap(N int, b *testing.B) (*HashMap[string, int], []string) {
	m := NewHashMap[string, int](N / 2)
	keys := make([]string, N)

	for i := 0; i < N; i++ {
		key := "key_" + strconv.Itoa(i)
		m.Put(key, i)
		keys[i] = key
	}
	return m, keys
}

var benchmarkSizes = []int{
	100,
	1000,
	10000,
	100000,
}

func BenchmarkHashMap_Put(b *testing.B) {
	for _, size := range benchmarkSizes {
		b.Run("N="+strconv.Itoa(size), func(b *testing.B) {
			m := NewHashMap[string, int](size / 2)

			b.StopTimer()
			keys := make([]string, size)
			for i := 0; i < size; i++ {
				keys[i] = "key_" + strconv.Itoa(i)
			}

			b.StartTimer()

			for i := 0; i < b.N; i++ {
				m.Put(keys[i%size], i)
			}
		})
	}
}

func BenchmarkHashMap_Get(b *testing.B) {
	for _, size := range benchmarkSizes {
		b.Run("N="+strconv.Itoa(size), func(b *testing.B) {
			m, keys := setupMap(size, b)

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_, _ = m.Get(keys[i%size])
			}
		})
	}
}

func BenchmarkHashMap_Delete(b *testing.B) {
	for _, size := range benchmarkSizes {
		b.Run("N="+strconv.Itoa(size), func(b *testing.B) {
			b.StopTimer()

			keys := make([]string, size)
			for i := 0; i < size; i++ {
				keys[i] = "key_" + strconv.Itoa(i)
			}

			b.StartTimer()
			for i := 0; i < b.N; i++ {
				m := NewHashMap[string, int](size / 2)
				for j := 0; j < size; j++ {
					m.Put(keys[j], j)
				}

				b.StopTimer()
				keyToDelete := keys[i%size]
				b.StartTimer()

				m.Delete(keyToDelete)
			}
		})
	}
}
