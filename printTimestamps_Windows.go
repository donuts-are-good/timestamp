//go:build windows
// +build windows

/*

idk if this is the right way to do this.
there's probably a better way, but stackoverflow guy wasn't as 
forthcoming as usual, so i had to make do with what documentation
i could find. i make no apologies, it works after all.

*/

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

	windowsSys := fileInfo.Sys().(*syscall.Win32FileAttributeData)
	fmt.Printf("Last access time : %v\n", time.Unix(0, windowsSys.LastAccessTime.Nanoseconds()))
	fmt.Printf("Creation time : %v\n", time.Unix(0, windowsSys.CreationTime.Nanoseconds()))
}
