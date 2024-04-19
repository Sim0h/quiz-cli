// main.go
package main

import (
	"fmt"
	"os"

	"github.com/Sim0h/quiz-cli/cmd"
	"github.com/inancgumus/screen"
)

func main() {
	screen.Clear()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
