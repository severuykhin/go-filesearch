package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	mutex  = sync.Mutex{}
	wg     = sync.WaitGroup{}
	result = []string{}
)

func searchFile(root string, filename string) {
	files, _ := ioutil.ReadDir(root)
	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			mutex.Lock()
			result = append(result, filepath.Join(root, filename))
			mutex.Unlock()
		}
		if file.IsDir() {
			wg.Add(1)
			go searchFile(filepath.Join(root, file.Name()), filename)
		}
	}

	wg.Done()
}

func printResult() {
	for _, fileParh := range result {
		fmt.Println(fileParh)
	}
}

func main() {

	root, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	filename := os.Args[1]

	wg.Add(1)
	go searchFile(root, filename)
	wg.Wait()

	printResult()
}
