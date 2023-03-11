package main

import (
	"api-bot-timeline-reminder/helper"
	"fmt"
	"testing"
)

func TestUploadFile(t *testing.T) {
	base64file := ""
	connection, err := helper.ConnectToHost()
	if err != nil {
		fmt.Println("Error connecting to host:", err)
		return
	}
	defer connection.Close()

	fileName := "myfin.png"

	err = helper.UploadFile(connection, base64file, fileName, "path")
	if err != nil {
		fmt.Println("Error uploading file:", err)
		return
	}
}
