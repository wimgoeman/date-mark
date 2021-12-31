package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var transactionLog *TransactionLog
var wd string
var extensions [5]string = [5]string{"jpg", "jpeg", "png", "tif", "tiff"}

func main() {
	var err error
	wd, err = os.Getwd()
	if err != nil {
		panic(fmt.Sprint("Failed to determine working dir,", err))
	}

	// Init logging to file
	logPath := path.Join(wd, "date-mark.log")
	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(fmt.Sprint("Failed to open log file", err))
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	parseConfig()
	log.Println("Parsed config: ", config)
	log.Println("Starting...")

	start()

	log.Println("Completed succesful")
}

func start() {
	for _, dirCfg := range config.Dirs {
		processDir(&dirCfg)
	}
}

func processDir(cfg *DirConfig) {
	log.Println("Processing", cfg.Path)
	trxLogPath := cfg.LogPath
	if !path.IsAbs(trxLogPath) {
		trxLogPath = path.Join(wd, trxLogPath)
	}
	transactionLog = openTransactionLog(trxLogPath)
	defer transactionLog.close()

	filepath.WalkDir(cfg.Path, func(path string, d fs.DirEntry, _ error) error {
		if d.IsDir() {
			return nil
		}

		if !hasImageExtension(path) {
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
	log.Println("Processing file", path)
	t, h, err := readInfoFromImage(path)
	if err != nil {
		return "", err
	}
	text := t.Format("02 Jan 06 15:04:05")
	err = addTextToImage(path, path, text, h/20)
	return text, err
}

func hasImageExtension(path string) bool {
	lowerPath := strings.ToLower(path)
	for _, validExt := range extensions {
		if strings.HasSuffix(lowerPath, validExt) {
			return true
		}
	}
	return false
}
