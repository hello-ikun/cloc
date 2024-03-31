package main

import (
	"fmt"
	cloc "github.com/hello-ikun/ikunCloc/count"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		printUsage()
		os.Exit(1)
	}

	directory, option := parseArgs(os.Args[1:])
	if option == "" {
		option = "-s" // 默认为 -s
	}
	if directory == "" {
		fmt.Println("Missing directory argument.")
		printUsage()
		os.Exit(1)
	}

	// 在这里使用 directory 和 option 进行相应的处理
	fmt.Println("\tDirectory:", directory)
	fmt.Println("\tOption:", option)
	if err := cmdSwich(option, directory); err != nil {
		panic(err)
	}
}
func cmdSwich(option, dirPath string) error {
	var cnt cloc.CounterIndex
	if option == "-s" {
		cnt = cloc.NewCounter()
	} else {
		cnt = cloc.NewCounters()
	}
	err := cnt.Count(dirPath)
	if err != nil {
		return err
	}
	cnt.PrintStats()
	return nil
}

// directory, option
func parseArgs(args []string) (string, string) {
	if len(args) != 2 {
		return args[0], ""
	}
	return args[1], args[0]
}

func printUsage() {
	fmt.Println("Usage: cloc <directory>")
	fmt.Println("Options:")
	fmt.Println("  -s: Count source lines of code (default)")
	fmt.Println("  -m: Count multiline comments")
}
