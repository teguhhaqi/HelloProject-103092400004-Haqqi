package main

import "fmt"

const arrMax int = 999

type produk struct {
	kodeProduk  int
	namaProduk  string
	stokProduk  int
	hargaProduk int
}

type bon struct {
	noTransaksi  int
	member       string
	belanja      [arrMax]produk
	totalBelanja int
}

type tabtab struct {
	tabProduk [arrMax]produk
	tabOmzet  [arrMax]bon
}

var data = tabtab{
	tabProduk: [arrMax]produk{
		{kodeProduk: 1, namaProduk: "MamaLemon", stokProduk: 30, hargaProduk: 12000},    // Sabun cuci piring
		{kodeProduk: 2, namaProduk: "Sania", stokProduk: 40, hargaProduk: 35000},        // Minyak goreng
		{kodeProduk: 3, namaProduk: "Indomie", stokProduk: 80, hargaProduk: 3000},       // Mi instan
		{kodeProduk: 4, namaProduk: "SoKlin", stokProduk: 50, hargaProduk: 14000},       // Detergen bubuk
		{kodeProduk: 5, namaProduk: "Rinso", stokProduk: 40, hargaProduk: 28000},        // Detergen cair
		{kodeProduk: 6, namaProduk: "Pepsodent", stokProduk: 70, hargaProduk: 15000},    // Pasta gigi
		{kodeProduk: 7, namaProduk: "Dancow", stokProduk: 40, hargaProduk: 80000},       // Susu bubuk
		{kodeProduk: 8, namaProduk: "KapalApi", stokProduk: 50, hargaProduk: 10000},     // Kopi bubuk
		{kodeProduk: 9, namaProduk: "ABC", stokProduk: 60, hargaProduk: 10000},          // Saus sambal/tomat
		{kodeProduk: 10, namaProduk: "SariRoti", stokProduk: 60, hargaProduk: 18000},    // Roti kemasan
		{kodeProduk: 11, namaProduk: "Aqua", stokProduk: 100, hargaProduk: 4000},        // Air mineral
		{kodeProduk: 12, namaProduk: "GoodDay", stokProduk: 40, hargaProduk: 5000},      // Kopi siap minum
		{kodeProduk: 13, namaProduk: "UltraMilk", stokProduk: 40, hargaProduk: 7000},    // Susu cair
		{kodeProduk: 14, namaProduk: "TehBotol", stokProduk: 45, hargaProduk: 5000},     // Minuman teh
		{kodeProduk: 15, namaProduk: "Chitato", stokProduk: 30, hargaProduk: 12000},     // Snack
		{kodeProduk: 16, namaProduk: "Taro", stokProduk: 35, hargaProduk: 10000},        // Snack ringan
		{kodeProduk: 17, namaProduk: "SilverQueen", stokProduk: 25, hargaProduk: 25000}, // Cokelat
		{kodeProduk: 18, namaProduk: "Milo", stokProduk: 30, hargaProduk: 35000},        // Minuman serbuk
		{kodeProduk: 19, namaProduk: "Nuvo", stokProduk: 55, hargaProduk: 25000},        // Sabun cair
		{kodeProduk: 20, namaProduk: "Clear", stokProduk: 45, hargaProduk: 30000},       // Shampoo
	},
}

func main() {
	//var data tabtab
	loginKaryawan(&data)
}

func listMenu() {
	fmt.Println("========== L I S T * M E N U ==========")
	fmt.Println("1. Transaksi")
	fmt.Println("2. Tambah Produk")
	fmt.Println("3. Hapus Produk")
	fmt.Println("4. Edit Produk")
	fmt.Println("5. Daftar Produk")
	fmt.Println("6. Omzet")
	fmt.Println("7. EXIT")
	fmt.Println("=======================================")
	fmt.Println("")
}

func listCustomer() {
	fmt.Println("========== M E N U * C U S T O M E R ==========")
	fmt.Println("1. Member")
	fmt.Println("2. Non-Member")
	fmt.Println("3. BACK")
	fmt.Println("===============================================")
	fmt.Println("")
}

