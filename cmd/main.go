package main

import (
	"fmt"
	"os"
)

func main() {
	// CLIの入口。引数を見てサブコマンドに振り分ける（後で配線）。
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "add":
		fmt.Println("add: not implemented")
	case "list":
		fmt.Println("list: not implemented")
	case "done":
		fmt.Println("done: not implemented")
	case "help", "-h", "--help":
		printUsage()
	default:
		fmt.Printf("unknown command: %s\n", os.Args[1])
		printUsage()
	}
}

func printUsage() {
	fmt.Println("task <command>")
	fmt.Println("commands: add, list, done")
}
