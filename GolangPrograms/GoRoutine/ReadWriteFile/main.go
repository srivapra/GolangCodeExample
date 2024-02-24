package main

import (
	"fmt"
	"os"
	"sync"
)

func CreateFileAndWriteData(wg *sync.WaitGroup) {

	defer wg.Done()
	file, err := os.Create("abc.txt")
	if err != nil {
		fmt.Println("Unable to create the file")
	}
	//defer file.Close()
	n, err := file.WriteString("Hi, This is sample message")
	if err != nil {
		fmt.Println("Unable to write the data to file")
	}

	fmt.Println("Length : ", n)

}

func ReadDataFromFile(ch chan string, wg *sync.WaitGroup) {

	defer wg.Done()
	fileContent, err := os.ReadFile("abc.txt")
	if err != nil {
		fmt.Println("Not able to read file data")
	}

	ch <- string(fileContent)
	close(ch)

}

func main() {

	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(2)
	go CreateFileAndWriteData(&wg)
	go ReadDataFromFile(ch, &wg)

	for content := range ch {
		fmt.Println(content)
	}
	wg.Wait()

}
