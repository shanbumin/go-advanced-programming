package main

import (
	"github.com/zieckey/etcdsync"
	"log"
)


//etcdsync是Go中使用etcd的分布式锁库。 像sync.Mutex一样易于使用。
//
//实际上，有许多类似的实现都已过时，具体取决于已正式标记为已弃用的库github.com/coreos/go-etcd/etcd，并且使用有些复杂。 否则，该库非常非常简单。 用法很简单，代码很简单。


func main() {
	//声明一个锁
	m, err := etcdsync.New("/lock", 10, []string{"http://127.0.0.1:2379"})
	if m == nil || err != nil {
		log.Printf("etcdsync.New failed")
		return
	}
	//上锁
	err = m.Lock()
	if err != nil {
		log.Printf("etcdsync.Lock failed")
		return
	}

	log.Printf("etcdsync.Lock OK")
	log.Printf("Get the lock. Do something here.")

	//解锁
	err = m.Unlock()
	if err != nil {
		log.Printf("etcdsync.Unlock failed")
	} else {
		log.Printf("etcdsync.Unlock OK")
	}
}

