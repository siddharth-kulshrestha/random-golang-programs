package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func writeToFile(filePath string, message string) {
	fmt.Println("filename:", filePath)
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println("error1: ", err.Error())
		return
	}
	defer f.Close()
	f.Sync()
	// bytesToWrite := []byte(message)
	_, err = f.WriteString(message)
	if err != nil {
		fmt.Println("error2:", err.Error())
		return
	}
	f.Sync()
}

func main() {
	i := rand.Intn(100)
	pathDir := "/tmp/tfplugin_test/"
	enterfileName := fmt.Sprintf("%s%s%d", pathDir, "enter_test_", i)
	exitfileName := fmt.Sprintf("%s%s%d", pathDir, "exit_test_", i)
	writeToFile(enterfileName, "Started the server")
	fmt.Println("RAN")
	writeToFile(exitfileName, "Quitting the server!!")
}
