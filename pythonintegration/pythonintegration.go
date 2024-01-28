// pythonintegration.go

package pythonintegration

import (
	"fmt"
	"os/exec"
)

// DBInteractionCallback is a callback function for handling database interactions.
type DBInteractionCallback func() error

// RunPythonScript runs a Python script from Golang.
func RunPythonScript(scriptPath string, dbCallback DBInteractionCallback, args ...string) error {
	cmd := exec.Command("python", append([]string{scriptPath}, args...)...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error running Python script: %v\nOutput:\n%s", err, output)
	}

	// Invoke the database interaction callback
	if dbCallback != nil {
		if err := dbCallback(); err != nil {
			return fmt.Errorf("error performing database interaction: %v", err)
		}
	}

	fmt.Println("Python script executed successfully.")
	return nil
}
