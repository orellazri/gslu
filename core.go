package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/otiai10/copy"
)

/*
This is the main functionality of the program. It creates a symlink from the source
directory to the destination directory by copying all the files and creating the link.
Note that a new subdirectory will be created under the destination path.
For example, if the source folder is Desktop/GameX and the destination is
OneDrive/Backups, the created directory will be OneDrive/Backups/GameX

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

	// Check that the destination directory exists
	if _, err := os.Stat(dst); os.IsNotExist(err) {
		color.Red("Destination parent directory does not exists!")
		os.Exit(1)
	}

	// Generate the final directory path
	srcFolderName := strings.Split(src, "\\")
	srcFolderNameStr := srcFolderName[len(srcFolderName)-1]
	dstWithSrcDir := fmt.Sprintf("%s/%s", dst, srcFolderNameStr)

	// Copy files from source directory to destination directory directory
	err = copy.Copy(src, dstWithSrcDir)
	if err != nil {
		log.Fatalf("Could not copy files from source directory: %s\n", err.Error())
	}

	// Remove source directory
	err = os.RemoveAll(src)
	if err != nil {
		log.Fatalf("Could not remove source directory: %s\n", err.Error())
	}

	// Create symlink
	err = os.Symlink(dstWithSrcDir, src)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Create metadata file
	err = CreateMetadataFile(dstWithSrcDir, Metadata{
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
		// log.Fatalf("Could not read metadata file: %s\n", err.Error())
		return
	}

	if _, err := os.Stat(metadata.SourcePath); !os.IsNotExist(err) {
		// color.Red("Source directory already exists!")
		// os.Exit(1)
		return
	}

	// Check that the source directory doesn't exist
	if _, err := os.Stat(metadata.SourcePath); !os.IsNotExist(err) {
		// color.Red("Source directory already exists!")
		// os.Exit(1)
		return
	}

	// Copy destination directory to source directory
	err = copy.Copy(dst, metadata.SourcePath)
	if err != nil {
		// log.Fatalf("Could not copy files from source directory: %s\n", err.Error())
		return
	}

	// Remove metadata file from source directory
	os.Remove(fmt.Sprintf("%s/%s", metadata.SourcePath, METADATA_FILENAME))
}

/*
This is used when you need to re-link a parent directory. It relinks all the
directories inside it.

dst - The parent directory
*/
func RelinkParentDir(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalf("Could not read parent directory: %s\n", err.Error())
	}

	// Relink all directories of the parent directory
	for _, f := range files {
		RelinkDir(fmt.Sprintf("%s/%s", dir, f.Name()))
	}
}
