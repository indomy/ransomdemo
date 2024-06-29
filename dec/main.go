package main

import (
	"os"
	"path/filepath"
	"strings"

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
		//kalo bukan file yang diencrypt maka di abaikan
		if !strings.Contains(file, ".croot") {
			return nil
		}
		// Baca file
		data, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		// Dekripsi data menggunakan AES
		key := []byte("indonesiacodeacademykeyto32bytes") // 32 bytes key for AES-256
		decryptedData, err := pkg.Decrypt(data, key)
		if err != nil {
			return err
		}

		// Simpan data terenkripsi ke file
		outputFile := strings.Replace(file, ".croot", "", 1)
		err = os.WriteFile(outputFile, decryptedData, 0644)
		if err != nil {
			return err
		}
		println(file)
		//hapus file aslinya
		err = os.Remove(file)
		if err != nil {
			println(err.Error())
		}

		return nil
	})

	if err != nil {
		println(err.Error())
		//panic(err)
	}

	println("File berhasil didekripsi")
}
