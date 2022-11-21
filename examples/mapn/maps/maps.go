package maps

import ()

type Map1[K0 comparable, V any] struct {
	inner Map[K0, V]
}

func (m *Map1[K0, V]) Delete0(key K0) {
	m.inner.Delete(key)
}

func (m *Map1[K0, V]) Load0(key K0) (value V, ok bool) {
	return m.inner.Load(key)
}

func (m *Map1[K0, V]) LoadAndDelete0(key K0) (value V, loaded bool) {
	return m.inner.LoadAndDelete(key)
}

func (m *Map1[K0, V]) LoadOrStore0(key K0, value V) (actual V, loaded bool) {
	return m.inner.LoadOrStore(key, value)
}

func (m *Map1[K0, V]) Range0(f func(key K0, value V) bool) {
	m.inner.Range(f)
}

func (m *Map1[K0, V]) Store0(key K0, value V) {
	m.inner.Store(key, value)
}

type Map2[K0 comparable, K1 comparable, V any] struct {
	inner *Map1[K0, *Map1[K1, V]]
}

func (m *Map2[K0, K1, V]) Load(k0 K0, k1 K1,
) (value V, ok bool) {
	return m.Load1(k0, k1)
}

func (m *Map2[K0, K1, V]) Delete(k0 K0, k1 K1,
) {
	m.Delete1(k0, k1)
}

func (m *Map2[K0, K1, V]) Range(k0 K0,
	f func(K1, V) bool) {
	m.Range1(k0, f)
}

func (m *Map2[K0, K1, V]) LoadAndDelete(k0 K0, k1 K1,
) (value V, loaded bool) {
	return m.LoadAndDelete1(k0, k1)
}

func (m *Map2[K0, K1, V]) LoadOrStore(k0 K0, k1 K1,
	v V) (value V, loaded bool) {
	return m.LoadOrStore1(k0, k1,
		v)
}

func (m *Map2[K0, K1, V]) Store(k0 K0, k1 K1,
	value V) {
	m.Store1(k0, k1,
		value)
}

func (m *Map2[K0, K1, V]) Load0(k0 K0,
) (value *Map1[K1, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	return item0.Load0(k0)
}

func (m *Map2[K0, K1, V]) Delete0(k0 K0,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	item0.Delete0(k0)
}

func (m *Map2[K0, K1, V]) Range0(
	f func(K0, *Map1[K1, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	item0.Range0(f)
}

func (m *Map2[K0, K1, V]) LoadAndDelete0(k0 K0,
) (value *Map1[K1, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	return item0.LoadAndDelete0(k0)
}

func (m *Map2[K0, K1, V]) LoadOrStore0(k0 K0,
	v *Map1[K1, V]) (value *Map1[K1, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map1[K1, V]])
	}
	item0 := m.inner

	return item0.LoadOrStore0(k0, v)
}

func (m *Map2[K0, K1, V]) Store0(k0 K0,
	value *Map1[K1, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map1[K1, V]])
	}
	item0 := m.inner

	item0.Store0(k0, value)
}

func (m *Map2[K0, K1, V]) Load1(k0 K0, k1 K1,
) (value V, ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map1[K1, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	return item1.Load0(k1)
}

func (m *Map2[K0, K1, V]) Delete1(k0 K0, k1 K1,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map1[K1, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	item1.Delete0(k1)
}

func (m *Map2[K0, K1, V]) Range1(k0 K0,
	f func(K1, V) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map1[K1, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	item1.Range0(f)
}

func (m *Map2[K0, K1, V]) LoadAndDelete1(k0 K0, k1 K1,
) (value V, loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map1[K1, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	return item1.LoadAndDelete0(k1)
}

func (m *Map2[K0, K1, V]) LoadOrStore1(k0 K0, k1 K1,
	v V) (value V, loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map1[K1, V]])
	}
	item0 := m.inner

	var item1 *Map1[K1, V]
	item1, _ = item0.LoadOrStore0(k0, &Map1[K1, V]{})

	return item1.LoadOrStore0(k1, v)
}

func (m *Map2[K0, K1, V]) Store1(k0 K0, k1 K1,
	value V) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map1[K1, V]])
	}
	item0 := m.inner

	var item1 *Map1[K1, V]
	item1, _ = item0.LoadOrStore0(k0, &Map1[K1, V]{})

	item1.Store0(k1, value)
}

type Map3[K0 comparable, K1 comparable, K2 comparable, V any] struct {
	inner *Map1[K0, *Map2[K1, K2, V]]
}

func (m *Map3[K0, K1, K2, V]) Load(k0 K0, k1 K1, k2 K2,
) (value V, ok bool) {
	return m.Load2(k0, k1, k2)
}

func (m *Map3[K0, K1, K2, V]) Delete(k0 K0, k1 K1, k2 K2,
) {
	m.Delete2(k0, k1, k2)
}

func (m *Map3[K0, K1, K2, V]) Range(k0 K0, k1 K1,
	f func(K2, V) bool) {
	m.Range2(k0, k1, f)
}

func (m *Map3[K0, K1, K2, V]) LoadAndDelete(k0 K0, k1 K1, k2 K2,
) (value V, loaded bool) {
	return m.LoadAndDelete2(k0, k1, k2)
}

func (m *Map3[K0, K1, K2, V]) LoadOrStore(k0 K0, k1 K1, k2 K2,
	v V) (value V, loaded bool) {
	return m.LoadOrStore2(k0, k1, k2,
		v)
}

func (m *Map3[K0, K1, K2, V]) Store(k0 K0, k1 K1, k2 K2,
	value V) {
	m.Store2(k0, k1, k2,
		value)
}

func (m *Map3[K0, K1, K2, V]) Load0(k0 K0,
) (value *Map2[K1, K2, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	return item0.Load0(k0)
}

func (m *Map3[K0, K1, K2, V]) Delete0(k0 K0,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	item0.Delete0(k0)
}

func (m *Map3[K0, K1, K2, V]) Range0(
	f func(K0, *Map2[K1, K2, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	item0.Range0(f)
}

func (m *Map3[K0, K1, K2, V]) LoadAndDelete0(k0 K0,
) (value *Map2[K1, K2, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	return item0.LoadAndDelete0(k0)
}

func (m *Map3[K0, K1, K2, V]) LoadOrStore0(k0 K0,
	v *Map2[K1, K2, V]) (value *Map2[K1, K2, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map2[K1, K2, V]])
	}
	item0 := m.inner

	return item0.LoadOrStore0(k0, v)
}

func (m *Map3[K0, K1, K2, V]) Store0(k0 K0,
	value *Map2[K1, K2, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map2[K1, K2, V]])
	}
	item0 := m.inner

	item0.Store0(k0, value)
}

