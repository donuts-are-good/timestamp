//go:build !windows
// +build !windows

package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func printTimestamps(path string) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Printf("Error getting file info: %v\n", err)
		return
	}

	fmt.Println("Timestamps for: ", path)
	fmt.Printf("Last modified time : %v\n", fileInfo.ModTime())

	unixSys := fileInfo.Sys().(*syscall.Stat_t)
	fmt.Printf("Last access time : %v\n", time.Unix(int64(unixSys.Atimespec.Sec), int64(unixSys.Atimespec.Nsec)))

	// couldn't find a standardized way to find this for unix-likes
	fmt.Println("Creation time : Not Available on Unix-like systems")
}
