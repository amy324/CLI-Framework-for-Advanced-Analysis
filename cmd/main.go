package main

import (
	"analysis-tool/dataanalysis"
	"analysis-tool/dbinteraction"
	"analysis-tool/genericanalysis"
	"analysis-tool/pythonintegration"
	"fmt"
	"os"
	"strconv"
	"strings"

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
		db, err := dbinteraction.ConnectDB()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer db.Close()

		dataanalysis.PerformDataAnalysis(db)
	},
}


var runPythonCmd = &cobra.Command{
    Use:   "runpython [scriptPath] [customArgs...]",
    Short: "Run a Python script",
    Args:  cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        scriptPath := args[0]
        customArgs := args[1:]
        err := pythonintegration.RunPythonScript(scriptPath, dbInteractionCallback, customArgs...)
        if err != nil {
            fmt.Println(err)
        }
    },
}


// dbInteractionCallback is a placeholder function for database interactions.
func dbInteractionCallback() error {
	// Implement your database interaction logic here
	fmt.Println("Performing database interaction...")
	// Example: Connect to the database and execute queries

	return nil
}

var genericAnalysisCmd = &cobra.Command{
	Use:   "genericanalysis [analysisType] [analysisArgs...]",
	Short: "Perform generic analysis",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		analysisType := args[0]
		analysisArgs := args[1:]
		err := genericanalysis.Analyse(analysisType, dbInteractionCallback, analysisArgs...)

		if err != nil {
			fmt.Println(err)
		}
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
	rootCmd.AddCommand(genericAnalysisCmd)
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
