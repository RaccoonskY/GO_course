package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	printCircleAreaIs(22 * 33)
}

func printCircleAreaIs(circleRad int) {

	circleArea, err := circleAreaCalculate(circleRad) // in Go we have module "errors" for throwing\catching different errors
	if err != nil {                                   //checking if is there any error happened
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Radius of circle: %d sm.\n", circleRad)
	fmt.Printf("Circle area: %f32 sm^2 .\n", circleArea)
}

func circleAreaCalculate(radius int) (float32, error) { //multiple arguments returned! Type 'error' is for ehmmm errors...
	if radius <= 0 {

		return float32(0), errors.New("Radius cannot be negative value!!!")
	}

	return float32(math.Pow(float64(radius), 2)) * math.Pi, nil // nil means NULL, nullptr, NOTHING VALUE U KNOW
}
