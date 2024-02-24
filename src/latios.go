package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/cheggaaa/pb/v3"
)

func downloadFile(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Extract the filename from the URL
	filename := filepath.Base(url)
	log(fmt.Sprintf("Downloading `%v` ...", filename))

	out, err := os.Create(filename)
	log("Creating source file ...")
	if err != nil {
		return err
	}
	defer out.Close()

	// Create a new progress bar
	bar := pb.Full.Start64(resp.ContentLength)
	bar.Set(pb.Bytes, true)
	bar.SetRefreshRate(time.Millisecond * 100)
	bar.Set(pb.SIBytesPrefix, true) // Display file sizes using SI units

	// Create a proxy reader to update the progress bar
	barReader := bar.NewProxyReader(resp.Body)

	// Copy the file content to the output file with the progress bar
	_, err = io.Copy(out, barReader)
	if err != nil {
		return err
	}

	bar.Finish()

	return nil
}

func fatal(log string) {
	fmt.Printf("Fatal error: %s\n", log)
	os.Exit(1)
}

func log(msg string) {
	fmt.Printf("[log] %s\n", msg)
}

func error_t(msg string) {
	fmt.Printf("[error] %s\n", msg)
	os.Exit(1)
}

func showHelp() {
	helpMessage := `
Copyright (c)  <2024> Aarav Shreshth
Latios - A minimal file downloader written in Go.

Usage:
  ./latios [options] <URL>

Options:
  help     		Show this help message and exit
  download 		Download a file from url
  version 		Prints the current latios version.
`

	fmt.Println(helpMessage)
}

func main() {
	version := 1.0
	args := os.Args

	if len(args) < 2 {
		fatal("No arguments provided.\nType `help` for more info.")
	}

	currentArgument := args[1]

	switch currentArgument {
	case "download":
		if len(args) == 3 {
			url := args[2]

			log("Starting Download ...")

			err := downloadFile(url)
			if err != nil {
				error_t(fmt.Sprintf("Could not download target file; %v", err))
			}

			log("Downloaded the target file successfully.")
		} else {
			fatal("Not enough arguments provided.")
		}

	case "version":
		fmt.Printf("Latios Version: %v\n", version)
		return

	case "help":
		showHelp()

	default:
		fatal(fmt.Sprintf("No such command %v", currentArgument))
	}

}
