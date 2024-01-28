
package statsanalysis

//import "fmt"

// Mean calculates the mean of a slice of numbers.
func Mean(numbers []float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}
	return total / float64(len(numbers))
}
