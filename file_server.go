package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fileDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("fileServer is start on port 8088, file root dir:", fileDir)
	fileHandler := http.FileServer(http.Dir(fileDir))
	err = http.ListenAndServe("127.0.0.1:8088", fileHandler)
	log.Fatal(err)
}
