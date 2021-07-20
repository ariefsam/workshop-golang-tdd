package main

import "fmt"

func Cetak(addr Alamat) (output string) {
	output = fmt.Sprintf("%s\nKelurahan %s\nKecamatan %s\nKota %s\nProvinsi %s\n",
		addr.AlamatLengkap,
		addr.Kelurahan,
		addr.Kecamatan,
		addr.Kota,
		addr.Provinsi,
	)
	return
}
