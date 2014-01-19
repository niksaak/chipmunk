package vect

import ( 
	"math"
	"testing"
)

type vectTest struct {
	in1, in2 V
	out      V
}

var addTests = []vectTest{
	{V{0, 0}, V{0, 0}, V{0, 0}},
	{V{0, 1}, V{0, 0}, V{0, 1}},
	{V{1, 0}, V{0, 0}, V{1, 0}},
	{V{1, 2}, V{0, 0}, V{1, 2}},
	{V{0, 0}, V{0, 1}, V{0, 1}},
	{V{0, 0}, V{1, 0}, V{1, 0}},
	{V{0, 0}, V{1, 2}, V{1, 2}},
	{V{2, 4}, V{1, 3}, V{3, 7}},
	{V{3, 1}, V{4, 2}, V{7, 3}},
	{V{2, 4}, V{2, 4}, V{4, 8}},
	{V{5, 5}, V{2, 2}, V{7, 7}},
}

func TestAdd(t *testing.T) {
	for _, at := range addTests {
		v := Add(at.in1, at.in2)
		if !Equal(at.out, v) {
			t.Errorf("Add(%v, %v) = %v, want %v.",
				at.in1, at.in2, v, at.out)
		}
	}
}

var minTests = []vectTest{
	{V{0, 0}, V{0, 0}, V{0, 0}},
	{V{1, 2}, V{9, 9}, V{1, 2}},
	{V{9, 9}, V{1, 2}, V{1, 2}},
	{V{5, 2}, V{1, 4}, V{1, 2}},
	{V{9, 6}, V{7, 8}, V{7, 6}},
}

func TestMin(t *testing.T) {
	for _, at := range minTests {
		v := Min(at.in1, at.in2)
		if !Equal(at.out, v) {
			t.Errorf("Min(%v, %v) = %v, want %v.",
				at.in1, at.in2, v, at.out)
		}
	}
}

var maxTests = []vectTest{
	{V{0, 0}, V{0, 0}, V{0, 0}},
	{V{1, 2}, V{9, 9}, V{9, 9}},
	{V{9, 9}, V{1, 2}, V{9, 9}},
	{V{5, 2}, V{1, 4}, V{5, 4}},
	{V{9, 6}, V{7, 8}, V{9, 8}},
}

func TestMax(t *testing.T) {
	for _, at := range maxTests {
		v := Max(at.in1, at.in2)
		if !Equal(at.out, v) {
			t.Errorf("Max(%v, %v) = %v, want %v.",
				at.in1, at.in2, v, at.out)
		}
	}
}

type distTest struct {
	in1, in2 V
	out      float64
}


var distTests = []distTest{
	{V{0, 0}, V{0, 0}, 0},
	{V{0, 2}, V{0, 0}, 2},
	{V{2, 0}, V{0, 0}, 2},
	{V{0, 0}, V{4, 0}, 4},
	{V{0, 0}, V{0, 4}, 4},
	{V{1, 1}, V{0, 0}, math.Sqrt(2)},
	{V{1, 1}, V{2, 2}, math.Sqrt(2)},
}

func TestDist(t *testing.T) {
	for _, at := range distTests {
		v := Distance(at.in1, at.in2)
		if at.out != v {
			t.Errorf("Dist(%v, %v) = %v, want %v.", at.in1, at.in2, v, at.out)
		}
	}
}
