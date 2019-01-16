package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"./learner"
)

const (
	DataFileName = "iris.data"
)

var (
	irusNameToValue = map[string]float64{
		"Iris-setosa":     learner.IrusSetosa,
		"Iris-versicolor": learner.IrusVersicolour,
		"Iris-virginica":  learner.IrusVirginica,
	}
)

var (
	l = learner.New()
)

func parsef(inp string) float64 {
	v, err := strconv.ParseFloat(inp, 32)
	if err != nil {
		panic(fmt.Errorf("Failed to parse %s to float. %s", inp, err))
	}
	return v
}

func loadData() {
	f, err := os.Open(DataFileName)
	if err != nil {
		panic(fmt.Errorf("Failed to open file %s - %s", DataFileName, err))
	}

	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()

	for _, data := range records {
		if err != nil {
			panic(fmt.Errorf("CSV error - %s", err))
		}
		fmt.Println(data)

		l.Add(learner.Irus{
			SeptalLength: parsef(data[0]),
			SeptalWidth:  parsef(data[1]),
			PetalLength:  parsef(data[2]),
			PetalWidth:   parsef(data[3]),
			Breed:        irusNameToValue[data[4]],
		})
	}
}

func main() {
	loadData()

	for i := 0; i < 6; i++ {
		fmt.Println(l.Train())
	}

	fmt.Println(l.Predict(learner.Irus{
		5.1,
		3.5,
		1.4,
		0.2,
		0,
	}))
}
