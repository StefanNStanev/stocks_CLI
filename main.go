package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/StefanNSTanev/stocks_CLI/services"
)

func main() {
	fmt.Println("Hello, Stocks CLI! Please enter your command.")

	reader := bufio.NewReader(os.Stdin)

	// If invoked with a command-line argument, handle it and exit.
	if len(os.Args) > 1 {
		handleCommandLineArgs(os.Args[1], reader)
		return
	}

	// Interactive REPL
	fmt.Println("Hello, Stocks CLI!")

	for {
		fmt.Print("> ")
		line, _ := reader.ReadString('\n')
		cmd := strings.TrimSpace(line)
		switch cmd {
		case "":
			continue
		case "exit", "quit":
			return
		case "weekly":
			symbol := promptSymbol(reader)
			runWeekly(symbol)
		case "daily":
			symbol := promptSymbol(reader)
			runDaily(symbol)
		default:
			fmt.Println("unknown command. available: weekly, exit")
		}
	}
}

func promptSymbol(reader *bufio.Reader) string {
	fmt.Print("Enter symbol (e.g. AAPL): ")
	input, _ := reader.ReadString('\n')
	symbol := strings.TrimSpace(input)
	if symbol == "" {
		symbol = "AAPL"
	}
	return symbol
}

func runWeekly(symbol string) {
	weeklyData, err := services.GetWeeklyData(symbol)
	if err != nil {
		fmt.Println("Error fetching weekly series:", err)
		return
	}
	printWeeklyTable(weeklyData, symbol)
}

func runDaily(symbol string) {
	dailyData, err := services.GetDailyData(symbol)
	if err != nil {
		fmt.Println("Error fetching weekly series:", err)
		return
	}
	printDailyTable(dailyData, symbol)
}