func listTransaksi() {
	fmt.Println("========== M E N U * T R A N S A K S I ==========")
	fmt.Println("1. Bayar Belanja")
	fmt.Println("2. Tambah Belanja")
	fmt.Println("3. Hapus Belanja")
	fmt.Println("4. Edit jumlah produk belanja")
	fmt.Println("5. Daftar Produk")
	fmt.Println("6. BACK")
	fmt.Println("=================================================")
	fmt.Println("")
}

func listOmzet() {
	fmt.Println("========== M E N U * O M Z E T ==========")
	fmt.Println("1. List Transaksi")
	fmt.Println("2. Hapus Transaksi")
	fmt.Println("3. Total Omzet hari ini")
	fmt.Println("3. Back")
	fmt.Println("=========================================")
	fmt.Println("")
}

func listDaftarProduk() {
	fmt.Println("========== M E N U * L I S T ==========")
	fmt.Println("1. List berdasarkan Harga")
	fmt.Println("2. List berdasarkan Stok")
	fmt.Println("3. Cari Produk")
	fmt.Println("4. BACK")
	fmt.Println("=======================================")
	fmt.Println("")
}

func listAscendingDescending() {
	fmt.Println("========== A S C E N D I N G * D E S C E N D I N G ==========")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	fmt.Println("3. BACK")
	fmt.Println("=============================================================")
	fmt.Println("")
}

func listCariData() {
	fmt.Println("========== M E N U * C A R I ==========")
	fmt.Println("1. Kode Produk")
	fmt.Println("2. Nama Produk")
	fmt.Println("3. BACK")
	fmt.Println("========================================")
	fmt.Println("")
}

func loginKaryawan(A *tabtab) {
	var akunPass string
	var login bool = true
	for login == true {
		fmt.Println("Input akun admin: (nick_password)")
		fmt.Scan(&akunPass)
		login = akunAdmin(akunPass)
		if login == true {
			fmt.Println("Login berhasil")
			menu(A)
		} else if akunPass == "none" {
			fmt.Println("~ ~ ~ PROGRAM SELESAI, TERIMAKASIH ~ ~ ~")
			fmt.Println("")
			login = false
		} else {
			fmt.Println("Login gagal")
			fmt.Println("")
		}
	}
}

func akunAdmin(akun string) bool {
	if (akun == "admin123_pass123") || (akun == "adminadmin_passpass") || (akun == "123admin_pass321") {
		return true
	} else {
		return false
	}
}

func menu(A *tabtab) {
	var input int
	var on bool = true
	for on == true {
		listMenu()
		fmt.Println("Input: ")
		fmt.Scan(&input)
		if input == 1 {
			transaksi(A)
		} else if input == 2 {
			tambahProduk(A)
		} else if input == 3 {
			hapusProduk(A)
		} else if input == 4 {
			editProduk(A)
		} else if input == 5 {
			daftarProduk(A)
		} else if input == 6 {
			omzet(A)
		} else if input == 7 {
			on = false
		} else {
			fmt.Println("Input Salah")
		}
	}
}

func transaksi(A *tabtab) {
	var input int
	var on bool = true
	for on == true {
		listCustomer()
		fmt.Println("Input: ")
		fmt.Scan(&input)
		if input == 1 {
			transaksiMember(A)
			on = false
		} else if input == 2 {
			transaksiNonMember(A)
			on = false
		} else if input == 3 {
			on = false
		} else {
			fmt.Println("Input Salah")
		}
	}
}

