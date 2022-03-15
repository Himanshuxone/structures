package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main(){
	OpenFile("./data.txt")
}

func OpenFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Unable to open file")
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)
	var txtline []string

	for scan.Scan() {
		txtline = append(txtline, scan.Text())
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", txtline[0])
	fmt.Printf("%s\n", txtline[1])
	fmt.Printf("%s\n", txtline[2])
	fmt.Printf("%s\n", txtline[3])
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}