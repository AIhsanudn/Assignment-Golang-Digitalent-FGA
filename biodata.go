package main

import (
	"fmt"
	"os"
	"strconv"
)

// Struct untuk merepresentasikan data teman
type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// Function untuk mendapatkan data teman berdasarkan nomor absen
func getDataTeman(absen int) Teman {

	var teman Teman

	switch absen {
	case 1:
		teman = Teman{Nama: "Jainudin", Alamat: "Jl. Kemayoran No. 123", Pekerjaan: "Devloper", Alasan: "Ingin mendalami bahasa Go"}
	case 2:
		teman = Teman{Nama: "Ahmad", Alamat: "Jl. Kebagiusan No. 456", Pekerjaan: "Designer", Alasan: "Tertarik dengan performa Go"}
	case 3:
		teman = Teman{Nama: "Raisa", Alamat: "Jl. Jakasempurna No. 888", Pekerjaan: "Penyanyi", Alasan: "Ingin mempelajari hal baru"}
	default:
		fmt.Println("Teman dengan nomor absen tersebut tidak ditemukan")
		os.Exit(1)
	}

	return teman
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run biodata.go <nomor_absen>")
		os.Exit(1)
	}

	absen, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Nomor absen harus berupa bilangan bulat")
		os.Exit(1)
	}

	teman := getDataTeman(absen)

	// Menampilkan data teman
	fmt.Println("Nama:", teman.Nama)
	fmt.Println("Alamat:", teman.Alamat)
	fmt.Println("Pekerjaan:", teman.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", teman.Alasan)
}
