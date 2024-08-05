package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go AddToMap(data, i, group)
	}

	group.Wait()
	data.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})
}

func AddToMap(data *sync.Map, nilai int, group *sync.WaitGroup) {
	defer group.Done()
	group.Add(1)
	data.Store(nilai, nilai)
}
