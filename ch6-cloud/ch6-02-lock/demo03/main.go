package main

import (
	"sync"
)


// ----------------------Lock尝试锁 ---------------------------------
type Lock struct {
	c chan struct{}
}

//锁住尝试锁返回加锁结果
func (l Lock) Lock() bool {
	lockResult := false
	select {
	case <-l.c:
		lockResult = true
	default:
	}
	return lockResult
}

//解锁尝试锁
func (l Lock) Unlock() {
	l.c <- struct{}{}
}
//----------------------------------------------------------------
// NewLock 生成一个尝试锁
func NewLock() Lock {
	var l Lock
	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}
	return l
}


var counter int
//因为我们的逻辑限定每个goroutine只有成功执行了`Lock`才会继续执行后续逻辑，
//因此在`Unlock`时可以保证Lock结构体中的channel一定是空，从而不会阻塞，也不会失败。
//上面的代码使用了大小为1的channel来模拟trylock，理论上还可以使用标准库中的CAS来实现相同的功能且成本更低，读者可以自行尝试。
func main() {
	var l = NewLock()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !l.Lock() {
				// log error
				println("lock failed")
				return
			}
			counter++
			println("current counter", counter)
			l.Unlock()
		}()
	}
	wg.Wait()
}