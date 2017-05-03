package tasks

type Task interface {
}

type Input interface {
}

type Output interface {
}

type FloatTask struct {
	inputs  []Input
	outputs []Output

	env Environment

	op func(in []float64) []float64
}