func transaksiMember(A *tabtab) {
	var input, kodeP, tempIdx, tempIdxDua, tempStok, tempCount, jumlahBeliProduk, total, j, countBelanja int
	var namaP string
	var belanja [arrMax]produk
	var on bool = true
	var tambah bool
	var produkInginDiBeli bool
	var onEdit bool
	for on == true {
		total = 0
		tambah = false
		fmt.Printf("%-15s %-15s %-15s %-15s\n", "Kode Produk", "Nama Produk", "Jumlah", "Harga")
		for i := 0; i < arrMax-1; i++ {
			if belanja[i].namaProduk != "" {
				fmt.Printf("%-15d %-15s %-15d %-15d\n", belanja[i].kodeProduk, belanja[i].namaProduk, belanja[i].stokProduk, belanja[i].hargaProduk)
				total = (belanja[i].hargaProduk * belanja[i].stokProduk) + total
			}
		}
		fmt.Printf("Total transaksi: %d\n", total)
		listTransaksi()
		fmt.Println("Input: ")
		fmt.Scan(&input)
		if input == 1 { //bayar belanja
			bayarMember(A, belanja)

			on = false
		} else if input == 2 { //tambah belanja
			fmt.Println("Input kode produk: ")
			fmt.Scan(&kodeP)
			tempIdx = findProdukSquentialKode(*A, kodeP) //tempHarga <- untuk nampung function tambah yang me return idx serta harga produk
			if tempIdx == -1 {
				fmt.Println("Produk tidak tersedia")
			} else {
				for i := 0; i < arrMax-1 && tambah == false; i++ {
					if belanja[i].kodeProduk == kodeP {
						produkInginDiBeli = true
						for produkInginDiBeli {
							fmt.Println("Input jumlah produk yang ingin dibeli: ")
							fmt.Scan(&jumlahBeliProduk)
							if A.tabProduk[kodeP-1].stokProduk >= jumlahBeliProduk {
								belanja[i].kodeProduk = kodeP
								belanja[i].namaProduk = A.tabProduk[kodeP-1].namaProduk
								belanja[i].stokProduk = belanja[i].stokProduk + jumlahBeliProduk
								belanja[i].hargaProduk = A.tabProduk[tempIdx].hargaProduk - (A.tabProduk[tempIdx].hargaProduk / 10)
								tambah = true
								produkInginDiBeli = false
							} else {
								fmt.Println("Stok tidak cukup")
							}
						}
					} else if belanja[i].namaProduk == "" {
						produkInginDiBeli = true
						for produkInginDiBeli {
							fmt.Println("Input jumlah produk yang ingin dibeli: ")
							fmt.Scan(&jumlahBeliProduk)
							if A.tabProduk[kodeP-1].stokProduk >= jumlahBeliProduk {
								belanja[i].kodeProduk = kodeP
								belanja[i].namaProduk = A.tabProduk[kodeP-1].namaProduk
								belanja[i].stokProduk = jumlahBeliProduk
								belanja[i].hargaProduk = A.tabProduk[tempIdx].hargaProduk - (A.tabProduk[tempIdx].hargaProduk / 10)
								tambah = true
								produkInginDiBeli = false
							} else {
								fmt.Println("Stok tidak cukup")
							}
						}
					}
				}
			}
		} else if input == 3 { //hapus belanja (Sequential Search)
			fmt.Println("Input nama produk yang ingin dihapus: ")
			fmt.Scan(&namaP)
			countBelanja = 0
			for i := 0; i < arrMax-1; i++ {
				if belanja[i].namaProduk != "" {
					countBelanja++
				}
			}
			for i := 0; i < arrMax; i++ {
				if belanja[i].namaProduk == namaP {
					j = i
					for j < countBelanja {
						belanja[j] = belanja[j+1]
						j++
					}
					belanja[countBelanja-1] = produk{0, "", 0, 0}
				}
			}
		} else if input == 4 { //edit jumlah produk belanja
			fmt.Println("Input kode produk yang ingin di edit: ")
			fmt.Scan(&kodeP)
			fmt.Println("Input jumlah fix: ")
			fmt.Scan(&tempStok)

			tempCount = 0
			onEdit = true
			for i := 0; i < arrMax && onEdit; i++ {
				if belanja[i].namaProduk != "" {
					tempCount++
				} else {
					onEdit = false
				}
			}

			tempIdx = -1
			for i := 0; i < tempCount; i++ {
				if belanja[i].kodeProduk == kodeP {
					tempIdx = i
				}
			}

			tempIdxDua = findProdukSquentialKode(*A, kodeP)
			if tempIdx != -1 && A.tabProduk[tempIdxDua].stokProduk >= tempStok {
				belanja[tempIdx].stokProduk = tempStok
			} else if tempIdx != -1 && A.tabProduk[tempIdxDua].stokProduk < tempStok {
				fmt.Println("Stok tidak cukup")
			} else {
				fmt.Println("Produk tidak ditemukan")
			}
		} else if input == 5 { //daftar produk
			daftarProduk(A)
		} else if input == 6 { //back
			on = false
		}
	}
}

