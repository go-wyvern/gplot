package gplot

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

var (
	// const pi 3.1415926
	Pi = math.Pi

	Sin__0 = math.Sin
	Cos__0 = math.Cos
	Tan__0 = math.Tan
	Log__0 = math.Log
)

func Sin__1(v Vector) Vector {
	vec := mat.NewVecDense(v.Len(), nil)
	for i := 0; i < v.Len(); i++ {
		vec.SetVec(i, math.Sin(v.AtVec(i)))
	}
	return Vector{vec}
}

func Cos__1(v Vector) Vector {
	vec := mat.NewVecDense(v.Len(), nil)
	for i := 0; i < v.Len(); i++ {
		vec.SetVec(i, math.Cos(v.AtVec(i)))
	}
	return Vector{vec}
}

func Tan__1(v Vector) Vector {
	vec := mat.NewVecDense(v.Len(), nil)
	for i := 0; i < v.Len(); i++ {
		vec.SetVec(i, math.Tan(v.AtVec(i)))
	}
	return Vector{vec}
}

func Log__1(v Vector) Vector {
	vec := mat.NewVecDense(v.Len(), nil)
	for i := 0; i < v.Len(); i++ {
		vec.SetVec(i, math.Log(v.AtVec(i)))
	}
	return Vector{vec}
}
