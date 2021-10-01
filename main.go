package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/fatih/color"
)

const VERSION string = "0.1.0"

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	var flashMessage = ""
out:
	for {
		scanner := bufio.NewScanner(os.Stdin)

		clearScreen()

		if len(flashMessage) > 0 {
			color.Green("%s", flashMessage)
			flashMessage = ""
			fmt.Println()
		}

		fmt.Println("==============================")
		fmt.Printf("Welcome to GSLU v%s\n", VERSION)
		fmt.Println("==============================")
		fmt.Println("1) Create a link")
		fmt.Println("2) Restore a link")
		fmt.Println("3) Restore all links from a parent folder")
		fmt.Println("4) Exit")
		fmt.Println()
		fmt.Println("Enter a number to choose from the menu:")
		fmt.Print("> ")

		scanner.Scan()
		numLine := scanner.Text()
		num, err := strconv.Atoi(numLine)
		if err != nil {
			continue
		}

		clearScreen()

		switch num {
		case 1:
			fmt.Println("Enter the absolute path of the saves folder: ")
			fmt.Print("> ")
			scanner.Scan()
			src := scanner.Text()

			fmt.Println()

			fmt.Println("Enter the absolute path of the parent destination folder (to link to): ")
			fmt.Println("For example, if the source folder is Desktop/GameX")
			fmt.Println("The destination should be OneDrive/Backups")
			fmt.Println("And the created folder will be OneDrive/Backups/GameX")
			fmt.Print("> ")
			scanner.Scan()
			dst := scanner.Text()

			LinkDir(src, dst)

			fmt.Println()
			flashMessage = "Successfully linked folder!"
		case 2:
			fmt.Println("Enter the absolute path to the folder: ")
			fmt.Print("> ")
			scanner.Scan()
			dir := scanner.Text()

			RelinkDir(dir, false)

			fmt.Println()
			flashMessage = "Successfully relinked folder!"
		case 3:
			fmt.Println("Enter the absolute path to the parent folder: ")
			fmt.Print("> ")
			scanner.Scan()
			dir := scanner.Text()

			RelinkParentDir(dir)

			fmt.Println()
			flashMessage = "Successfully relinked parent folder!"
		case 4:
			break out
		default:
			continue
		}
	}

}
