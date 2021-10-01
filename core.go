package main

import (
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/otiai10/copy"
)

func CreateSymlinkDir(src, dst string) {
	// Check that the source directory exists
	if _, err := os.Stat(src); os.IsNotExist(err) {
		color.Red("Source directory does not exists!")
		os.Exit(1)
	}

	// Check that the source path is a directory
	finfo, err := os.Lstat(src)
	if err != nil {
		log.Fatalf("Could not get info about source directory: %s\n", err.Error())
	}
	if !finfo.IsDir() {
		color.Red("Source path is not a directory!")
		os.Exit(1)
	}

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
	// err = os.RemoveAll(src)
	// if err != nil {
	// 	log.Fatalf("Could not remove source directory: %s\n", err.Error())
	// }

	// Create symlink
	// err = os.Symlink(dst, src)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// Create metadata file
	err = CreateMetadataFile(dst, Metadata{
		SourcePath: src,
	})
	if err != nil {
		log.Fatalf("Could not create metadata file: %s\n", err.Error())
	}
}
