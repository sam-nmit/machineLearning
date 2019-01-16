package main

import (
	"fmt"
	"math"
)

const (
	DATASET_LENGTH = 3
)

var (
	in = [DATASET_LENGTH]float64{
		1.0,
		2.0,
		4.0,
	}

	out = [DATASET_LENGTH]float64{
		2.0,
		4.0,
		5.0,
	}
)

var (
	wG = -7.65 //0.67
	bG = -3.63 //2.37
)

func wags(pats float64) float64 {
	return (pats * wG) + bG
}

func distance(pats, expected float64) float64 {
	return math.Pow(wags(pats)-expected, 2)
}

// dcdw returns the cost value relitive to w
func dcdw(w, b float64) float64 {
	var c float64
	for i := 0; i < DATASET_LENGTH; i++ {
		c += (2 * in[i]) * ((in[i] * w) + b - out[i])
	}
	return c

}

// dcdb returns the cost value relitive to b
func dcdb(w, b float64) float64 {
	var c float64
	for i := 0; i < DATASET_LENGTH; i++ {
		c += 2 * ((in[i] * w) + b - out[i])
	}
	return c
}

func cost() float64 {
	var c float64
	for i := 0; i < DATASET_LENGTH; i++ {
		c += distance(in[i], out[i])
	}
	return c
}

func calculate() {
	fmt.Printf("In\tOut\tExp\tDist\n")

	for i := range in {
		fmt.Printf("%.2f\t%.2f\t%.2f\t%.2f\n", in[i], wags(in[i]), out[i], distance(in[i], out[i]))
	}

	weightMod := dcdw(wG, bG) * 0.01
	biasMod := dcdb(bG, bG) * 0.01

	fmt.Println("Error", cost())
	fmt.Printf("Weigth: %f \t\tModify By: %f\n", wG, weightMod)
	fmt.Printf("bias: %f \t\tModify By: %f\n", bG, biasMod)
	fmt.Println()

	wG -= weightMod
	bG -= biasMod
}

func main() {

	for i := 0; i < 5; i++ {
		calculate()
	}

	fmt.Println("Predictions:")
	for _, v := range []float64{
		8,
		63,
		22,
		77,
	} {
		fmt.Println(v, wags(v))
	}
}
