package maps

import (
)


type Map2[K0 comparable, K1 comparable, V any] struct {
  inner Map1[K0, Map1[K1,V]]
}

func (m *Map2[K0, K1, V]) Load(key K0) (value Map1[K1, V], ok bool) {
  return m.inner.Load(key)
}



func (m *Map2[K0, K1, V]) Load1(k0 K0, k1 K1, 
) (value V, ok bool) {
  item0 := m.inner

  var item1  Map1[K1, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  return item1.Load(k1)
}



type Map3[K0 comparable, K1 comparable, K2 comparable, V any] struct {
  inner Map1[K0, Map2[K1,K2,V]]
}

func (m *Map3[K0, K1, K2, V]) Load(key K0) (value Map2[K1, K2, V], ok bool) {
  return m.inner.Load(key)
}



func (m *Map3[K0, K1, K2, V]) Load1(k0 K0, k1 K1, 
) (value  Map1[K2, V] , ok bool) {
  item0 := m.inner

  var item1  Map2[K1, K2, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  return item1.Load(k1)
}


func (m *Map3[K0, K1, K2, V]) Load2(k0 K0, k1 K1, k2 K2, 
) (value V, ok bool) {
  item0 := m.inner

  var item1  Map2[K1, K2, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map1[K2, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  return item2.Load(k2)
}


