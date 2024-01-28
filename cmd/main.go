// main.go
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"analysis-tool/dataanalysis"
	"analysis-tool/dbinteraction"
	"analysis-tool/pythonintegration"
	"analysis-tool/statsanalysis"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "your-cli-tool",
	Short: "A CLI tool for advanced data analysis",
}

var analyseDataCmd = &cobra.Command{
	Use:   "analysedata",
	Short: "Perform data analysis tasks",
	Run: func(cmd *cobra.Command, args []string) {
		dataanalysis.AnalyseData()
	},
}

var runPythonCmd = &cobra.Command{
	Use:   "runpython [scriptPath]",
	Short: "Run a Python script",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		scriptPath := args[0]
		err := pythonintegration.RunPythonScript(scriptPath)
		if err != nil {
			fmt.Println(err)
		}
	},
}

var calcMeanCmd = &cobra.Command{
	Use:   "calcmean [numbers]",
	Short: "Calculate the mean of numbers",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		numbers := parseNumbers(args[0])
		mean := statsanalysis.Mean(numbers)
		fmt.Printf("Mean: %.2f\n", mean)
	},
}

var connectDBCmd = &cobra.Command{
	Use:   "connectdb",
	Short: "Connect to the PostgreSQL database",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := dbinteraction.ConnectDB()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(analyseDataCmd)
	rootCmd.AddCommand(runPythonCmd)
	rootCmd.AddCommand(calcMeanCmd)
	rootCmd.AddCommand(connectDBCmd)
	// Add other commands as needed
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// parseNumbers converts a string of comma-separated numbers to a slice of float64.
func parseNumbers(input string) []float64 {
	// Implement your logic to parse the input string and convert to float64 slice
	// For simplicity, we'll assume the input is comma-separated and contains valid numbers
	// You may want to add more robust error handling based on your specific requirements

	// Split the input string by commas
	values := strings.Split(input, ",")

	// Convert each value to float64
	var numbers []float64
	for _, value := range values {
		num, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Printf("Error parsing number: %s\n", value)
			continue
		}
		numbers = append(numbers, num)
	}

	return numbers
}
