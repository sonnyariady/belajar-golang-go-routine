package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Nilai x=", x)
}

type AkunBank struct {
	RWMutex sync.RWMutex
	Saldo   int
}

func (acc *AkunBank) TambahSaldo(nominal int) {
	acc.RWMutex.Lock()
	acc.Saldo = acc.Saldo + nominal
	acc.RWMutex.Unlock()
}

func (acc *AkunBank) KurangSaldo(nominal int) {
	acc.RWMutex.Lock()
	acc.Saldo = acc.Saldo - nominal
	acc.RWMutex.Unlock()
}

func (acc *AkunBank) GetBalance() int {
	acc.RWMutex.RLock()
	balance := acc.Saldo
	acc.RWMutex.RUnlock()
	return balance
}

func TestReadWriteMutex(t *testing.T) {
	akun := AkunBank{}
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				akun.TambahSaldo(1)
				fmt.Println(akun.GetBalance())
			}

		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Saldo akhir:", akun.GetBalance())
}

type UserBalance struct {
	sync.Mutex
	Nama  string
	Saldo int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Saldo = user.Saldo + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1:", user1.Nama)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2:", user2.Nama)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Nama:  "Budi",
		Saldo: 100000,
	}
	user2 := UserBalance{
		Nama:  "Dedi",
		Saldo: 100000,
	}

	go Transfer(&user1, &user2, 2000)
	go Transfer(&user2, &user1, 2000)

	time.Sleep(10 * time.Second)

	fmt.Println("Saldo dari ", user1.Nama, "adalah", user1.Saldo)
	fmt.Println("Saldo dari ", user2.Nama, "adalah", user2.Saldo)
}
