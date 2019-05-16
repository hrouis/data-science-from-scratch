package data_science

import (
	"fmt"
	"testing"
)

func TestMatMul(t *testing.T) {

	matA := new(Matrix)
	matB := new(Matrix)
	matA.columnDim = 2
	matA.rowDim = 2
	matA.values = []Vector{{2,6}, {4,8}}

	matB.columnDim = 2
	matB.rowDim = 2
	matB.values = []Vector{{1,2}, {0,1}}

	res := MatMul(*matB, *matA)
	fmt.Println(res)
}