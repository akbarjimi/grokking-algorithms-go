package Chapter05

import (
	"testing"
)

func TestHashMap_BasicOperations(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		key       string
		val       int
		wantVal   int
		wantFound bool
	}{
		{"Insert and Retrieve", "apple", 10, 10, true},
		{"Key Not Found", "banana", 0, 0, false},
		{"Update Existing", "apple", 20, 20, true},
	}

	m := NewHashMap[string, int](16)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Key Not Found" {
				// Just check retrieval
			} else {
				m.Put(tt.key, tt.val)
			}

			gotVal, gotFound := m.Get(tt.key)

			if gotFound != tt.wantFound {
				t.Errorf("Get(%q) found = %v, want %v", tt.key, gotFound, tt.wantFound)
			}
			if gotVal != tt.wantVal {
				t.Errorf("Get(%q) value = %v, want %v", tt.key, gotVal, tt.wantVal)
			}
		})
	}
}

func TestHashMap_CollisionHandling(t *testing.T) {
	t.Parallel()

	m := NewHashMap[string, int](1)

	keys := []string{"alpha", "beta", "gamma", "delta"}

	for i, k := range keys {
		m.Put(k, i)
	}

	if m.count != 4 {
		t.Errorf("Expected count 4, got %d", m.count)
	}

	for i, k := range keys {
		val, ok := m.Get(k)
		if !ok || val != i {
			t.Errorf("Failed to retrieve collided key %s: got %d, want %d", k, val, i)
		}
	}
}

func TestHashMap_Delete(t *testing.T) {
	t.Parallel()

	m := NewHashMap[string, int](16)
	m.Put("key1", 100)
	m.Put("key2", 200)

	tests := []struct {
		name          string
		keyToDelete   string
		wantDeleted   bool
		expectedCount int
	}{
		{"Delete Existing", "key1", true, 1},
		{"Delete Non-Existent", "ghost", false, 1},
		{"Delete Remaining", "key2", true, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleted := m.Delete(tt.keyToDelete)

			if deleted != tt.wantDeleted {
				t.Errorf("Delete(%q) = %v, want %v", tt.keyToDelete, deleted, tt.wantDeleted)
			}

			if m.count != tt.expectedCount {
				t.Errorf("Count after delete = %d, want %d", m.count, tt.expectedCount)
			}

			if _, ok := m.Get(tt.keyToDelete); ok {
				t.Errorf("Key %q should have been deleted but was found", tt.keyToDelete)
			}
		})
	}
}
