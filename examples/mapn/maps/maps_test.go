package maps

import (
	"sync"
	"testing"
)

func TestNestedStoreLoad(t *testing.T) {
	var m Map3[string, int, string, int]
	m.Store("a", 1, "x", 42)
	if v, ok := m.Load("a", 1, "x"); !ok || v != 42 {
		t.Fatalf("got %v %v", v, ok)
	}
	if _, ok := m.Load("a", 2, "x"); ok {
		t.Fatal("unexpected hit")
	}
	m.Delete("a", 1, "x")
	if _, ok := m.Load("a", 1, "x"); ok {
		t.Fatal("expected delete")
	}
}

func TestLoadOrStoreReturnsStored(t *testing.T) {
	var m Map[string, int]
	if v, loaded := m.LoadOrStore("a", 5); loaded || v != 5 {
		t.Fatalf("got %v %v", v, loaded)
	}
	if v, loaded := m.LoadOrStore("a", 6); !loaded || v != 5 {
		t.Fatalf("got %v %v", v, loaded)
	}
}

func TestConcurrentStore(t *testing.T) {
	var m Map2[int, int, int]
	var wg sync.WaitGroup
	for i := range 50 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m.Store(0, i, i)
		}()
	}
	wg.Wait()
	for i := range 50 {
		if v, ok := m.Load(0, i); !ok || v != i {
			t.Fatalf("lost write %d: %v %v", i, v, ok)
		}
	}
}
