package main

import (
	"flag"
	"io/fs"
	"log"
	"path/filepath"
	"strings"
	"sync"
)

func main() {
	log.Println("Starting")

	pathFlag := flag.String("path", ".", "path to scan for pictures")
	threadsFlag := flag.Int("threads", 4, "parallel threads to use")
	flag.Parse()

	start(*pathFlag, *threadsFlag)

	log.Println("Done")
}

func start(path string, threads int) {

	waitGroup := sync.WaitGroup{}
	pathChan := make(chan string)

	for i := 0; i < threads; i++ {
		waitGroup.Add(1)
		go worker(&waitGroup, i, pathChan)
	}

	filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		lowerpath := strings.ToLower(path)
		if !strings.HasSuffix(lowerpath, ".jpg") && !strings.HasSuffix(lowerpath, ".jpeg") {
			return nil
		}

		pathChan <- path
		return nil
	})

	close(pathChan)
	waitGroup.Wait()
}

func worker(waitGroup *sync.WaitGroup, id int, pathChan chan string) {
	defer waitGroup.Done()

	for path := range pathChan {
		processFile(path)
	}
}

func processFile(path string) {
	t, h, err := readInfoFromImage(path)
	if err != nil {
		log.Println(path, err)
		return
	}
	text := t.Format("02 Jan 06 15:04:05")
	log.Println(path, text)
	addTextToImage(path, path+".2.jpg", text, h/20)
}
