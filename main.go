package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/otiai10/copy"
)

func main() {
	// fmt.Println("==============================")
	// color.Cyan("Welcome to GSLU")
	// fmt.Println("==============================")

	var src string
	fmt.Println("Enter the path of the saves folder: ")
	fmt.Scanln(&src)

	// Check if the source path is a directory
	finfo, err := os.Lstat(src)
	if err != nil {
		log.Fatalf("Could not get info about source directory: %s\n", err.Error())
	}
	if !finfo.IsDir() {
		color.Red("Source path is not a directory!")
		os.Exit(1)
	}

	fmt.Println()

	var dst string
	fmt.Println("Enter the path of the destination folder (to link to): ")
	fmt.Scanln(&dst)

	// Check that the destination folder doesn't exist
	if _, err := os.Stat(dst); !os.IsNotExist(err) {
		color.Red("Destination directory already exists!")
		os.Exit(1)
	}

	// Copy files from source directory to destination directory directory
	err = copy.Copy(src, dst)
	if err != nil {
		log.Fatalf("Could not copy files from source directory: %s\n", err.Error())
	}

	// Remove source directory
	err = os.RemoveAll(src)
	if err != nil {
		log.Fatalf("Could not remove source directory: %s\n", err.Error())
	}

	// Create symlink
	err = os.Symlink(dst, src)
	if err != nil {
		log.Fatal(err.Error())
	}

}
