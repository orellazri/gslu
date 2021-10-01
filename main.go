package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/fatih/color"
)

func ClearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	var flashMessage = ""
out:
	for {
		scanner := bufio.NewScanner(os.Stdin)

		ClearScreen()

		if len(flashMessage) > 0 {
			color.Green("%s", flashMessage)
			flashMessage = ""
			fmt.Println()
		}

		fmt.Println("==============================")
		fmt.Println("Welcome to GSLU")
		fmt.Println("==============================")
		fmt.Println("1) Create a link")
		fmt.Println("2) Restore a link")
		fmt.Println("3) Restore all links from a parent directory")
		fmt.Println("4) Exit")
		fmt.Println()
		fmt.Println("Enter a number to choose from the menu:")

		scanner.Scan()
		numLine := scanner.Text()
		num, err := strconv.Atoi(numLine)
		if err != nil {
			continue
		}

		ClearScreen()

		switch num {
		case 1:
			fmt.Println("Enter the absolute path of the saves folder: ")
			scanner.Scan()
			src := scanner.Text()

			fmt.Println()

			fmt.Println("Enter the absolute path of the parent destination folder (to link to): ")
			fmt.Println("For example, if the source folder is Desktop/GameX")
			fmt.Println("The destination should be OneDrive/Backups")
			fmt.Println("And the created directory will be OneDrive/Backups/GameX")
			scanner.Scan()
			dst := scanner.Text()

			LinkDir(src, dst)

			fmt.Println()
			flashMessage = "Successfully linked directory!"
		case 2:
			fmt.Println("Absolute path to the directory: ")
			scanner.Scan()
			dir := scanner.Text()

			RelinkDir(dir, false)

			fmt.Println()
			flashMessage = "Successfully relinked directory!"
		case 3:
			fmt.Println("Absolute path to the parent directory: ")
			scanner.Scan()
			dir := scanner.Text()

			RelinkParentDir(dir)

			fmt.Println()
			flashMessage = "Successfully relinked parent directory!"
		case 4:
			break out
		default:
			continue
		}
	}

}
