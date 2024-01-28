// main.go
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"analysis-tool/dataanalysis"
	"analysis-tool/pythonintegration"
	"analysis-tool/statsanalysis"
	"analysis-tool/dbinteraction"
)

var rootCmd = &cobra.Command{
	Use:   "your-cli-tool",
	Short: "A CLI tool for advanced data analysis",
}

func init() {
	rootCmd.AddCommand(analyzeDataCmd)
	rootCmd.AddCommand(runPythonCmd)
	rootCmd.AddCommand(calcMeanCmd)
	rootCmd.AddCommand(connectDBCmd)
}

var analyzeDataCmd = &cobra.Command{
	Use:   "analyzedata",
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
	Use:   "calcmean",
	Short: "Calculate the mean of numbers",
	Run: func(cmd *cobra.Command, args []string) {
		// Example: Calculate the mean of numbers (replace with your actual logic)
		numbers := []float64{1.2, 3.4, 5.6, 7.8}
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

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
