package main

import (
	"sync"
	"fmt"
)
//muutex
//UnLock() 必须在Lock() 之后使用，如果UnLock() 在Lock()之前使用，报错(fatal error: sync: unlock of unlocked mutex)
//如进行多次加锁或者说 调用两次lock()就会报fatal error: all goroutines are asleep - deadlock!



func main() {


//type Mutex
//func(m *Mutex) Lock()
//func (m *Mutex) UnLock()
//
//	type RWMutex
//	func (rw *RWMutex) Lock()
//	func (rw *RWMutex) RLock()
//	func (rw *RWMutex) RLocker() Locker
//	func (rw *RWMutex) RUnlock()
//	func (rw *RWMutex) Unlock()

	var l * sync.Mutex
	l =new(sync.Mutex)
	l.Lock()
	l.Lock()
	l.Unlock()
	fmt.Println("1")
}
