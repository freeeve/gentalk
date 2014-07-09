package main

import "sync"

// START OMIT
type ThreadSafeHashMap struct {
	// ... 
	m    map[interface{}]interface{} // OMIT
	mut  *sync.Mutex // OMIT
	less func(a, b interface{}) bool 
}

func NewThreadSafeHashMap(less func(a, b interface{}) bool) ThreadSafeHashMap { // ...
	return ThreadSafeHashMap{map[interface{}]interface{}{}, &sync.Mutex{}, less} // OMIT
} // OMIT

func (hm ThreadSafeHashMap) Put(k interface{}, v interface{}) {
	hm.mut.Lock() // OMIT
	hm.m[k] = v  // OMIT
	hm.mut.Unlock() // OMIT
} // OMIT

func (hm ThreadSafeHashMap) Get(k interface{}) interface{}, bool {
	hm.mut.Lock() // OMIT
	defer hm.mut.Unlock() // OMIT
	return hm.m[k] // OMIT
} // OMIT
// END OMIT

func main() {
	m := NewThreadSafeHashMap(func(a, b int) bool { a < b })

	m.Put(1, 3)
	x, ok := m.Get(1)

	xi := x.(int)
	// do something with x as an int
	// ... does this remind anyone of Java 1.4?
}
