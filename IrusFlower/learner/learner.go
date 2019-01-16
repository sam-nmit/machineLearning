package learner

import (
	"fmt"
	"math/rand"
)

type Learner struct {
	LearnRate float64

	data []Irus

	slw float64 //SeptalLength length
	sww float64 //SeptalWidth weight
	plw float64 //PetalLength weight
	pww float64 //PetalWidth weight

	bias float64
}

//New Learner with initilized values
func New() *Learner {
	return &Learner{
		data:      make([]Irus, 1),
		LearnRate: 0.01,

		slw:  rand.Float64(),
		sww:  rand.Float64(),
		plw:  rand.Float64(),
		pww:  rand.Float64(),
		bias: rand.Float64(),
	}
}

//Add a new irus dataset into the learner
func (l *Learner) Add(i Irus) {
	l.data = append(l.data, i)
}

func (l *Learner) Predict(i Irus) float64 {
	v := l.getSpecies(i.SeptalLength, i.SeptalWidth, i.PetalLength, i.PetalWidth)
	return v
}

func (l *Learner) getSpecies(sepLen, sepWid, petLen, petWid float64) float64 {

	val :=
		(sepLen*l.slw +
			sepWid*l.sww +
			petLen*l.plw +
			petWid*l.pww) + l.bias

	return val
}

//dcrBase is base equasion for petal derivitives
func (l *Learner) dcrBase(fieldWeight float64) float64 {
	cost := 0.0
	for _, data := range l.data {
		cost -= (2 * fieldWeight) *
			((data.SeptalLength*l.slw +
				data.SeptalWidth*l.sww +
				data.PetalLength*l.plw +
				data.PetalWidth*l.pww) + l.bias - data.Breed)
	}
	return cost
}

func (l *Learner) getFlowerDirivitives() (sl, sw, pl, pw float64) {
	for _, data := range l.data {
		base := ((data.SeptalLength*l.slw +
			data.SeptalWidth*l.sww +
			data.PetalLength*l.plw +
			data.PetalWidth*l.pww) + l.bias - data.Breed)

		sl += (l.slw * 2) * base
		sw += (l.sww * 2) * base
		pl += (l.plw * 2) * base
		pw += (l.pww * 2) * base
	}

	return
}

func (l *Learner) getBiasDirivitive() float64 {
	cost := 0.0
	for _, data := range l.data {
		cost += ((data.SeptalLength*l.slw +
			data.SeptalWidth*l.sww +
			data.PetalLength*l.plw +
			data.PetalWidth*l.pww) + l.bias - data.Breed)
	}
	return cost
}

//Train processes the infomation added and returns the current distance.
func (l *Learner) Train() float64 {
	sl, sw, pl, pw := l.getFlowerDirivitives()
	b := l.getBiasDirivitive()

	l.slw -= l.LearnRate * sl
	l.sww -= l.LearnRate * sw
	l.plw -= l.LearnRate * pl
	l.pww -= l.LearnRate * pw
	l.bias -= l.LearnRate * b
	fmt.Println("dbg", sl, sw, pl, pw, b)

	return l.slw + l.sww + l.plw + l.pww
}
