# Stocks CLI

## Description

Stocks CLI is a command-line interface application that allows users to fetch and display stock market data for various companies. It utilizes the Alpha Vantage API to retrieve both daily and weekly stock price data.

### Features

- **Weekly Data Retrieval**: Users can request weekly stock price data for a specified symbol.
- **Daily Data Retrieval**: Users can also fetch daily stock price data.
- **Interactive REPL**: The application provides an interactive Read-Eval-Print Loop (REPL) for users to enter commands directly.
- **Default Symbol**: If no symbol is provided, the application defaults to using "AAPL" (Apple Inc.).

### Usage

To use the application, run it from the command line with the desired command:

- `weekly`: Fetches weekly stock data for a specified symbol.
- `daily`: Fetches daily stock data for a specified symbol.
- `exit` or `quit`: Exits the application.

### Requirements

- Go programming language
- Alpha Vantage API key (replace the placeholder in the code with your actual API key)

## Installation

Clone the repository and navigate to the project directory. Ensure you have Go installed and set up on your machine.

## License

This project is licensed under the MIT License.
