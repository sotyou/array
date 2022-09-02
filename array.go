package array

import (
	"math"
)


type Array []float64

type D2Array [][]float64

type FuncMap func(float64) []interface{}

type FuncReduce func(interface{}, float64, int) interface{}

type FuncFilter func(float64) bool

func (a Array) Add(b Array, rate float64) Array {
	res := make(Array, len(a))
	for i, v := range a {
		res[i] = (v + b[i]) * rate
	}
	return res
}

func (a Array) Minus(b Array, rate float64) Array {
	res := make(Array, len(a))
	for i, v := range a {
		res[i] = (v - b[i]) * rate
	}
	return res
}

func (a Array) Norm(power float64) float64 {
	res := 0.0
	for _, v := range a {
		res += math.Pow(v, power)
	}
	return math.Pow(res, 1.0/power)
}

func (a Array) Std() float64 {
	mean := a.Mean()
	return math.Sqrt(a.Reduce(Array{}, func(acc interface{}, val float64, index int) interface{} {
		res := acc.(Array)
		res = append(res, (val-mean)*(val-mean))
		return res
	}).(Array).Reduce(0.0, func(acc interface{}, val float64, index int) interface{} {
		res := acc.(float64)
		res += val
		return res
	}).(float64) / float64(len(a)))
}

func (a Array) Diff() Array {
	len := len(a)
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

func (a Array) Regularize(length int) Array {
	std := a[:length].Std()
	mean := a[:length].Mean()
	res := a[:length].Reduce(Array{}, func(acc interface{}, val float64, index int) interface{} {
		res := acc.(Array)
		res = append(res, (val-mean)/std)
		return res
	}).(Array)
	return append(res, a[length:]...)
}

func (a Array) Map(f FuncMap) []interface{} {
	res := make([]interface{}, len(a))
	for i, v := range a {
		res[i] = f(v)
	}
	return res
}

func (a Array) Reduce(initV interface{}, f FuncReduce) interface{} {
	res := initV
	for i, v := range a {
		res = f(res, v, i)
	}
	return res
}

func (a Array) Filter(f FuncFilter) (res Array) {
	for _, v := range a {
		if f(v) {
			res = append(res, v)
		}
	}
	return
}