func (m *Map3[K0, K1, K2, V]) Load1(k0 K0, k1 K1,
) (value *Map1[K2, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map2[K1, K2, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	return item1.Load0(k1)
}

func (m *Map3[K0, K1, K2, V]) Delete1(k0 K0, k1 K1,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map2[K1, K2, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	item1.Delete0(k1)
}

func (m *Map3[K0, K1, K2, V]) Range1(k0 K0,
	f func(K1, *Map1[K2, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map2[K1, K2, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	item1.Range0(f)
}

func (m *Map3[K0, K1, K2, V]) LoadAndDelete1(k0 K0, k1 K1,
) (value *Map1[K2, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map2[K1, K2, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	return item1.LoadAndDelete0(k1)
}

func (m *Map3[K0, K1, K2, V]) LoadOrStore1(k0 K0, k1 K1,
	v *Map1[K2, V]) (value *Map1[K2, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map2[K1, K2, V]])
	}
	item0 := m.inner

	var item1 *Map2[K1, K2, V]
	item1, _ = item0.LoadOrStore0(k0, &Map2[K1, K2, V]{})

	return item1.LoadOrStore0(k1, v)
}

func (m *Map3[K0, K1, K2, V]) Store1(k0 K0, k1 K1,
	value *Map1[K2, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map2[K1, K2, V]])
	}
	item0 := m.inner

	var item1 *Map2[K1, K2, V]
	item1, _ = item0.LoadOrStore0(k0, &Map2[K1, K2, V]{})

	item1.Store0(k1, value)
}

func (m *Map3[K0, K1, K2, V]) Load2(k0 K0, k1 K1, k2 K2,
) (value V, ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map2[K1, K2, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map1[K2, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	return item2.Load0(k2)
}

func (m *Map3[K0, K1, K2, V]) Delete2(k0 K0, k1 K1, k2 K2,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map2[K1, K2, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map1[K2, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	item2.Delete0(k2)
}

func (m *Map3[K0, K1, K2, V]) Range2(k0 K0, k1 K1,
	f func(K2, V) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map2[K1, K2, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map1[K2, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	item2.Range0(f)
}

func (m *Map3[K0, K1, K2, V]) LoadAndDelete2(k0 K0, k1 K1, k2 K2,
) (value V, loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map2[K1, K2, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map1[K2, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	return item2.LoadAndDelete0(k2)
}

func (m *Map3[K0, K1, K2, V]) LoadOrStore2(k0 K0, k1 K1, k2 K2,
	v V) (value V, loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map2[K1, K2, V]])
	}
	item0 := m.inner

	var item1 *Map2[K1, K2, V]
	item1, _ = item0.LoadOrStore0(k0, &Map2[K1, K2, V]{})

	var item2 *Map1[K2, V]
	item2, _ = item1.LoadOrStore0(k1, &Map1[K2, V]{})

	return item2.LoadOrStore0(k2, v)
}

func (m *Map3[K0, K1, K2, V]) Store2(k0 K0, k1 K1, k2 K2,
	value V) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map2[K1, K2, V]])
	}
	item0 := m.inner

	var item1 *Map2[K1, K2, V]
	item1, _ = item0.LoadOrStore0(k0, &Map2[K1, K2, V]{})

	var item2 *Map1[K2, V]
	item2, _ = item1.LoadOrStore0(k1, &Map1[K2, V]{})

	item2.Store0(k2, value)
}

type Map4[K0 comparable, K1 comparable, K2 comparable, K3 comparable, V any] struct {
	inner *Map1[K0, *Map3[K1, K2, K3, V]]
}

func (m *Map4[K0, K1, K2, K3, V]) Load(k0 K0, k1 K1, k2 K2, k3 K3,
) (value V, ok bool) {
	return m.Load3(k0, k1, k2, k3)
}

func (m *Map4[K0, K1, K2, K3, V]) Delete(k0 K0, k1 K1, k2 K2, k3 K3,
) {
	m.Delete3(k0, k1, k2, k3)
}

func (m *Map4[K0, K1, K2, K3, V]) Range(k0 K0, k1 K1, k2 K2,
	f func(K3, V) bool) {
	m.Range3(k0, k1, k2, f)
}

func (m *Map4[K0, K1, K2, K3, V]) LoadAndDelete(k0 K0, k1 K1, k2 K2, k3 K3,
) (value V, loaded bool) {
	return m.LoadAndDelete3(k0, k1, k2, k3)
}

func (m *Map4[K0, K1, K2, K3, V]) LoadOrStore(k0 K0, k1 K1, k2 K2, k3 K3,
	v V) (value V, loaded bool) {
	return m.LoadOrStore3(k0, k1, k2, k3,
		v)
}

func (m *Map4[K0, K1, K2, K3, V]) Store(k0 K0, k1 K1, k2 K2, k3 K3,
	value V) {
	m.Store3(k0, k1, k2, k3,
		value)
}

func (m *Map4[K0, K1, K2, K3, V]) Load0(k0 K0,
) (value *Map3[K1, K2, K3, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	return item0.Load0(k0)
}

func (m *Map4[K0, K1, K2, K3, V]) Delete0(k0 K0,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	item0.Delete0(k0)
}

func (m *Map4[K0, K1, K2, K3, V]) Range0(
	f func(K0, *Map3[K1, K2, K3, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	item0.Range0(f)
}

func (m *Map4[K0, K1, K2, K3, V]) LoadAndDelete0(k0 K0,
) (value *Map3[K1, K2, K3, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	return item0.LoadAndDelete0(k0)
}

func (m *Map4[K0, K1, K2, K3, V]) LoadOrStore0(k0 K0,
	v *Map3[K1, K2, K3, V]) (value *Map3[K1, K2, K3, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map3[K1, K2, K3, V]])
	}
	item0 := m.inner

	return item0.LoadOrStore0(k0, v)
}

func (m *Map4[K0, K1, K2, K3, V]) Store0(k0 K0,
	value *Map3[K1, K2, K3, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map3[K1, K2, K3, V]])
	}
	item0 := m.inner

	item0.Store0(k0, value)
}

func (m *Map4[K0, K1, K2, K3, V]) Load1(k0 K0, k1 K1,
) (value *Map2[K2, K3, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map3[K1, K2, K3, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	return item1.Load0(k1)
}

func (m *Map4[K0, K1, K2, K3, V]) Delete1(k0 K0, k1 K1,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map3[K1, K2, K3, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	item1.Delete0(k1)
}

func (m *Map4[K0, K1, K2, K3, V]) Range1(k0 K0,
	f func(K1, *Map2[K2, K3, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map3[K1, K2, K3, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	item1.Range0(f)
}

func (m *Map4[K0, K1, K2, K3, V]) LoadAndDelete1(k0 K0, k1 K1,
) (value *Map2[K2, K3, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map3[K1, K2, K3, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	return item1.LoadAndDelete0(k1)
}

func (m *Map4[K0, K1, K2, K3, V]) LoadOrStore1(k0 K0, k1 K1,
	v *Map2[K2, K3, V]) (value *Map2[K2, K3, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map3[K1, K2, K3, V]])
	}
	item0 := m.inner

	var item1 *Map3[K1, K2, K3, V]
	item1, _ = item0.LoadOrStore0(k0, &Map3[K1, K2, K3, V]{})

	return item1.LoadOrStore0(k1, v)
}

func (m *Map4[K0, K1, K2, K3, V]) Store1(k0 K0, k1 K1,
	value *Map2[K2, K3, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map3[K1, K2, K3, V]])
	}
	item0 := m.inner

	var item1 *Map3[K1, K2, K3, V]
	item1, _ = item0.LoadOrStore0(k0, &Map3[K1, K2, K3, V]{})

	item1.Store0(k1, value)
}

func (m *Map4[K0, K1, K2, K3, V]) Load2(k0 K0, k1 K1, k2 K2,
) (value *Map1[K3, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map3[K1, K2, K3, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map2[K2, K3, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	return item2.Load0(k2)
}

func (m *Map4[K0, K1, K2, K3, V]) Delete2(k0 K0, k1 K1, k2 K2,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map3[K1, K2, K3, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map2[K2, K3, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	item2.Delete0(k2)
}

func (m *Map4[K0, K1, K2, K3, V]) Range2(k0 K0, k1 K1,
	f func(K2, *Map1[K3, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map3[K1, K2, K3, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map2[K2, K3, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	item2.Range0(f)
}

func (m *Map4[K0, K1, K2, K3, V]) LoadAndDelete2(k0 K0, k1 K1, k2 K2,
) (value *Map1[K3, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map3[K1, K2, K3, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map2[K2, K3, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	return item2.LoadAndDelete0(k2)
}

func (m *Map4[K0, K1, K2, K3, V]) LoadOrStore2(k0 K0, k1 K1, k2 K2,
	v *Map1[K3, V]) (value *Map1[K3, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map3[K1, K2, K3, V]])
	}
	item0 := m.inner

	var item1 *Map3[K1, K2, K3, V]
	item1, _ = item0.LoadOrStore0(k0, &Map3[K1, K2, K3, V]{})

	var item2 *Map2[K2, K3, V]
	item2, _ = item1.LoadOrStore0(k1, &Map2[K2, K3, V]{})

	return item2.LoadOrStore0(k2, v)
}

func (m *Map4[K0, K1, K2, K3, V]) Store2(k0 K0, k1 K1, k2 K2,
	value *Map1[K3, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map3[K1, K2, K3, V]])
	}
	item0 := m.inner

	var item1 *Map3[K1, K2, K3, V]
	item1, _ = item0.LoadOrStore0(k0, &Map3[K1, K2, K3, V]{})

	var item2 *Map2[K2, K3, V]
	item2, _ = item1.LoadOrStore0(k1, &Map2[K2, K3, V]{})

	item2.Store0(k2, value)
}

func (m *Map4[K0, K1, K2, K3, V]) Load3(k0 K0, k1 K1, k2 K2, k3 K3,
) (value V, ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map3[K1, K2, K3, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map2[K2, K3, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map1[K3, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	return item3.Load0(k3)
}

func (m *Map4[K0, K1, K2, K3, V]) Delete3(k0 K0, k1 K1, k2 K2, k3 K3,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map3[K1, K2, K3, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map2[K2, K3, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map1[K3, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	item3.Delete0(k3)
}

func (m *Map4[K0, K1, K2, K3, V]) Range3(k0 K0, k1 K1, k2 K2,
	f func(K3, V) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map3[K1, K2, K3, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map2[K2, K3, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map1[K3, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	item3.Range0(f)
}

func (m *Map4[K0, K1, K2, K3, V]) LoadAndDelete3(k0 K0, k1 K1, k2 K2, k3 K3,
) (value V, loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map3[K1, K2, K3, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map2[K2, K3, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map1[K3, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	return item3.LoadAndDelete0(k3)
}

func (m *Map4[K0, K1, K2, K3, V]) LoadOrStore3(k0 K0, k1 K1, k2 K2, k3 K3,
	v V) (value V, loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map3[K1, K2, K3, V]])
	}
	item0 := m.inner

	var item1 *Map3[K1, K2, K3, V]
	item1, _ = item0.LoadOrStore0(k0, &Map3[K1, K2, K3, V]{})

	var item2 *Map2[K2, K3, V]
	item2, _ = item1.LoadOrStore0(k1, &Map2[K2, K3, V]{})

	var item3 *Map1[K3, V]
	item3, _ = item2.LoadOrStore0(k2, &Map1[K3, V]{})

	return item3.LoadOrStore0(k3, v)
}

func (m *Map4[K0, K1, K2, K3, V]) Store3(k0 K0, k1 K1, k2 K2, k3 K3,
	value V) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map3[K1, K2, K3, V]])
	}
	item0 := m.inner

	var item1 *Map3[K1, K2, K3, V]
	item1, _ = item0.LoadOrStore0(k0, &Map3[K1, K2, K3, V]{})

	var item2 *Map2[K2, K3, V]
	item2, _ = item1.LoadOrStore0(k1, &Map2[K2, K3, V]{})

	var item3 *Map1[K3, V]
	item3, _ = item2.LoadOrStore0(k2, &Map1[K3, V]{})

	item3.Store0(k3, value)
}

type Map5[K0 comparable, K1 comparable, K2 comparable, K3 comparable, K4 comparable, V any] struct {
	inner *Map1[K0, *Map4[K1, K2, K3, K4, V]]
}

func (m *Map5[K0, K1, K2, K3, K4, V]) Load(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
) (value V, ok bool) {
	return m.Load4(k0, k1, k2, k3, k4)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) Delete(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
) {
	m.Delete4(k0, k1, k2, k3, k4)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) Range(k0 K0, k1 K1, k2 K2, k3 K3,
	f func(K4, V) bool) {
	m.Range4(k0, k1, k2, k3, f)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) LoadAndDelete(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
) (value V, loaded bool) {
	return m.LoadAndDelete4(k0, k1, k2, k3, k4)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) LoadOrStore(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
	v V) (value V, loaded bool) {
	return m.LoadOrStore4(k0, k1, k2, k3, k4,
		v)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) Store(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
	value V) {
	m.Store4(k0, k1, k2, k3, k4,
		value)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) Load0(k0 K0,
) (value *Map4[K1, K2, K3, K4, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	return item0.Load0(k0)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) Delete0(k0 K0,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	item0.Delete0(k0)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) Range0(
	f func(K0, *Map4[K1, K2, K3, K4, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	item0.Range0(f)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) LoadAndDelete0(k0 K0,
) (value *Map4[K1, K2, K3, K4, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	return item0.LoadAndDelete0(k0)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) LoadOrStore0(k0 K0,
	v *Map4[K1, K2, K3, K4, V]) (value *Map4[K1, K2, K3, K4, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map4[K1, K2, K3, K4, V]])
	}
	item0 := m.inner

	return item0.LoadOrStore0(k0, v)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) Store0(k0 K0,
	value *Map4[K1, K2, K3, K4, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map4[K1, K2, K3, K4, V]])
	}
	item0 := m.inner

	item0.Store0(k0, value)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) Load1(k0 K0, k1 K1,
) (value *Map3[K2, K3, K4, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map4[K1, K2, K3, K4, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	return item1.Load0(k1)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) Delete1(k0 K0, k1 K1,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map4[K1, K2, K3, K4, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	item1.Delete0(k1)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) Range1(k0 K0,
	f func(K1, *Map3[K2, K3, K4, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map4[K1, K2, K3, K4, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	item1.Range0(f)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) LoadAndDelete1(k0 K0, k1 K1,
) (value *Map3[K2, K3, K4, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map4[K1, K2, K3, K4, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	return item1.LoadAndDelete0(k1)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) LoadOrStore1(k0 K0, k1 K1,
	v *Map3[K2, K3, K4, V]) (value *Map3[K2, K3, K4, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map4[K1, K2, K3, K4, V]])
	}
	item0 := m.inner

	var item1 *Map4[K1, K2, K3, K4, V]
	item1, _ = item0.LoadOrStore0(k0, &Map4[K1, K2, K3, K4, V]{})

	return item1.LoadOrStore0(k1, v)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) Store1(k0 K0, k1 K1,
	value *Map3[K2, K3, K4, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map4[K1, K2, K3, K4, V]])
	}
	item0 := m.inner

	var item1 *Map4[K1, K2, K3, K4, V]
	item1, _ = item0.LoadOrStore0(k0, &Map4[K1, K2, K3, K4, V]{})

	item1.Store0(k1, value)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) Load2(k0 K0, k1 K1, k2 K2,
) (value *Map2[K3, K4, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map4[K1, K2, K3, K4, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map3[K2, K3, K4, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	return item2.Load0(k2)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) Delete2(k0 K0, k1 K1, k2 K2,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map4[K1, K2, K3, K4, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map3[K2, K3, K4, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	item2.Delete0(k2)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) Range2(k0 K0, k1 K1,
	f func(K2, *Map2[K3, K4, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map4[K1, K2, K3, K4, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map3[K2, K3, K4, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	item2.Range0(f)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) LoadAndDelete2(k0 K0, k1 K1, k2 K2,
) (value *Map2[K3, K4, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map4[K1, K2, K3, K4, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map3[K2, K3, K4, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	return item2.LoadAndDelete0(k2)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) LoadOrStore2(k0 K0, k1 K1, k2 K2,
	v *Map2[K3, K4, V]) (value *Map2[K3, K4, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map4[K1, K2, K3, K4, V]])
	}
	item0 := m.inner

	var item1 *Map4[K1, K2, K3, K4, V]
	item1, _ = item0.LoadOrStore0(k0, &Map4[K1, K2, K3, K4, V]{})

	var item2 *Map3[K2, K3, K4, V]
	item2, _ = item1.LoadOrStore0(k1, &Map3[K2, K3, K4, V]{})

	return item2.LoadOrStore0(k2, v)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) Store2(k0 K0, k1 K1, k2 K2,
	value *Map2[K3, K4, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map4[K1, K2, K3, K4, V]])
	}
	item0 := m.inner

	var item1 *Map4[K1, K2, K3, K4, V]
	item1, _ = item0.LoadOrStore0(k0, &Map4[K1, K2, K3, K4, V]{})

	var item2 *Map3[K2, K3, K4, V]
	item2, _ = item1.LoadOrStore0(k1, &Map3[K2, K3, K4, V]{})

	item2.Store0(k2, value)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) Load3(k0 K0, k1 K1, k2 K2, k3 K3,
) (value *Map1[K4, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map4[K1, K2, K3, K4, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map3[K2, K3, K4, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map2[K3, K4, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	return item3.Load0(k3)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) Delete3(k0 K0, k1 K1, k2 K2, k3 K3,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map4[K1, K2, K3, K4, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map3[K2, K3, K4, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map2[K3, K4, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	item3.Delete0(k3)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) Range3(k0 K0, k1 K1, k2 K2,
	f func(K3, *Map1[K4, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map4[K1, K2, K3, K4, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map3[K2, K3, K4, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map2[K3, K4, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	item3.Range0(f)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) LoadAndDelete3(k0 K0, k1 K1, k2 K2, k3 K3,
) (value *Map1[K4, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map4[K1, K2, K3, K4, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map3[K2, K3, K4, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map2[K3, K4, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	return item3.LoadAndDelete0(k3)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) LoadOrStore3(k0 K0, k1 K1, k2 K2, k3 K3,
	v *Map1[K4, V]) (value *Map1[K4, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map4[K1, K2, K3, K4, V]])
	}
	item0 := m.inner

	var item1 *Map4[K1, K2, K3, K4, V]
	item1, _ = item0.LoadOrStore0(k0, &Map4[K1, K2, K3, K4, V]{})

	var item2 *Map3[K2, K3, K4, V]
	item2, _ = item1.LoadOrStore0(k1, &Map3[K2, K3, K4, V]{})

	var item3 *Map2[K3, K4, V]
	item3, _ = item2.LoadOrStore0(k2, &Map2[K3, K4, V]{})

	return item3.LoadOrStore0(k3, v)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) Store3(k0 K0, k1 K1, k2 K2, k3 K3,
	value *Map1[K4, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map4[K1, K2, K3, K4, V]])
	}
	item0 := m.inner

	var item1 *Map4[K1, K2, K3, K4, V]
	item1, _ = item0.LoadOrStore0(k0, &Map4[K1, K2, K3, K4, V]{})

	var item2 *Map3[K2, K3, K4, V]
	item2, _ = item1.LoadOrStore0(k1, &Map3[K2, K3, K4, V]{})

	var item3 *Map2[K3, K4, V]
	item3, _ = item2.LoadOrStore0(k2, &Map2[K3, K4, V]{})

	item3.Store0(k3, value)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) Load4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
) (value V, ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map4[K1, K2, K3, K4, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map3[K2, K3, K4, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map2[K3, K4, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map1[K4, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	return item4.Load0(k4)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) Delete4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map4[K1, K2, K3, K4, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map3[K2, K3, K4, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map2[K3, K4, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map1[K4, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	item4.Delete0(k4)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) Range4(k0 K0, k1 K1, k2 K2, k3 K3,
	f func(K4, V) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map4[K1, K2, K3, K4, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map3[K2, K3, K4, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map2[K3, K4, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map1[K4, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	item4.Range0(f)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) LoadAndDelete4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
) (value V, loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map4[K1, K2, K3, K4, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map3[K2, K3, K4, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map2[K3, K4, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	var item4 *Map1[K4, V]
	item4, loaded = item3.Load0(k3)
	if !loaded {
		return
	}

	return item4.LoadAndDelete0(k4)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) LoadOrStore4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
	v V) (value V, loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map4[K1, K2, K3, K4, V]])
	}
	item0 := m.inner

	var item1 *Map4[K1, K2, K3, K4, V]
	item1, _ = item0.LoadOrStore0(k0, &Map4[K1, K2, K3, K4, V]{})

	var item2 *Map3[K2, K3, K4, V]
	item2, _ = item1.LoadOrStore0(k1, &Map3[K2, K3, K4, V]{})

	var item3 *Map2[K3, K4, V]
	item3, _ = item2.LoadOrStore0(k2, &Map2[K3, K4, V]{})

	var item4 *Map1[K4, V]
	item4, _ = item3.LoadOrStore0(k3, &Map1[K4, V]{})

	return item4.LoadOrStore0(k4, v)
}

func (m *Map5[K0, K1, K2, K3, K4, V]) Store4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
	value V) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map4[K1, K2, K3, K4, V]])
	}
	item0 := m.inner

	var item1 *Map4[K1, K2, K3, K4, V]
	item1, _ = item0.LoadOrStore0(k0, &Map4[K1, K2, K3, K4, V]{})

	var item2 *Map3[K2, K3, K4, V]
	item2, _ = item1.LoadOrStore0(k1, &Map3[K2, K3, K4, V]{})

	var item3 *Map2[K3, K4, V]
	item3, _ = item2.LoadOrStore0(k2, &Map2[K3, K4, V]{})

	var item4 *Map1[K4, V]
	item4, _ = item3.LoadOrStore0(k3, &Map1[K4, V]{})

	item4.Store0(k4, value)
}

type Map6[K0 comparable, K1 comparable, K2 comparable, K3 comparable, K4 comparable, K5 comparable, V any] struct {
	inner *Map1[K0, *Map5[K1, K2, K3, K4, K5, V]]
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Load(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
) (value V, ok bool) {
	return m.Load5(k0, k1, k2, k3, k4, k5)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Delete(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
) {
	m.Delete5(k0, k1, k2, k3, k4, k5)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Range(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
	f func(K5, V) bool) {
	m.Range5(k0, k1, k2, k3, k4, f)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) LoadAndDelete(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
) (value V, loaded bool) {
	return m.LoadAndDelete5(k0, k1, k2, k3, k4, k5)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) LoadOrStore(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
	v V) (value V, loaded bool) {
	return m.LoadOrStore5(k0, k1, k2, k3, k4, k5,
		v)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Store(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
	value V) {
	m.Store5(k0, k1, k2, k3, k4, k5,
		value)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Load0(k0 K0,
) (value *Map5[K1, K2, K3, K4, K5, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	return item0.Load0(k0)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Delete0(k0 K0,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	item0.Delete0(k0)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Range0(
	f func(K0, *Map5[K1, K2, K3, K4, K5, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	item0.Range0(f)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) LoadAndDelete0(k0 K0,
) (value *Map5[K1, K2, K3, K4, K5, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	return item0.LoadAndDelete0(k0)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) LoadOrStore0(k0 K0,
	v *Map5[K1, K2, K3, K4, K5, V]) (value *Map5[K1, K2, K3, K4, K5, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map5[K1, K2, K3, K4, K5, V]])
	}
	item0 := m.inner

	return item0.LoadOrStore0(k0, v)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Store0(k0 K0,
	value *Map5[K1, K2, K3, K4, K5, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map5[K1, K2, K3, K4, K5, V]])
	}
	item0 := m.inner

	item0.Store0(k0, value)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Load1(k0 K0, k1 K1,
) (value *Map4[K2, K3, K4, K5, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	return item1.Load0(k1)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Delete1(k0 K0, k1 K1,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	item1.Delete0(k1)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Range1(k0 K0,
	f func(K1, *Map4[K2, K3, K4, K5, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	item1.Range0(f)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) LoadAndDelete1(k0 K0, k1 K1,
) (value *Map4[K2, K3, K4, K5, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	return item1.LoadAndDelete0(k1)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) LoadOrStore1(k0 K0, k1 K1,
	v *Map4[K2, K3, K4, K5, V]) (value *Map4[K2, K3, K4, K5, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map5[K1, K2, K3, K4, K5, V]])
	}
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, _ = item0.LoadOrStore0(k0, &Map5[K1, K2, K3, K4, K5, V]{})

	return item1.LoadOrStore0(k1, v)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Store1(k0 K0, k1 K1,
	value *Map4[K2, K3, K4, K5, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map5[K1, K2, K3, K4, K5, V]])
	}
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, _ = item0.LoadOrStore0(k0, &Map5[K1, K2, K3, K4, K5, V]{})

	item1.Store0(k1, value)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Load2(k0 K0, k1 K1, k2 K2,
) (value *Map3[K3, K4, K5, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map4[K2, K3, K4, K5, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	return item2.Load0(k2)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Delete2(k0 K0, k1 K1, k2 K2,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map4[K2, K3, K4, K5, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	item2.Delete0(k2)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Range2(k0 K0, k1 K1,
	f func(K2, *Map3[K3, K4, K5, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map4[K2, K3, K4, K5, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	item2.Range0(f)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) LoadAndDelete2(k0 K0, k1 K1, k2 K2,
) (value *Map3[K3, K4, K5, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map4[K2, K3, K4, K5, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	return item2.LoadAndDelete0(k2)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) LoadOrStore2(k0 K0, k1 K1, k2 K2,
	v *Map3[K3, K4, K5, V]) (value *Map3[K3, K4, K5, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map5[K1, K2, K3, K4, K5, V]])
	}
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, _ = item0.LoadOrStore0(k0, &Map5[K1, K2, K3, K4, K5, V]{})

	var item2 *Map4[K2, K3, K4, K5, V]
	item2, _ = item1.LoadOrStore0(k1, &Map4[K2, K3, K4, K5, V]{})

	return item2.LoadOrStore0(k2, v)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Store2(k0 K0, k1 K1, k2 K2,
	value *Map3[K3, K4, K5, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map5[K1, K2, K3, K4, K5, V]])
	}
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, _ = item0.LoadOrStore0(k0, &Map5[K1, K2, K3, K4, K5, V]{})

	var item2 *Map4[K2, K3, K4, K5, V]
	item2, _ = item1.LoadOrStore0(k1, &Map4[K2, K3, K4, K5, V]{})

	item2.Store0(k2, value)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Load3(k0 K0, k1 K1, k2 K2, k3 K3,
) (value *Map2[K4, K5, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map4[K2, K3, K4, K5, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map3[K3, K4, K5, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	return item3.Load0(k3)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Delete3(k0 K0, k1 K1, k2 K2, k3 K3,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map4[K2, K3, K4, K5, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map3[K3, K4, K5, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	item3.Delete0(k3)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Range3(k0 K0, k1 K1, k2 K2,
	f func(K3, *Map2[K4, K5, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map4[K2, K3, K4, K5, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map3[K3, K4, K5, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	item3.Range0(f)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) LoadAndDelete3(k0 K0, k1 K1, k2 K2, k3 K3,
) (value *Map2[K4, K5, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map4[K2, K3, K4, K5, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map3[K3, K4, K5, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	return item3.LoadAndDelete0(k3)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) LoadOrStore3(k0 K0, k1 K1, k2 K2, k3 K3,
	v *Map2[K4, K5, V]) (value *Map2[K4, K5, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map5[K1, K2, K3, K4, K5, V]])
	}
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, _ = item0.LoadOrStore0(k0, &Map5[K1, K2, K3, K4, K5, V]{})

	var item2 *Map4[K2, K3, K4, K5, V]
	item2, _ = item1.LoadOrStore0(k1, &Map4[K2, K3, K4, K5, V]{})

	var item3 *Map3[K3, K4, K5, V]
	item3, _ = item2.LoadOrStore0(k2, &Map3[K3, K4, K5, V]{})

	return item3.LoadOrStore0(k3, v)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Store3(k0 K0, k1 K1, k2 K2, k3 K3,
	value *Map2[K4, K5, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map5[K1, K2, K3, K4, K5, V]])
	}
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, _ = item0.LoadOrStore0(k0, &Map5[K1, K2, K3, K4, K5, V]{})

	var item2 *Map4[K2, K3, K4, K5, V]
	item2, _ = item1.LoadOrStore0(k1, &Map4[K2, K3, K4, K5, V]{})

	var item3 *Map3[K3, K4, K5, V]
	item3, _ = item2.LoadOrStore0(k2, &Map3[K3, K4, K5, V]{})

	item3.Store0(k3, value)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Load4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
) (value *Map1[K5, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map4[K2, K3, K4, K5, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map3[K3, K4, K5, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map2[K4, K5, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	return item4.Load0(k4)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Delete4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map4[K2, K3, K4, K5, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map3[K3, K4, K5, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map2[K4, K5, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	item4.Delete0(k4)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Range4(k0 K0, k1 K1, k2 K2, k3 K3,
	f func(K4, *Map1[K5, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map4[K2, K3, K4, K5, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map3[K3, K4, K5, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map2[K4, K5, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	item4.Range0(f)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) LoadAndDelete4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
) (value *Map1[K5, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map4[K2, K3, K4, K5, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map3[K3, K4, K5, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	var item4 *Map2[K4, K5, V]
	item4, loaded = item3.Load0(k3)
	if !loaded {
		return
	}

	return item4.LoadAndDelete0(k4)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) LoadOrStore4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
	v *Map1[K5, V]) (value *Map1[K5, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map5[K1, K2, K3, K4, K5, V]])
	}
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, _ = item0.LoadOrStore0(k0, &Map5[K1, K2, K3, K4, K5, V]{})

	var item2 *Map4[K2, K3, K4, K5, V]
	item2, _ = item1.LoadOrStore0(k1, &Map4[K2, K3, K4, K5, V]{})

	var item3 *Map3[K3, K4, K5, V]
	item3, _ = item2.LoadOrStore0(k2, &Map3[K3, K4, K5, V]{})

	var item4 *Map2[K4, K5, V]
	item4, _ = item3.LoadOrStore0(k3, &Map2[K4, K5, V]{})

	return item4.LoadOrStore0(k4, v)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Store4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
	value *Map1[K5, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map5[K1, K2, K3, K4, K5, V]])
	}
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, _ = item0.LoadOrStore0(k0, &Map5[K1, K2, K3, K4, K5, V]{})

	var item2 *Map4[K2, K3, K4, K5, V]
	item2, _ = item1.LoadOrStore0(k1, &Map4[K2, K3, K4, K5, V]{})

	var item3 *Map3[K3, K4, K5, V]
	item3, _ = item2.LoadOrStore0(k2, &Map3[K3, K4, K5, V]{})

	var item4 *Map2[K4, K5, V]
	item4, _ = item3.LoadOrStore0(k3, &Map2[K4, K5, V]{})

	item4.Store0(k4, value)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Load5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
) (value V, ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map4[K2, K3, K4, K5, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map3[K3, K4, K5, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map2[K4, K5, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map1[K5, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	return item5.Load0(k5)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Delete5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map4[K2, K3, K4, K5, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map3[K3, K4, K5, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map2[K4, K5, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map1[K5, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	item5.Delete0(k5)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Range5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
	f func(K5, V) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map4[K2, K3, K4, K5, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map3[K3, K4, K5, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map2[K4, K5, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map1[K5, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	item5.Range0(f)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) LoadAndDelete5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
) (value V, loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map4[K2, K3, K4, K5, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map3[K3, K4, K5, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	var item4 *Map2[K4, K5, V]
	item4, loaded = item3.Load0(k3)
	if !loaded {
		return
	}

	var item5 *Map1[K5, V]
	item5, loaded = item4.Load0(k4)
	if !loaded {
		return
	}

	return item5.LoadAndDelete0(k5)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) LoadOrStore5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
	v V) (value V, loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map5[K1, K2, K3, K4, K5, V]])
	}
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, _ = item0.LoadOrStore0(k0, &Map5[K1, K2, K3, K4, K5, V]{})

	var item2 *Map4[K2, K3, K4, K5, V]
	item2, _ = item1.LoadOrStore0(k1, &Map4[K2, K3, K4, K5, V]{})

	var item3 *Map3[K3, K4, K5, V]
	item3, _ = item2.LoadOrStore0(k2, &Map3[K3, K4, K5, V]{})

	var item4 *Map2[K4, K5, V]
	item4, _ = item3.LoadOrStore0(k3, &Map2[K4, K5, V]{})

	var item5 *Map1[K5, V]
	item5, _ = item4.LoadOrStore0(k4, &Map1[K5, V]{})

	return item5.LoadOrStore0(k5, v)
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Store5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
	value V) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map5[K1, K2, K3, K4, K5, V]])
	}
	item0 := m.inner

	var item1 *Map5[K1, K2, K3, K4, K5, V]
	item1, _ = item0.LoadOrStore0(k0, &Map5[K1, K2, K3, K4, K5, V]{})

	var item2 *Map4[K2, K3, K4, K5, V]
	item2, _ = item1.LoadOrStore0(k1, &Map4[K2, K3, K4, K5, V]{})

	var item3 *Map3[K3, K4, K5, V]
	item3, _ = item2.LoadOrStore0(k2, &Map3[K3, K4, K5, V]{})

	var item4 *Map2[K4, K5, V]
	item4, _ = item3.LoadOrStore0(k3, &Map2[K4, K5, V]{})

	var item5 *Map1[K5, V]
	item5, _ = item4.LoadOrStore0(k4, &Map1[K5, V]{})

	item5.Store0(k5, value)
}

type Map7[K0 comparable, K1 comparable, K2 comparable, K3 comparable, K4 comparable, K5 comparable, K6 comparable, V any] struct {
	inner *Map1[K0, *Map6[K1, K2, K3, K4, K5, K6, V]]
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Load(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
) (value V, ok bool) {
	return m.Load6(k0, k1, k2, k3, k4, k5, k6)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Delete(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
) {
	m.Delete6(k0, k1, k2, k3, k4, k5, k6)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Range(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
	f func(K6, V) bool) {
	m.Range6(k0, k1, k2, k3, k4, k5, f)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) LoadAndDelete(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
) (value V, loaded bool) {
	return m.LoadAndDelete6(k0, k1, k2, k3, k4, k5, k6)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) LoadOrStore(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
	v V) (value V, loaded bool) {
	return m.LoadOrStore6(k0, k1, k2, k3, k4, k5, k6,
		v)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Store(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
	value V) {
	m.Store6(k0, k1, k2, k3, k4, k5, k6,
		value)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Load0(k0 K0,
) (value *Map6[K1, K2, K3, K4, K5, K6, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	return item0.Load0(k0)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Delete0(k0 K0,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	item0.Delete0(k0)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Range0(
	f func(K0, *Map6[K1, K2, K3, K4, K5, K6, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	item0.Range0(f)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) LoadAndDelete0(k0 K0,
) (value *Map6[K1, K2, K3, K4, K5, K6, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	return item0.LoadAndDelete0(k0)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) LoadOrStore0(k0 K0,
	v *Map6[K1, K2, K3, K4, K5, K6, V]) (value *Map6[K1, K2, K3, K4, K5, K6, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map6[K1, K2, K3, K4, K5, K6, V]])
	}
	item0 := m.inner

	return item0.LoadOrStore0(k0, v)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Store0(k0 K0,
	value *Map6[K1, K2, K3, K4, K5, K6, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map6[K1, K2, K3, K4, K5, K6, V]])
	}
	item0 := m.inner

	item0.Store0(k0, value)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Load1(k0 K0, k1 K1,
) (value *Map5[K2, K3, K4, K5, K6, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	return item1.Load0(k1)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Delete1(k0 K0, k1 K1,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	item1.Delete0(k1)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Range1(k0 K0,
	f func(K1, *Map5[K2, K3, K4, K5, K6, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	item1.Range0(f)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) LoadAndDelete1(k0 K0, k1 K1,
) (value *Map5[K2, K3, K4, K5, K6, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	return item1.LoadAndDelete0(k1)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) LoadOrStore1(k0 K0, k1 K1,
	v *Map5[K2, K3, K4, K5, K6, V]) (value *Map5[K2, K3, K4, K5, K6, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map6[K1, K2, K3, K4, K5, K6, V]])
	}
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, _ = item0.LoadOrStore0(k0, &Map6[K1, K2, K3, K4, K5, K6, V]{})

	return item1.LoadOrStore0(k1, v)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Store1(k0 K0, k1 K1,
	value *Map5[K2, K3, K4, K5, K6, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map6[K1, K2, K3, K4, K5, K6, V]])
	}
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, _ = item0.LoadOrStore0(k0, &Map6[K1, K2, K3, K4, K5, K6, V]{})

	item1.Store0(k1, value)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Load2(k0 K0, k1 K1, k2 K2,
) (value *Map4[K3, K4, K5, K6, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	return item2.Load0(k2)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Delete2(k0 K0, k1 K1, k2 K2,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	item2.Delete0(k2)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Range2(k0 K0, k1 K1,
	f func(K2, *Map4[K3, K4, K5, K6, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	item2.Range0(f)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) LoadAndDelete2(k0 K0, k1 K1, k2 K2,
) (value *Map4[K3, K4, K5, K6, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	return item2.LoadAndDelete0(k2)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) LoadOrStore2(k0 K0, k1 K1, k2 K2,
	v *Map4[K3, K4, K5, K6, V]) (value *Map4[K3, K4, K5, K6, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map6[K1, K2, K3, K4, K5, K6, V]])
	}
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, _ = item0.LoadOrStore0(k0, &Map6[K1, K2, K3, K4, K5, K6, V]{})

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, _ = item1.LoadOrStore0(k1, &Map5[K2, K3, K4, K5, K6, V]{})

	return item2.LoadOrStore0(k2, v)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Store2(k0 K0, k1 K1, k2 K2,
	value *Map4[K3, K4, K5, K6, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map6[K1, K2, K3, K4, K5, K6, V]])
	}
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, _ = item0.LoadOrStore0(k0, &Map6[K1, K2, K3, K4, K5, K6, V]{})

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, _ = item1.LoadOrStore0(k1, &Map5[K2, K3, K4, K5, K6, V]{})

	item2.Store0(k2, value)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Load3(k0 K0, k1 K1, k2 K2, k3 K3,
) (value *Map3[K4, K5, K6, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map4[K3, K4, K5, K6, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	return item3.Load0(k3)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Delete3(k0 K0, k1 K1, k2 K2, k3 K3,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map4[K3, K4, K5, K6, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	item3.Delete0(k3)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Range3(k0 K0, k1 K1, k2 K2,
	f func(K3, *Map3[K4, K5, K6, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map4[K3, K4, K5, K6, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	item3.Range0(f)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) LoadAndDelete3(k0 K0, k1 K1, k2 K2, k3 K3,
) (value *Map3[K4, K5, K6, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map4[K3, K4, K5, K6, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	return item3.LoadAndDelete0(k3)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) LoadOrStore3(k0 K0, k1 K1, k2 K2, k3 K3,
	v *Map3[K4, K5, K6, V]) (value *Map3[K4, K5, K6, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map6[K1, K2, K3, K4, K5, K6, V]])
	}
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, _ = item0.LoadOrStore0(k0, &Map6[K1, K2, K3, K4, K5, K6, V]{})

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, _ = item1.LoadOrStore0(k1, &Map5[K2, K3, K4, K5, K6, V]{})

	var item3 *Map4[K3, K4, K5, K6, V]
	item3, _ = item2.LoadOrStore0(k2, &Map4[K3, K4, K5, K6, V]{})

	return item3.LoadOrStore0(k3, v)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Store3(k0 K0, k1 K1, k2 K2, k3 K3,
	value *Map3[K4, K5, K6, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map6[K1, K2, K3, K4, K5, K6, V]])
	}
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, _ = item0.LoadOrStore0(k0, &Map6[K1, K2, K3, K4, K5, K6, V]{})

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, _ = item1.LoadOrStore0(k1, &Map5[K2, K3, K4, K5, K6, V]{})

	var item3 *Map4[K3, K4, K5, K6, V]
	item3, _ = item2.LoadOrStore0(k2, &Map4[K3, K4, K5, K6, V]{})

	item3.Store0(k3, value)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Load4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
) (value *Map2[K5, K6, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map4[K3, K4, K5, K6, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map3[K4, K5, K6, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	return item4.Load0(k4)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Delete4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map4[K3, K4, K5, K6, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map3[K4, K5, K6, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	item4.Delete0(k4)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Range4(k0 K0, k1 K1, k2 K2, k3 K3,
	f func(K4, *Map2[K5, K6, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map4[K3, K4, K5, K6, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map3[K4, K5, K6, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	item4.Range0(f)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) LoadAndDelete4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
) (value *Map2[K5, K6, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map4[K3, K4, K5, K6, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	var item4 *Map3[K4, K5, K6, V]
	item4, loaded = item3.Load0(k3)
	if !loaded {
		return
	}

	return item4.LoadAndDelete0(k4)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) LoadOrStore4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
	v *Map2[K5, K6, V]) (value *Map2[K5, K6, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map6[K1, K2, K3, K4, K5, K6, V]])
	}
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, _ = item0.LoadOrStore0(k0, &Map6[K1, K2, K3, K4, K5, K6, V]{})

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, _ = item1.LoadOrStore0(k1, &Map5[K2, K3, K4, K5, K6, V]{})

	var item3 *Map4[K3, K4, K5, K6, V]
	item3, _ = item2.LoadOrStore0(k2, &Map4[K3, K4, K5, K6, V]{})

	var item4 *Map3[K4, K5, K6, V]
	item4, _ = item3.LoadOrStore0(k3, &Map3[K4, K5, K6, V]{})

	return item4.LoadOrStore0(k4, v)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Store4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
	value *Map2[K5, K6, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map6[K1, K2, K3, K4, K5, K6, V]])
	}
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, _ = item0.LoadOrStore0(k0, &Map6[K1, K2, K3, K4, K5, K6, V]{})

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, _ = item1.LoadOrStore0(k1, &Map5[K2, K3, K4, K5, K6, V]{})

	var item3 *Map4[K3, K4, K5, K6, V]
	item3, _ = item2.LoadOrStore0(k2, &Map4[K3, K4, K5, K6, V]{})

	var item4 *Map3[K4, K5, K6, V]
	item4, _ = item3.LoadOrStore0(k3, &Map3[K4, K5, K6, V]{})

	item4.Store0(k4, value)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Load5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
) (value *Map1[K6, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map4[K3, K4, K5, K6, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map3[K4, K5, K6, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map2[K5, K6, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	return item5.Load0(k5)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Delete5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map4[K3, K4, K5, K6, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map3[K4, K5, K6, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map2[K5, K6, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	item5.Delete0(k5)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Range5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
	f func(K5, *Map1[K6, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map4[K3, K4, K5, K6, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map3[K4, K5, K6, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map2[K5, K6, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	item5.Range0(f)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) LoadAndDelete5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
) (value *Map1[K6, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map4[K3, K4, K5, K6, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	var item4 *Map3[K4, K5, K6, V]
	item4, loaded = item3.Load0(k3)
	if !loaded {
		return
	}

	var item5 *Map2[K5, K6, V]
	item5, loaded = item4.Load0(k4)
	if !loaded {
		return
	}

	return item5.LoadAndDelete0(k5)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) LoadOrStore5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
	v *Map1[K6, V]) (value *Map1[K6, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map6[K1, K2, K3, K4, K5, K6, V]])
	}
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, _ = item0.LoadOrStore0(k0, &Map6[K1, K2, K3, K4, K5, K6, V]{})

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, _ = item1.LoadOrStore0(k1, &Map5[K2, K3, K4, K5, K6, V]{})

	var item3 *Map4[K3, K4, K5, K6, V]
	item3, _ = item2.LoadOrStore0(k2, &Map4[K3, K4, K5, K6, V]{})

	var item4 *Map3[K4, K5, K6, V]
	item4, _ = item3.LoadOrStore0(k3, &Map3[K4, K5, K6, V]{})

	var item5 *Map2[K5, K6, V]
	item5, _ = item4.LoadOrStore0(k4, &Map2[K5, K6, V]{})

	return item5.LoadOrStore0(k5, v)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Store5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
	value *Map1[K6, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map6[K1, K2, K3, K4, K5, K6, V]])
	}
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, _ = item0.LoadOrStore0(k0, &Map6[K1, K2, K3, K4, K5, K6, V]{})

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, _ = item1.LoadOrStore0(k1, &Map5[K2, K3, K4, K5, K6, V]{})

	var item3 *Map4[K3, K4, K5, K6, V]
	item3, _ = item2.LoadOrStore0(k2, &Map4[K3, K4, K5, K6, V]{})

	var item4 *Map3[K4, K5, K6, V]
	item4, _ = item3.LoadOrStore0(k3, &Map3[K4, K5, K6, V]{})

	var item5 *Map2[K5, K6, V]
	item5, _ = item4.LoadOrStore0(k4, &Map2[K5, K6, V]{})

	item5.Store0(k5, value)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Load6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
) (value V, ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map4[K3, K4, K5, K6, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map3[K4, K5, K6, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map2[K5, K6, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map1[K6, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	return item6.Load0(k6)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Delete6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map4[K3, K4, K5, K6, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map3[K4, K5, K6, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map2[K5, K6, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map1[K6, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	item6.Delete0(k6)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Range6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
	f func(K6, V) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map4[K3, K4, K5, K6, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map3[K4, K5, K6, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map2[K5, K6, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map1[K6, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	item6.Range0(f)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) LoadAndDelete6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
) (value V, loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map4[K3, K4, K5, K6, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	var item4 *Map3[K4, K5, K6, V]
	item4, loaded = item3.Load0(k3)
	if !loaded {
		return
	}

	var item5 *Map2[K5, K6, V]
	item5, loaded = item4.Load0(k4)
	if !loaded {
		return
	}

	var item6 *Map1[K6, V]
	item6, loaded = item5.Load0(k5)
	if !loaded {
		return
	}

	return item6.LoadAndDelete0(k6)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) LoadOrStore6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
	v V) (value V, loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map6[K1, K2, K3, K4, K5, K6, V]])
	}
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, _ = item0.LoadOrStore0(k0, &Map6[K1, K2, K3, K4, K5, K6, V]{})

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, _ = item1.LoadOrStore0(k1, &Map5[K2, K3, K4, K5, K6, V]{})

	var item3 *Map4[K3, K4, K5, K6, V]
	item3, _ = item2.LoadOrStore0(k2, &Map4[K3, K4, K5, K6, V]{})

	var item4 *Map3[K4, K5, K6, V]
	item4, _ = item3.LoadOrStore0(k3, &Map3[K4, K5, K6, V]{})

	var item5 *Map2[K5, K6, V]
	item5, _ = item4.LoadOrStore0(k4, &Map2[K5, K6, V]{})

	var item6 *Map1[K6, V]
	item6, _ = item5.LoadOrStore0(k5, &Map1[K6, V]{})

	return item6.LoadOrStore0(k6, v)
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Store6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
	value V) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map6[K1, K2, K3, K4, K5, K6, V]])
	}
	item0 := m.inner

	var item1 *Map6[K1, K2, K3, K4, K5, K6, V]
	item1, _ = item0.LoadOrStore0(k0, &Map6[K1, K2, K3, K4, K5, K6, V]{})

	var item2 *Map5[K2, K3, K4, K5, K6, V]
	item2, _ = item1.LoadOrStore0(k1, &Map5[K2, K3, K4, K5, K6, V]{})

	var item3 *Map4[K3, K4, K5, K6, V]
	item3, _ = item2.LoadOrStore0(k2, &Map4[K3, K4, K5, K6, V]{})

	var item4 *Map3[K4, K5, K6, V]
	item4, _ = item3.LoadOrStore0(k3, &Map3[K4, K5, K6, V]{})

	var item5 *Map2[K5, K6, V]
	item5, _ = item4.LoadOrStore0(k4, &Map2[K5, K6, V]{})

	var item6 *Map1[K6, V]
	item6, _ = item5.LoadOrStore0(k5, &Map1[K6, V]{})

	item6.Store0(k6, value)
}

type Map8[K0 comparable, K1 comparable, K2 comparable, K3 comparable, K4 comparable, K5 comparable, K6 comparable, K7 comparable, V any] struct {
	inner *Map1[K0, *Map7[K1, K2, K3, K4, K5, K6, K7, V]]
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Load(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
) (value V, ok bool) {
	return m.Load7(k0, k1, k2, k3, k4, k5, k6, k7)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Delete(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
) {
	m.Delete7(k0, k1, k2, k3, k4, k5, k6, k7)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Range(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
	f func(K7, V) bool) {
	m.Range7(k0, k1, k2, k3, k4, k5, k6, f)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) LoadAndDelete(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
) (value V, loaded bool) {
	return m.LoadAndDelete7(k0, k1, k2, k3, k4, k5, k6, k7)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) LoadOrStore(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
	v V) (value V, loaded bool) {
	return m.LoadOrStore7(k0, k1, k2, k3, k4, k5, k6, k7,
		v)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Store(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
	value V) {
	m.Store7(k0, k1, k2, k3, k4, k5, k6, k7,
		value)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Load0(k0 K0,
) (value *Map7[K1, K2, K3, K4, K5, K6, K7, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	return item0.Load0(k0)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Delete0(k0 K0,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	item0.Delete0(k0)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Range0(
	f func(K0, *Map7[K1, K2, K3, K4, K5, K6, K7, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	item0.Range0(f)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) LoadAndDelete0(k0 K0,
) (value *Map7[K1, K2, K3, K4, K5, K6, K7, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	return item0.LoadAndDelete0(k0)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) LoadOrStore0(k0 K0,
	v *Map7[K1, K2, K3, K4, K5, K6, K7, V]) (value *Map7[K1, K2, K3, K4, K5, K6, K7, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map7[K1, K2, K3, K4, K5, K6, K7, V]])
	}
	item0 := m.inner

	return item0.LoadOrStore0(k0, v)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Store0(k0 K0,
	value *Map7[K1, K2, K3, K4, K5, K6, K7, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map7[K1, K2, K3, K4, K5, K6, K7, V]])
	}
	item0 := m.inner

	item0.Store0(k0, value)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Load1(k0 K0, k1 K1,
) (value *Map6[K2, K3, K4, K5, K6, K7, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	return item1.Load0(k1)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Delete1(k0 K0, k1 K1,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	item1.Delete0(k1)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Range1(k0 K0,
	f func(K1, *Map6[K2, K3, K4, K5, K6, K7, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	item1.Range0(f)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) LoadAndDelete1(k0 K0, k1 K1,
) (value *Map6[K2, K3, K4, K5, K6, K7, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	return item1.LoadAndDelete0(k1)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) LoadOrStore1(k0 K0, k1 K1,
	v *Map6[K2, K3, K4, K5, K6, K7, V]) (value *Map6[K2, K3, K4, K5, K6, K7, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map7[K1, K2, K3, K4, K5, K6, K7, V]])
	}
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, _ = item0.LoadOrStore0(k0, &Map7[K1, K2, K3, K4, K5, K6, K7, V]{})

	return item1.LoadOrStore0(k1, v)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Store1(k0 K0, k1 K1,
	value *Map6[K2, K3, K4, K5, K6, K7, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map7[K1, K2, K3, K4, K5, K6, K7, V]])
	}
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, _ = item0.LoadOrStore0(k0, &Map7[K1, K2, K3, K4, K5, K6, K7, V]{})

	item1.Store0(k1, value)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Load2(k0 K0, k1 K1, k2 K2,
) (value *Map5[K3, K4, K5, K6, K7, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	return item2.Load0(k2)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Delete2(k0 K0, k1 K1, k2 K2,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	item2.Delete0(k2)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Range2(k0 K0, k1 K1,
	f func(K2, *Map5[K3, K4, K5, K6, K7, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	item2.Range0(f)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) LoadAndDelete2(k0 K0, k1 K1, k2 K2,
) (value *Map5[K3, K4, K5, K6, K7, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	return item2.LoadAndDelete0(k2)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) LoadOrStore2(k0 K0, k1 K1, k2 K2,
	v *Map5[K3, K4, K5, K6, K7, V]) (value *Map5[K3, K4, K5, K6, K7, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map7[K1, K2, K3, K4, K5, K6, K7, V]])
	}
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, _ = item0.LoadOrStore0(k0, &Map7[K1, K2, K3, K4, K5, K6, K7, V]{})

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, _ = item1.LoadOrStore0(k1, &Map6[K2, K3, K4, K5, K6, K7, V]{})

	return item2.LoadOrStore0(k2, v)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Store2(k0 K0, k1 K1, k2 K2,
	value *Map5[K3, K4, K5, K6, K7, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map7[K1, K2, K3, K4, K5, K6, K7, V]])
	}
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, _ = item0.LoadOrStore0(k0, &Map7[K1, K2, K3, K4, K5, K6, K7, V]{})

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, _ = item1.LoadOrStore0(k1, &Map6[K2, K3, K4, K5, K6, K7, V]{})

	item2.Store0(k2, value)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Load3(k0 K0, k1 K1, k2 K2, k3 K3,
) (value *Map4[K4, K5, K6, K7, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	return item3.Load0(k3)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Delete3(k0 K0, k1 K1, k2 K2, k3 K3,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	item3.Delete0(k3)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Range3(k0 K0, k1 K1, k2 K2,
	f func(K3, *Map4[K4, K5, K6, K7, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	item3.Range0(f)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) LoadAndDelete3(k0 K0, k1 K1, k2 K2, k3 K3,
) (value *Map4[K4, K5, K6, K7, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	return item3.LoadAndDelete0(k3)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) LoadOrStore3(k0 K0, k1 K1, k2 K2, k3 K3,
	v *Map4[K4, K5, K6, K7, V]) (value *Map4[K4, K5, K6, K7, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map7[K1, K2, K3, K4, K5, K6, K7, V]])
	}
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, _ = item0.LoadOrStore0(k0, &Map7[K1, K2, K3, K4, K5, K6, K7, V]{})

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, _ = item1.LoadOrStore0(k1, &Map6[K2, K3, K4, K5, K6, K7, V]{})

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, _ = item2.LoadOrStore0(k2, &Map5[K3, K4, K5, K6, K7, V]{})

	return item3.LoadOrStore0(k3, v)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Store3(k0 K0, k1 K1, k2 K2, k3 K3,
	value *Map4[K4, K5, K6, K7, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map7[K1, K2, K3, K4, K5, K6, K7, V]])
	}
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, _ = item0.LoadOrStore0(k0, &Map7[K1, K2, K3, K4, K5, K6, K7, V]{})

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, _ = item1.LoadOrStore0(k1, &Map6[K2, K3, K4, K5, K6, K7, V]{})

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, _ = item2.LoadOrStore0(k2, &Map5[K3, K4, K5, K6, K7, V]{})

	item3.Store0(k3, value)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Load4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
) (value *Map3[K5, K6, K7, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map4[K4, K5, K6, K7, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	return item4.Load0(k4)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Delete4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map4[K4, K5, K6, K7, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	item4.Delete0(k4)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Range4(k0 K0, k1 K1, k2 K2, k3 K3,
	f func(K4, *Map3[K5, K6, K7, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map4[K4, K5, K6, K7, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	item4.Range0(f)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) LoadAndDelete4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
) (value *Map3[K5, K6, K7, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	var item4 *Map4[K4, K5, K6, K7, V]
	item4, loaded = item3.Load0(k3)
	if !loaded {
		return
	}

	return item4.LoadAndDelete0(k4)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) LoadOrStore4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
	v *Map3[K5, K6, K7, V]) (value *Map3[K5, K6, K7, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map7[K1, K2, K3, K4, K5, K6, K7, V]])
	}
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, _ = item0.LoadOrStore0(k0, &Map7[K1, K2, K3, K4, K5, K6, K7, V]{})

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, _ = item1.LoadOrStore0(k1, &Map6[K2, K3, K4, K5, K6, K7, V]{})

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, _ = item2.LoadOrStore0(k2, &Map5[K3, K4, K5, K6, K7, V]{})

	var item4 *Map4[K4, K5, K6, K7, V]
	item4, _ = item3.LoadOrStore0(k3, &Map4[K4, K5, K6, K7, V]{})

	return item4.LoadOrStore0(k4, v)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Store4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
	value *Map3[K5, K6, K7, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map7[K1, K2, K3, K4, K5, K6, K7, V]])
	}
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, _ = item0.LoadOrStore0(k0, &Map7[K1, K2, K3, K4, K5, K6, K7, V]{})

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, _ = item1.LoadOrStore0(k1, &Map6[K2, K3, K4, K5, K6, K7, V]{})

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, _ = item2.LoadOrStore0(k2, &Map5[K3, K4, K5, K6, K7, V]{})

	var item4 *Map4[K4, K5, K6, K7, V]
	item4, _ = item3.LoadOrStore0(k3, &Map4[K4, K5, K6, K7, V]{})

	item4.Store0(k4, value)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Load5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
) (value *Map2[K6, K7, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map4[K4, K5, K6, K7, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map3[K5, K6, K7, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	return item5.Load0(k5)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Delete5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map4[K4, K5, K6, K7, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map3[K5, K6, K7, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	item5.Delete0(k5)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Range5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
	f func(K5, *Map2[K6, K7, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map4[K4, K5, K6, K7, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map3[K5, K6, K7, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	item5.Range0(f)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) LoadAndDelete5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
) (value *Map2[K6, K7, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	var item4 *Map4[K4, K5, K6, K7, V]
	item4, loaded = item3.Load0(k3)
	if !loaded {
		return
	}

	var item5 *Map3[K5, K6, K7, V]
	item5, loaded = item4.Load0(k4)
	if !loaded {
		return
	}

	return item5.LoadAndDelete0(k5)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) LoadOrStore5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
	v *Map2[K6, K7, V]) (value *Map2[K6, K7, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map7[K1, K2, K3, K4, K5, K6, K7, V]])
	}
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, _ = item0.LoadOrStore0(k0, &Map7[K1, K2, K3, K4, K5, K6, K7, V]{})

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, _ = item1.LoadOrStore0(k1, &Map6[K2, K3, K4, K5, K6, K7, V]{})

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, _ = item2.LoadOrStore0(k2, &Map5[K3, K4, K5, K6, K7, V]{})

	var item4 *Map4[K4, K5, K6, K7, V]
	item4, _ = item3.LoadOrStore0(k3, &Map4[K4, K5, K6, K7, V]{})

	var item5 *Map3[K5, K6, K7, V]
	item5, _ = item4.LoadOrStore0(k4, &Map3[K5, K6, K7, V]{})

	return item5.LoadOrStore0(k5, v)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Store5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
	value *Map2[K6, K7, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map7[K1, K2, K3, K4, K5, K6, K7, V]])
	}
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, _ = item0.LoadOrStore0(k0, &Map7[K1, K2, K3, K4, K5, K6, K7, V]{})

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, _ = item1.LoadOrStore0(k1, &Map6[K2, K3, K4, K5, K6, K7, V]{})

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, _ = item2.LoadOrStore0(k2, &Map5[K3, K4, K5, K6, K7, V]{})

	var item4 *Map4[K4, K5, K6, K7, V]
	item4, _ = item3.LoadOrStore0(k3, &Map4[K4, K5, K6, K7, V]{})

	var item5 *Map3[K5, K6, K7, V]
	item5, _ = item4.LoadOrStore0(k4, &Map3[K5, K6, K7, V]{})

	item5.Store0(k5, value)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Load6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
) (value *Map1[K7, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map4[K4, K5, K6, K7, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map3[K5, K6, K7, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map2[K6, K7, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	return item6.Load0(k6)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Delete6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map4[K4, K5, K6, K7, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map3[K5, K6, K7, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map2[K6, K7, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	item6.Delete0(k6)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Range6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
	f func(K6, *Map1[K7, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map4[K4, K5, K6, K7, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map3[K5, K6, K7, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map2[K6, K7, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	item6.Range0(f)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) LoadAndDelete6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
) (value *Map1[K7, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	var item4 *Map4[K4, K5, K6, K7, V]
	item4, loaded = item3.Load0(k3)
	if !loaded {
		return
	}

	var item5 *Map3[K5, K6, K7, V]
	item5, loaded = item4.Load0(k4)
	if !loaded {
		return
	}

	var item6 *Map2[K6, K7, V]
	item6, loaded = item5.Load0(k5)
	if !loaded {
		return
	}

	return item6.LoadAndDelete0(k6)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) LoadOrStore6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
	v *Map1[K7, V]) (value *Map1[K7, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map7[K1, K2, K3, K4, K5, K6, K7, V]])
	}
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, _ = item0.LoadOrStore0(k0, &Map7[K1, K2, K3, K4, K5, K6, K7, V]{})

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, _ = item1.LoadOrStore0(k1, &Map6[K2, K3, K4, K5, K6, K7, V]{})

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, _ = item2.LoadOrStore0(k2, &Map5[K3, K4, K5, K6, K7, V]{})

	var item4 *Map4[K4, K5, K6, K7, V]
	item4, _ = item3.LoadOrStore0(k3, &Map4[K4, K5, K6, K7, V]{})

	var item5 *Map3[K5, K6, K7, V]
	item5, _ = item4.LoadOrStore0(k4, &Map3[K5, K6, K7, V]{})

	var item6 *Map2[K6, K7, V]
	item6, _ = item5.LoadOrStore0(k5, &Map2[K6, K7, V]{})

	return item6.LoadOrStore0(k6, v)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Store6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
	value *Map1[K7, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map7[K1, K2, K3, K4, K5, K6, K7, V]])
	}
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, _ = item0.LoadOrStore0(k0, &Map7[K1, K2, K3, K4, K5, K6, K7, V]{})

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, _ = item1.LoadOrStore0(k1, &Map6[K2, K3, K4, K5, K6, K7, V]{})

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, _ = item2.LoadOrStore0(k2, &Map5[K3, K4, K5, K6, K7, V]{})

	var item4 *Map4[K4, K5, K6, K7, V]
	item4, _ = item3.LoadOrStore0(k3, &Map4[K4, K5, K6, K7, V]{})

	var item5 *Map3[K5, K6, K7, V]
	item5, _ = item4.LoadOrStore0(k4, &Map3[K5, K6, K7, V]{})

	var item6 *Map2[K6, K7, V]
	item6, _ = item5.LoadOrStore0(k5, &Map2[K6, K7, V]{})

	item6.Store0(k6, value)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Load7(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
) (value V, ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map4[K4, K5, K6, K7, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map3[K5, K6, K7, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map2[K6, K7, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map1[K7, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	return item7.Load0(k7)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Delete7(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map4[K4, K5, K6, K7, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map3[K5, K6, K7, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map2[K6, K7, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map1[K7, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	item7.Delete0(k7)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Range7(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
	f func(K7, V) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map4[K4, K5, K6, K7, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map3[K5, K6, K7, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map2[K6, K7, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map1[K7, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	item7.Range0(f)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) LoadAndDelete7(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
) (value V, loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	var item4 *Map4[K4, K5, K6, K7, V]
	item4, loaded = item3.Load0(k3)
	if !loaded {
		return
	}

	var item5 *Map3[K5, K6, K7, V]
	item5, loaded = item4.Load0(k4)
	if !loaded {
		return
	}

	var item6 *Map2[K6, K7, V]
	item6, loaded = item5.Load0(k5)
	if !loaded {
		return
	}

	var item7 *Map1[K7, V]
	item7, loaded = item6.Load0(k6)
	if !loaded {
		return
	}

	return item7.LoadAndDelete0(k7)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) LoadOrStore7(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
	v V) (value V, loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map7[K1, K2, K3, K4, K5, K6, K7, V]])
	}
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, _ = item0.LoadOrStore0(k0, &Map7[K1, K2, K3, K4, K5, K6, K7, V]{})

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, _ = item1.LoadOrStore0(k1, &Map6[K2, K3, K4, K5, K6, K7, V]{})

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, _ = item2.LoadOrStore0(k2, &Map5[K3, K4, K5, K6, K7, V]{})

	var item4 *Map4[K4, K5, K6, K7, V]
	item4, _ = item3.LoadOrStore0(k3, &Map4[K4, K5, K6, K7, V]{})

	var item5 *Map3[K5, K6, K7, V]
	item5, _ = item4.LoadOrStore0(k4, &Map3[K5, K6, K7, V]{})

	var item6 *Map2[K6, K7, V]
	item6, _ = item5.LoadOrStore0(k5, &Map2[K6, K7, V]{})

	var item7 *Map1[K7, V]
	item7, _ = item6.LoadOrStore0(k6, &Map1[K7, V]{})

	return item7.LoadOrStore0(k7, v)
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Store7(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
	value V) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map7[K1, K2, K3, K4, K5, K6, K7, V]])
	}
	item0 := m.inner

	var item1 *Map7[K1, K2, K3, K4, K5, K6, K7, V]
	item1, _ = item0.LoadOrStore0(k0, &Map7[K1, K2, K3, K4, K5, K6, K7, V]{})

	var item2 *Map6[K2, K3, K4, K5, K6, K7, V]
	item2, _ = item1.LoadOrStore0(k1, &Map6[K2, K3, K4, K5, K6, K7, V]{})

	var item3 *Map5[K3, K4, K5, K6, K7, V]
	item3, _ = item2.LoadOrStore0(k2, &Map5[K3, K4, K5, K6, K7, V]{})

	var item4 *Map4[K4, K5, K6, K7, V]
	item4, _ = item3.LoadOrStore0(k3, &Map4[K4, K5, K6, K7, V]{})

	var item5 *Map3[K5, K6, K7, V]
	item5, _ = item4.LoadOrStore0(k4, &Map3[K5, K6, K7, V]{})

	var item6 *Map2[K6, K7, V]
	item6, _ = item5.LoadOrStore0(k5, &Map2[K6, K7, V]{})

	var item7 *Map1[K7, V]
	item7, _ = item6.LoadOrStore0(k6, &Map1[K7, V]{})

	item7.Store0(k7, value)
}

type Map9[K0 comparable, K1 comparable, K2 comparable, K3 comparable, K4 comparable, K5 comparable, K6 comparable, K7 comparable, K8 comparable, V any] struct {
	inner *Map1[K0, *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]]
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Load(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8,
) (value V, ok bool) {
	return m.Load8(k0, k1, k2, k3, k4, k5, k6, k7, k8)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Delete(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8,
) {
	m.Delete8(k0, k1, k2, k3, k4, k5, k6, k7, k8)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Range(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
	f func(K8, V) bool) {
	m.Range8(k0, k1, k2, k3, k4, k5, k6, k7, f)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) LoadAndDelete(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8,
) (value V, loaded bool) {
	return m.LoadAndDelete8(k0, k1, k2, k3, k4, k5, k6, k7, k8)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) LoadOrStore(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8,
	v V) (value V, loaded bool) {
	return m.LoadOrStore8(k0, k1, k2, k3, k4, k5, k6, k7, k8,
		v)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Store(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8,
	value V) {
	m.Store8(k0, k1, k2, k3, k4, k5, k6, k7, k8,
		value)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Load0(k0 K0,
) (value *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	return item0.Load0(k0)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Delete0(k0 K0,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	item0.Delete0(k0)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Range0(
	f func(K0, *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	item0.Range0(f)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) LoadAndDelete0(k0 K0,
) (value *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	return item0.LoadAndDelete0(k0)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) LoadOrStore0(k0 K0,
	v *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]) (value *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]])
	}
	item0 := m.inner

	return item0.LoadOrStore0(k0, v)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Store0(k0 K0,
	value *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]])
	}
	item0 := m.inner

	item0.Store0(k0, value)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Load1(k0 K0, k1 K1,
) (value *Map7[K2, K3, K4, K5, K6, K7, K8, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	return item1.Load0(k1)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Delete1(k0 K0, k1 K1,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	item1.Delete0(k1)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Range1(k0 K0,
	f func(K1, *Map7[K2, K3, K4, K5, K6, K7, K8, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	item1.Range0(f)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) LoadAndDelete1(k0 K0, k1 K1,
) (value *Map7[K2, K3, K4, K5, K6, K7, K8, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	return item1.LoadAndDelete0(k1)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) LoadOrStore1(k0 K0, k1 K1,
	v *Map7[K2, K3, K4, K5, K6, K7, K8, V]) (value *Map7[K2, K3, K4, K5, K6, K7, K8, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]])
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, _ = item0.LoadOrStore0(k0, &Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]{})

	return item1.LoadOrStore0(k1, v)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Store1(k0 K0, k1 K1,
	value *Map7[K2, K3, K4, K5, K6, K7, K8, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]])
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, _ = item0.LoadOrStore0(k0, &Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]{})

	item1.Store0(k1, value)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Load2(k0 K0, k1 K1, k2 K2,
) (value *Map6[K3, K4, K5, K6, K7, K8, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	return item2.Load0(k2)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Delete2(k0 K0, k1 K1, k2 K2,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	item2.Delete0(k2)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Range2(k0 K0, k1 K1,
	f func(K2, *Map6[K3, K4, K5, K6, K7, K8, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	item2.Range0(f)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) LoadAndDelete2(k0 K0, k1 K1, k2 K2,
) (value *Map6[K3, K4, K5, K6, K7, K8, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	return item2.LoadAndDelete0(k2)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) LoadOrStore2(k0 K0, k1 K1, k2 K2,
	v *Map6[K3, K4, K5, K6, K7, K8, V]) (value *Map6[K3, K4, K5, K6, K7, K8, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]])
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, _ = item0.LoadOrStore0(k0, &Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]{})

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, _ = item1.LoadOrStore0(k1, &Map7[K2, K3, K4, K5, K6, K7, K8, V]{})

	return item2.LoadOrStore0(k2, v)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Store2(k0 K0, k1 K1, k2 K2,
	value *Map6[K3, K4, K5, K6, K7, K8, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]])
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, _ = item0.LoadOrStore0(k0, &Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]{})

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, _ = item1.LoadOrStore0(k1, &Map7[K2, K3, K4, K5, K6, K7, K8, V]{})

	item2.Store0(k2, value)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Load3(k0 K0, k1 K1, k2 K2, k3 K3,
) (value *Map5[K4, K5, K6, K7, K8, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	return item3.Load0(k3)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Delete3(k0 K0, k1 K1, k2 K2, k3 K3,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	item3.Delete0(k3)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Range3(k0 K0, k1 K1, k2 K2,
	f func(K3, *Map5[K4, K5, K6, K7, K8, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	item3.Range0(f)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) LoadAndDelete3(k0 K0, k1 K1, k2 K2, k3 K3,
) (value *Map5[K4, K5, K6, K7, K8, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	return item3.LoadAndDelete0(k3)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) LoadOrStore3(k0 K0, k1 K1, k2 K2, k3 K3,
	v *Map5[K4, K5, K6, K7, K8, V]) (value *Map5[K4, K5, K6, K7, K8, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]])
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, _ = item0.LoadOrStore0(k0, &Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]{})

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, _ = item1.LoadOrStore0(k1, &Map7[K2, K3, K4, K5, K6, K7, K8, V]{})

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, _ = item2.LoadOrStore0(k2, &Map6[K3, K4, K5, K6, K7, K8, V]{})

	return item3.LoadOrStore0(k3, v)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Store3(k0 K0, k1 K1, k2 K2, k3 K3,
	value *Map5[K4, K5, K6, K7, K8, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]])
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, _ = item0.LoadOrStore0(k0, &Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]{})

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, _ = item1.LoadOrStore0(k1, &Map7[K2, K3, K4, K5, K6, K7, K8, V]{})

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, _ = item2.LoadOrStore0(k2, &Map6[K3, K4, K5, K6, K7, K8, V]{})

	item3.Store0(k3, value)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Load4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
) (value *Map4[K5, K6, K7, K8, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	return item4.Load0(k4)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Delete4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	item4.Delete0(k4)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Range4(k0 K0, k1 K1, k2 K2, k3 K3,
	f func(K4, *Map4[K5, K6, K7, K8, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	item4.Range0(f)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) LoadAndDelete4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
) (value *Map4[K5, K6, K7, K8, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, loaded = item3.Load0(k3)
	if !loaded {
		return
	}

	return item4.LoadAndDelete0(k4)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) LoadOrStore4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
	v *Map4[K5, K6, K7, K8, V]) (value *Map4[K5, K6, K7, K8, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]])
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, _ = item0.LoadOrStore0(k0, &Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]{})

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, _ = item1.LoadOrStore0(k1, &Map7[K2, K3, K4, K5, K6, K7, K8, V]{})

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, _ = item2.LoadOrStore0(k2, &Map6[K3, K4, K5, K6, K7, K8, V]{})

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, _ = item3.LoadOrStore0(k3, &Map5[K4, K5, K6, K7, K8, V]{})

	return item4.LoadOrStore0(k4, v)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Store4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
	value *Map4[K5, K6, K7, K8, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]])
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, _ = item0.LoadOrStore0(k0, &Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]{})

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, _ = item1.LoadOrStore0(k1, &Map7[K2, K3, K4, K5, K6, K7, K8, V]{})

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, _ = item2.LoadOrStore0(k2, &Map6[K3, K4, K5, K6, K7, K8, V]{})

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, _ = item3.LoadOrStore0(k3, &Map5[K4, K5, K6, K7, K8, V]{})

	item4.Store0(k4, value)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Load5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
) (value *Map3[K6, K7, K8, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map4[K5, K6, K7, K8, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	return item5.Load0(k5)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Delete5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map4[K5, K6, K7, K8, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	item5.Delete0(k5)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Range5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
	f func(K5, *Map3[K6, K7, K8, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map4[K5, K6, K7, K8, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	item5.Range0(f)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) LoadAndDelete5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
) (value *Map3[K6, K7, K8, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, loaded = item3.Load0(k3)
	if !loaded {
		return
	}

	var item5 *Map4[K5, K6, K7, K8, V]
	item5, loaded = item4.Load0(k4)
	if !loaded {
		return
	}

	return item5.LoadAndDelete0(k5)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) LoadOrStore5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
	v *Map3[K6, K7, K8, V]) (value *Map3[K6, K7, K8, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]])
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, _ = item0.LoadOrStore0(k0, &Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]{})

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, _ = item1.LoadOrStore0(k1, &Map7[K2, K3, K4, K5, K6, K7, K8, V]{})

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, _ = item2.LoadOrStore0(k2, &Map6[K3, K4, K5, K6, K7, K8, V]{})

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, _ = item3.LoadOrStore0(k3, &Map5[K4, K5, K6, K7, K8, V]{})

	var item5 *Map4[K5, K6, K7, K8, V]
	item5, _ = item4.LoadOrStore0(k4, &Map4[K5, K6, K7, K8, V]{})

	return item5.LoadOrStore0(k5, v)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Store5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
	value *Map3[K6, K7, K8, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]])
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, _ = item0.LoadOrStore0(k0, &Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]{})

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, _ = item1.LoadOrStore0(k1, &Map7[K2, K3, K4, K5, K6, K7, K8, V]{})

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, _ = item2.LoadOrStore0(k2, &Map6[K3, K4, K5, K6, K7, K8, V]{})

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, _ = item3.LoadOrStore0(k3, &Map5[K4, K5, K6, K7, K8, V]{})

	var item5 *Map4[K5, K6, K7, K8, V]
	item5, _ = item4.LoadOrStore0(k4, &Map4[K5, K6, K7, K8, V]{})

	item5.Store0(k5, value)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Load6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
) (value *Map2[K7, K8, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map4[K5, K6, K7, K8, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map3[K6, K7, K8, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	return item6.Load0(k6)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Delete6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map4[K5, K6, K7, K8, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map3[K6, K7, K8, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	item6.Delete0(k6)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Range6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
	f func(K6, *Map2[K7, K8, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map4[K5, K6, K7, K8, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map3[K6, K7, K8, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	item6.Range0(f)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) LoadAndDelete6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
) (value *Map2[K7, K8, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, loaded = item3.Load0(k3)
	if !loaded {
		return
	}

	var item5 *Map4[K5, K6, K7, K8, V]
	item5, loaded = item4.Load0(k4)
	if !loaded {
		return
	}

	var item6 *Map3[K6, K7, K8, V]
	item6, loaded = item5.Load0(k5)
	if !loaded {
		return
	}

	return item6.LoadAndDelete0(k6)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) LoadOrStore6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
	v *Map2[K7, K8, V]) (value *Map2[K7, K8, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]])
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, _ = item0.LoadOrStore0(k0, &Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]{})

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, _ = item1.LoadOrStore0(k1, &Map7[K2, K3, K4, K5, K6, K7, K8, V]{})

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, _ = item2.LoadOrStore0(k2, &Map6[K3, K4, K5, K6, K7, K8, V]{})

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, _ = item3.LoadOrStore0(k3, &Map5[K4, K5, K6, K7, K8, V]{})

	var item5 *Map4[K5, K6, K7, K8, V]
	item5, _ = item4.LoadOrStore0(k4, &Map4[K5, K6, K7, K8, V]{})

	var item6 *Map3[K6, K7, K8, V]
	item6, _ = item5.LoadOrStore0(k5, &Map3[K6, K7, K8, V]{})

	return item6.LoadOrStore0(k6, v)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Store6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
	value *Map2[K7, K8, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]])
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, _ = item0.LoadOrStore0(k0, &Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]{})

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, _ = item1.LoadOrStore0(k1, &Map7[K2, K3, K4, K5, K6, K7, K8, V]{})

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, _ = item2.LoadOrStore0(k2, &Map6[K3, K4, K5, K6, K7, K8, V]{})

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, _ = item3.LoadOrStore0(k3, &Map5[K4, K5, K6, K7, K8, V]{})

	var item5 *Map4[K5, K6, K7, K8, V]
	item5, _ = item4.LoadOrStore0(k4, &Map4[K5, K6, K7, K8, V]{})

	var item6 *Map3[K6, K7, K8, V]
	item6, _ = item5.LoadOrStore0(k5, &Map3[K6, K7, K8, V]{})

	item6.Store0(k6, value)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Load7(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
) (value *Map1[K8, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map4[K5, K6, K7, K8, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map3[K6, K7, K8, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map2[K7, K8, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	return item7.Load0(k7)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Delete7(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map4[K5, K6, K7, K8, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map3[K6, K7, K8, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map2[K7, K8, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	item7.Delete0(k7)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Range7(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
	f func(K7, *Map1[K8, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map4[K5, K6, K7, K8, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map3[K6, K7, K8, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map2[K7, K8, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	item7.Range0(f)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) LoadAndDelete7(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
) (value *Map1[K8, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, loaded = item3.Load0(k3)
	if !loaded {
		return
	}

	var item5 *Map4[K5, K6, K7, K8, V]
	item5, loaded = item4.Load0(k4)
	if !loaded {
		return
	}

	var item6 *Map3[K6, K7, K8, V]
	item6, loaded = item5.Load0(k5)
	if !loaded {
		return
	}

	var item7 *Map2[K7, K8, V]
	item7, loaded = item6.Load0(k6)
	if !loaded {
		return
	}

	return item7.LoadAndDelete0(k7)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) LoadOrStore7(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
	v *Map1[K8, V]) (value *Map1[K8, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]])
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, _ = item0.LoadOrStore0(k0, &Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]{})

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, _ = item1.LoadOrStore0(k1, &Map7[K2, K3, K4, K5, K6, K7, K8, V]{})

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, _ = item2.LoadOrStore0(k2, &Map6[K3, K4, K5, K6, K7, K8, V]{})

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, _ = item3.LoadOrStore0(k3, &Map5[K4, K5, K6, K7, K8, V]{})

	var item5 *Map4[K5, K6, K7, K8, V]
	item5, _ = item4.LoadOrStore0(k4, &Map4[K5, K6, K7, K8, V]{})

	var item6 *Map3[K6, K7, K8, V]
	item6, _ = item5.LoadOrStore0(k5, &Map3[K6, K7, K8, V]{})

	var item7 *Map2[K7, K8, V]
	item7, _ = item6.LoadOrStore0(k6, &Map2[K7, K8, V]{})

	return item7.LoadOrStore0(k7, v)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Store7(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
	value *Map1[K8, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]])
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, _ = item0.LoadOrStore0(k0, &Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]{})

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, _ = item1.LoadOrStore0(k1, &Map7[K2, K3, K4, K5, K6, K7, K8, V]{})

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, _ = item2.LoadOrStore0(k2, &Map6[K3, K4, K5, K6, K7, K8, V]{})

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, _ = item3.LoadOrStore0(k3, &Map5[K4, K5, K6, K7, K8, V]{})

	var item5 *Map4[K5, K6, K7, K8, V]
	item5, _ = item4.LoadOrStore0(k4, &Map4[K5, K6, K7, K8, V]{})

	var item6 *Map3[K6, K7, K8, V]
	item6, _ = item5.LoadOrStore0(k5, &Map3[K6, K7, K8, V]{})

	var item7 *Map2[K7, K8, V]
	item7, _ = item6.LoadOrStore0(k6, &Map2[K7, K8, V]{})

	item7.Store0(k7, value)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Load8(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8,
) (value V, ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map4[K5, K6, K7, K8, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map3[K6, K7, K8, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map2[K7, K8, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	var item8 *Map1[K8, V]
	item8, ok = item7.Load0(k7)
	if !ok {
		return
	}

	return item8.Load0(k8)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Delete8(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map4[K5, K6, K7, K8, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map3[K6, K7, K8, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map2[K7, K8, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	var item8 *Map1[K8, V]
	item8, ok = item7.Load0(k7)
	if !ok {
		return
	}

	item8.Delete0(k8)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Range8(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
	f func(K8, V) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map4[K5, K6, K7, K8, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map3[K6, K7, K8, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map2[K7, K8, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	var item8 *Map1[K8, V]
	item8, ok = item7.Load0(k7)
	if !ok {
		return
	}

	item8.Range0(f)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) LoadAndDelete8(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8,
) (value V, loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, loaded = item3.Load0(k3)
	if !loaded {
		return
	}

	var item5 *Map4[K5, K6, K7, K8, V]
	item5, loaded = item4.Load0(k4)
	if !loaded {
		return
	}

	var item6 *Map3[K6, K7, K8, V]
	item6, loaded = item5.Load0(k5)
	if !loaded {
		return
	}

	var item7 *Map2[K7, K8, V]
	item7, loaded = item6.Load0(k6)
	if !loaded {
		return
	}

	var item8 *Map1[K8, V]
	item8, loaded = item7.Load0(k7)
	if !loaded {
		return
	}

	return item8.LoadAndDelete0(k8)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) LoadOrStore8(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8,
	v V) (value V, loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]])
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, _ = item0.LoadOrStore0(k0, &Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]{})

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, _ = item1.LoadOrStore0(k1, &Map7[K2, K3, K4, K5, K6, K7, K8, V]{})

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, _ = item2.LoadOrStore0(k2, &Map6[K3, K4, K5, K6, K7, K8, V]{})

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, _ = item3.LoadOrStore0(k3, &Map5[K4, K5, K6, K7, K8, V]{})

	var item5 *Map4[K5, K6, K7, K8, V]
	item5, _ = item4.LoadOrStore0(k4, &Map4[K5, K6, K7, K8, V]{})

	var item6 *Map3[K6, K7, K8, V]
	item6, _ = item5.LoadOrStore0(k5, &Map3[K6, K7, K8, V]{})

	var item7 *Map2[K7, K8, V]
	item7, _ = item6.LoadOrStore0(k6, &Map2[K7, K8, V]{})

	var item8 *Map1[K8, V]
	item8, _ = item7.LoadOrStore0(k7, &Map1[K8, V]{})

	return item8.LoadOrStore0(k8, v)
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Store8(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8,
	value V) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]])
	}
	item0 := m.inner

	var item1 *Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
	item1, _ = item0.LoadOrStore0(k0, &Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]{})

	var item2 *Map7[K2, K3, K4, K5, K6, K7, K8, V]
	item2, _ = item1.LoadOrStore0(k1, &Map7[K2, K3, K4, K5, K6, K7, K8, V]{})

	var item3 *Map6[K3, K4, K5, K6, K7, K8, V]
	item3, _ = item2.LoadOrStore0(k2, &Map6[K3, K4, K5, K6, K7, K8, V]{})

	var item4 *Map5[K4, K5, K6, K7, K8, V]
	item4, _ = item3.LoadOrStore0(k3, &Map5[K4, K5, K6, K7, K8, V]{})

	var item5 *Map4[K5, K6, K7, K8, V]
	item5, _ = item4.LoadOrStore0(k4, &Map4[K5, K6, K7, K8, V]{})

	var item6 *Map3[K6, K7, K8, V]
	item6, _ = item5.LoadOrStore0(k5, &Map3[K6, K7, K8, V]{})

	var item7 *Map2[K7, K8, V]
	item7, _ = item6.LoadOrStore0(k6, &Map2[K7, K8, V]{})

	var item8 *Map1[K8, V]
	item8, _ = item7.LoadOrStore0(k7, &Map1[K8, V]{})

	item8.Store0(k8, value)
}

type Map10[K0 comparable, K1 comparable, K2 comparable, K3 comparable, K4 comparable, K5 comparable, K6 comparable, K7 comparable, K8 comparable, K9 comparable, V any] struct {
	inner *Map1[K0, *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]]
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Load(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9,
) (value V, ok bool) {
	return m.Load9(k0, k1, k2, k3, k4, k5, k6, k7, k8, k9)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Delete(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9,
) {
	m.Delete9(k0, k1, k2, k3, k4, k5, k6, k7, k8, k9)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Range(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8,
	f func(K9, V) bool) {
	m.Range9(k0, k1, k2, k3, k4, k5, k6, k7, k8, f)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) LoadAndDelete(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9,
) (value V, loaded bool) {
	return m.LoadAndDelete9(k0, k1, k2, k3, k4, k5, k6, k7, k8, k9)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) LoadOrStore(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9,
	v V) (value V, loaded bool) {
	return m.LoadOrStore9(k0, k1, k2, k3, k4, k5, k6, k7, k8, k9,
		v)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Store(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9,
	value V) {
	m.Store9(k0, k1, k2, k3, k4, k5, k6, k7, k8, k9,
		value)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Load0(k0 K0,
) (value *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	return item0.Load0(k0)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Delete0(k0 K0,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	item0.Delete0(k0)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Range0(
	f func(K0, *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	item0.Range0(f)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) LoadAndDelete0(k0 K0,
) (value *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	return item0.LoadAndDelete0(k0)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) LoadOrStore0(k0 K0,
	v *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) (value *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]])
	}
	item0 := m.inner

	return item0.LoadOrStore0(k0, v)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Store0(k0 K0,
	value *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]])
	}
	item0 := m.inner

	item0.Store0(k0, value)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Load1(k0 K0, k1 K1,
) (value *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	return item1.Load0(k1)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Delete1(k0 K0, k1 K1,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	item1.Delete0(k1)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Range1(k0 K0,
	f func(K1, *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	item1.Range0(f)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) LoadAndDelete1(k0 K0, k1 K1,
) (value *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	return item1.LoadAndDelete0(k1)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) LoadOrStore1(k0 K0, k1 K1,
	v *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]) (value *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]])
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, _ = item0.LoadOrStore0(k0, &Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	return item1.LoadOrStore0(k1, v)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Store1(k0 K0, k1 K1,
	value *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]])
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, _ = item0.LoadOrStore0(k0, &Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	item1.Store0(k1, value)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Load2(k0 K0, k1 K1, k2 K2,
) (value *Map7[K3, K4, K5, K6, K7, K8, K9, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	return item2.Load0(k2)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Delete2(k0 K0, k1 K1, k2 K2,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	item2.Delete0(k2)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Range2(k0 K0, k1 K1,
	f func(K2, *Map7[K3, K4, K5, K6, K7, K8, K9, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	item2.Range0(f)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) LoadAndDelete2(k0 K0, k1 K1, k2 K2,
) (value *Map7[K3, K4, K5, K6, K7, K8, K9, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	return item2.LoadAndDelete0(k2)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) LoadOrStore2(k0 K0, k1 K1, k2 K2,
	v *Map7[K3, K4, K5, K6, K7, K8, K9, V]) (value *Map7[K3, K4, K5, K6, K7, K8, K9, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]])
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, _ = item0.LoadOrStore0(k0, &Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, _ = item1.LoadOrStore0(k1, &Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	return item2.LoadOrStore0(k2, v)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Store2(k0 K0, k1 K1, k2 K2,
	value *Map7[K3, K4, K5, K6, K7, K8, K9, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]])
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, _ = item0.LoadOrStore0(k0, &Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, _ = item1.LoadOrStore0(k1, &Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	item2.Store0(k2, value)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Load3(k0 K0, k1 K1, k2 K2, k3 K3,
) (value *Map6[K4, K5, K6, K7, K8, K9, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	return item3.Load0(k3)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Delete3(k0 K0, k1 K1, k2 K2, k3 K3,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	item3.Delete0(k3)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Range3(k0 K0, k1 K1, k2 K2,
	f func(K3, *Map6[K4, K5, K6, K7, K8, K9, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	item3.Range0(f)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) LoadAndDelete3(k0 K0, k1 K1, k2 K2, k3 K3,
) (value *Map6[K4, K5, K6, K7, K8, K9, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	return item3.LoadAndDelete0(k3)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) LoadOrStore3(k0 K0, k1 K1, k2 K2, k3 K3,
	v *Map6[K4, K5, K6, K7, K8, K9, V]) (value *Map6[K4, K5, K6, K7, K8, K9, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]])
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, _ = item0.LoadOrStore0(k0, &Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, _ = item1.LoadOrStore0(k1, &Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, _ = item2.LoadOrStore0(k2, &Map7[K3, K4, K5, K6, K7, K8, K9, V]{})

	return item3.LoadOrStore0(k3, v)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Store3(k0 K0, k1 K1, k2 K2, k3 K3,
	value *Map6[K4, K5, K6, K7, K8, K9, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]])
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, _ = item0.LoadOrStore0(k0, &Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, _ = item1.LoadOrStore0(k1, &Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, _ = item2.LoadOrStore0(k2, &Map7[K3, K4, K5, K6, K7, K8, K9, V]{})

	item3.Store0(k3, value)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Load4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
) (value *Map5[K5, K6, K7, K8, K9, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	return item4.Load0(k4)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Delete4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	item4.Delete0(k4)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Range4(k0 K0, k1 K1, k2 K2, k3 K3,
	f func(K4, *Map5[K5, K6, K7, K8, K9, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	item4.Range0(f)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) LoadAndDelete4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
) (value *Map5[K5, K6, K7, K8, K9, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, loaded = item3.Load0(k3)
	if !loaded {
		return
	}

	return item4.LoadAndDelete0(k4)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) LoadOrStore4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
	v *Map5[K5, K6, K7, K8, K9, V]) (value *Map5[K5, K6, K7, K8, K9, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]])
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, _ = item0.LoadOrStore0(k0, &Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, _ = item1.LoadOrStore0(k1, &Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, _ = item2.LoadOrStore0(k2, &Map7[K3, K4, K5, K6, K7, K8, K9, V]{})

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, _ = item3.LoadOrStore0(k3, &Map6[K4, K5, K6, K7, K8, K9, V]{})

	return item4.LoadOrStore0(k4, v)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Store4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
	value *Map5[K5, K6, K7, K8, K9, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]])
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, _ = item0.LoadOrStore0(k0, &Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, _ = item1.LoadOrStore0(k1, &Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, _ = item2.LoadOrStore0(k2, &Map7[K3, K4, K5, K6, K7, K8, K9, V]{})

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, _ = item3.LoadOrStore0(k3, &Map6[K4, K5, K6, K7, K8, K9, V]{})

	item4.Store0(k4, value)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Load5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
) (value *Map4[K6, K7, K8, K9, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	return item5.Load0(k5)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Delete5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	item5.Delete0(k5)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Range5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
	f func(K5, *Map4[K6, K7, K8, K9, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	item5.Range0(f)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) LoadAndDelete5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
) (value *Map4[K6, K7, K8, K9, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, loaded = item3.Load0(k3)
	if !loaded {
		return
	}

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, loaded = item4.Load0(k4)
	if !loaded {
		return
	}

	return item5.LoadAndDelete0(k5)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) LoadOrStore5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
	v *Map4[K6, K7, K8, K9, V]) (value *Map4[K6, K7, K8, K9, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]])
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, _ = item0.LoadOrStore0(k0, &Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, _ = item1.LoadOrStore0(k1, &Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, _ = item2.LoadOrStore0(k2, &Map7[K3, K4, K5, K6, K7, K8, K9, V]{})

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, _ = item3.LoadOrStore0(k3, &Map6[K4, K5, K6, K7, K8, K9, V]{})

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, _ = item4.LoadOrStore0(k4, &Map5[K5, K6, K7, K8, K9, V]{})

	return item5.LoadOrStore0(k5, v)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Store5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
	value *Map4[K6, K7, K8, K9, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]])
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, _ = item0.LoadOrStore0(k0, &Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, _ = item1.LoadOrStore0(k1, &Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, _ = item2.LoadOrStore0(k2, &Map7[K3, K4, K5, K6, K7, K8, K9, V]{})

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, _ = item3.LoadOrStore0(k3, &Map6[K4, K5, K6, K7, K8, K9, V]{})

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, _ = item4.LoadOrStore0(k4, &Map5[K5, K6, K7, K8, K9, V]{})

	item5.Store0(k5, value)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Load6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
) (value *Map3[K7, K8, K9, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map4[K6, K7, K8, K9, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	return item6.Load0(k6)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Delete6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map4[K6, K7, K8, K9, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	item6.Delete0(k6)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Range6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
	f func(K6, *Map3[K7, K8, K9, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map4[K6, K7, K8, K9, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	item6.Range0(f)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) LoadAndDelete6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
) (value *Map3[K7, K8, K9, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, loaded = item3.Load0(k3)
	if !loaded {
		return
	}

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, loaded = item4.Load0(k4)
	if !loaded {
		return
	}

	var item6 *Map4[K6, K7, K8, K9, V]
	item6, loaded = item5.Load0(k5)
	if !loaded {
		return
	}

	return item6.LoadAndDelete0(k6)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) LoadOrStore6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
	v *Map3[K7, K8, K9, V]) (value *Map3[K7, K8, K9, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]])
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, _ = item0.LoadOrStore0(k0, &Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, _ = item1.LoadOrStore0(k1, &Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, _ = item2.LoadOrStore0(k2, &Map7[K3, K4, K5, K6, K7, K8, K9, V]{})

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, _ = item3.LoadOrStore0(k3, &Map6[K4, K5, K6, K7, K8, K9, V]{})

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, _ = item4.LoadOrStore0(k4, &Map5[K5, K6, K7, K8, K9, V]{})

	var item6 *Map4[K6, K7, K8, K9, V]
	item6, _ = item5.LoadOrStore0(k5, &Map4[K6, K7, K8, K9, V]{})

	return item6.LoadOrStore0(k6, v)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Store6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
	value *Map3[K7, K8, K9, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]])
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, _ = item0.LoadOrStore0(k0, &Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, _ = item1.LoadOrStore0(k1, &Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, _ = item2.LoadOrStore0(k2, &Map7[K3, K4, K5, K6, K7, K8, K9, V]{})

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, _ = item3.LoadOrStore0(k3, &Map6[K4, K5, K6, K7, K8, K9, V]{})

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, _ = item4.LoadOrStore0(k4, &Map5[K5, K6, K7, K8, K9, V]{})

	var item6 *Map4[K6, K7, K8, K9, V]
	item6, _ = item5.LoadOrStore0(k5, &Map4[K6, K7, K8, K9, V]{})

	item6.Store0(k6, value)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Load7(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
) (value *Map2[K8, K9, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map4[K6, K7, K8, K9, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map3[K7, K8, K9, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	return item7.Load0(k7)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Delete7(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map4[K6, K7, K8, K9, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map3[K7, K8, K9, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	item7.Delete0(k7)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Range7(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
	f func(K7, *Map2[K8, K9, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map4[K6, K7, K8, K9, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map3[K7, K8, K9, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	item7.Range0(f)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) LoadAndDelete7(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
) (value *Map2[K8, K9, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, loaded = item3.Load0(k3)
	if !loaded {
		return
	}

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, loaded = item4.Load0(k4)
	if !loaded {
		return
	}

	var item6 *Map4[K6, K7, K8, K9, V]
	item6, loaded = item5.Load0(k5)
	if !loaded {
		return
	}

	var item7 *Map3[K7, K8, K9, V]
	item7, loaded = item6.Load0(k6)
	if !loaded {
		return
	}

	return item7.LoadAndDelete0(k7)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) LoadOrStore7(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
	v *Map2[K8, K9, V]) (value *Map2[K8, K9, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]])
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, _ = item0.LoadOrStore0(k0, &Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, _ = item1.LoadOrStore0(k1, &Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, _ = item2.LoadOrStore0(k2, &Map7[K3, K4, K5, K6, K7, K8, K9, V]{})

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, _ = item3.LoadOrStore0(k3, &Map6[K4, K5, K6, K7, K8, K9, V]{})

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, _ = item4.LoadOrStore0(k4, &Map5[K5, K6, K7, K8, K9, V]{})

	var item6 *Map4[K6, K7, K8, K9, V]
	item6, _ = item5.LoadOrStore0(k5, &Map4[K6, K7, K8, K9, V]{})

	var item7 *Map3[K7, K8, K9, V]
	item7, _ = item6.LoadOrStore0(k6, &Map3[K7, K8, K9, V]{})

	return item7.LoadOrStore0(k7, v)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Store7(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
	value *Map2[K8, K9, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]])
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, _ = item0.LoadOrStore0(k0, &Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, _ = item1.LoadOrStore0(k1, &Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, _ = item2.LoadOrStore0(k2, &Map7[K3, K4, K5, K6, K7, K8, K9, V]{})

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, _ = item3.LoadOrStore0(k3, &Map6[K4, K5, K6, K7, K8, K9, V]{})

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, _ = item4.LoadOrStore0(k4, &Map5[K5, K6, K7, K8, K9, V]{})

	var item6 *Map4[K6, K7, K8, K9, V]
	item6, _ = item5.LoadOrStore0(k5, &Map4[K6, K7, K8, K9, V]{})

	var item7 *Map3[K7, K8, K9, V]
	item7, _ = item6.LoadOrStore0(k6, &Map3[K7, K8, K9, V]{})

	item7.Store0(k7, value)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Load8(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8,
) (value *Map1[K9, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map4[K6, K7, K8, K9, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map3[K7, K8, K9, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	var item8 *Map2[K8, K9, V]
	item8, ok = item7.Load0(k7)
	if !ok {
		return
	}

	return item8.Load0(k8)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Delete8(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map4[K6, K7, K8, K9, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map3[K7, K8, K9, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	var item8 *Map2[K8, K9, V]
	item8, ok = item7.Load0(k7)
	if !ok {
		return
	}

	item8.Delete0(k8)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Range8(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
	f func(K8, *Map1[K9, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map4[K6, K7, K8, K9, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map3[K7, K8, K9, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	var item8 *Map2[K8, K9, V]
	item8, ok = item7.Load0(k7)
	if !ok {
		return
	}

	item8.Range0(f)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) LoadAndDelete8(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8,
) (value *Map1[K9, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, loaded = item3.Load0(k3)
	if !loaded {
		return
	}

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, loaded = item4.Load0(k4)
	if !loaded {
		return
	}

	var item6 *Map4[K6, K7, K8, K9, V]
	item6, loaded = item5.Load0(k5)
	if !loaded {
		return
	}

	var item7 *Map3[K7, K8, K9, V]
	item7, loaded = item6.Load0(k6)
	if !loaded {
		return
	}

	var item8 *Map2[K8, K9, V]
	item8, loaded = item7.Load0(k7)
	if !loaded {
		return
	}

	return item8.LoadAndDelete0(k8)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) LoadOrStore8(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8,
	v *Map1[K9, V]) (value *Map1[K9, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]])
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, _ = item0.LoadOrStore0(k0, &Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, _ = item1.LoadOrStore0(k1, &Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, _ = item2.LoadOrStore0(k2, &Map7[K3, K4, K5, K6, K7, K8, K9, V]{})

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, _ = item3.LoadOrStore0(k3, &Map6[K4, K5, K6, K7, K8, K9, V]{})

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, _ = item4.LoadOrStore0(k4, &Map5[K5, K6, K7, K8, K9, V]{})

	var item6 *Map4[K6, K7, K8, K9, V]
	item6, _ = item5.LoadOrStore0(k5, &Map4[K6, K7, K8, K9, V]{})

	var item7 *Map3[K7, K8, K9, V]
	item7, _ = item6.LoadOrStore0(k6, &Map3[K7, K8, K9, V]{})

	var item8 *Map2[K8, K9, V]
	item8, _ = item7.LoadOrStore0(k7, &Map2[K8, K9, V]{})

	return item8.LoadOrStore0(k8, v)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Store8(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8,
	value *Map1[K9, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]])
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, _ = item0.LoadOrStore0(k0, &Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, _ = item1.LoadOrStore0(k1, &Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, _ = item2.LoadOrStore0(k2, &Map7[K3, K4, K5, K6, K7, K8, K9, V]{})

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, _ = item3.LoadOrStore0(k3, &Map6[K4, K5, K6, K7, K8, K9, V]{})

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, _ = item4.LoadOrStore0(k4, &Map5[K5, K6, K7, K8, K9, V]{})

	var item6 *Map4[K6, K7, K8, K9, V]
	item6, _ = item5.LoadOrStore0(k5, &Map4[K6, K7, K8, K9, V]{})

	var item7 *Map3[K7, K8, K9, V]
	item7, _ = item6.LoadOrStore0(k6, &Map3[K7, K8, K9, V]{})

	var item8 *Map2[K8, K9, V]
	item8, _ = item7.LoadOrStore0(k7, &Map2[K8, K9, V]{})

	item8.Store0(k8, value)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Load9(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9,
) (value V, ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map4[K6, K7, K8, K9, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map3[K7, K8, K9, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	var item8 *Map2[K8, K9, V]
	item8, ok = item7.Load0(k7)
	if !ok {
		return
	}

	var item9 *Map1[K9, V]
	item9, ok = item8.Load0(k8)
	if !ok {
		return
	}

	return item9.Load0(k9)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Delete9(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map4[K6, K7, K8, K9, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map3[K7, K8, K9, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	var item8 *Map2[K8, K9, V]
	item8, ok = item7.Load0(k7)
	if !ok {
		return
	}

	var item9 *Map1[K9, V]
	item9, ok = item8.Load0(k8)
	if !ok {
		return
	}

	item9.Delete0(k9)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Range9(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8,
	f func(K9, V) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map4[K6, K7, K8, K9, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map3[K7, K8, K9, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	var item8 *Map2[K8, K9, V]
	item8, ok = item7.Load0(k7)
	if !ok {
		return
	}

	var item9 *Map1[K9, V]
	item9, ok = item8.Load0(k8)
	if !ok {
		return
	}

	item9.Range0(f)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) LoadAndDelete9(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9,
) (value V, loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, loaded = item3.Load0(k3)
	if !loaded {
		return
	}

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, loaded = item4.Load0(k4)
	if !loaded {
		return
	}

	var item6 *Map4[K6, K7, K8, K9, V]
	item6, loaded = item5.Load0(k5)
	if !loaded {
		return
	}

	var item7 *Map3[K7, K8, K9, V]
	item7, loaded = item6.Load0(k6)
	if !loaded {
		return
	}

	var item8 *Map2[K8, K9, V]
	item8, loaded = item7.Load0(k7)
	if !loaded {
		return
	}

	var item9 *Map1[K9, V]
	item9, loaded = item8.Load0(k8)
	if !loaded {
		return
	}

	return item9.LoadAndDelete0(k9)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) LoadOrStore9(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9,
	v V) (value V, loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]])
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, _ = item0.LoadOrStore0(k0, &Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, _ = item1.LoadOrStore0(k1, &Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, _ = item2.LoadOrStore0(k2, &Map7[K3, K4, K5, K6, K7, K8, K9, V]{})

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, _ = item3.LoadOrStore0(k3, &Map6[K4, K5, K6, K7, K8, K9, V]{})

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, _ = item4.LoadOrStore0(k4, &Map5[K5, K6, K7, K8, K9, V]{})

	var item6 *Map4[K6, K7, K8, K9, V]
	item6, _ = item5.LoadOrStore0(k5, &Map4[K6, K7, K8, K9, V]{})

	var item7 *Map3[K7, K8, K9, V]
	item7, _ = item6.LoadOrStore0(k6, &Map3[K7, K8, K9, V]{})

	var item8 *Map2[K8, K9, V]
	item8, _ = item7.LoadOrStore0(k7, &Map2[K8, K9, V]{})

	var item9 *Map1[K9, V]
	item9, _ = item8.LoadOrStore0(k8, &Map1[K9, V]{})

	return item9.LoadOrStore0(k9, v)
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Store9(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9,
	value V) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]])
	}
	item0 := m.inner

	var item1 *Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
	item1, _ = item0.LoadOrStore0(k0, &Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item2 *Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
	item2, _ = item1.LoadOrStore0(k1, &Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]{})

	var item3 *Map7[K3, K4, K5, K6, K7, K8, K9, V]
	item3, _ = item2.LoadOrStore0(k2, &Map7[K3, K4, K5, K6, K7, K8, K9, V]{})

	var item4 *Map6[K4, K5, K6, K7, K8, K9, V]
	item4, _ = item3.LoadOrStore0(k3, &Map6[K4, K5, K6, K7, K8, K9, V]{})

	var item5 *Map5[K5, K6, K7, K8, K9, V]
	item5, _ = item4.LoadOrStore0(k4, &Map5[K5, K6, K7, K8, K9, V]{})

	var item6 *Map4[K6, K7, K8, K9, V]
	item6, _ = item5.LoadOrStore0(k5, &Map4[K6, K7, K8, K9, V]{})

	var item7 *Map3[K7, K8, K9, V]
	item7, _ = item6.LoadOrStore0(k6, &Map3[K7, K8, K9, V]{})

	var item8 *Map2[K8, K9, V]
	item8, _ = item7.LoadOrStore0(k7, &Map2[K8, K9, V]{})

	var item9 *Map1[K9, V]
	item9, _ = item8.LoadOrStore0(k8, &Map1[K9, V]{})

	item9.Store0(k9, value)
}

type Map11[K0 comparable, K1 comparable, K2 comparable, K3 comparable, K4 comparable, K5 comparable, K6 comparable, K7 comparable, K8 comparable, K9 comparable, K10 comparable, V any] struct {
	inner *Map1[K0, *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]]
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Load(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9, k10 K10,
) (value V, ok bool) {
	return m.Load10(k0, k1, k2, k3, k4, k5, k6, k7, k8, k9, k10)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Delete(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9, k10 K10,
) {
	m.Delete10(k0, k1, k2, k3, k4, k5, k6, k7, k8, k9, k10)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Range(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9,
	f func(K10, V) bool) {
	m.Range10(k0, k1, k2, k3, k4, k5, k6, k7, k8, k9, f)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) LoadAndDelete(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9, k10 K10,
) (value V, loaded bool) {
	return m.LoadAndDelete10(k0, k1, k2, k3, k4, k5, k6, k7, k8, k9, k10)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) LoadOrStore(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9, k10 K10,
	v V) (value V, loaded bool) {
	return m.LoadOrStore10(k0, k1, k2, k3, k4, k5, k6, k7, k8, k9, k10,
		v)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Store(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9, k10 K10,
	value V) {
	m.Store10(k0, k1, k2, k3, k4, k5, k6, k7, k8, k9, k10,
		value)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Load0(k0 K0,
) (value *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	return item0.Load0(k0)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Delete0(k0 K0,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	item0.Delete0(k0)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Range0(
	f func(K0, *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	item0.Range0(f)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) LoadAndDelete0(k0 K0,
) (value *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	return item0.LoadAndDelete0(k0)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) LoadOrStore0(k0 K0,
	v *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) (value *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]])
	}
	item0 := m.inner

	return item0.LoadOrStore0(k0, v)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Store0(k0 K0,
	value *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]])
	}
	item0 := m.inner

	item0.Store0(k0, value)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Load1(k0 K0, k1 K1,
) (value *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	return item1.Load0(k1)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Delete1(k0 K0, k1 K1,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	item1.Delete0(k1)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Range1(k0 K0,
	f func(K1, *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	item1.Range0(f)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) LoadAndDelete1(k0 K0, k1 K1,
) (value *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	return item1.LoadAndDelete0(k1)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) LoadOrStore1(k0 K0, k1 K1,
	v *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) (value *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]])
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, _ = item0.LoadOrStore0(k0, &Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	return item1.LoadOrStore0(k1, v)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Store1(k0 K0, k1 K1,
	value *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]])
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, _ = item0.LoadOrStore0(k0, &Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	item1.Store0(k1, value)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Load2(k0 K0, k1 K1, k2 K2,
) (value *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	return item2.Load0(k2)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Delete2(k0 K0, k1 K1, k2 K2,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	item2.Delete0(k2)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Range2(k0 K0, k1 K1,
	f func(K2, *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	item2.Range0(f)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) LoadAndDelete2(k0 K0, k1 K1, k2 K2,
) (value *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	return item2.LoadAndDelete0(k2)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) LoadOrStore2(k0 K0, k1 K1, k2 K2,
	v *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]) (value *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]])
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, _ = item0.LoadOrStore0(k0, &Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, _ = item1.LoadOrStore0(k1, &Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	return item2.LoadOrStore0(k2, v)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Store2(k0 K0, k1 K1, k2 K2,
	value *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]])
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, _ = item0.LoadOrStore0(k0, &Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, _ = item1.LoadOrStore0(k1, &Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	item2.Store0(k2, value)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Load3(k0 K0, k1 K1, k2 K2, k3 K3,
) (value *Map7[K4, K5, K6, K7, K8, K9, K10, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	return item3.Load0(k3)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Delete3(k0 K0, k1 K1, k2 K2, k3 K3,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	item3.Delete0(k3)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Range3(k0 K0, k1 K1, k2 K2,
	f func(K3, *Map7[K4, K5, K6, K7, K8, K9, K10, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	item3.Range0(f)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) LoadAndDelete3(k0 K0, k1 K1, k2 K2, k3 K3,
) (value *Map7[K4, K5, K6, K7, K8, K9, K10, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	return item3.LoadAndDelete0(k3)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) LoadOrStore3(k0 K0, k1 K1, k2 K2, k3 K3,
	v *Map7[K4, K5, K6, K7, K8, K9, K10, V]) (value *Map7[K4, K5, K6, K7, K8, K9, K10, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]])
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, _ = item0.LoadOrStore0(k0, &Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, _ = item1.LoadOrStore0(k1, &Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, _ = item2.LoadOrStore0(k2, &Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	return item3.LoadOrStore0(k3, v)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Store3(k0 K0, k1 K1, k2 K2, k3 K3,
	value *Map7[K4, K5, K6, K7, K8, K9, K10, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]])
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, _ = item0.LoadOrStore0(k0, &Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, _ = item1.LoadOrStore0(k1, &Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, _ = item2.LoadOrStore0(k2, &Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	item3.Store0(k3, value)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Load4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
) (value *Map6[K5, K6, K7, K8, K9, K10, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	return item4.Load0(k4)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Delete4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	item4.Delete0(k4)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Range4(k0 K0, k1 K1, k2 K2, k3 K3,
	f func(K4, *Map6[K5, K6, K7, K8, K9, K10, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	item4.Range0(f)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) LoadAndDelete4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
) (value *Map6[K5, K6, K7, K8, K9, K10, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, loaded = item3.Load0(k3)
	if !loaded {
		return
	}

	return item4.LoadAndDelete0(k4)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) LoadOrStore4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
	v *Map6[K5, K6, K7, K8, K9, K10, V]) (value *Map6[K5, K6, K7, K8, K9, K10, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]])
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, _ = item0.LoadOrStore0(k0, &Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, _ = item1.LoadOrStore0(k1, &Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, _ = item2.LoadOrStore0(k2, &Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, _ = item3.LoadOrStore0(k3, &Map7[K4, K5, K6, K7, K8, K9, K10, V]{})

	return item4.LoadOrStore0(k4, v)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Store4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
	value *Map6[K5, K6, K7, K8, K9, K10, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]])
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, _ = item0.LoadOrStore0(k0, &Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, _ = item1.LoadOrStore0(k1, &Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, _ = item2.LoadOrStore0(k2, &Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, _ = item3.LoadOrStore0(k3, &Map7[K4, K5, K6, K7, K8, K9, K10, V]{})

	item4.Store0(k4, value)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Load5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
) (value *Map5[K6, K7, K8, K9, K10, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	return item5.Load0(k5)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Delete5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	item5.Delete0(k5)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Range5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4,
	f func(K5, *Map5[K6, K7, K8, K9, K10, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	item5.Range0(f)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) LoadAndDelete5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
) (value *Map5[K6, K7, K8, K9, K10, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, loaded = item3.Load0(k3)
	if !loaded {
		return
	}

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, loaded = item4.Load0(k4)
	if !loaded {
		return
	}

	return item5.LoadAndDelete0(k5)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) LoadOrStore5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
	v *Map5[K6, K7, K8, K9, K10, V]) (value *Map5[K6, K7, K8, K9, K10, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]])
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, _ = item0.LoadOrStore0(k0, &Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, _ = item1.LoadOrStore0(k1, &Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, _ = item2.LoadOrStore0(k2, &Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, _ = item3.LoadOrStore0(k3, &Map7[K4, K5, K6, K7, K8, K9, K10, V]{})

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, _ = item4.LoadOrStore0(k4, &Map6[K5, K6, K7, K8, K9, K10, V]{})

	return item5.LoadOrStore0(k5, v)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Store5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
	value *Map5[K6, K7, K8, K9, K10, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]])
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, _ = item0.LoadOrStore0(k0, &Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, _ = item1.LoadOrStore0(k1, &Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, _ = item2.LoadOrStore0(k2, &Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, _ = item3.LoadOrStore0(k3, &Map7[K4, K5, K6, K7, K8, K9, K10, V]{})

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, _ = item4.LoadOrStore0(k4, &Map6[K5, K6, K7, K8, K9, K10, V]{})

	item5.Store0(k5, value)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Load6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
) (value *Map4[K7, K8, K9, K10, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	return item6.Load0(k6)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Delete6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	item6.Delete0(k6)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Range6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5,
	f func(K6, *Map4[K7, K8, K9, K10, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	item6.Range0(f)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) LoadAndDelete6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
) (value *Map4[K7, K8, K9, K10, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, loaded = item3.Load0(k3)
	if !loaded {
		return
	}

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, loaded = item4.Load0(k4)
	if !loaded {
		return
	}

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, loaded = item5.Load0(k5)
	if !loaded {
		return
	}

	return item6.LoadAndDelete0(k6)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) LoadOrStore6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
	v *Map4[K7, K8, K9, K10, V]) (value *Map4[K7, K8, K9, K10, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]])
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, _ = item0.LoadOrStore0(k0, &Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, _ = item1.LoadOrStore0(k1, &Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, _ = item2.LoadOrStore0(k2, &Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, _ = item3.LoadOrStore0(k3, &Map7[K4, K5, K6, K7, K8, K9, K10, V]{})

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, _ = item4.LoadOrStore0(k4, &Map6[K5, K6, K7, K8, K9, K10, V]{})

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, _ = item5.LoadOrStore0(k5, &Map5[K6, K7, K8, K9, K10, V]{})

	return item6.LoadOrStore0(k6, v)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Store6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
	value *Map4[K7, K8, K9, K10, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]])
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, _ = item0.LoadOrStore0(k0, &Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, _ = item1.LoadOrStore0(k1, &Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, _ = item2.LoadOrStore0(k2, &Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, _ = item3.LoadOrStore0(k3, &Map7[K4, K5, K6, K7, K8, K9, K10, V]{})

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, _ = item4.LoadOrStore0(k4, &Map6[K5, K6, K7, K8, K9, K10, V]{})

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, _ = item5.LoadOrStore0(k5, &Map5[K6, K7, K8, K9, K10, V]{})

	item6.Store0(k6, value)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Load7(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
) (value *Map3[K8, K9, K10, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map4[K7, K8, K9, K10, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	return item7.Load0(k7)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Delete7(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map4[K7, K8, K9, K10, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	item7.Delete0(k7)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Range7(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6,
	f func(K7, *Map3[K8, K9, K10, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map4[K7, K8, K9, K10, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	item7.Range0(f)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) LoadAndDelete7(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
) (value *Map3[K8, K9, K10, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, loaded = item3.Load0(k3)
	if !loaded {
		return
	}

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, loaded = item4.Load0(k4)
	if !loaded {
		return
	}

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, loaded = item5.Load0(k5)
	if !loaded {
		return
	}

	var item7 *Map4[K7, K8, K9, K10, V]
	item7, loaded = item6.Load0(k6)
	if !loaded {
		return
	}

	return item7.LoadAndDelete0(k7)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) LoadOrStore7(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
	v *Map3[K8, K9, K10, V]) (value *Map3[K8, K9, K10, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]])
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, _ = item0.LoadOrStore0(k0, &Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, _ = item1.LoadOrStore0(k1, &Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, _ = item2.LoadOrStore0(k2, &Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, _ = item3.LoadOrStore0(k3, &Map7[K4, K5, K6, K7, K8, K9, K10, V]{})

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, _ = item4.LoadOrStore0(k4, &Map6[K5, K6, K7, K8, K9, K10, V]{})

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, _ = item5.LoadOrStore0(k5, &Map5[K6, K7, K8, K9, K10, V]{})

	var item7 *Map4[K7, K8, K9, K10, V]
	item7, _ = item6.LoadOrStore0(k6, &Map4[K7, K8, K9, K10, V]{})

	return item7.LoadOrStore0(k7, v)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Store7(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
	value *Map3[K8, K9, K10, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]])
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, _ = item0.LoadOrStore0(k0, &Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, _ = item1.LoadOrStore0(k1, &Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, _ = item2.LoadOrStore0(k2, &Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, _ = item3.LoadOrStore0(k3, &Map7[K4, K5, K6, K7, K8, K9, K10, V]{})

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, _ = item4.LoadOrStore0(k4, &Map6[K5, K6, K7, K8, K9, K10, V]{})

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, _ = item5.LoadOrStore0(k5, &Map5[K6, K7, K8, K9, K10, V]{})

	var item7 *Map4[K7, K8, K9, K10, V]
	item7, _ = item6.LoadOrStore0(k6, &Map4[K7, K8, K9, K10, V]{})

	item7.Store0(k7, value)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Load8(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8,
) (value *Map2[K9, K10, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map4[K7, K8, K9, K10, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	var item8 *Map3[K8, K9, K10, V]
	item8, ok = item7.Load0(k7)
	if !ok {
		return
	}

	return item8.Load0(k8)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Delete8(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map4[K7, K8, K9, K10, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	var item8 *Map3[K8, K9, K10, V]
	item8, ok = item7.Load0(k7)
	if !ok {
		return
	}

	item8.Delete0(k8)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Range8(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7,
	f func(K8, *Map2[K9, K10, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map4[K7, K8, K9, K10, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	var item8 *Map3[K8, K9, K10, V]
	item8, ok = item7.Load0(k7)
	if !ok {
		return
	}

	item8.Range0(f)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) LoadAndDelete8(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8,
) (value *Map2[K9, K10, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, loaded = item3.Load0(k3)
	if !loaded {
		return
	}

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, loaded = item4.Load0(k4)
	if !loaded {
		return
	}

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, loaded = item5.Load0(k5)
	if !loaded {
		return
	}

	var item7 *Map4[K7, K8, K9, K10, V]
	item7, loaded = item6.Load0(k6)
	if !loaded {
		return
	}

	var item8 *Map3[K8, K9, K10, V]
	item8, loaded = item7.Load0(k7)
	if !loaded {
		return
	}

	return item8.LoadAndDelete0(k8)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) LoadOrStore8(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8,
	v *Map2[K9, K10, V]) (value *Map2[K9, K10, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]])
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, _ = item0.LoadOrStore0(k0, &Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, _ = item1.LoadOrStore0(k1, &Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, _ = item2.LoadOrStore0(k2, &Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, _ = item3.LoadOrStore0(k3, &Map7[K4, K5, K6, K7, K8, K9, K10, V]{})

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, _ = item4.LoadOrStore0(k4, &Map6[K5, K6, K7, K8, K9, K10, V]{})

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, _ = item5.LoadOrStore0(k5, &Map5[K6, K7, K8, K9, K10, V]{})

	var item7 *Map4[K7, K8, K9, K10, V]
	item7, _ = item6.LoadOrStore0(k6, &Map4[K7, K8, K9, K10, V]{})

	var item8 *Map3[K8, K9, K10, V]
	item8, _ = item7.LoadOrStore0(k7, &Map3[K8, K9, K10, V]{})

	return item8.LoadOrStore0(k8, v)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Store8(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8,
	value *Map2[K9, K10, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]])
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, _ = item0.LoadOrStore0(k0, &Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, _ = item1.LoadOrStore0(k1, &Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, _ = item2.LoadOrStore0(k2, &Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, _ = item3.LoadOrStore0(k3, &Map7[K4, K5, K6, K7, K8, K9, K10, V]{})

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, _ = item4.LoadOrStore0(k4, &Map6[K5, K6, K7, K8, K9, K10, V]{})

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, _ = item5.LoadOrStore0(k5, &Map5[K6, K7, K8, K9, K10, V]{})

	var item7 *Map4[K7, K8, K9, K10, V]
	item7, _ = item6.LoadOrStore0(k6, &Map4[K7, K8, K9, K10, V]{})

	var item8 *Map3[K8, K9, K10, V]
	item8, _ = item7.LoadOrStore0(k7, &Map3[K8, K9, K10, V]{})

	item8.Store0(k8, value)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Load9(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9,
) (value *Map1[K10, V], ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map4[K7, K8, K9, K10, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	var item8 *Map3[K8, K9, K10, V]
	item8, ok = item7.Load0(k7)
	if !ok {
		return
	}

	var item9 *Map2[K9, K10, V]
	item9, ok = item8.Load0(k8)
	if !ok {
		return
	}

	return item9.Load0(k9)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Delete9(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map4[K7, K8, K9, K10, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	var item8 *Map3[K8, K9, K10, V]
	item8, ok = item7.Load0(k7)
	if !ok {
		return
	}

	var item9 *Map2[K9, K10, V]
	item9, ok = item8.Load0(k8)
	if !ok {
		return
	}

	item9.Delete0(k9)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Range9(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8,
	f func(K9, *Map1[K10, V]) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map4[K7, K8, K9, K10, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	var item8 *Map3[K8, K9, K10, V]
	item8, ok = item7.Load0(k7)
	if !ok {
		return
	}

	var item9 *Map2[K9, K10, V]
	item9, ok = item8.Load0(k8)
	if !ok {
		return
	}

	item9.Range0(f)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) LoadAndDelete9(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9,
) (value *Map1[K10, V], loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, loaded = item3.Load0(k3)
	if !loaded {
		return
	}

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, loaded = item4.Load0(k4)
	if !loaded {
		return
	}

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, loaded = item5.Load0(k5)
	if !loaded {
		return
	}

	var item7 *Map4[K7, K8, K9, K10, V]
	item7, loaded = item6.Load0(k6)
	if !loaded {
		return
	}

	var item8 *Map3[K8, K9, K10, V]
	item8, loaded = item7.Load0(k7)
	if !loaded {
		return
	}

	var item9 *Map2[K9, K10, V]
	item9, loaded = item8.Load0(k8)
	if !loaded {
		return
	}

	return item9.LoadAndDelete0(k9)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) LoadOrStore9(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9,
	v *Map1[K10, V]) (value *Map1[K10, V], loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]])
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, _ = item0.LoadOrStore0(k0, &Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, _ = item1.LoadOrStore0(k1, &Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, _ = item2.LoadOrStore0(k2, &Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, _ = item3.LoadOrStore0(k3, &Map7[K4, K5, K6, K7, K8, K9, K10, V]{})

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, _ = item4.LoadOrStore0(k4, &Map6[K5, K6, K7, K8, K9, K10, V]{})

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, _ = item5.LoadOrStore0(k5, &Map5[K6, K7, K8, K9, K10, V]{})

	var item7 *Map4[K7, K8, K9, K10, V]
	item7, _ = item6.LoadOrStore0(k6, &Map4[K7, K8, K9, K10, V]{})

	var item8 *Map3[K8, K9, K10, V]
	item8, _ = item7.LoadOrStore0(k7, &Map3[K8, K9, K10, V]{})

	var item9 *Map2[K9, K10, V]
	item9, _ = item8.LoadOrStore0(k8, &Map2[K9, K10, V]{})

	return item9.LoadOrStore0(k9, v)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Store9(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9,
	value *Map1[K10, V]) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]])
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, _ = item0.LoadOrStore0(k0, &Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, _ = item1.LoadOrStore0(k1, &Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, _ = item2.LoadOrStore0(k2, &Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, _ = item3.LoadOrStore0(k3, &Map7[K4, K5, K6, K7, K8, K9, K10, V]{})

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, _ = item4.LoadOrStore0(k4, &Map6[K5, K6, K7, K8, K9, K10, V]{})

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, _ = item5.LoadOrStore0(k5, &Map5[K6, K7, K8, K9, K10, V]{})

	var item7 *Map4[K7, K8, K9, K10, V]
	item7, _ = item6.LoadOrStore0(k6, &Map4[K7, K8, K9, K10, V]{})

	var item8 *Map3[K8, K9, K10, V]
	item8, _ = item7.LoadOrStore0(k7, &Map3[K8, K9, K10, V]{})

	var item9 *Map2[K9, K10, V]
	item9, _ = item8.LoadOrStore0(k8, &Map2[K9, K10, V]{})

	item9.Store0(k9, value)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Load10(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9, k10 K10,
) (value V, ok bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map4[K7, K8, K9, K10, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	var item8 *Map3[K8, K9, K10, V]
	item8, ok = item7.Load0(k7)
	if !ok {
		return
	}

	var item9 *Map2[K9, K10, V]
	item9, ok = item8.Load0(k8)
	if !ok {
		return
	}

	var item10 *Map1[K10, V]
	item10, ok = item9.Load0(k9)
	if !ok {
		return
	}

	return item10.Load0(k10)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Delete10(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9, k10 K10,
) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map4[K7, K8, K9, K10, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	var item8 *Map3[K8, K9, K10, V]
	item8, ok = item7.Load0(k7)
	if !ok {
		return
	}

	var item9 *Map2[K9, K10, V]
	item9, ok = item8.Load0(k8)
	if !ok {
		return
	}

	var item10 *Map1[K10, V]
	item10, ok = item9.Load0(k9)
	if !ok {
		return
	}

	item10.Delete0(k10)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Range10(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9,
	f func(K10, V) bool) {
	if m.inner == nil {
		return
	}
	var ok bool
	_ = ok
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, ok = item0.Load0(k0)
	if !ok {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, ok = item1.Load0(k1)
	if !ok {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, ok = item2.Load0(k2)
	if !ok {
		return
	}

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, ok = item3.Load0(k3)
	if !ok {
		return
	}

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, ok = item4.Load0(k4)
	if !ok {
		return
	}

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, ok = item5.Load0(k5)
	if !ok {
		return
	}

	var item7 *Map4[K7, K8, K9, K10, V]
	item7, ok = item6.Load0(k6)
	if !ok {
		return
	}

	var item8 *Map3[K8, K9, K10, V]
	item8, ok = item7.Load0(k7)
	if !ok {
		return
	}

	var item9 *Map2[K9, K10, V]
	item9, ok = item8.Load0(k8)
	if !ok {
		return
	}

	var item10 *Map1[K10, V]
	item10, ok = item9.Load0(k9)
	if !ok {
		return
	}

	item10.Range0(f)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) LoadAndDelete10(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9, k10 K10,
) (value V, loaded bool) {
	if m.inner == nil {
		return
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, loaded = item0.Load0(k0)
	if !loaded {
		return
	}

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, loaded = item1.Load0(k1)
	if !loaded {
		return
	}

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, loaded = item2.Load0(k2)
	if !loaded {
		return
	}

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, loaded = item3.Load0(k3)
	if !loaded {
		return
	}

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, loaded = item4.Load0(k4)
	if !loaded {
		return
	}

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, loaded = item5.Load0(k5)
	if !loaded {
		return
	}

	var item7 *Map4[K7, K8, K9, K10, V]
	item7, loaded = item6.Load0(k6)
	if !loaded {
		return
	}

	var item8 *Map3[K8, K9, K10, V]
	item8, loaded = item7.Load0(k7)
	if !loaded {
		return
	}

	var item9 *Map2[K9, K10, V]
	item9, loaded = item8.Load0(k8)
	if !loaded {
		return
	}

	var item10 *Map1[K10, V]
	item10, loaded = item9.Load0(k9)
	if !loaded {
		return
	}

	return item10.LoadAndDelete0(k10)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) LoadOrStore10(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9, k10 K10,
	v V) (value V, loaded bool) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]])
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, _ = item0.LoadOrStore0(k0, &Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, _ = item1.LoadOrStore0(k1, &Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, _ = item2.LoadOrStore0(k2, &Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, _ = item3.LoadOrStore0(k3, &Map7[K4, K5, K6, K7, K8, K9, K10, V]{})

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, _ = item4.LoadOrStore0(k4, &Map6[K5, K6, K7, K8, K9, K10, V]{})

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, _ = item5.LoadOrStore0(k5, &Map5[K6, K7, K8, K9, K10, V]{})

	var item7 *Map4[K7, K8, K9, K10, V]
	item7, _ = item6.LoadOrStore0(k6, &Map4[K7, K8, K9, K10, V]{})

	var item8 *Map3[K8, K9, K10, V]
	item8, _ = item7.LoadOrStore0(k7, &Map3[K8, K9, K10, V]{})

	var item9 *Map2[K9, K10, V]
	item9, _ = item8.LoadOrStore0(k8, &Map2[K9, K10, V]{})

	var item10 *Map1[K10, V]
	item10, _ = item9.LoadOrStore0(k9, &Map1[K10, V]{})

	return item10.LoadOrStore0(k10, v)
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Store10(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9, k10 K10,
	value V) {
	if m.inner == nil {
		m.inner = new(Map1[K0, *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]])
	}
	item0 := m.inner

	var item1 *Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item1, _ = item0.LoadOrStore0(k0, &Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item2 *Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
	item2, _ = item1.LoadOrStore0(k1, &Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item3 *Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
	item3, _ = item2.LoadOrStore0(k2, &Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]{})

	var item4 *Map7[K4, K5, K6, K7, K8, K9, K10, V]
	item4, _ = item3.LoadOrStore0(k3, &Map7[K4, K5, K6, K7, K8, K9, K10, V]{})

	var item5 *Map6[K5, K6, K7, K8, K9, K10, V]
	item5, _ = item4.LoadOrStore0(k4, &Map6[K5, K6, K7, K8, K9, K10, V]{})

	var item6 *Map5[K6, K7, K8, K9, K10, V]
	item6, _ = item5.LoadOrStore0(k5, &Map5[K6, K7, K8, K9, K10, V]{})

	var item7 *Map4[K7, K8, K9, K10, V]
	item7, _ = item6.LoadOrStore0(k6, &Map4[K7, K8, K9, K10, V]{})

	var item8 *Map3[K8, K9, K10, V]
	item8, _ = item7.LoadOrStore0(k7, &Map3[K8, K9, K10, V]{})

	var item9 *Map2[K9, K10, V]
	item9, _ = item8.LoadOrStore0(k8, &Map2[K9, K10, V]{})

	var item10 *Map1[K10, V]
	item10, _ = item9.LoadOrStore0(k9, &Map1[K10, V]{})

	item10.Store0(k10, value)
}
