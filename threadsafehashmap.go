package main

import "sync"

type ThreadSafeHashMap struct {
	m    map[interface{}]interface{}
	mut  *sync.Mutex
	less func(a, b interface{}) bool
}

func NewThreadSafeHashMap(less func(a, b interface{}) bool) ThreadSafeHashMap { // ...
	return ThreadSafeHashMap{map[interface{}]interface{}{}, &sync.Mutex{}, less} // OMIT
} // OMIT

func (hm ThreadSafeHashMap) Put(k interface{}, v interface{}) {
	hm.mut.Lock()
	hm.m[k] = v
	hm.mut.Unlock()
}

func (hm ThreadSafeHashMap) Get(k interface{}) interface{} {
	hm.mut.Lock()
	defer hm.mut.Unlock()
	return hm.m[k]
}

func main() {

}
