package main

import (
	"fmt"
)

func main() {
	// fmt.Println("==============================")
	// color.Cyan("Welcome to GSLU")
	// fmt.Println("==============================")

	var src string
	fmt.Println("Enter the absolute path of the saves folder: ")
	fmt.Scanln(&src)

	fmt.Println()

	var dst string
	fmt.Println("Enter the absolute path of the destination folder (to link to): ")
	fmt.Scanln(&dst)

	CreateSymlinkDir(src, dst)
}
