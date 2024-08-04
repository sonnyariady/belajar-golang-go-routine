package belajar_golang_goroutines

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

type Mahasiswa struct {
	Nama    string
	Jurusan string
}

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Siapa kah ini"
		fmt.Println("Selesai kirim")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func TestChannelParameter(t *testing.T) {
	channel := make(chan string)

	go SapaAgus(channel)

	go SapaJoko(channel)

	pesan1 := <-channel
	fmt.Println("Pesan1: ", pesan1)
	pesan2 := <-channel
	fmt.Println("Pesan2: ", pesan2)

	time.Sleep(5 * time.Second)
}

func SapaAgus(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Dari agus nih"
}

func SapaJoko(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Dari Joko"
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan Mahasiswa)
	defer close(channel)

	go MhsChanIn(channel)
	go MhsChanOut(channel)

	time.Sleep(3 * time.Second)

}

func MhsChanIn(channel chan<- Mahasiswa) {
	time.Sleep(2 * time.Second)

	rand.Seed(time.Now().UnixNano())

	// Generate a random integer between 0 and 99
	acakint := rand.Intn(4)

	var Mhs Mahasiswa

	switch acakint {
	case 0:
		Mhs.Nama = "Ihsan"
		Mhs.Jurusan = "Ekonomi"
	case 1:
		Mhs.Nama = "Dicky"
		Mhs.Jurusan = "Telektro"
	case 2:
		Mhs.Nama = "Ahmad"
		Mhs.Jurusan = "Akuntansi"
	default:
		Mhs.Nama = "Dedi"
		Mhs.Jurusan = "Sistem Informasi"
	}

	channel <- Mhs

}

func MhsChanOut(channel <-chan Mahasiswa) {
	data := <-channel
	fmt.Println(data)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan Mahasiswa, 3)
	defer close(channel)
	fmt.Println(cap(channel))
	fmt.Println(len(channel))

	go func() {
		channel <- Mahasiswa{"Budi", "Sosiologi"}
		channel <- Mahasiswa{"Jaka", "Ilmu Hukum"}
		channel <- Mahasiswa{"Dani", "Pendidikan Olahraga"}
	}()

	fmt.Println("Panjang terisi:", len(channel))

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()
	//fmt.Println(<-channel) //melebihi, potensi deadlock

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Data ke-" + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Terima : ", data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan Mahasiswa)
	channel2 := make(chan Mahasiswa)
	channel3 := make(chan Mahasiswa)

	defer close(channel1)
	defer close(channel2)
	defer close(channel3)

	go func() {
		time.Sleep(2 * time.Second)
		channel1 <- Mahasiswa{Nama: "Andi", Jurusan: "Informatika"}
	}()

	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- Mahasiswa{Nama: "Budi", Jurusan: "Ekonomi"}
	}()

	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- Mahasiswa{Nama: "Johan", Jurusan: "Fisika"}
	}()

	counter := 0
	for {
		select {
		case msg1 := <-channel1:
			fmt.Println("Received:", msg1)
			counter++
		case msg2 := <-channel2:
			fmt.Println("Received:", msg2)
			counter++
		case msg3 := <-channel3:
			fmt.Println("Received:", msg3)
			counter++
			//default:
			//	fmt.Println("Menunggu data")
		}
		if counter == 3 {
			break
		}

	}
	fmt.Println("Selesai")
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan Mahasiswa)
	channel2 := make(chan Mahasiswa)
	channel3 := make(chan Mahasiswa)

	defer close(channel1)
	defer close(channel2)
	defer close(channel3)

	go func() {
		time.Sleep(2 * time.Second)
		channel1 <- Mahasiswa{Nama: "Andi", Jurusan: "Informatika"}
	}()

	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- Mahasiswa{Nama: "Budi", Jurusan: "Ekonomi"}
	}()

	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- Mahasiswa{Nama: "Johan", Jurusan: "Fisika"}
	}()

	counter := 0
	for {
		select {
		case msg1 := <-channel1:
			fmt.Println("Received:", msg1)
			counter++
		case msg2 := <-channel2:
			fmt.Println("Received:", msg2)
			counter++
		case msg3 := <-channel3:
			fmt.Println("Received:", msg3)
			counter++
		default:
			fmt.Println("Menunggu data")
		}
		if counter == 3 {
			break
		}

	}
	fmt.Println("Selesai")
}
