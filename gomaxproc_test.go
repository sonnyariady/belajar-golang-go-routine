package belajar_golang_goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGomaxProc(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1) //menambah jumlah go routine. ada 102 krn 100 dari iterasi dan 2 (jalankan program dan garbage collection)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}
	totalcpu := runtime.NumCPU()
	fmt.Println("Total CPU:", totalcpu)

	jmlthread := runtime.GOMAXPROCS(-1) //kasih 0 kebawah agar tidak mengubah
	fmt.Println("Total Thread:", jmlthread)

	totalgoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine:", totalgoroutine)

	group.Wait()
}

func TestChangeThreadNumber(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1) //menambah jumlah go routine. ada 102 krn 100 dari iterasi dan 2 (jalankan program dan garbage collection)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}
	totalcpu := runtime.NumCPU()
	fmt.Println("Total CPU:", totalcpu)

	runtime.GOMAXPROCS(20)
	jmlthread := runtime.GOMAXPROCS(-1) //kasih 0 kebawah agar tidak mengubah
	fmt.Println("Total Thread:", jmlthread)

	totalgoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine:", totalgoroutine)

	group.Wait()
}
