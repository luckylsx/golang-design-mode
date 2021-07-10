package singletonpattern

import "sync"

type Singleton interface {
	showMessage() string
}

type singleton struct{}

var	instance  *singleton

func (s *singleton) showMessage() string {
	return "singleton message"
}

func getInstance() *singleton {
	var once sync.Once
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}
