package helper

import (
	"encoding/base64"
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
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

func UploadPdfDocument(pdfBase64 string) string {
	// Mengubah data file PDF dari format base64 menjadi bytes
	pdfBytes, err := base64.StdEncoding.DecodeString(strings.Split(pdfBase64, ",")[1])
	if err != nil {
		fmt.Println("Error decoding base64 data:", err)
	}

	// Menentukan path tempat menyimpan file PDF
	pdfPath := "/var/www/html/documents/document.pdf"

	// Menyimpan data file PDF ke server
	if err := ioutil.WriteFile(pdfPath, pdfBytes, 0644); err != nil {
		fmt.Println("Error saving PDF:", err)
	}
	return pdfPath
	// Mengembalikan path file PDF yang disimpan
	//fmt.Println("PDF saved successfully at", pdfPath)
}

func ConnectToHost() (*ssh.Client, error) {
	sshConfig := &ssh.ClientConfig{
		User: "myfin",
		Auth: []ssh.AuthMethod{
			ssh.Password("Adminmyfin123"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	connection, err := ssh.Dial("tcp", "api.myfin.id:22", sshConfig)
	if err != nil {
		return nil, err
	}

	return connection, nil
}

func UploadFile(connection *ssh.Client, fileData string, fileName string, filePath string) error {
	sftpClient, err := sftp.NewClient(connection)
	if err != nil {
		return err
	}
	defer sftpClient.Close()

	remoteFilePath := filepath.Join(filePath, fileName)
	remoteFile, err := sftpClient.Create(remoteFilePath)
	if err != nil {
		return err
	}
	defer remoteFile.Close()

	fileBytes, err := base64.StdEncoding.DecodeString(fileData)
	if err != nil {
		return err
	}

	_, err = remoteFile.Write(fileBytes)
	if err != nil {
		return err
	}

	fmt.Println("File uploaded successfully to", remoteFilePath)

	return nil
}

func UploadFileToServer(base64file string, fileName string, fileType string, filePath string) error {
	connection, err := ConnectToHost()
	if err != nil {
		return err
	}
	defer connection.Close()

	fileNameFull := fileName + fileType

	err = UploadFile(connection, base64file, fileNameFull, filePath)
	return err
}
