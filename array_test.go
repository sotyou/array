package array

import (
	"math"
	"testing"
)

func Test_Add(t *testing.T) {
	a := Array{1, 2, 3, 4, 5}
	b := Array{5, 4, 3, 2, 1}
	r := a.Add(b, 1)
	for i, v := range r {
		if v != 6 {
			t.Errorf("Add [1,2,3,4,5] and [5,4,3,2,1] failed. Got %d at %f, expected 6.", i, v)
		}
	}
}

func Test_Mimus(t *testing.T) {
	a := Array{1, 2, 3, 4, 5}
	b := Array{1, 2, 3, 4, 5}
	r := a.Minus(b, 1)
	for i, v := range r {
		if v != 0 {
			t.Errorf("Add [1,2,3,4,5] and [5,4,3,2,1] failed. Got %d at %f, expected 6.", i, v)
		}
	}
}

func Test_NormDiff(t *testing.T) {
	a := Array{6, 7}
	b := Array{3, 3}
	res := a.Minus(b, 1).Norm(2)
	if res != 5 {
		t.Errorf("norm of diff [6, 7, 5, 4] and [3, 3, 2, 0] failed. Got %f, expected 5.", res)
	}
}

func Test_Reduce(t *testing.T) {
	a := Array{1, 2, 3, 4, 5}

	r1 := a.Reduce(Array{}, func(acc interface{}, val float64, index int) interface{} {
		res := acc.(Array)
		res = append(res, val*val)
		return res
	}).(Array)
	r2 := a.Reduce(0.0, func(acc interface{}, val float64, index int) interface{} {
		res := acc.(float64)
		res += val * val
		return res
	}).(float64)
	for i, v := range r1 {
		vT := a[i] * a[i]
		if v != vT {
			t.Errorf("Reduce [1,2,3,4,5] failed. Got %f at %d, expected %f.", v, i, vT)
		}
	}
	if r2 != 55 {
		t.Errorf("Reduce [1,2,3,4,5] failed. Got %f, expected %d.", r2, 55)
	}
}

func Test_Std(t *testing.T) {
	a := Array{1, 2, 3, 4, 5}
	r := a.Std()
	if math.Abs(r-1.414) > 0.01 {
		t.Errorf("Reduce [1,2,3,4,5] failed. Got %f, expected 1.414214", r)
	}
}

func Test_Regularize(t *testing.T) {
	a := Array{1,2,3,4,5}
	r := a.Regularize(4)
	var vS float64
	for i, v := range r {
		switch i {
		case 0:
			vS = -1.3416
		case 1:
			vS = -0.4472
		case 2:
			vS = 0.4472
		case 3:
			vS = 1.3416
		case 4:
			vS = 5
		}
		if math.Abs(v - vS) > 0.01 {
			t.Errorf("Regularize [1,2,3,4,5] failed. Got %f at %d, expected %f", v, i, vS)
		}
	}
}
