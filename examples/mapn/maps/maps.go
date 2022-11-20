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



type Map4[K0 comparable, K1 comparable, K2 comparable, K3 comparable, V any] struct {
  inner Map1[K0, Map3[K1,K2,K3,V]]
}

func (m *Map4[K0, K1, K2, K3, V]) Load(key K0) (value Map3[K1, K2, K3, V], ok bool) {
  return m.inner.Load(key)
}



func (m *Map4[K0, K1, K2, K3, V]) Load1(k0 K0, k1 K1, 
) (value  Map2[K2, K3, V] , ok bool) {
  item0 := m.inner

  var item1  Map3[K1, K2, K3, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  return item1.Load(k1)
}


func (m *Map4[K0, K1, K2, K3, V]) Load2(k0 K0, k1 K1, k2 K2, 
) (value  Map1[K3, V] , ok bool) {
  item0 := m.inner

  var item1  Map3[K1, K2, K3, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map2[K2, K3, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  return item2.Load(k2)
}


func (m *Map4[K0, K1, K2, K3, V]) Load3(k0 K0, k1 K1, k2 K2, k3 K3, 
) (value V, ok bool) {
  item0 := m.inner

  var item1  Map3[K1, K2, K3, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map2[K2, K3, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map1[K3, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  return item3.Load(k3)
}



type Map5[K0 comparable, K1 comparable, K2 comparable, K3 comparable, K4 comparable, V any] struct {
  inner Map1[K0, Map4[K1,K2,K3,K4,V]]
}

func (m *Map5[K0, K1, K2, K3, K4, V]) Load(key K0) (value Map4[K1, K2, K3, K4, V], ok bool) {
  return m.inner.Load(key)
}



func (m *Map5[K0, K1, K2, K3, K4, V]) Load1(k0 K0, k1 K1, 
) (value  Map3[K2, K3, K4, V] , ok bool) {
  item0 := m.inner

  var item1  Map4[K1, K2, K3, K4, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  return item1.Load(k1)
}


func (m *Map5[K0, K1, K2, K3, K4, V]) Load2(k0 K0, k1 K1, k2 K2, 
) (value  Map2[K3, K4, V] , ok bool) {
  item0 := m.inner

  var item1  Map4[K1, K2, K3, K4, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map3[K2, K3, K4, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  return item2.Load(k2)
}


func (m *Map5[K0, K1, K2, K3, K4, V]) Load3(k0 K0, k1 K1, k2 K2, k3 K3, 
) (value  Map1[K4, V] , ok bool) {
  item0 := m.inner

  var item1  Map4[K1, K2, K3, K4, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map3[K2, K3, K4, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map2[K3, K4, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  return item3.Load(k3)
}


func (m *Map5[K0, K1, K2, K3, K4, V]) Load4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, 
) (value V, ok bool) {
  item0 := m.inner

  var item1  Map4[K1, K2, K3, K4, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map3[K2, K3, K4, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map2[K3, K4, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  var item4  Map1[K4, V]
  item4, ok = item3.Load(k3)
  if !ok {
    return
  }

  return item4.Load(k4)
}



type Map6[K0 comparable, K1 comparable, K2 comparable, K3 comparable, K4 comparable, K5 comparable, V any] struct {
  inner Map1[K0, Map5[K1,K2,K3,K4,K5,V]]
}

func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Load(key K0) (value Map5[K1, K2, K3, K4, K5, V], ok bool) {
  return m.inner.Load(key)
}



func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Load1(k0 K0, k1 K1, 
) (value  Map4[K2, K3, K4, K5, V] , ok bool) {
  item0 := m.inner

  var item1  Map5[K1, K2, K3, K4, K5, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  return item1.Load(k1)
}


func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Load2(k0 K0, k1 K1, k2 K2, 
) (value  Map3[K3, K4, K5, V] , ok bool) {
  item0 := m.inner

  var item1  Map5[K1, K2, K3, K4, K5, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map4[K2, K3, K4, K5, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  return item2.Load(k2)
}


func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Load3(k0 K0, k1 K1, k2 K2, k3 K3, 
) (value  Map2[K4, K5, V] , ok bool) {
  item0 := m.inner

  var item1  Map5[K1, K2, K3, K4, K5, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map4[K2, K3, K4, K5, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map3[K3, K4, K5, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  return item3.Load(k3)
}


func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Load4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, 
) (value  Map1[K5, V] , ok bool) {
  item0 := m.inner

  var item1  Map5[K1, K2, K3, K4, K5, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map4[K2, K3, K4, K5, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map3[K3, K4, K5, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  var item4  Map2[K4, K5, V]
  item4, ok = item3.Load(k3)
  if !ok {
    return
  }

  return item4.Load(k4)
}


func (m *Map6[K0, K1, K2, K3, K4, K5, V]) Load5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, 
) (value V, ok bool) {
  item0 := m.inner

  var item1  Map5[K1, K2, K3, K4, K5, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map4[K2, K3, K4, K5, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map3[K3, K4, K5, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  var item4  Map2[K4, K5, V]
  item4, ok = item3.Load(k3)
  if !ok {
    return
  }

  var item5  Map1[K5, V]
  item5, ok = item4.Load(k4)
  if !ok {
    return
  }

  return item5.Load(k5)
}



type Map7[K0 comparable, K1 comparable, K2 comparable, K3 comparable, K4 comparable, K5 comparable, K6 comparable, V any] struct {
  inner Map1[K0, Map6[K1,K2,K3,K4,K5,K6,V]]
}

func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Load(key K0) (value Map6[K1, K2, K3, K4, K5, K6, V], ok bool) {
  return m.inner.Load(key)
}



func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Load1(k0 K0, k1 K1, 
) (value  Map5[K2, K3, K4, K5, K6, V] , ok bool) {
  item0 := m.inner

  var item1  Map6[K1, K2, K3, K4, K5, K6, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  return item1.Load(k1)
}


func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Load2(k0 K0, k1 K1, k2 K2, 
) (value  Map4[K3, K4, K5, K6, V] , ok bool) {
  item0 := m.inner

  var item1  Map6[K1, K2, K3, K4, K5, K6, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map5[K2, K3, K4, K5, K6, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  return item2.Load(k2)
}


func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Load3(k0 K0, k1 K1, k2 K2, k3 K3, 
) (value  Map3[K4, K5, K6, V] , ok bool) {
  item0 := m.inner

  var item1  Map6[K1, K2, K3, K4, K5, K6, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map5[K2, K3, K4, K5, K6, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map4[K3, K4, K5, K6, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  return item3.Load(k3)
}


func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Load4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, 
) (value  Map2[K5, K6, V] , ok bool) {
  item0 := m.inner

  var item1  Map6[K1, K2, K3, K4, K5, K6, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map5[K2, K3, K4, K5, K6, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map4[K3, K4, K5, K6, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  var item4  Map3[K4, K5, K6, V]
  item4, ok = item3.Load(k3)
  if !ok {
    return
  }

  return item4.Load(k4)
}


func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Load5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, 
) (value  Map1[K6, V] , ok bool) {
  item0 := m.inner

  var item1  Map6[K1, K2, K3, K4, K5, K6, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map5[K2, K3, K4, K5, K6, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map4[K3, K4, K5, K6, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  var item4  Map3[K4, K5, K6, V]
  item4, ok = item3.Load(k3)
  if !ok {
    return
  }

  var item5  Map2[K5, K6, V]
  item5, ok = item4.Load(k4)
  if !ok {
    return
  }

  return item5.Load(k5)
}


func (m *Map7[K0, K1, K2, K3, K4, K5, K6, V]) Load6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, 
) (value V, ok bool) {
  item0 := m.inner

  var item1  Map6[K1, K2, K3, K4, K5, K6, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map5[K2, K3, K4, K5, K6, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map4[K3, K4, K5, K6, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  var item4  Map3[K4, K5, K6, V]
  item4, ok = item3.Load(k3)
  if !ok {
    return
  }

  var item5  Map2[K5, K6, V]
  item5, ok = item4.Load(k4)
  if !ok {
    return
  }

  var item6  Map1[K6, V]
  item6, ok = item5.Load(k5)
  if !ok {
    return
  }

  return item6.Load(k6)
}



type Map8[K0 comparable, K1 comparable, K2 comparable, K3 comparable, K4 comparable, K5 comparable, K6 comparable, K7 comparable, V any] struct {
  inner Map1[K0, Map7[K1,K2,K3,K4,K5,K6,K7,V]]
}

func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Load(key K0) (value Map7[K1, K2, K3, K4, K5, K6, K7, V], ok bool) {
  return m.inner.Load(key)
}



func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Load1(k0 K0, k1 K1, 
) (value  Map6[K2, K3, K4, K5, K6, K7, V] , ok bool) {
  item0 := m.inner

  var item1  Map7[K1, K2, K3, K4, K5, K6, K7, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  return item1.Load(k1)
}


func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Load2(k0 K0, k1 K1, k2 K2, 
) (value  Map5[K3, K4, K5, K6, K7, V] , ok bool) {
  item0 := m.inner

  var item1  Map7[K1, K2, K3, K4, K5, K6, K7, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map6[K2, K3, K4, K5, K6, K7, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  return item2.Load(k2)
}


func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Load3(k0 K0, k1 K1, k2 K2, k3 K3, 
) (value  Map4[K4, K5, K6, K7, V] , ok bool) {
  item0 := m.inner

  var item1  Map7[K1, K2, K3, K4, K5, K6, K7, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map6[K2, K3, K4, K5, K6, K7, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map5[K3, K4, K5, K6, K7, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  return item3.Load(k3)
}


func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Load4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, 
) (value  Map3[K5, K6, K7, V] , ok bool) {
  item0 := m.inner

  var item1  Map7[K1, K2, K3, K4, K5, K6, K7, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map6[K2, K3, K4, K5, K6, K7, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map5[K3, K4, K5, K6, K7, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  var item4  Map4[K4, K5, K6, K7, V]
  item4, ok = item3.Load(k3)
  if !ok {
    return
  }

  return item4.Load(k4)
}


func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Load5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, 
) (value  Map2[K6, K7, V] , ok bool) {
  item0 := m.inner

  var item1  Map7[K1, K2, K3, K4, K5, K6, K7, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map6[K2, K3, K4, K5, K6, K7, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map5[K3, K4, K5, K6, K7, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  var item4  Map4[K4, K5, K6, K7, V]
  item4, ok = item3.Load(k3)
  if !ok {
    return
  }

  var item5  Map3[K5, K6, K7, V]
  item5, ok = item4.Load(k4)
  if !ok {
    return
  }

  return item5.Load(k5)
}


func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Load6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, 
) (value  Map1[K7, V] , ok bool) {
  item0 := m.inner

  var item1  Map7[K1, K2, K3, K4, K5, K6, K7, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map6[K2, K3, K4, K5, K6, K7, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map5[K3, K4, K5, K6, K7, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  var item4  Map4[K4, K5, K6, K7, V]
  item4, ok = item3.Load(k3)
  if !ok {
    return
  }

  var item5  Map3[K5, K6, K7, V]
  item5, ok = item4.Load(k4)
  if !ok {
    return
  }

  var item6  Map2[K6, K7, V]
  item6, ok = item5.Load(k5)
  if !ok {
    return
  }

  return item6.Load(k6)
}


func (m *Map8[K0, K1, K2, K3, K4, K5, K6, K7, V]) Load7(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, 
) (value V, ok bool) {
  item0 := m.inner

  var item1  Map7[K1, K2, K3, K4, K5, K6, K7, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map6[K2, K3, K4, K5, K6, K7, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map5[K3, K4, K5, K6, K7, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  var item4  Map4[K4, K5, K6, K7, V]
  item4, ok = item3.Load(k3)
  if !ok {
    return
  }

  var item5  Map3[K5, K6, K7, V]
  item5, ok = item4.Load(k4)
  if !ok {
    return
  }

  var item6  Map2[K6, K7, V]
  item6, ok = item5.Load(k5)
  if !ok {
    return
  }

  var item7  Map1[K7, V]
  item7, ok = item6.Load(k6)
  if !ok {
    return
  }

  return item7.Load(k7)
}



type Map9[K0 comparable, K1 comparable, K2 comparable, K3 comparable, K4 comparable, K5 comparable, K6 comparable, K7 comparable, K8 comparable, V any] struct {
  inner Map1[K0, Map8[K1,K2,K3,K4,K5,K6,K7,K8,V]]
}

func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Load(key K0) (value Map8[K1, K2, K3, K4, K5, K6, K7, K8, V], ok bool) {
  return m.inner.Load(key)
}



func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Load1(k0 K0, k1 K1, 
) (value  Map7[K2, K3, K4, K5, K6, K7, K8, V] , ok bool) {
  item0 := m.inner

  var item1  Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  return item1.Load(k1)
}


func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Load2(k0 K0, k1 K1, k2 K2, 
) (value  Map6[K3, K4, K5, K6, K7, K8, V] , ok bool) {
  item0 := m.inner

  var item1  Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map7[K2, K3, K4, K5, K6, K7, K8, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  return item2.Load(k2)
}


func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Load3(k0 K0, k1 K1, k2 K2, k3 K3, 
) (value  Map5[K4, K5, K6, K7, K8, V] , ok bool) {
  item0 := m.inner

  var item1  Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map7[K2, K3, K4, K5, K6, K7, K8, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map6[K3, K4, K5, K6, K7, K8, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  return item3.Load(k3)
}


func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Load4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, 
) (value  Map4[K5, K6, K7, K8, V] , ok bool) {
  item0 := m.inner

  var item1  Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map7[K2, K3, K4, K5, K6, K7, K8, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map6[K3, K4, K5, K6, K7, K8, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  var item4  Map5[K4, K5, K6, K7, K8, V]
  item4, ok = item3.Load(k3)
  if !ok {
    return
  }

  return item4.Load(k4)
}


func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Load5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, 
) (value  Map3[K6, K7, K8, V] , ok bool) {
  item0 := m.inner

  var item1  Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map7[K2, K3, K4, K5, K6, K7, K8, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map6[K3, K4, K5, K6, K7, K8, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  var item4  Map5[K4, K5, K6, K7, K8, V]
  item4, ok = item3.Load(k3)
  if !ok {
    return
  }

  var item5  Map4[K5, K6, K7, K8, V]
  item5, ok = item4.Load(k4)
  if !ok {
    return
  }

  return item5.Load(k5)
}


func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Load6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, 
) (value  Map2[K7, K8, V] , ok bool) {
  item0 := m.inner

  var item1  Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map7[K2, K3, K4, K5, K6, K7, K8, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map6[K3, K4, K5, K6, K7, K8, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  var item4  Map5[K4, K5, K6, K7, K8, V]
  item4, ok = item3.Load(k3)
  if !ok {
    return
  }

  var item5  Map4[K5, K6, K7, K8, V]
  item5, ok = item4.Load(k4)
  if !ok {
    return
  }

  var item6  Map3[K6, K7, K8, V]
  item6, ok = item5.Load(k5)
  if !ok {
    return
  }

  return item6.Load(k6)
}


func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Load7(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, 
) (value  Map1[K8, V] , ok bool) {
  item0 := m.inner

  var item1  Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map7[K2, K3, K4, K5, K6, K7, K8, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map6[K3, K4, K5, K6, K7, K8, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  var item4  Map5[K4, K5, K6, K7, K8, V]
  item4, ok = item3.Load(k3)
  if !ok {
    return
  }

  var item5  Map4[K5, K6, K7, K8, V]
  item5, ok = item4.Load(k4)
  if !ok {
    return
  }

  var item6  Map3[K6, K7, K8, V]
  item6, ok = item5.Load(k5)
  if !ok {
    return
  }

  var item7  Map2[K7, K8, V]
  item7, ok = item6.Load(k6)
  if !ok {
    return
  }

  return item7.Load(k7)
}


func (m *Map9[K0, K1, K2, K3, K4, K5, K6, K7, K8, V]) Load8(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, 
) (value V, ok bool) {
  item0 := m.inner

  var item1  Map8[K1, K2, K3, K4, K5, K6, K7, K8, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map7[K2, K3, K4, K5, K6, K7, K8, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map6[K3, K4, K5, K6, K7, K8, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  var item4  Map5[K4, K5, K6, K7, K8, V]
  item4, ok = item3.Load(k3)
  if !ok {
    return
  }

  var item5  Map4[K5, K6, K7, K8, V]
  item5, ok = item4.Load(k4)
  if !ok {
    return
  }

  var item6  Map3[K6, K7, K8, V]
  item6, ok = item5.Load(k5)
  if !ok {
    return
  }

  var item7  Map2[K7, K8, V]
  item7, ok = item6.Load(k6)
  if !ok {
    return
  }

  var item8  Map1[K8, V]
  item8, ok = item7.Load(k7)
  if !ok {
    return
  }

  return item8.Load(k8)
}



type Map10[K0 comparable, K1 comparable, K2 comparable, K3 comparable, K4 comparable, K5 comparable, K6 comparable, K7 comparable, K8 comparable, K9 comparable, V any] struct {
  inner Map1[K0, Map9[K1,K2,K3,K4,K5,K6,K7,K8,K9,V]]
}

func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Load(key K0) (value Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V], ok bool) {
  return m.inner.Load(key)
}



func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Load1(k0 K0, k1 K1, 
) (value  Map8[K2, K3, K4, K5, K6, K7, K8, K9, V] , ok bool) {
  item0 := m.inner

  var item1  Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  return item1.Load(k1)
}


func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Load2(k0 K0, k1 K1, k2 K2, 
) (value  Map7[K3, K4, K5, K6, K7, K8, K9, V] , ok bool) {
  item0 := m.inner

  var item1  Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  return item2.Load(k2)
}


func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Load3(k0 K0, k1 K1, k2 K2, k3 K3, 
) (value  Map6[K4, K5, K6, K7, K8, K9, V] , ok bool) {
  item0 := m.inner

  var item1  Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map7[K3, K4, K5, K6, K7, K8, K9, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  return item3.Load(k3)
}


func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Load4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, 
) (value  Map5[K5, K6, K7, K8, K9, V] , ok bool) {
  item0 := m.inner

  var item1  Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map7[K3, K4, K5, K6, K7, K8, K9, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  var item4  Map6[K4, K5, K6, K7, K8, K9, V]
  item4, ok = item3.Load(k3)
  if !ok {
    return
  }

  return item4.Load(k4)
}


func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Load5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, 
) (value  Map4[K6, K7, K8, K9, V] , ok bool) {
  item0 := m.inner

  var item1  Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map7[K3, K4, K5, K6, K7, K8, K9, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  var item4  Map6[K4, K5, K6, K7, K8, K9, V]
  item4, ok = item3.Load(k3)
  if !ok {
    return
  }

  var item5  Map5[K5, K6, K7, K8, K9, V]
  item5, ok = item4.Load(k4)
  if !ok {
    return
  }

  return item5.Load(k5)
}


func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Load6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, 
) (value  Map3[K7, K8, K9, V] , ok bool) {
  item0 := m.inner

  var item1  Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map7[K3, K4, K5, K6, K7, K8, K9, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  var item4  Map6[K4, K5, K6, K7, K8, K9, V]
  item4, ok = item3.Load(k3)
  if !ok {
    return
  }

  var item5  Map5[K5, K6, K7, K8, K9, V]
  item5, ok = item4.Load(k4)
  if !ok {
    return
  }

  var item6  Map4[K6, K7, K8, K9, V]
  item6, ok = item5.Load(k5)
  if !ok {
    return
  }

  return item6.Load(k6)
}


func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Load7(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, 
) (value  Map2[K8, K9, V] , ok bool) {
  item0 := m.inner

  var item1  Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map7[K3, K4, K5, K6, K7, K8, K9, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  var item4  Map6[K4, K5, K6, K7, K8, K9, V]
  item4, ok = item3.Load(k3)
  if !ok {
    return
  }

  var item5  Map5[K5, K6, K7, K8, K9, V]
  item5, ok = item4.Load(k4)
  if !ok {
    return
  }

  var item6  Map4[K6, K7, K8, K9, V]
  item6, ok = item5.Load(k5)
  if !ok {
    return
  }

  var item7  Map3[K7, K8, K9, V]
  item7, ok = item6.Load(k6)
  if !ok {
    return
  }

  return item7.Load(k7)
}


func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Load8(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, 
) (value  Map1[K9, V] , ok bool) {
  item0 := m.inner

  var item1  Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map7[K3, K4, K5, K6, K7, K8, K9, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  var item4  Map6[K4, K5, K6, K7, K8, K9, V]
  item4, ok = item3.Load(k3)
  if !ok {
    return
  }

  var item5  Map5[K5, K6, K7, K8, K9, V]
  item5, ok = item4.Load(k4)
  if !ok {
    return
  }

  var item6  Map4[K6, K7, K8, K9, V]
  item6, ok = item5.Load(k5)
  if !ok {
    return
  }

  var item7  Map3[K7, K8, K9, V]
  item7, ok = item6.Load(k6)
  if !ok {
    return
  }

  var item8  Map2[K8, K9, V]
  item8, ok = item7.Load(k7)
  if !ok {
    return
  }

  return item8.Load(k8)
}


func (m *Map10[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, V]) Load9(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9, 
) (value V, ok bool) {
  item0 := m.inner

  var item1  Map9[K1, K2, K3, K4, K5, K6, K7, K8, K9, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map8[K2, K3, K4, K5, K6, K7, K8, K9, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map7[K3, K4, K5, K6, K7, K8, K9, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  var item4  Map6[K4, K5, K6, K7, K8, K9, V]
  item4, ok = item3.Load(k3)
  if !ok {
    return
  }

  var item5  Map5[K5, K6, K7, K8, K9, V]
  item5, ok = item4.Load(k4)
  if !ok {
    return
  }

  var item6  Map4[K6, K7, K8, K9, V]
  item6, ok = item5.Load(k5)
  if !ok {
    return
  }

  var item7  Map3[K7, K8, K9, V]
  item7, ok = item6.Load(k6)
  if !ok {
    return
  }

  var item8  Map2[K8, K9, V]
  item8, ok = item7.Load(k7)
  if !ok {
    return
  }

  var item9  Map1[K9, V]
  item9, ok = item8.Load(k8)
  if !ok {
    return
  }

  return item9.Load(k9)
}



type Map11[K0 comparable, K1 comparable, K2 comparable, K3 comparable, K4 comparable, K5 comparable, K6 comparable, K7 comparable, K8 comparable, K9 comparable, K10 comparable, V any] struct {
  inner Map1[K0, Map10[K1,K2,K3,K4,K5,K6,K7,K8,K9,K10,V]]
}

func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Load(key K0) (value Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V], ok bool) {
  return m.inner.Load(key)
}



func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Load1(k0 K0, k1 K1, 
) (value  Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V] , ok bool) {
  item0 := m.inner

  var item1  Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  return item1.Load(k1)
}


func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Load2(k0 K0, k1 K1, k2 K2, 
) (value  Map8[K3, K4, K5, K6, K7, K8, K9, K10, V] , ok bool) {
  item0 := m.inner

  var item1  Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  return item2.Load(k2)
}


func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Load3(k0 K0, k1 K1, k2 K2, k3 K3, 
) (value  Map7[K4, K5, K6, K7, K8, K9, K10, V] , ok bool) {
  item0 := m.inner

  var item1  Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  return item3.Load(k3)
}


func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Load4(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, 
) (value  Map6[K5, K6, K7, K8, K9, K10, V] , ok bool) {
  item0 := m.inner

  var item1  Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  var item4  Map7[K4, K5, K6, K7, K8, K9, K10, V]
  item4, ok = item3.Load(k3)
  if !ok {
    return
  }

  return item4.Load(k4)
}


func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Load5(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, 
) (value  Map5[K6, K7, K8, K9, K10, V] , ok bool) {
  item0 := m.inner

  var item1  Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  var item4  Map7[K4, K5, K6, K7, K8, K9, K10, V]
  item4, ok = item3.Load(k3)
  if !ok {
    return
  }

  var item5  Map6[K5, K6, K7, K8, K9, K10, V]
  item5, ok = item4.Load(k4)
  if !ok {
    return
  }

  return item5.Load(k5)
}


func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Load6(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, 
) (value  Map4[K7, K8, K9, K10, V] , ok bool) {
  item0 := m.inner

  var item1  Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  var item4  Map7[K4, K5, K6, K7, K8, K9, K10, V]
  item4, ok = item3.Load(k3)
  if !ok {
    return
  }

  var item5  Map6[K5, K6, K7, K8, K9, K10, V]
  item5, ok = item4.Load(k4)
  if !ok {
    return
  }

  var item6  Map5[K6, K7, K8, K9, K10, V]
  item6, ok = item5.Load(k5)
  if !ok {
    return
  }

  return item6.Load(k6)
}


func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Load7(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, 
) (value  Map3[K8, K9, K10, V] , ok bool) {
  item0 := m.inner

  var item1  Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  var item4  Map7[K4, K5, K6, K7, K8, K9, K10, V]
  item4, ok = item3.Load(k3)
  if !ok {
    return
  }

  var item5  Map6[K5, K6, K7, K8, K9, K10, V]
  item5, ok = item4.Load(k4)
  if !ok {
    return
  }

  var item6  Map5[K6, K7, K8, K9, K10, V]
  item6, ok = item5.Load(k5)
  if !ok {
    return
  }

  var item7  Map4[K7, K8, K9, K10, V]
  item7, ok = item6.Load(k6)
  if !ok {
    return
  }

  return item7.Load(k7)
}


func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Load8(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, 
) (value  Map2[K9, K10, V] , ok bool) {
  item0 := m.inner

  var item1  Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  var item4  Map7[K4, K5, K6, K7, K8, K9, K10, V]
  item4, ok = item3.Load(k3)
  if !ok {
    return
  }

  var item5  Map6[K5, K6, K7, K8, K9, K10, V]
  item5, ok = item4.Load(k4)
  if !ok {
    return
  }

  var item6  Map5[K6, K7, K8, K9, K10, V]
  item6, ok = item5.Load(k5)
  if !ok {
    return
  }

  var item7  Map4[K7, K8, K9, K10, V]
  item7, ok = item6.Load(k6)
  if !ok {
    return
  }

  var item8  Map3[K8, K9, K10, V]
  item8, ok = item7.Load(k7)
  if !ok {
    return
  }

  return item8.Load(k8)
}


func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Load9(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9, 
) (value  Map1[K10, V] , ok bool) {
  item0 := m.inner

  var item1  Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  var item4  Map7[K4, K5, K6, K7, K8, K9, K10, V]
  item4, ok = item3.Load(k3)
  if !ok {
    return
  }

  var item5  Map6[K5, K6, K7, K8, K9, K10, V]
  item5, ok = item4.Load(k4)
  if !ok {
    return
  }

  var item6  Map5[K6, K7, K8, K9, K10, V]
  item6, ok = item5.Load(k5)
  if !ok {
    return
  }

  var item7  Map4[K7, K8, K9, K10, V]
  item7, ok = item6.Load(k6)
  if !ok {
    return
  }

  var item8  Map3[K8, K9, K10, V]
  item8, ok = item7.Load(k7)
  if !ok {
    return
  }

  var item9  Map2[K9, K10, V]
  item9, ok = item8.Load(k8)
  if !ok {
    return
  }

  return item9.Load(k9)
}


func (m *Map11[K0, K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]) Load10(k0 K0, k1 K1, k2 K2, k3 K3, k4 K4, k5 K5, k6 K6, k7 K7, k8 K8, k9 K9, k10 K10, 
) (value V, ok bool) {
  item0 := m.inner

  var item1  Map10[K1, K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
  item1, ok = item0.Load(k0)
  if !ok {
    return
  }

  var item2  Map9[K2, K3, K4, K5, K6, K7, K8, K9, K10, V]
  item2, ok = item1.Load(k1)
  if !ok {
    return
  }

  var item3  Map8[K3, K4, K5, K6, K7, K8, K9, K10, V]
  item3, ok = item2.Load(k2)
  if !ok {
    return
  }

  var item4  Map7[K4, K5, K6, K7, K8, K9, K10, V]
  item4, ok = item3.Load(k3)
  if !ok {
    return
  }

  var item5  Map6[K5, K6, K7, K8, K9, K10, V]
  item5, ok = item4.Load(k4)
  if !ok {
    return
  }

  var item6  Map5[K6, K7, K8, K9, K10, V]
  item6, ok = item5.Load(k5)
  if !ok {
    return
  }

  var item7  Map4[K7, K8, K9, K10, V]
  item7, ok = item6.Load(k6)
  if !ok {
    return
  }

  var item8  Map3[K8, K9, K10, V]
  item8, ok = item7.Load(k7)
  if !ok {
    return
  }

  var item9  Map2[K9, K10, V]
  item9, ok = item8.Load(k8)
  if !ok {
    return
  }

  var item10  Map1[K10, V]
  item10, ok = item9.Load(k9)
  if !ok {
    return
  }

  return item10.Load(k10)
}


