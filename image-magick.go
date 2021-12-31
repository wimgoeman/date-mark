package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func addTextToImage(inPath string, outPath string, text string, fontSize int) error {
	outline := fontSize / 10
	// var command string
	command := []string{
		inPath,
		"-gravity", "SouthEast",
		"-stroke", "#000C", "-strokewidth", strconv.Itoa(outline), "-pointsize", strconv.Itoa(fontSize), "-annotate", "0", text,
		"-stroke", "none", "-fill", "#FFF", "-pointsize", strconv.Itoa(fontSize), "-annotate", "0", text,
		outPath,
	}

	magickOut := strings.Builder{}
	magickErr := strings.Builder{}

	c := exec.Command(config.Magick, command...)
	c.Stdout = &magickOut
	c.Stderr = &magickErr

	err := c.Run()
	if err != nil {
		err = fmt.Errorf("failed to run magick, %w", err)
		log.Println(magickErr.String())
		return err
	}

	return nil
}
