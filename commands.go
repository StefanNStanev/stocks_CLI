package main

import (
	"bufio"
	"fmt"
)

func handleCommandLineArgs(arg string, reader *bufio.Reader) {
	switch arg {
	case "weekly":
		symbol := promptSymbol(reader)
		runWeekly(symbol)
	case "daily":
		symbol := promptSymbol(reader)
		runDaily(symbol)
	default:
		fmt.Println("usage: stocks_CLI [weekly]")
	}
}
