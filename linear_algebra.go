package data_science

import "math"

type Vector []float64

func (vec Vector) Len() int {
	return len(vec)
}

func (vec Vector) Less(i, j int) bool {
	return vec[i] < vec[j]
}

func (vec Vector) Swap(i, j int) {
	tmp := vec[i]
	vec[i] = vec[j]
	vec [j] = tmp
}

// Matrix structure definition.
type Matrix struct {
	values    []Vector
	rowDim    int
	columnDim int
}

type EntryFn = func(i int, j int) float64

// Matrix constructor.
func NewMatrix(rowDim int, columnDim int) *Matrix {
	matrix := new(Matrix)
	matrix.columnDim = columnDim
	matrix.rowDim = rowDim
	matrix.values = make([]Vector, rowDim, columnDim)
	return matrix
}

// Return row at index.
func (matrix Matrix) getRow(index int) Vector {
	return matrix.values[index]
}

// Return column at index.
func (matrix Matrix) getColumn(index int) Vector {
	res := make(Vector, 0)
	for _, vec := range matrix.values {
		res = append(res, vec[index])
	}
	return res
}

// Sum all elements of vector
func Sum(vector Vector) float64 {
	res := float64(0)
	for _, val := range vector {
		res += val
	}
	return res
}

// Add two vectors.
func VectorAdd(v Vector, w Vector) Vector {

	if len(v) != len(w) {
		panic("Error vectors does not have the same length")
	}

	res := make(Vector, 0)
	for i, val := range v {
		res = append(res, val+w[i])
	}
	return res
}

//  Substract two vectors.
func VectorSubstract(v Vector, w Vector) Vector {

	if len(v) != len(w) {
		panic("Error vectors does not have the same length")
	}

	res := make(Vector, 0)
	for i, val := range v {
		res = append(res, val-w[i])
	}
	return res
}

// Sum all elements in vectors,
func VectorSum(vectors ... Vector) Vector {
	res := vectors[0]
	for _, vec := range vectors[1:] {
		res = VectorAdd(res, vec)
	}
	return res
}

// Multiply vector by a scalar.
func ScalarMultiply(scalar float64, vector Vector) Vector {
	res := make(Vector, 0)
	for i, val := range vector {
		res[i] = val * scalar
	}
	return res
}

// Mean of n vectors
func VectorMean(vectors ... Vector) Vector {
	n := len(vectors)
	return ScalarMultiply(float64(1/n), VectorSum(vectors...))
}

/* Dot product of two vectors
   V * W = v1 * w1 + v2 * w2 ....+ vn * wn
 */
func Dot(v Vector, w Vector) float64 {
	var result float64

	for i, val := range v {
		result += val * w[i]
	}
	return result
}

func SumOfSquares(vector Vector) float64 {
	return Dot(vector, vector)
}

func Magnitude(vector Vector) float64 {
	return math.Sqrt(SumOfSquares(vector))
}

/* Squared distance between two vectors */
func SquaredDistance(v Vector, w Vector) float64 {
	return SumOfSquares(VectorSubstract(v, w))
}

/* Distance between two vectors */
func Distance(v Vector, w Vector) float64 {
	return math.Sqrt(SquaredDistance(v, w))
}

func MakeMatrix(columnDim int, rowDim int, filler EntryFn) *Matrix {
	matrix := new(Matrix)
	matrix.columnDim = columnDim
	matrix.rowDim = rowDim
	matrix.values = make([]Vector, rowDim, columnDim)
	for i := 0; i < rowDim; i++ {
		vector := make(Vector, columnDim)
		for j := 0; j < columnDim; j++ {
			vector[j] = filler(i, j)
		}
		matrix.values[i] = vector
	}
	return matrix
}

func isDiagonal(i int, j int) float64 {
	if i == j {
		return 1
	}
	return 0
}

/*
    Multiply Matrices.
 */
func MatMul(matA Matrix, matB Matrix) *Matrix {
	rowA := matA.rowDim
	colA := matA.columnDim
	rowB := matB.rowDim
	colB := matB.columnDim

	if colA != rowB {
		panic(" error can not multiply matrices")
	}
	var filler = func(i int, j int) float64 {
		return Dot(matA.getRow(i), matB.getColumn(j))
	}
	return MakeMatrix(colB, rowA, filler)
}
