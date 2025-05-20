package main

import "fmt"

type anggota struct {
	nama, peran string
}

type startup struct {
	nama, bidang  string
	tahunBerdiri  int
	dana          float64
	anggota       [10]anggota
	jumlahAnggota int
}

var dataStartup [100]startup
var jumlahdataStartup int

func main() {
	var pilihanMenu int
	memilihMenu(&pilihanMenu)
}

func memilihMenu(pilihanMenu *int) {
	for *pilihanMenu != 10 {
		fmt.Println()
		fmt.Println(" ===== Selamat Datang =====")
		fmt.Println(" Aplikasi Manajemen Startup")
		fmt.Println("------------------------------")
		fmt.Println("1. Tambah Startup")
		fmt.Println("2. Ubah Data Startup")
		fmt.Println("3. Hapus Data Startup")
		fmt.Println("4. Tambah Anggota Tim")
		fmt.Println("5. Cari startup")
		fmt.Println("6. Urutan Berdasarkan Tahun Berdiri")
		fmt.Println("7. Urutan Berdasarkan Pendanaan")
		fmt.Println("8. Tampilkan Laporan")
		fmt.Println("9. Tampilkan Semua Startup")
		fmt.Println("10. Exit")
		fmt.Println("------------------------------")
		fmt.Print("Silahkan Pilih: 1/2/3/4/5/6/7/8/9/10? ")
		fmt.Scan(pilihanMenu)

		if *pilihanMenu == 1 {
			tambahStartup()
		} else if *pilihanMenu == 2 {
			ubahData()
		} else if *pilihanMenu == 3 {
			hapusData()
			//	} else if *pilihanMenu == 4 {
			//	tambahAnggota()
		} else if *pilihanMenu == 5 {
			cariStartup()
		} else if *pilihanMenu == 6 {
			startupTahun()
		} else if *pilihanMenu == 7 {
			startupDana()
		} else if *pilihanMenu == 8 {
			//		tampilanLaporan()
		} else if *pilihanMenu == 9 {
			tampilanSemuaStartup()
		} else if *pilihanMenu == 10 {
			fmt.Println("Terima Kasih telah menggunakan Aplikasi manajemen startup sederhana")
		}
	}
}

func tambahStartup() {
	//Fungsi Untuk Menambahkan Startup dan juga bidangnya
	var i, nStartup int
	fmt.Println("Berapa Startup yang ingin anda tambahkan?")
	fmt.Scan(&nStartup)
	if jumlahdataStartup >= 100 {
		fmt.Print("Data startup penuh. Tidak dapat menambahkan startup baru")
		return
	}
	for i = 0; i < nStartup; i++ {
		fmt.Printf("\nStartup ke-%d:\n", i+1)

		fmt.Print("Nama Startup: ")
		fmt.Scan(&dataStartup[jumlahdataStartup].nama)

		fmt.Print("Bidang Usaha: ")
		fmt.Scan(&dataStartup[jumlahdataStartup].bidang)

		fmt.Print("Tahun Berdiri: ")
		fmt.Scan(&dataStartup[jumlahdataStartup].tahunBerdiri)

		fmt.Print("Total Pendanaan: ")
		fmt.Scan(&dataStartup[jumlahdataStartup].dana)

		jumlahdataStartup++
		fmt.Println("Startup Berhasil ditambahkan")
	}
}

func ubahData() {
	//Fungsi untuk mengubah startup yang ingin di ubah
	var idx int //idx adalah yang ingin diubah
	fmt.Print("Memilih startup yang ingin di ubah: ")
	fmt.Scan(&idx)
	idx = idx - 1
	if idx < 0 || idx >= jumlahdataStartup {
		fmt.Println("Data Tidak Ditemukan.")
		return
	}
	fmt.Print("Nama Baru Startup: ")
	fmt.Scan(&dataStartup[idx].nama)

	fmt.Print("Bidang Baru Startup: ")
	fmt.Scan(&dataStartup[idx].bidang)

	fmt.Print("Tahun Baru Startup: ")
	fmt.Scan(&dataStartup[idx].tahunBerdiri)

	fmt.Print("Dana Baru Startup: ")
	fmt.Scan(&dataStartup[idx].dana)

	fmt.Println("Data Startup berhasil diubah.")
}
func hapusData() {
	//Fungsi untuk menghapus data
	var i, idx int //idx adalah memilih mana data yang ingin di hapus
	fmt.Print("Memilih startup yang ingin dihapus: ")
	fmt.Scan(&idx)
	idx = idx - 1
	if idx < 0 || idx >= jumlahdataStartup {
		fmt.Println("Data Tidak Ditemukan.")
		return
	}
	for i = idx; i < jumlahdataStartup-1; i++ {
		dataStartup[i] = dataStartup[i+1]
	}
	jumlahdataStartup--
	fmt.Println("Data Startup berhasil di hapus")
}

