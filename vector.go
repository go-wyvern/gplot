package gplot

import (
	"gonum.org/v1/gonum/mat"
)

type Vector struct {
	*mat.VecDense
}

func Vec(data []float64) Vector {
	return newVec(data)
}

func newVec(data []float64) Vector {
	vec := mat.NewVecDense(len(data), data)
	return Vector{vec}
}

// Gop_Add: func (a vector) + (b vector) vector
func (a Vector) Gop_Add(b Vector) *Vector {
	vec := mat.NewVecDense(a.Len(), nil)
	vec.AddVec(a, b)
	return &Vector{vec}
}

// Gop_Sub: func (a vector) - (b vector) vector
func (a Vector) Gop_Sub(b Vector) *Vector {
	vec := mat.NewVecDense(a.Len(), nil)
	vec.SubVec(a, b)
	return &Vector{vec}
}

// Gop_Mul: func (a vector) * (b vector) vector
func (a Vector) Gop_Mul(b Vector) Vector {
	vec := mat.NewVecDense(a.Len(), nil)
	vec.MulElemVec(a, b)
	return Vector{vec}
}

// Gop_Quo: func (a vector) / (b vector) vector {
func (a Vector) Gop_Quo(b Vector) Vector {
	vec := mat.NewVecDense(a.Len(), nil)
	vec.DivElemVec(a, b)
	return Vector{vec}
}

// Scale a vector
func Scale(a float64, b Vector) Vector {
	vec := mat.NewVecDense(b.Len(), nil)
	vec.ScaleVec(float64(a), b)
	return Vector{vec}
}