func transaksiNonMember(A *tabtab) {
	var input, kodeP, tempIdx, tempIdxDua, tempStok, tempCount, jumlahBeliProduk, total, j, countBelanja int
	var namaP string
	var belanja [arrMax]produk
	var on bool = true
	var tambah bool
	var produkInginDiBeli bool
	var onEdit bool
	for on == true {
		total = 0
		tambah = false
		fmt.Printf("%-15s %-15s %-15s %-15s\n", "Kode Produk", "Nama Produk", "Jumlah", "Harga")
		for i := 0; i < arrMax-1; i++ {
			if belanja[i].namaProduk != "" {
				fmt.Printf("%-15d %-15s %-15d %-15d\n", belanja[i].kodeProduk, belanja[i].namaProduk, belanja[i].stokProduk, belanja[i].hargaProduk)
				total = (belanja[i].hargaProduk * belanja[i].stokProduk) + total
			}
		}
		fmt.Printf("Total transaksi: %d\n", total)
		listTransaksi()
		fmt.Println("Input: ")
		fmt.Scan(&input)
		if input == 1 { //bayar belanja
			bayarNonMember(A, belanja)
			on = false
		} else if input == 2 { //tambah belanja
			fmt.Println("Input kode produk: ")
			fmt.Scan(&kodeP)
			tempIdx = findProdukSquentialKode(*A, kodeP) //tempHarga <- untuk nampung function tambah yang me return idx serta harga produk
			if tempIdx == -1 {
				fmt.Println("Produk tidak tersedia")
			} else {
				for i := 0; i < arrMax-1 && tambah == false; i++ {
					if belanja[i].kodeProduk == kodeP {
						produkInginDiBeli = true
						for produkInginDiBeli {
							fmt.Println("Input jumlah produk yang ingin dibeli: ")
							fmt.Scan(&jumlahBeliProduk)
							if A.tabProduk[kodeP-1].stokProduk >= jumlahBeliProduk {
								belanja[i].kodeProduk = kodeP
								belanja[i].namaProduk = A.tabProduk[kodeP-1].namaProduk
								belanja[i].stokProduk = belanja[i].stokProduk + jumlahBeliProduk
								belanja[i].hargaProduk = A.tabProduk[tempIdx].hargaProduk
								tambah = true
								produkInginDiBeli = false
							} else {
								fmt.Println("Stok tidak cukup")
							}
						}
					} else if belanja[i].namaProduk == "" {
						produkInginDiBeli = true
						for produkInginDiBeli {
							fmt.Println("Input jumlah produk yang ingin dibeli: ")
							fmt.Scan(&jumlahBeliProduk)
							if A.tabProduk[kodeP-1].stokProduk >= jumlahBeliProduk {
								belanja[i].kodeProduk = kodeP
								belanja[i].namaProduk = A.tabProduk[kodeP-1].namaProduk
								belanja[i].stokProduk = jumlahBeliProduk
								belanja[i].hargaProduk = A.tabProduk[tempIdx].hargaProduk
								tambah = true
								produkInginDiBeli = false
							} else {
								fmt.Println("Stok tidak cukup")
							}
						}
					}
				}
			}
		} else if input == 3 { //hapus belanja (Sequential Search)
			fmt.Println("Input nama produk yang ingin dihapus: ")
			fmt.Scan(&namaP)
			countBelanja = 0
			for i := 0; i < arrMax-1; i++ {
				if belanja[i].namaProduk != "" {
					countBelanja++
				}
			}
			for i := 0; i < arrMax; i++ {
				if belanja[i].namaProduk == namaP {
					j = i
					for j < countBelanja {
						belanja[j] = belanja[j+1]
						j++
					}
					belanja[countBelanja-1] = produk{0, "", 0, 0}
				}
			}
		} else if input == 4 { //edit jumlah produk belanja
			fmt.Println("Input kode produk yang ingin di edit: ")
			fmt.Scan(&kodeP)
			fmt.Println("Input jumlah fix: ")
			fmt.Scan(&tempStok)

			tempCount = 0
			onEdit = true
			for i := 0; i < arrMax && onEdit; i++ {
				if belanja[i].namaProduk != "" {
					tempCount++
				} else {
					onEdit = false
				}
			}

			tempIdx = -1
			for i := 0; i < tempCount; i++ {
				if belanja[i].kodeProduk == kodeP {
					tempIdx = i
				}
			}

			tempIdxDua = findProdukSquentialKode(*A, kodeP)
			if tempIdx != -1 && A.tabProduk[tempIdxDua].stokProduk >= tempStok {
				belanja[tempIdx].stokProduk = tempStok
			} else if tempIdx != -1 && A.tabProduk[tempIdxDua].stokProduk < tempStok {
				fmt.Println("Stok tidak cukup")
			} else {
				fmt.Println("Produk tidak ditemukan")
			}
		} else if input == 5 { //daftar produk
			daftarProduk(A)
		} else if input == 6 { //back
			on = false
		}
	}
}