func tampilanSemuaStartup() {
	//menampilkan semua data startup
	var i int
	if jumlahdataStartup == 0 {
		fmt.Print("Tidak ada data startup.")
		return
	}

	for i = 0; i < jumlahdataStartup; i++ {
		fmt.Printf("\nStartup ke-%d:\n", i+1)
		fmt.Println("Nama   :", dataStartup[i].nama)
		fmt.Println("Bidang :", dataStartup[i].bidang)
		fmt.Println("Tahun  :", dataStartup[i].tahunBerdiri)
		fmt.Println("Dana   :", dataStartup[i].dana)
	}
}

func cariStartup() {
	var pilihan int
	var idx int
	var cariNama, cariBidang string
	fmt.Println("Cari Startup Berdasarkan:")
	fmt.Println("1. Nama")
	fmt.Println("2. Bidang Usaha")
	fmt.Print("Pilih (1/2): ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		fmt.Print("Masukkan nama startup yang dicari: ")
		fmt.Scan(&cariNama)
		urutkanNama()
		idx = binarySearchNama(cariNama)
		if idx != -1 {
			fmt.Println("Startup ditemukan:")
			fmt.Println("Nama   :", dataStartup[idx].nama)
			fmt.Println("Bidang :", dataStartup[idx].bidang)
			fmt.Println("Tahun  :", dataStartup[idx].tahunBerdiri)
			fmt.Println("Dana   :", dataStartup[idx].dana)
		} else {
			fmt.Println("Startup tidak ditemukan")
		}
	} else if pilihan == 2 {
		fmt.Print("Masukkan bidang usaha yang dicari: ")
		fmt.Scan(&cariBidang)
		idx = sequentialSearchBidang(cariBidang)
		if idx != -1 {
			fmt.Println("Startup ditemukan:")
			fmt.Println("Nama   :", dataStartup[idx].nama)
			fmt.Println("Bidang :", dataStartup[idx].bidang)
			fmt.Println("Tahun  :", dataStartup[idx].tahunBerdiri)
			fmt.Println("Dana   :", dataStartup[idx].dana)
		} else {
			fmt.Println("Bidang usaha tidak ditemukan")
		}
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func sequentialSearchBidang(x string) int {
	var ketemu, k int

	ketemu = -1
	k = 0

	for ketemu == -1 && k < jumlahdataStartup {
		if dataStartup[k].bidang == x {
			ketemu = k
		}
		k = k + 1
	}

	return ketemu
}

func binarySearchNama(x string) int {
	var left, right, mid int
	var found int

	left = 0
	right = jumlahdataStartup - 1
	found = -1

	for left <= right && found == -1 {
		mid = (left + right) / 2
		if x < dataStartup[mid].nama {
			right = mid - 1
		} else if x > dataStartup[mid].nama {
			left = mid + 1
		} else {
			found = mid
		}
	}
	return found
}

func urutkanNama() {
	var i, j int
	var temp startup
	for i = 1; i < jumlahdataStartup; i++ {
		temp = dataStartup[i]
		j = i - 1
		for j >= 0 && dataStartup[j].nama > temp.nama {
			dataStartup[j+1] = dataStartup[j]
			j--
		}
		dataStartup[j+1] = temp
	}
}

func startupTahun() {
	var i int
	if jumlahdataStartup == 0 {
		fmt.Println("Data startup masih kosong")
	}

	selectionSort()

	fmt.Println("Berikut merupakan list startup dari yang tahun berdiri nya paling baru sampai ke yang paling lama :")
	for i = 0; i < jumlahdataStartup; i++ {
		fmt.Printf("\nStartup ke-%d:\n", i)
		fmt.Println("Nama   :", dataStartup[i].nama)
		fmt.Println("Bidang :", dataStartup[i].bidang)
		fmt.Println("Tahun  :", dataStartup[i].tahunBerdiri)
		fmt.Println("Dana   :", dataStartup[i].dana)
	}

}

func startupDana() {
	var i int
	if jumlahdataStartup == 0 {
		fmt.Println("Data startup masih kosong")
	}
	insertionSort()
	fmt.Println("Berikut merupakan urutan startup berdasarkan dana dari yang terkecil ke paling besar")
	for i = 0; i < jumlahdataStartup; i++ {
		fmt.Printf("\nStartup ke-%d:\n", i)
		fmt.Println("Nama   :", dataStartup[i].nama)
		fmt.Println("Bidang :", dataStartup[i].bidang)
		fmt.Println("Tahun  :", dataStartup[i].tahunBerdiri)
		fmt.Println("Dana   :", dataStartup[i].dana)
	}

}

func selectionSort() {
	var i, idx, pass int
	var temp startup

	pass = 1
	for pass < jumlahdataStartup {
		idx = pass - 1
		i = pass
		for i < jumlahdataStartup {
			if dataStartup[i].tahunBerdiri > dataStartup[idx].tahunBerdiri {
				idx = i
			}
			i = i + 1
		}
		temp = dataStartup[pass-1]
		dataStartup[pass-1] = dataStartup[idx]
		dataStartup[idx] = temp
	}
}

func insertionSort() {
	var i, pass int
	var temp startup

	pass = 1
	for pass <= jumlahdataStartup-1 {
		i = pass
		temp = dataStartup[pass]

		for i > 0 && temp.dana < dataStartup[i-1].dana {
			dataStartup[i] = dataStartup[i-1]
			i = i - 1
		}
		dataStartup[i] = temp
	}
}
