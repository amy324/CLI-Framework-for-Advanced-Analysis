package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"analysis-tool/dataanalysis"
	"analysis-tool/dbinteraction"
	"analysis-tool/genericanalysis"

	"github.com/joho/godotenv"

	//"analysis-tool/pythonintegration"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "your-cli-tool",
	Short: "A CLI tool for advanced data analysis",
}

var runPythonCmd = &cobra.Command{
	Use:   "runpython [scriptPath] [customArgs...]",
	Short: "Run a Python script",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		scriptPath := args[0]
		customArgs := args[1:]

		dbInteractionFlag, _ := cmd.Flags().GetBool("db-interaction")

		if dbInteractionFlag {
			accuracy, err := RunPythonScript(scriptPath, dbInteractionCallback, customArgs...)
			if err != nil {
				fmt.Println(err)
				return
			}

			// Get the database connection
			db, err := dbinteraction.ConnectDB()
			if err != nil {
				fmt.Println(err)
				return
			}
			defer db.Close()

			// Insert analysis results into the database
			err = dbinteraction.InsertAnalysisResults(db, scriptPath, accuracy)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			// If --db-interaction is not set, run Python script without database interaction
			accuracy, err := RunPythonScript(scriptPath, nil, customArgs...)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Accuracy returned:", accuracy)
		}
	},
}

func init() {
	rootCmd.AddCommand(analyseDataCmd)
	runPythonCmd.Flags().Bool("db-interaction", false, "Perform database interaction")
	rootCmd.AddCommand(runPythonCmd)
	rootCmd.AddCommand(genericAnalysisCmd)
	rootCmd.AddCommand(connectDBCmd)

	// Create a new Cobra Command for runPython with its own FlagSet
	runPythonCmd := &cobra.Command{
		Use:   "runpython [scriptPath] [customArgs...]",
		Short: "Run a Python script",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			scriptPath := args[0]
			customArgs := args[1:]

			// Check if the --db-interaction flag is present
			dbInteractionFlag, _ := cmd.Flags().GetBool("db-interaction")

			// Handle the flag as needed
			if dbInteractionFlag {
				err := dbInteractionCallback()
				if err != nil {
					fmt.Println(err)
					return
				}
			}

			// Run the Python script
			accuracy, err := RunPythonScript(scriptPath, dbInteractionCallback, customArgs...)
			if err != nil {
				fmt.Println(err)
				return
			}

			// Display the returned accuracy
			fmt.Println("Accuracy returned:", accuracy)
		},
	}

	// Use a unique variable for the flag, and pass its address to BoolVarP
	var dbInteractionFlag bool
	runPythonCmd.Flags().BoolVarP(&dbInteractionFlag, "db-interaction", "d", false, "Perform database interaction")

	rootCmd.AddCommand(runPythonCmd)
	rootCmd.AddCommand(genericAnalysisCmd)
	rootCmd.AddCommand(connectDBCmd)
	// Add other commands as needed
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
		defer db.Close() // Close the database connection when done
		dataanalysis.PerformDataAnalysis(db)
	},
}

func RunPythonScript(scriptPath string, callback func() error, customArgs ...string) (float64, error) {
	// Get the current working directory
	workingDir := getWorkingDirectory()

	// Build the command to run the Python script
	cmd := exec.Command("python", append([]string{scriptPath}, customArgs...)...)
	cmd.Dir = workingDir // Set the working directory for the command

	// Run the command and capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		return 0.0, fmt.Errorf("error running Python script: %v\nOutput:\n%s", err, output)
	}

	// Parse the accuracy value from the script output
	accuracy, err := parseAccuracyFromOutput(string(output))
	if err != nil {
		return 0.0, err
	}

	// Invoke the callback
	if callback != nil {
		err = callback()
		if err != nil {
			return 0.0, err
		}
	}

	fmt.Println("Script Path:", scriptPath)
	fmt.Println("Accuracy from Python script:", accuracy)

	return accuracy, nil
}

func parseAccuracyFromOutput(output string) (float64, error) {
	fmt.Println("Python script output:", output) // Print the output for debugging

	// Split the output into lines and find the line containing the accuracy
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Accuracy:") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				// Try to parse the second part as a float
				accuracy, err := strconv.ParseFloat(parts[1], 64)
				if err != nil {
					return 0.0, fmt.Errorf("error parsing accuracy from script output: %v", err)
				}
				return accuracy, nil
			}
		}
	}

	return 0.0, fmt.Errorf("accuracy not found in script output")
}

// getWorkingDirectory gets the current working directory.
func getWorkingDirectory() string {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
	}
	return wd
}

func dbInteractionCallback() error {
	// ...
	return nil
}

var genericAnalysisCmd = &cobra.Command{
	Use:   "genericanalysis [analysisType] [analysisArgs...]",
	Short: "Perform generic analysis",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		analysisType := args[0]
		analysisArgs := args[1:]

		// Check if the --db-interaction flag is present
		dbInteractionFlag, _ := cmd.Flags().GetBool("db-interaction")

		// Connect to the database
		db, err := dbinteraction.ConnectDB()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer db.Close() // Close the database connection when done

		// Define a callback function for database interaction
		dbInteractionCallback := func() error {
			// Implement your database interaction logic here
			fmt.Println("Performing database interaction...")
			// Example: Connect to the database and execute queries
			return nil
		}

		// Handle the flag as needed
		if dbInteractionFlag {
			err := dbInteractionCallback()
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		// Continue with the generic analysis
		err = genericanalysis.Analyse(db, analysisType, dbInteractionCallback, analysisArgs...)
		if err != nil {
			fmt.Println(err)
		}
	},
}

var connectDBCmd = &cobra.Command{
	Use:   "connectdb",
	Short: "Connect to the PostgreSQL database",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := dbinteraction.ConnectDB()
		if err != nil {
			fmt.Println("Error connecting to the database:", err)
			return
		}
		defer db.Close()

		fmt.Println("Connected to the database successfully!")

		// You can also print additional information about the database connection here
	},
}


// func init() {
// 	rootCmd.AddCommand(analyseDataCmd)
// 	runPythonCmd.Flags().Bool("db-interaction", false, "Perform database interaction")
// 	rootCmd.AddCommand(runPythonCmd)
// 	rootCmd.AddCommand(genericAnalysisCmd)
// 	rootCmd.AddCommand(connectDBCmd)
// 	// Add other commands as needed
// }

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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
