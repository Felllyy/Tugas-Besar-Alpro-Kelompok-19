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

type dataStartup [100]startup

var jumlahdataStartup int

func main() {
	var pilihanMenu int
	memilihMenu(&pilihanMenu)
}

func memilihMenu(pilihanMenu *int) {
	var data dataStartup
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
			tambahStartup(&data)
		} else if *pilihanMenu == 2 {
			ubahData(&data)
		} else if *pilihanMenu == 3 {
			hapusData(&data)
		} else if *pilihanMenu == 4 {
			tambahAnggota(&data)
		} else if *pilihanMenu == 5 {
			cariStartup(data)
		} else if *pilihanMenu == 6 {
			startupTahun(data)
		} else if *pilihanMenu == 7 {
			startupDana(data)
		} else if *pilihanMenu == 8 {
			tampilanLaporan(data)
		} else if *pilihanMenu == 9 {
			tampilanSemuaStartup(data)
		} else if *pilihanMenu == 10 {
			fmt.Println("TERIMA KASIH TELAH MENGGUNAKAN APLIKASI MANAJEMEN STARTUP SEDERHANA")
		}
	}
}

func tambahStartup(T *dataStartup) {
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
		fmt.Scan(&T[jumlahdataStartup].nama)

		fmt.Print("Bidang Usaha: ")
		fmt.Scan(&T[jumlahdataStartup].bidang)

		fmt.Print("Tahun Berdiri: ")
		fmt.Scan(&T[jumlahdataStartup].tahunBerdiri)

		fmt.Print("Total Pendanaan: ")
		fmt.Scan(&T[jumlahdataStartup].dana)

		jumlahdataStartup++
		fmt.Println("Startup Berhasil ditambahkan")
	}
}

func ubahData(T *dataStartup) {
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
	fmt.Scan(&T[idx].nama)

	fmt.Print("Bidang Baru Startup: ")
	fmt.Scan(&T[idx].bidang)

	fmt.Print("Tahun Baru Startup: ")
	fmt.Scan(&T[idx].tahunBerdiri)

	fmt.Print("Dana Baru Startup: ")
	fmt.Scan(&T[idx].dana)

	fmt.Println("Data Startup berhasil diubah.")
}
func hapusData(T *dataStartup) {
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
		T[i] = T[i+1]
	}
	jumlahdataStartup--
	fmt.Println("Data Startup berhasil di hapus")
}

func tampilanSemuaStartup(T dataStartup) {
	//menampilkan semua data startup
	var i int
	if jumlahdataStartup == 0 {
		fmt.Print("Tidak ada data startup.")
		return
	}

	for i = 0; i < jumlahdataStartup; i++ {
		fmt.Printf("\nStartup ke-%d:\n", i+1)
		fmt.Println("Nama   :", T[i].nama)
		fmt.Println("Bidang :", T[i].bidang)
		fmt.Println("Tahun  :", T[i].tahunBerdiri)
		fmt.Println("Dana   :", T[i].dana)
	}
}

