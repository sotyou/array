package array

import (
	"math"
)

type Array []float64

func (a Array) Add(b Array) Array {
	res := make(Array, len(a))
	for i, v := range a {
		res[i] = v + b[i]
	}
	return res
}

func (a Array) Array() Array {
	return []float64(a)
}

func (a Array) Diff() Array {
	len := len(a.Array())
	res := make([]float64, len)
	res[0] = math.NaN()
	for i := 1; i < len; i++ {
		res[i] = a[i] - a[i-1]
	}
	return res
}

func (a Array) Mean() (res float64) {
	for _, v := range a {
		res += v
	}
	return res / float64(len(a))
}

func (a Array) Map(f func(v float64) float64) Array {
	res := make(Array, len(a))
	for i, v := range a {
		res[i] = f(v)
	}
	return res
}

func (a Array) Reduce(initV interface{}, f func(acc interface{}, val float64, index int) interface{}) interface{} {
	res := initV
	for i, v := range a {
		res = f(res, v, i)
	}
	return res
}
