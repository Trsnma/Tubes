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
		kosong         bool
		namaPembeli    string
	)

	const (
		diskon5  float32 = 0.05
		diskon10 float32 = 0.10
		diskon20 float32 = 0.20
		diskon30 float32 = 0.30
	)

	fmt.Println("\n==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==")
	fmt.Println("==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==")
	fmt.Println("==                                      WELCOME TO GTR MART                                     ==")
	fmt.Println("==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==")
	fmt.Println("==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==  ==")

	fmt.Println("\n>> Selamat datang di GTR MART by Muhammad Ghassan Denaprinsyah and Muhammad Trisnatama Derdiansyah.\n>> Silahkan masukkan harga dari setiap barang yang anda beli\n>> Masukkan harga tanpa menggunakan tanda titik atau koma (desimal)\n>> Masukkan harga 0 (nol) untuk membatalkan pembelian")

	// --- Input Jumlah Barang ---
	fmt.Print("\nMasukkan jumlah barang yang anda beli : ")
	jumlahBarang = -1
	fmt.Scan(&jumlahBarang)

	// --- Validasi Jumlah Barang ---
	for jumlahBarang < 0 {
		fmt.Println(">> [ERROR] Jumlah barang tidak boleh negatif dan berupa karakter")
		fmt.Print("Masukkan jumlah barang yang anda beli : ")
		fmt.Scan(&jumlahBarang)
	}

	if jumlahBarang == 0 {
		kosong = true
	}

	// --- Input Harga Barang ---
	if jumlahBarang > 0 {
		totalPembelian = 0
		i := 0
		for i < jumlahBarang {
			fmt.Printf("Masukkan harga barang ke - %d : ", i+1)
			fmt.Scan(&inputHarga)

			if inputHarga > 0 {
				harga = append(harga, inputHarga)
				totalPembelian += inputHarga
				i++
			} else if inputHarga < 0 {
				fmt.Println(">> [ERROR] Harga barang tidak boleh negatif dan berupa karakter, silahkan masukkan kembali")
			} else if inputHarga == 0 {
				i = jumlahBarang
			}
		}

		// --- Pengurutan Harga (Bubble Sort) ---
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
		} else if totalPembelian >= 300000 && totalPembelian < 500000 {
			diskon = diskon20
		} else if totalPembelian >= 500000 {
			diskon = diskon30
		} else if totalPembelian == 0 || totalPembelian < 100000 {
			diskon = 0
		}

	}
	setelahDiskon = float32(totalPembelian) - (float32(totalPembelian) * diskon)

	// --- Output ---
	fmt.Println("\n==================================================================================================")
	// --- Kondisi Pembatalan Pembelian ---
	if kosong == true {
		fmt.Println("\n>> ANDA TIDAK MEMBELI BARANG APAPUN ")
	} else if inputHarga == 0 {
		fmt.Println("\n>> PEMBELIAN ANDA DIBATALKAN")
	} else if inputHarga > 0 {
		fmt.Println("\nUrutan Harga :", harga)
		fmt.Printf("Total Pembelian Anda : %d\n", totalPembelian)
		fmt.Println("Diskon yang anda peroleh :", diskon*100, "%")
		fmt.Printf("Total pembelian setelah diskon : %.f\n", setelahDiskon)
		fmt.Print("Masukkan nama pembeli : ")
		fmt.Scan(&namaPembeli)
		fmt.Println("Terimakasih telah berbelanja di GTR Mart", namaPembeli)
	}

	fmt.Println("\n==================================================================================================")

}
