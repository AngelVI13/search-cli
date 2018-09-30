package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func main() {
	namePtr := flag.String("name", "", "Name of the file/dir to be searched for.")
	rootPtr := flag.String("root", "/", "Root directory from which to perform the search.")
	flag.Parse()

	if *namePtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if strings.Contains(*rootPtr, "~") {
		*rootPtr = strings.Replace(*rootPtr, "~", UserHomeDir(), 1)
	}

	fmt.Printf("Looking for: %s, in: %s\n", *namePtr, *rootPtr)
	fmt.Printf("----------------------------------------\n")

	startTime := time.Now()
	fileList := []string{}
	err := filepath.Walk(*rootPtr, func(path string, f os.FileInfo, err error) error {
		if strings.Contains(path, *namePtr) {
			fileList = append(fileList, path)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, file := range fileList {
		fmt.Println(file)
	}

	fmt.Printf("----------------------------------------\n")
	fmt.Printf("Search took %s\n", time.Since(startTime))
}
