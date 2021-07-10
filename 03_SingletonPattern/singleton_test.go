package singletonpattern

import (
	"sync"
	"testing"
)

func TestSingleton1(t *testing.T) {
	s1 := getInstance()
	s2 := getInstance()
	if s1 != s2 {
		t.Fatal("error")
	}
}

const parCount = 100

func TestParallelSingleton(t *testing.T) {
	start := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instances := [parCount]Singleton{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			//协程阻塞，等待channel被关闭才能继续运行
			<-start
			instances[index] = getInstance()
			wg.Done()
		}(i)
	}
	// 关闭channel，所有协程同时开始运行，实现并行(parallel)
	close(start)
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}
