
package pythonintegration

import (
	"fmt"
	"os/exec"
)

// RunPythonScript runs a Python script from Golang.
func RunPythonScript(scriptPath string) error {
	cmd := exec.Command("python", scriptPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error running Python script: %v\nOutput:\n%s", err, output)
	}
	fmt.Println("Python script executed successfully.")
	return nil
}
