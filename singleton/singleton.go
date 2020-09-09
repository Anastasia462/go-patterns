package singleton

import "sync"

type Singleton struct {
	name string
}

var (
	instance *Singleton
	once     sync.Once
)

// GetInstance ...
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{name: "name"}
	})
	return instance
}

