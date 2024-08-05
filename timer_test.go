package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTicker(5 * time.Second)
	fmt.Println("Waktu saat ini:", time.Now())

	time := <-timer.C
	fmt.Println("Waktu setelah timer:", time)
}

func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println("Waktu saat ini:", time.Now())

	time := <-channel
	fmt.Println("Waktu setelah timer:", time)
}

func TestAfterFunction(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(2)

	time.AfterFunc(3*time.Second, func() {
		fmt.Println("Dijalankan setelah 3 detik :", time.Now())
		group.Done()
	})

	time.AfterFunc(2*time.Second, func() {
		fmt.Println("Dijalankan setelah 2 detik :", time.Now())
		group.Done()
	})

	fmt.Println("Waktu awal: ", time.Now())

	group.Wait()
}
