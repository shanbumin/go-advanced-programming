package main

import "sync"

var (
	instance *singleton
	once     sync.Once
)

type singleton struct {}




func Instance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}




func main() {
	s:= Instance()
	_ = s
}