func cariStartup(T dataStartup) {
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
		urutkanNama(&T)
		idx = binarySearchNama(cariNama, T)
		if idx != -1 {
			fmt.Println("Startup ditemukan:")
			fmt.Println("Nama   :", T[idx].nama)
			fmt.Println("Bidang :", T[idx].bidang)
			fmt.Println("Tahun  :", T[idx].tahunBerdiri)
			fmt.Println("Dana   :", T[idx].dana)
		} else {
			fmt.Println("Startup tidak ditemukan")
		}
	} else if pilihan == 2 {
		fmt.Print("Masukkan bidang usaha yang dicari: ")
		fmt.Scan(&cariBidang)
		idx = sequentialSearchBidang(cariBidang, T)
		if idx != -1 {
			fmt.Println("Startup ditemukan:")
			fmt.Println("Nama   :", T[idx].nama)
			fmt.Println("Bidang :", T[idx].bidang)
			fmt.Println("Tahun  :", T[idx].tahunBerdiri)
			fmt.Println("Dana   :", T[idx].dana)
		} else {
			fmt.Println("Bidang usaha tidak ditemukan")
		}
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func sequentialSearchBidang(x string, T dataStartup) int {
	var ketemu, k int

	ketemu = -1
	k = 0

	for ketemu == -1 && k < jumlahdataStartup {
		if T[k].bidang == x {
			ketemu = k
		}
		k = k + 1
	}

	return ketemu
}

func binarySearchNama(x string, T dataStartup) int {
	var left, right, mid int
	var found int

	left = 0
	right = jumlahdataStartup - 1
	found = -1

	for left <= right && found == -1 {
		mid = (left + right) / 2
		if x < T[mid].nama {
			right = mid - 1
		} else if x > T[mid].nama {
			left = mid + 1
		} else {
			found = mid
		}
	}
	return found
}

func urutkanNama(T *dataStartup) {
	var i, j int
	var temp startup
	for i = 1; i < jumlahdataStartup; i++ {
		temp = T[i]
		j = i - 1
		for j >= 0 && T[j].nama > temp.nama {
			T[j+1] = T[j]
			j--
		}
		T[j+1] = temp
	}
}

func startupTahun(T dataStartup) {
	var i int
	if jumlahdataStartup == 0 {
		fmt.Println("Data startup masih kosong")
		return
	}

	selectionSort(&T)

	fmt.Println("Berikut merupakan list startup dari yang tahun berdiri nya paling baru sampai ke yang paling lama :")
	for i = 0; i < jumlahdataStartup; i++ {
		fmt.Printf("\nStartup ke-%d:\n", i+1)
		fmt.Println("Nama   :", T[i].nama)
		fmt.Println("Bidang :", T[i].bidang)
		fmt.Println("Tahun  :", T[i].tahunBerdiri)
		fmt.Println("Dana   :", T[i].dana)
	}

}

func startupDana(T dataStartup) {
	var i int
	if jumlahdataStartup == 0 {
		fmt.Println("Data startup masih kosong")
		return
	}
	insertionSort(&T)
	fmt.Println("Berikut merupakan urutan startup berdasarkan dana dari yang terkecil ke paling besar")
	for i = 0; i < jumlahdataStartup; i++ {
		fmt.Printf("\nStartup ke-%d:\n", i+1)
		fmt.Println("Nama   :", T[i].nama)
		fmt.Println("Bidang :", T[i].bidang)
		fmt.Println("Tahun  :", T[i].tahunBerdiri)
		fmt.Println("Dana   :", T[i].dana)
	}

}

func tampilanLaporan(T dataStartup) {
	if jumlahdataStartup == 0 {
		fmt.Println("Tidak ada data startup untuk ditampilkan.")
		return
	}

	fmt.Println("\n=========== Laporan Lengkap Startup ===========")
	for i := 0; i < jumlahdataStartup; i++ {
		fmt.Printf("\nStartup ke-%d\n", i+1)
		fmt.Println("----------------------------------------")
		fmt.Println("Nama Startup  :", T[i].nama)
		fmt.Println("Bidang Usaha  :", T[i].bidang)
		fmt.Println("Tahun Berdiri :", T[i].tahunBerdiri)
		fmt.Println("Total Dana    :", T[i].dana)

		if T[i].jumlahAnggota == 0 {
			fmt.Println("Anggota       : Belum ada anggota terdaftar")
		} else {
			fmt.Printf("Anggota (%d):\n", T[i].jumlahAnggota)
			for j := 0; j < T[i].jumlahAnggota; j++ {
				fmt.Printf("  - %s (%s)\n", T[i].anggota[j].nama, T[i].anggota[j].peran)
			}
		}
	}
	fmt.Println("===============================================")
}

func selectionSort(T *dataStartup) {
	var i, idx, pass int
	var temp startup

	for pass = 0; pass < jumlahdataStartup-1; pass++ {
		idx = pass
		i = pass
		for i = pass + 1; i < jumlahdataStartup; i++ {
			if T[i].tahunBerdiri < T[idx].tahunBerdiri {
				idx = i
			}
		}
		temp = T[pass]
		T[pass] = T[idx]
		T[idx] = temp
	}
}

func insertionSort(T *dataStartup) {
	var i, pass int
	var temp startup

	pass = 1
	for pass <= jumlahdataStartup-1 {
		i = pass
		temp = T[pass]

		for i > 0 && temp.dana < T[i-1].dana {
			T[i] = T[i-1]
			i = i - 1
		}
		T[i] = temp
		pass++
	}
}

func tambahAnggota(T *dataStartup) {
	var idxStartup, jumlahAnggotaBaru int

	if jumlahdataStartup == 0 {
		fmt.Println("Belum ada startup yang terdaftar.")
		return
	}

	fmt.Println("\nDaftar Startup:")
	for i := 0; i < jumlahdataStartup; i++ {
		fmt.Printf("%d. %s (Anggota: %d/10)\n", i+1, T[i].nama, T[i].jumlahAnggota)
	}

	fmt.Print("\nPilih nomor startup: ")
	fmt.Scan(&idxStartup)
	idxStartup--

	if idxStartup < 0 || idxStartup >= jumlahdataStartup {
		fmt.Println("Nomor startup tidak valid!")
		return
	}

	startup := &T[idxStartup]

	fmt.Printf("\nJumlah anggota yang akan ditambahkan ke %s (tersisa %d slot): ", startup.nama, 10-startup.jumlahAnggota)
	fmt.Scan(&jumlahAnggotaBaru)

	if jumlahAnggotaBaru <= 0 {
		fmt.Println("Jumlah anggota harus lebih dari 0")
		return
	}

	if startup.jumlahAnggota+jumlahAnggotaBaru > 10 {
		fmt.Printf("Error: Hanya bisa menambahkan %d anggota\n", 10-startup.jumlahAnggota)
		return
	}

	fmt.Println("\nMasukkan data anggota:")
	for i := 0; i < jumlahAnggotaBaru; i++ {
		var nama, peran string
		fmt.Printf("\nAnggota #%d:\n", i+1)
		fmt.Print("Nama: ")
		fmt.Scan(&nama)
		fmt.Print("Peran: ")
		fmt.Scan(&peran)

		startup.anggota[startup.jumlahAnggota] = anggota{nama, peran}
		startup.jumlahAnggota++
	}

	fmt.Printf("\nBerhasil menambahkan %d anggota ke %s\n", jumlahAnggotaBaru, startup.nama)
}
