package learner

const (
	IrusSetosa = iota
	IrusVersicolour
	IrusVirginica
)

type Irus struct {
	SeptalLength float64
	SeptalWidth  float64
	PetalLength  float64
	PetalWidth   float64
	Breed        float64
}
