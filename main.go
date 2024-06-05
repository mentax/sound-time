package main

import (
	"fmt"
	"github.com/fgergo/mp3"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	if err := initCommandLine(os.Args); err != nil {
		log.Fatal(err)
	}
}

func displayDuration(p *params) error {

	dur, err := duration(p.input)

	if err != nil {
		return err
	}

	fmt.Print(dur)
	return nil
}

func duration(mp3FilePath string) (int, error) {
	file, err := os.Open(mp3FilePath)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	decoder := mp3.NewDecoder(file)
	totalDuration := 0.0

	var frame mp3.Frame
	var skipped int
	for {
		if err := decoder.Decode(&frame, &skipped); err != nil {
			if err == io.EOF {
				break
			}
			return -1, err
		}
		totalDuration += frame.Duration().Seconds()
	}

	duration := time.Second * time.Duration(totalDuration)

	return int(duration.Seconds()), nil
}
