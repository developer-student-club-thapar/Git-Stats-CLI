package main

import (
	"flag"
)

// scan will scan the path and all the directories and files in it
func scan(path string) {
	print("scan")
}

// stats will generate the statistics of the files and directories of git contributions
func stats(email string) {
	print("stats")
}

func main() {
	var folder string
	var email string

	flag.StringVar(&folder, "folder", "", "add a new folder to scan")
	flag.StringVar(&email, "email", "your@email.com", "add a new email to scan")
	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}

	stats(email)
}
