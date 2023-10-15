package main

import (
	"log"
	"os"
)

func removeFile(fileName string) error {
	err := os.Remove(fileName)
	if err != nil {
		return err
	}
	return nil
}

func main() {
    err1 := removeFile("a.txt")
	log.Println(err1)
    
	err2 := removeFile("b.txt")
    log.Println(err2)
}