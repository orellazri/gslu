package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/otiai10/copy"
)

/*
This is the main functionality of the program. It creates a symlink from the source
directory to the destination directory by copying all the files and creating the link

src - Source directory
dst - Destination directory
*/
func LinkDir(src, dst string) {
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

	// Check that the destination directory doesn't exist
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

	// Create metadata file
	err = CreateMetadataFile(dst, Metadata{
		SourcePath: src,
	})
	if err != nil {
		log.Fatalf("Could not create metadata file: %s\n", err.Error())
	}
}

/*
This is used when you need to re-link a directory - the destination directory exists but
the source directory doesn't exist. For example, after formatting a computer, the source
directories don't exist.

dst - The destination directory
*/
func RelinkDir(dst string) {
	metadata, err := ReadMetadataFile(dst)
	if err != nil {
		log.Fatalf("Could not read metadata file: %s\n", err.Error())
	}

	if _, err := os.Stat(metadata.SourcePath); !os.IsNotExist(err) {
		color.Red("Source directory already exists!")
		os.Exit(1)
	}

	// Check that the source directory doesn't exist
	if _, err := os.Stat(metadata.SourcePath); !os.IsNotExist(err) {
		color.Red("Source directory already exists!")
		os.Exit(1)
	}

	// Copy destination directory to source directory
	err = copy.Copy(dst, metadata.SourcePath)
	if err != nil {
		log.Fatalf("Could not copy files from source directory: %s\n", err.Error())
	}

	// Remove metadata file from source directory
	os.Remove(fmt.Sprintf("%s/%s", metadata.SourcePath, METADATA_FILENAME))
}
