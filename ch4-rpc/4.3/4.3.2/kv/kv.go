package kv

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type KVStoreService struct {
	m      map[string]string //是一个map类型，用于存储KV数据。
	filter map[string]func(key string) //对应每个Watch调用时定义的过滤器函数列表。
	mu     sync.Mutex //互斥锁，用于在多个Goroutine访问或修改时对其它成员提供保护。

}


func (p *KVStoreService) Get(key string, value *string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if v, ok := p.m[key]; ok {
		*value = v
		return nil
	}

	return fmt.Errorf("not found")
}

//在Set方法中，输入参数是key和value组成的数组，用一个匿名的空结构体表示忽略了输出参数。
//当修改某个key对应的值时会调用每一个过滤器函数。
func (p *KVStoreService) Set(kv [2]string, reply *struct{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	key, value := kv[0], kv[1]

	if oldValue := p.m[key]; oldValue != value {
		for _, fn := range p.filter {
			fn(key)
		}
	}

	p.m[key] = value
	return nil
}

func (p *KVStoreService) Watch(timeoutSecond int, keyChanged *string) error {
	id := fmt.Sprintf("watch-%s-%03d", time.Now(), rand.Int())
	ch := make(chan string, 10) // buffered

	p.mu.Lock()
	p.filter[id] = func(key string) { ch <- key }
	p.mu.Unlock()

	select {
	case <-time.After(time.Duration(timeoutSecond) * time.Second):
		return fmt.Errorf("timeout")
	case key := <-ch:
		*keyChanged = key
		return nil
	}

	return nil
}



func NewKVStoreService() *KVStoreService {
	return &KVStoreService{
		m:      make(map[string]string),
		filter: make(map[string]func(key string)),
	}
}