func findProdukSquentialKode(A tabtab, kode int) int { //Sequential Search
	var tempCount, idx int
	idx = -1
	tempCount = jumlahProduk(A)
	for i := 0; i < tempCount && idx == -1; i++ {
		if A.tabProduk[i].kodeProduk == kode {
			idx = i
		}
	}
	return idx
}

func findProdukSquentialNamaProduk(A tabtab, namaProduk string) int {
	var tempCount, idx int
	idx = -1
	tempCount = jumlahProduk(A)
	for i := 0; i < tempCount && idx == -1; i++ {
		if A.tabProduk[i].namaProduk == namaProduk {
			idx = i
		}
	}
	return idx
}

func jumlahProduk(A tabtab) int {
	var count int
	count = 0
	for i := 0; i < arrMax-1; i++ {
		if A.tabProduk[i].namaProduk != "" {
			count++
		}
	}
	return count
}

func daftarProduk(A *tabtab) {
	var input int
	var on bool = true
	var onDua bool = true
	for on {
		listDaftarProduk()
		fmt.Println("Input: ")
		fmt.Scan(&input)
		if input == 1 {
			for onDua {
				listAscendingDescending()
				fmt.Println("Input: ")
				fmt.Scan(&input)
				if input == 1 {
					sortingHargaAscending(A)
					onDua = false
					on = false
				} else if input == 2 {
					sortingHargaDescending(A)
					onDua = false
					on = false
				} else {
					fmt.Println("Input Salah")
				}
			}
		} else if input == 2 {
			for onDua {
				listAscendingDescending()
				fmt.Println("Input: ")
				fmt.Scan(&input)
				if input == 1 {
					sortingStokAscending(A)
					onDua = false
					on = false
				} else if input == 2 {
					sortingStokDescending(A)
					onDua = false
					on = false
				} else {
					fmt.Println("Input Salah")
				}
			}
		} else if input == 3 {
			cariProduk(A)
		} else if input == 4{
			on = false
		} else {
			fmt.Println("Input Salah")
		}
	}
}

func listProduk(A tabtab) {
	fmt.Println("====================================================================================================")
	fmt.Printf("%-15s %-15s %-15s %-15s\n", "Kode Produk", "Nama Produk", "Stok", "Harga")
	for i := 0; i < arrMax; i++ {
		if A.tabProduk[i].namaProduk != "" {
			fmt.Printf("%-15d %-15s %-15d %-15d\n", A.tabProduk[i].kodeProduk, A.tabProduk[i].namaProduk, A.tabProduk[i].stokProduk, A.tabProduk[i].hargaProduk)
		}
	}
	fmt.Println("====================================================================================================")
}

func sortingHargaAscending(A *tabtab) { //selection sort
	var i, iMin, pass, count int
	var temp produk
	count = jumlahProduk(*A)
	for pass = 0; pass < count; pass++ {
		iMin = pass
		i = pass + 1
		for i < count {
			if A.tabProduk[i].hargaProduk < A.tabProduk[iMin].hargaProduk {
				iMin = i
			}
			i++
		}
		temp = A.tabProduk[iMin]
		A.tabProduk[iMin] = A.tabProduk[pass]
		A.tabProduk[pass] = temp
	}
	listProduk(*A)
}

