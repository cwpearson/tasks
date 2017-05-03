package tasks

type Environment interface {
	Run(in chan interface{}) chan interface{}
	Input() chan interface{}
	Output() chan interface{}
}

type SimpleEnv struct {
	input, output chan interface{}
	program       []Task
}

func NewSimpleEnv() *SimpleEnv {
	return &SimpleEnv{
		input:  make(chan interface{}),
		output: make(chan interface{}),
	}
}
