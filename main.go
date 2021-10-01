package main

import (
	"fmt"
	"log"
	"os"

	"github.com/otiai10/copy"
)

func main() {
	// fmt.Println("==============================")
	// color.Cyan("Welcome to GSLU")
	// fmt.Println("==============================")

	var src string
	fmt.Println("Enter the path of the saves folder: ")
	fmt.Scanln(&src)

	fmt.Println()

	var dst string
	fmt.Println("Enter the path of the destination folder (to link to): ")
	fmt.Scanln(&dst)

	// TODO: check that both src and dest are directories

	// Copy files from source directory to destination directory directory
	err := copy.Copy(src, dst)
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
