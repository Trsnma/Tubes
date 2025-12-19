package main

import (
	"fmt"
)

func main() {

	var (
		harga          []int
		inputHarga     int
		jumlahBarang   int
		totalPembelian int
		setelahDiskon  float32
		diskon         float32
		dummy          string
	)

	const (
		batasBarang int     = 100
		diskon5  float32 = 0.05
		diskon10 float32 = 0.10
		diskon20 float32 = 0.20
	)

	fmt.Println(">> Selamat datang di GTR MART, silahkan masukkan harga dari setiap barang yang anda beli\n>> Masukkan harga HARUS menggunakan tanda titik (desimal)\n>> Input angka 0.0 untuk menyelesaikan\n")

	fmt.Print("\nMasukkan jumlah barang yang anda beli z: ")
	jumlahBarang = -1
	fmt.Scan(&jumlahBarang)

	if jumlahBarang > batasBarang {
		fmt.Println(">> [ERROR] Jumlah barang tidak boleh lebih dari 1000 ")
		fmt.Print("Masukkan jumlah barang yang anda beli : ")
		fmt.Scan(&jumlahBarang)
	}

	if jumlahBarang > 0 && jumlahBarang <= batasBarang {
		totalPembelian = 0
		i := 0
		for i < jumlahBarang {
			fmt.Printf("Masukkan harga barang ke - %d : ", i+1)
			inputHarga = -1
			fmt.Scan(&inputHarga)

			if inputHarga > 0 {
				harga = append(harga, inputHarga)
				totalPembelian += inputHarga
				i++
			} else if inputHarga == 0 {
				fmt.Println("Input harga tidak boleh 0 sebelum semua barang diinput")
			} else if inputHarga == 101 {
				i--
			} else {
				fmt.Println("Input harga tidak boleh negatif")
				fmt.Scanf("%s", &dummy)
			}
		}
	}

	l := len(harga)

	for i := 0; i < l-1; i++ {
		for j := 0; j < l-1-i; j++ {
			if harga[j] > harga[j+1] {
				harga[j], harga[j+1] = harga[j+1], harga[j]
			}
		}
	}

	// --- Perhitungan Diskon ---

	if totalPembelian >= 100000 && totalPembelian < 200000 {
		diskon = diskon5
	} else if totalPembelian >= 200000 && totalPembelian < 300000 {
		diskon = diskon10
	} else if totalPembelian >= 300000 {
		diskon = diskon20
	} else if totalPembelian == 0 || totalPembelian < 100000 {
		diskon = 0
	}

	setelahDiskon = float32(totalPembelian) - (float32(totalPembelian) * diskon)
	setelahDiskon = float32(int(setelahDiskon + 0.5)) // pembulatan ke atas

	fmt.Println("\n=========================================================================================")
	if jumlahBarang <= 0 {
		fmt.Println("\n>> KAMU TIDAK MEMBELI BARANG APAPUN ")
	} else if len(harga) != jumlahBarang {
		fmt.Println("\n>> PEMBELIAN ANDA DIBATALKAN")
	}
	fmt.Println("\nUrutan Harga :", harga)
	fmt.Printf("Total Pembelian Anda : %d\n", totalPembelian)
	fmt.Println("Diskon yang anda peroleh :", diskon*100, "%")
	fmt.Printf("Total pembelian setelah diskon : %.f\n", setelahDiskon)

	fmt.Println("\n=========================================================================================")
}
