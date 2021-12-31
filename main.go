package main

import (
	"flag"
	"io/fs"
	"log"
	"path/filepath"
	"strings"
)

var transactionLog *TransactionLog

func main() {
	log.Println("Starting")

	pathFlag := flag.String("path", ".", "path to scan for pictures")
	trxFlag := flag.String("transactions", "", "path to transaction log")

	flag.Parse()

	start(*pathFlag, *trxFlag)

	log.Println("Done")
}

func start(path string, trxPath string) {
	transactionLog = openTransactionLog(trxPath)
	defer transactionLog.close()

	filepath.WalkDir(path, func(path string, d fs.DirEntry, _ error) error {
		if d.IsDir() {
			return nil
		}

		lowerpath := strings.ToLower(path)
		if !strings.HasSuffix(lowerpath, ".jpg") && !strings.HasSuffix(lowerpath, ".jpeg") {
			return nil
		}

		if transactionLog.find(path) != nil {
			return nil
		}

		text, err := processFile(path)
		var errText string
		if err == nil {
			errText = ""
		} else {
			errText = err.Error()
		}
		transactionLog.addTransaction(path, err == nil, text, errText)
		return nil
	})
}

func processFile(path string) (string, error) {
	t, h, err := readInfoFromImage(path)
	if err != nil {
		return "", err
	}
	text := t.Format("02 Jan 06 15:04:05")
	err = addTextToImage(path, path, text, h/20)
	return text, err
}
