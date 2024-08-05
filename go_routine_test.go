package belajar_golang_goroutines

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func HaiHowAreYou(nama string) {
	fmt.Printf("Hai %s. Apa kabar?\n", nama)
}

func TestHaiHowAreYou(t *testing.T) {
	go HaiHowAreYou("Andre")
	fmt.Println("Coba dulu ya")

	time.Sleep(1 * time.Second)
}

func DisplayAngka(angka int) {
	fmt.Println("Hitung angka ke-", angka)
}

func TestBanyakGoRoutine(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go DisplayAngka(i)
	}

	//no goroutine 10rb = 4.61 s
	//no goroutine 10rb = 0.35 s
	time.Sleep(5 * time.Second)
	fmt.Println("Jumlah goroutine: ", runtime.NumGoroutine())
}
