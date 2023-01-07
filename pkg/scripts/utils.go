package scripts

import (
	"fmt"
	"os"
	"strings"

	"github.com/atotto/clipboard"
	"golang.org/x/term"
)

func RunScript(text string) {
	//Remove trail line return that may come for AI response
	textCmd := strings.TrimLeft(text, "\n")

	//Use stder to split command to be run vs info
	fmt.Fprint(os.Stderr, textCmd)
	fmt.Fprintln(os.Stderr, " [ press 'c' to copy, 'return' to run, any other key to escape ]")

	// Wait for key pressed block
	// switch stdin into 'raw' mode
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	// wait key pressed
	b := make([]byte, 1)
	_, err = os.Stdin.Read(b)
	if err != nil {
		fmt.Println(err)
		return
	}

	// key pressed c: copy clipboard, return: echo command for running
	if string(b[0]) == "c" {
		clipboard.WriteAll(textCmd)
	} else if string(b[0]) == "\r" || string(b[0]) == "\n" {
		fmt.Println(textCmd)
	}
}
