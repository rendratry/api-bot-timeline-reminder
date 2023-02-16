package helper

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func UploadImageProfile(imageBase64 string) (string, error) {
	// Mendapatkan data gambar dari request
	//imageBase64 := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAAuElEQVQ4T2NkQAOMgqf1rA7vG8yMnLAzGcgjKzUwMDYwMjIwMzAwMzQ2MjAwMzA9DfQ/LgY5/5IAHJwYk/GjZxNjlOaDpGS0B/cH5OzN0MHl+3q4fjGzhm8L0EycTA0cI/gSgQ2KjJATb/+n4fIq/x4eKjwYhZJFMBAoRgSCTAAABQABHRK9JQAAAABJRU5ErkJggg=="
	imageType := "png"

	// Mendapatkan direktori tempat menyimpan gambar
	imageDir := "/var/www/html/images/presensi-app/profile"

	// Menghapus prefix dan mendekode base64 menjadi byte array
	data := imageBase64[strings.IndexByte(imageBase64, ',')+1:]
	imageBytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		fmt.Println("Error decoding base64:", err)
		os.Exit(1)
	}

	// Simpan byte array ke file gambar
	imageName := "gambar." + imageType
	err = ioutil.WriteFile(filepath.Join(imageDir, imageName), imageBytes, 0644)
	if err != nil {
		fmt.Println("Error saving image:", err)
		os.Exit(1)
	}

	// Mengembalikan path gambar
	imagePath := "/images/" + imageName
	fmt.Println("Image path:", imagePath)

	return imagePath, nil
}
