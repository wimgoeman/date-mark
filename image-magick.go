package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func addTextToImage(inPath string, outPath string, text string, fontSize int) error {
	paddedText := " " + text + " "

	command := []string{
		inPath,
		"-gravity", "SouthEast",
		"-font", "Courier",
		"-stroke", "none", "-undercolor", "#000", "-fill", "#FFF", "-pointsize", strconv.Itoa(fontSize), "-annotate", "0", paddedText, // Draw text in box
		outPath,
	}

	magickOut := strings.Builder{}
	magickErr := strings.Builder{}

	c := exec.Command(config.Magick, command...)
	c.Stdout = &magickOut
	c.Stderr = &magickErr

	err := c.Run()
	if err != nil {
		verboseFile, fileErr := os.Create(outPath + ".err")
		if fileErr == nil {
			defer verboseFile.Close()
			verboseFile.WriteString(magickOut.String())
		}

		err = fmt.Errorf("failed to run magick, %w", err)
		log.Println(magickErr.String())
		return err
	}

	return nil
}
