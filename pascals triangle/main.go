package main

func generate(numRows int) [][]int {
	triangle := [][]int{{1}, {1, 1}}
	if numRows <= 2 {
		return triangle[:numRows]
	}

	for i := 2; i != numRows; i++ {
		triangle = append(triangle, make([]int, i+1))
		triangle[i][0] = 1
		triangle[i][i] = 1
		for j := 1; j != i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}
	return triangle
}