func sortingHargaDescending(A *tabtab) { //selection Sort
	var i, iMin, pass, count int
	var temp produk
	count = jumlahProduk(*A)
	for pass = 0; pass < count; pass++ {
		iMin = pass
		i = pass + 1
		for i < count {
			if A.tabProduk[i].hargaProduk > A.tabProduk[iMin].hargaProduk {
				iMin = i
			}
			i++
		}
		temp = A.tabProduk[iMin]
		A.tabProduk[iMin] = A.tabProduk[pass]
		A.tabProduk[pass] = temp
	}
	listProduk(*A)
}

func sortingStokAscending(A *tabtab) { //insertion sort
	var pass, i, count int
	var temp produk
	count = jumlahProduk(*A)
	for pass = 0; pass < count; pass++ {
		i = pass + 1
		temp = A.tabProduk[i]
		for i > 0 && A.tabProduk[i-1].stokProduk > temp.stokProduk {
			A.tabProduk[i] = A.tabProduk[i-1]
			i--
		}
		A.tabProduk[i] = temp
	}
	listProduk(*A)
}

func sortingStokDescending(A *tabtab) { //insertion sort
	var pass, i, count int
	var temp produk
	count = jumlahProduk(*A)
	for pass = 0; pass < count; pass++ {
		i = pass + 1
		temp = A.tabProduk[i]
		for i > 0 && A.tabProduk[i-1].stokProduk < temp.stokProduk {
			A.tabProduk[i] = A.tabProduk[i-1]
			i--
		}
		A.tabProduk[i] = temp
	}
	listProduk(*A)
}

func sortingKodeProdukAscending(A *tabtab) {
	var pass, i, count int
	var temp produk
	count = jumlahProduk(*A)
	for pass = 0; pass < count; pass++ {
		i = pass + 1
		temp = A.tabProduk[i]
		for i > 0 && A.tabProduk[i-1].kodeProduk > temp.kodeProduk {
			A.tabProduk[i] = A.tabProduk[i-1]
			i--
		}
		A.tabProduk[i] = temp
	}
}

func cariProduk(A *tabtab) {
	var input, tempIdx, kodeProduk int
	var namaProduk string
	var on bool = true
	for on {
		listCariData()
		fmt.Scan(&input)
		if input == 1 {
			fmt.Println("Input kode produk: ")
			fmt.Scan(&kodeProduk)
			tempIdx = findProdukBinary(A, kodeProduk)
			if tempIdx != -1 {
				fmt.Printf("Produk")
				fmt.Printf("Kode: %d \tNama: %s \tStok: %d \tHarga Reguler: %d \tHarga Member: %d\n", A.tabProduk[tempIdx].kodeProduk, A.tabProduk[tempIdx].namaProduk, A.tabProduk[tempIdx].stokProduk, A.tabProduk[tempIdx].hargaProduk, (A.tabProduk[tempIdx].hargaProduk - (A.tabProduk[tempIdx].hargaProduk / 10)))
				on = false
			} else {
				fmt.Println("Produk tidak ditemukan")
			}
		} else if input == 2 {
			fmt.Println("Input nama produk: ")
			fmt.Scan(&namaProduk)
			tempIdx = findProdukSquentialNamaProduk(*A, namaProduk)
			if tempIdx != -1 {
				fmt.Printf("Produk")
				fmt.Printf("Kode: %d \tNama: %s \tStok: %d \tHarga Reguler: %d \tHarga Member: %d\n", A.tabProduk[tempIdx].kodeProduk, A.tabProduk[tempIdx].namaProduk, A.tabProduk[tempIdx].stokProduk, A.tabProduk[tempIdx].hargaProduk, (A.tabProduk[tempIdx].hargaProduk - (A.tabProduk[tempIdx].hargaProduk / 10)))
				on = false
			} else {
				fmt.Println("Produk tidak ditemukan")
			}
		} else if input == 3{
			on = false
		} else {
			fmt.Println("Input Salah")
		}
	}
}

