package main

import (
	"fmt"
	"math"
)

const epochs_init int = 0               // initial epoch iteration
const step_dir_init float64 = 1.0       // initial step dir
const x_init float64 = 151.0            // initial y guess
const epochs float64 = 100              // number of iterations for algorithm
const alpha float64 = 0.01              // learning rate
const y_min_threshold float64 = 0.00001 // early y_min threshold

func foo() {
	fmt.Printf("foo() enter.\n")
}

// NA
// - need to nest within recursive function?
/*
func f(x float64) float64 {
	fmt.Printf("f(x) enter.\n")

	// simplify so we can derive gradient later
	// 10 * (x**2 - (4*x) - 4) + 7
	// (10 * x**2) - (40 * x) - 40 + 7
	// (10 * x**2) - (40 *64 x) - 47
	// NA
	// - static typing will pick up issues with math.Pow typed as float64
	var xSquared = math.Pow(x, 2.0)
	return 10.0*xSquared + 7.0
}
*/

func dumpEpoch(y_last float64, y_new float64, x_new float64, epoch int) {
	fmt.Printf("f(x) enter.\n")

	var xStepStr = func(y_last float64, y_new float64) string {
		if y_new < y_last {
			return "left"
		} else {
			return "right"
		}
	}

	fmt.Printf(
		"epoch: %d, y_last: %f, y_new: %f, y_new < y_last: %t, step_dir: %s, x_step: %f",
		epoch, y_last, y_new, (y_new < y_last), xStepStr(y_last, y_new), x_new)

}

func xStep(x float64) float64 {
	var x_new float64 = x - (alpha * gradient(x))
	return x_new
}

func yMinReached(y_last float64, y_new float64, x_new float64, epoch int) bool {
	if math.Abs(y_last-y_new) <= y_min_threshold {
		fmt.Printf(
			"*** y_min found early epoch %d for: y_new: %f, x_new: %f",
			epoch, y_new, x_new)
		return true
	} else {
		return false
	}
}

// solve, implement
// NA.
// - note *epoch base
// func findMin(f func(float64) float64, x float64, step_dir int, *epoch int) (int, float64,float64) {
func findMin(x float64, step_dir float64, epoch int) (int, float64, float64) {
	/*
	   - find y minimum by iterating using gradient descent:
	     - y function
	     - hyperparamters (y_init, alpha)
	     - loss
	     - y gradient
	*/
	var f = func(x float64) float64 {
		fmt.Printf("f(x) enter.\n")

		// simplify so we can derive gradient later
		// 10 * (x**2 - (4*x) - 4) + 7
		// (10 * x**2) - (40 * x) - 40 + 7
		// (10 * x**2) - (40 *64 x) - 47
		// NA
		// - static typing will pick up issues with math.Pow typed as float64
		var xSquared = math.Pow(x, 2.0)
		return 10.0*xSquared + 7.0
	}

	var epoch_new int = epoch + 1
	var y_last float64 = f(x)
	var x_new float64 = step_dir * xStep(x)
	var y_new float64 = f(x_new)
	dumpEpoch(y_last, y_new, x_new, epoch_new)

	/*
		if yMinReached(y_last, y_new, x_new, epoch_new) {
			return epoch_new, x_new, y_new
		} else if y_new < y_last {
			findMin(x_new, 1.0, epoch_new)
		} else {
			findMin(x, -1.0, epoch_new)
		}
	*/
	if yMinReached(y_last, y_new, x_new, epoch_new) {
		fmt.Printf("***yMinReached")
	} else if y_new < y_last {
		findMin(x_new, 1.0, epoch_new)
	} else {
		findMin(x, -1.0, epoch_new)
	}
	return epoch_new, x_new, y_new
}

func gradient(x float64) float64 {
	/*
	   - original polynomial expression
	   (10 * x**2) - (40 * x) - 47
	   - find derivative or gradient
	   (20 * x) - 40 - 0
	*/
	return (20.0 * x) - 40.0 - 0.0
}

func main() {
	fmt.Printf("main() enter.\n")
	// foo()
	// findMin(f, x_init, step_dir_init, epochs_init)
	findMin(x_init, step_dir_init, epochs_init)
}

// NA. Python to Golang notes
// - gomod, GOPATH and other set up for Makefile
// - rigid and fast compiler!
// - specific types from inferred int to exact float32
// - need to use Math.pow, Math.abs
// - verbose function typing now
// - different fmt.Printf
// - nested functions need var assigned anonymous inner function
// - found many static typing issues, global vars and return types
