package main

import (
	"fmt"
)

func main(){
	dataArray := []float64{21, 23, 24, 34, 45, 54}
	num := average(dataArray...)

	fmt.Println(num)
}

func average(args ...float64) float64 {
	var sum float64

	for _, v := range args {
		sum += v
	}

	return sum / float64(len(args))
} 