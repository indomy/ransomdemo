package main

import (
	"bytes"
	"compress/gzip"
	"io"
	"os"

	"github.com/indomy/ransomdemo/pkg"
)

func main() {
	// Baca file terenkripsi
	encryptedFile := "output.enc"
	encryptedData, err := os.ReadFile(encryptedFile)
	if err != nil {
		panic(err)
	}

	// Dekripsi data menggunakan AES
	key := []byte("passphrasewhichneedstobe32bytes!") // 32 bytes key for AES-256
	decryptedData, err := pkg.Decrypt(encryptedData, key)
	if err != nil {
		panic(err)
	}

	// Dekompresi data menggunakan gzip
	gzipReader, err := gzip.NewReader(bytes.NewReader(decryptedData))
	if err != nil {
		panic(err)
	}
	decompressedData, err := io.ReadAll(gzipReader)
	if err != nil {
		panic(err)
	}
	gzipReader.Close()

	// Simpan data yang sudah didekompresi ke file
	outputFile := "output.tar"
	err = os.WriteFile(outputFile, decompressedData, 0644)
	if err != nil {
		panic(err)
	}

	println("File berhasil didekripsi dan didekompresi!")
}