func findProdukBinary(A *tabtab, kode int) int { //binary search
	sortingKodeProdukAscending(A)
	var idx int
	var left, right, mid int
	left = 0
	right = jumlahProduk(*A) - 1
	idx = -1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if A.tabProduk[mid].kodeProduk == kode {
			return mid
		} else if A.tabProduk[mid].kodeProduk < kode {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return idx
}

func tambahProduk(A *tabtab) {
	var on bool = true
	for i := 0; i < arrMax && on; i++ {
		if A.tabProduk[i].namaProduk == ""{
			fmt.Println("Input nama, stok, dan harga produk: ")
			fmt.Scan(&A.tabProduk[i].namaProduk, &A.tabProduk[i].stokProduk, &A.tabProduk[i].hargaProduk)
			A.tabProduk[i].kodeProduk = i + 1
			on = false
		}
	}
}

func hapusProduk(A *tabtab) {
	var count, i int
	var namaProduk string
	fmt.Println("Input nama produk yang ingin dihapus: ")
	fmt.Scan(&namaProduk)
	count = jumlahProduk(*A)
	i = findProdukSquentialNamaProduk(*A, namaProduk)
	for i < count {
		A.tabProduk[i] = A.tabProduk[i+1]
		i++
	}
	A.tabProduk[count-1] = produk{0, "", 0, 0}
}

func editProduk(A *tabtab) {
	var i, input, stokProduk, hargaProduk int
	var namaProduk string
	var on bool = true
	fmt.Println("Input nama produk yang ingin dihapus: ")
	fmt.Scan(&namaProduk)
	i = findProdukSquentialNamaProduk(*A, namaProduk)
	for on {
		fmt.Println("1. Edit nama produk")
		fmt.Println("2. Edit stok produk")
		fmt.Println("3. Edit harga produk")
		fmt.Println("4. BACK")
		fmt.Println("Input: ")
		fmt.Scan(&input)
		if input == 1 {
			fmt.Println("Input nama produk baru: ")
			fmt.Scan(&namaProduk)
			A.tabProduk[i].namaProduk = namaProduk
		} else if input == 2 {
			fmt.Println("Input stok produk baru: ")
			fmt.Scan(&stokProduk)
			A.tabProduk[i].stokProduk = stokProduk
		} else if input == 3 {
			fmt.Println("Input harga produk baru: ")
			fmt.Scan(&hargaProduk)
			A.tabProduk[i].hargaProduk = hargaProduk
		} else if input == 4 {
			on = false
		} else {
			fmt.Println("Input Salah")
		}
	}
}

func bayarMember(A *tabtab, B [arrMax]produk) {
    var i, j, countI, countJ, tempTotal, idx int
    var namaMember string
    var on bool = true

    countI = jumlahProduk(*A)
    countJ = 0
    for x := 0; x < arrMax; x++ {
        if B[x].namaProduk != "" {
            countJ++
        }
    }

    for i = 0; i < countI && on; i++ {
        if A.tabOmzet[i].member == "" {
            fmt.Print("Input nama member: ")
            fmt.Scan(&namaMember)

            A.tabOmzet[i].noTransaksi = i + 1
            A.tabOmzet[i].member = namaMember
            tempTotal = 0

            for j = 0; j < countJ; j++ {
                idx = findProdukSquentialKode(*A, B[j].kodeProduk)
                if idx != -1 { 
                    A.tabOmzet[i].belanja[j] = B[j]
                    tempTotal += B[j].hargaProduk
                    A.tabProduk[idx].stokProduk -= B[j].stokProduk
                }
            }

            A.tabOmzet[i].totalBelanja = tempTotal
            on = false 
        }
    }
}

func bayarNonMember(A *tabtab, B [arrMax]produk) {
    var i, j, countI, countJ, tempTotal, idx int
    var on bool = true

    countI = jumlahProduk(*A)
    countJ = 0
    for x := 0; x < arrMax; x++ {
        if B[x].namaProduk != "" {
            countJ++
        }
    }

    for i = 0; i < countI && on; i++ {
        if A.tabOmzet[i].member == "" {
            A.tabOmzet[i].noTransaksi = i + 1
            A.tabOmzet[i].member = "Non-Member"
            tempTotal = 0

            for j = 0; j < countJ; j++ {
                idx = findProdukSquentialKode(*A, B[j].kodeProduk)
                if idx != -1 { 
                    A.tabOmzet[i].belanja[j] = B[j]
                    tempTotal += B[j].hargaProduk
                    A.tabProduk[idx].stokProduk -= B[j].stokProduk
                }
            }

            A.tabOmzet[i].totalBelanja = tempTotal
            on = false
        }
    }
}

func omzet(A *tabtab) {
	var input int
	var on bool = true
	for on {
		listOmzet()
		fmt.Println("Input")
		fmt.Scan(&input)
		if input == 1 {
			sortingOmzetNomor(A)
		} else if input == 2{
			hapusTransaksi(A)
		} else if input == 3{
			totalOmzet(*A)	
		} else if input == 4{
			on = false
		} else{
			fmt.Println("Input Salah")
		}
	}
}

func sortingOmzetNomor(A *tabtab) {
	var pass, i, count int
	var temp bon

	count = 0
	for j := 0; j < arrMax; j++ {
		if A.tabOmzet[j].member != "" {
			count++
		}
	}

	for pass = 1; pass < count; pass++ {
		temp = A.tabOmzet[pass]
		i = pass
		for i > 0 && A.tabOmzet[i-1].noTransaksi > temp.noTransaksi {
			A.tabOmzet[i] = A.tabOmzet[i-1] 
			i--
		}
		A.tabOmzet[i] = temp 
	}
	printOmzet(*A, count)
}

func cariTransaksi(A tabtab, kode int) int {
	var idx int
	idx = -1
	for i := 0; i < arrMax && idx == -1; i++ {
		if kode == A.tabOmzet[i].noTransaksi {
			idx = i
		}
	}
	return idx
}

func hapusTransaksi(A *tabtab){
	var count, i, kodeTransaksi int
	fmt.Println("Input no transaksi yang ingin dihapus: ")
	fmt.Scan(&kodeTransaksi)
	count = 0
	for j := 0; j < arrMax; j++ {
		if A.tabOmzet[j].member != "" {
			count++
		}
	}

	i = cariTransaksi(*A, kodeTransaksi)
	for i < count {
		A.tabOmzet[i] = A.tabOmzet[i+1]
		i++
	}
	A.tabProduk[count-1] = produk{0, "", 0, 0}	
}

func printOmzet(A tabtab, n int) {
    fmt.Println("=================================================================")
    fmt.Println("                   LAPORAN TRANSAKSI OMZET")
    fmt.Println("=================================================================")

    for i := 0; i < n; i++ {
        fmt.Println("No. Transaksi :", A.tabOmzet[i].noTransaksi)
        fmt.Println("Member        :", A.tabOmzet[i].member)
        fmt.Println("Daftar Belanja:")

        // Header tabel dengan alignment lebih baik
        fmt.Printf("%-15s %-20s %-10s %-15s\n", "Kode Produk", "Nama Produk", "Jumlah", "Harga")

        for j := 0; j < len(A.tabOmzet[i].belanja); j++ {
            if A.tabOmzet[i].belanja[j].namaProduk != "" {
                fmt.Printf("%-15d %-20s %-10d %-15d\n",
                    A.tabOmzet[i].belanja[j].kodeProduk,
                    A.tabOmzet[i].belanja[j].namaProduk,
                    A.tabOmzet[i].belanja[j].stokProduk,
                    A.tabOmzet[i].belanja[j].hargaProduk)
            }
        }

        fmt.Println("-----------------------------------------------------------------")
        fmt.Printf("TOTAL BELANJA  : Rp %d\n", A.tabOmzet[i].totalBelanja)
        fmt.Println("=================================================================")
        fmt.Println()
    }
}

func totalOmzet(A tabtab){
	var tempTotal int
	tempTotal = 0
	for i := 0; i < arrMax; i++{
		if A.tabOmzet[i].member != ""{
			tempTotal = tempTotal + A.tabOmzet[i].totalBelanja
		}
	}
	fmt.Println("Total Omzet: ", tempTotal)
}