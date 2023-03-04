package main

import (
	"fmt"
	"math"
)

func MeanMedian(arrayInput []float64) (float64, float64) {
	// your code here
	// menghitung mean 
	var sum float64
	for _, value := range arrayInput {
		sum += value
	} 
	mean := sum / float64(len(arrayInput))
	// menghitung median
	var median float64
	if len(arrayInput)%2 == 0 {
		// array memiliki jumlah elemen genap
		median = (arrayInput[len(arrayInput)/2-1] + arrayInput[len(arrayInput)/2]) / 2
	} else {
		// array memiliki jumlah elemen ganjil
		median = arrayInput[(len(arrayInput)-1)/2]
	}
	// pembulatan mean
	mean = math.Round(mean*10) / 10
	return mean, median
}

func main() {
	fmt.Println(MeanMedian([]float64{1, 2, 3, 4}))          // 2.5 2.5
	fmt.Println(MeanMedian([]float64{1, 2, 3, 4, 5}))       // 3 3
	fmt.Println(MeanMedian([]float64{7, 8, 9, 13, 15}))     // 10.4 9
	fmt.Println(MeanMedian([]float64{10, 20, 30, 40, 50}))  // 30 30
	fmt.Println(MeanMedian([]float64{15, 20, 30, 60, 120})) // 49 30
}
