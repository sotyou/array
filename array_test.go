package array

import "testing"

func Test_Add(t *testing.T) {
	a := Array{1, 2, 3, 4, 5}
	b := Array{5, 4, 3, 2, 1}
	r := a.Add(b)
	for i, v := range r {
		if v != 6 {
			t.Errorf("Add [1,2,3,4,5] and [5,4,3,2,1] failed. Got %d at %f, expected 6.", i, v)
		}
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
