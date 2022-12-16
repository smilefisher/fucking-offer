package main

import "fmt"

func main() {
	fmt.Println("log")
	a := NewArray([][]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {3, 3, 3, 3}, {4, 4, 4, 4}})
	fmt.Println(a.metrics)
	fmt.Println(a.preSum)
	//[1 1 1 1]
	//[2 2 2 2]
	//[3 3 3 3]
	//[4 4 4 4]

	//[0 0  0  0  0]
	//[0 1  2  3  4]   （row2+1， col2+1）- （row2+1，col1） - （row1 ，col2+1）+ （row1， col1）
	//[0 3  6  9  12]  (4,3) - (4, 1) - (0, 2) + (0,0)
	//[0 6  12 18 24]
	//[0 10 20 30 40]
	fmt.Println("result:", a.sumRegion(1, 1, 3, 2))

}

type Array struct {
	preSum  [][]int
	metrics [][]int
}

func NewArray(metrics [][]int) *Array {
	row := len(metrics)
	col := len(metrics[0])

	preSum := make([][]int, row+1)
	for i := 0; i < col+1; i++ {
		preSum[i] = make([]int, col+1, col+1)
	}

	for i := 1; i < row+1; i++ {
		for j := 1; j < col+1; j++ {
			preSum[i][j] = preSum[i][j-1] + preSum[i-1][j] + metrics[i-1][j-1] - preSum[i-1][j-1]
		}
	}

	return &Array{
		metrics: metrics,
		preSum:  preSum,
	}
}

func (a *Array) sumRegion(row1, col1, row2, col2 int) int {
	return a.preSum[row2+1][col2+1] - a.preSum[row2+1][col1] - a.preSum[row1][col2+1] + a.preSum[row1][col1]
}
