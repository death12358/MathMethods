package main

import (
	"fmt"

	"github.com/gonum/stat"
)

func main() {
	obs := []float64{200, 203, 204}
	exp := []float64{200, 202, 202}

	fmt.Println(ChiSquareTest(obs, exp))
}

func ChiSquareTest(obs, exp []float64) bool {
	if !CheckExpTimes_MoreThan5(exp) {
		return false
	}
	chiValue := stat.ChiSquare(obs, exp)
	df := len(obs) - 1
	if df > 30 {
		fmt.Printf("自由度>30未實裝")
		return false
	}
	fmt.Printf("Chi-square test df: %d, p-value: %f, critical value: %f", df, chiValue, ChiSquareCriticalValue_confidence95[df])
	return chiValue < ChiSquareCriticalValue_confidence95[df]
}

func CheckExpTimes_MoreThan5(exp []float64) bool {
	for i := 0; i < len(exp); i++ {
		if exp[i] < 5 {
			fmt.Println("期望次數小於5次, 效果不佳")
			return false
		}
	}
	return true
}

var ChiSquareCriticalValue_confidence95 = []float64{0, 0.004, 0.103, 0.352, 0.711, 1.145, 1.635, 2.167, 2.733, 3.325, 3.94, 4.575, 5.226, 5.892, 6.571, 7.261, 7.962, 8.672, 9.39, 10.117, 10.851, 11.591, 12.338, 13.091, 13.848, 14.611, 15.379, 16.151, 16.928, 17.708, 18.493}
