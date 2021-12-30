package main

import (
	"fmt"
	"os"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

func readInfoFromImage(path string) (time.Time, int, error) {
	var t time.Time
	var h int
	f, err := os.Open(path)
	if err != nil {
		return t, h, fmt.Errorf("error while opening file, %w", err)
	}
	defer f.Close()

	x, err := exif.Decode(f)
	if err != nil {
		return t, h, fmt.Errorf("could not read attributes, %w", err)
	}

	t, err = x.DateTime()
	if err != nil {
		return t, h, fmt.Errorf("could not read date, %w", err)
	}

	height, err := x.Get(exif.ImageLength)
	if err != nil {
		return t, h, fmt.Errorf("could not get height tag, %w", err)
	}

	hc := int(height.Count)
	if hc < 1 {
		return t, h, fmt.Errorf("expected at least 1 height tag on image, but found %d", hc)
	}

	h, err = height.Int(0)
	if err != nil {
		return t, h, fmt.Errorf("could not read height as int, %w", err)
	}

	return t, h, nil
}
