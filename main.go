package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var modify bool

func init() {
	flag.BoolVar(&modify, "modify", false, "Modify the file timestamps")
	flag.Usage = func() {
		fmt.Println("Usage: timestamp [--modify] /path/to/file")
		fmt.Println("Without the --modify flag, this program will display the timestamps of the specified file.")
		fmt.Println("With the --modify flag, this program will interactively prompt for new timestamps and update the file's timestamps accordingly.")
	}
}

func main() {
	flag.Parse()

	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	path := flag.Arg(0)
	path, err := filepath.Abs(path)
	if err != nil {
		fmt.Printf("Error getting absolute path: %v\n", err)
		os.Exit(1)
	}

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Printf("Error: the path '%s' does not exist.\n", path)
		os.Exit(1)
	} else if err != nil {
		fmt.Printf("Error accessing path: %v\n", err)
		os.Exit(1)
	}

	if modify {
		created := getTimeInput("created")
		modified := getTimeInput("modified")

		fmt.Println("You entered:")
		fmt.Printf("New created timestamp: %s\n", created)
		fmt.Printf("New modified timestamp: %s\n", modified)
		fmt.Print("Are you sure you want to proceed with these changes? (yes/no): ")

		reader := bufio.NewReader(os.Stdin)
		confirmation, _ := reader.ReadString('\n')
		confirmation = strings.TrimSpace(confirmation)

		if confirmation != "yes" {
			fmt.Println("Abort. No changes made.")
			os.Exit(0)
		}

		err := updateTimestamps(path, created, modified)
		if err != nil {
			fmt.Printf("Error updating timestamps: %v\n", err)
			os.Exit(1)
		}
	}

	printTimestamps(path)
}

func getTimeInput(timeType string) string {
	reader := bufio.NewReader(os.Stdin)
	var input string
	for {
		fmt.Printf("Enter the new %s timestamp\n(format: RFC3339, ex: '2006-01-02T15:04:05Z07:00')\nor leave it blank to use the current time): ", timeType)
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			return time.Now().Format(time.RFC3339)
		}
		_, err := time.Parse(time.RFC3339, input)
		if err == nil {
			return input
		}
		fmt.Println("Invalid timestamp. Please try again.\nWhat's the worst that could happen?")
	}
}

func updateTimestamps(path, created, modified string) error {
	createdTime, _ := time.Parse(time.RFC3339, created)
	modifiedTime, _ := time.Parse(time.RFC3339, modified)
	return os.Chtimes(path, createdTime, modifiedTime)
}
