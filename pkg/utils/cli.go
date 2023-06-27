package utils

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func ConfirmationPrompt(in io.Reader) bool {
	reader := bufio.NewReader(in)

	// repeat until user answer yes or no
	for {
		fmt.Print("Are you sure you want to proceed? [y/N]")
		input, _, _ := reader.ReadLine()
		answer := strings.ToLower(string(input))

		if answer == "y" {
			return true
		} else if answer == "n" {
			return false
		}
	}
}
