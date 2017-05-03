package main

import (
	"github.com/cwpearson/tasks"
)

func main() {

	env := tasks.NewSimpleEnv()

	data := []float64{4, 2, 0, 6}

	inputs := FloatSliceReader(data)

	a, b := Split(inputs)
	c := AddOne(a)
	d := Add(b, d)

}
