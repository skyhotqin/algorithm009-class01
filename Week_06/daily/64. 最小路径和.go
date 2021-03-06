package main

import (
	"fmt"
)

func minPathSum(grid [][]int) int {
	for i := len(grid) - 1; i >= 0; i-- {
		for j := len(grid[0]) - 1; j >= 0; j-- {
			if i == len(grid)-1 && j != len(grid[0])-1 {
				grid[i][j] = grid[i][j] + grid[i][j+1]

			} else if j == len(grid[0])-1 && i != len(grid)-1 {
				grid[i][j] = grid[i][j] + grid[i+1][j]

			} else if j != len(grid[0])-1 && i != len(grid)-1 {
				x := grid[i+1][j]
				y := grid[i][j+1]
				if x < y {
					grid[i][j] = grid[i][j] + x
				} else {
					grid[i][j] = grid[i][j] + y
				}
			}
		}
	}
	return grid[0][0]
}

//
//func dfs(nums [][]int, x, y int, m int, n int, ans int) int {
//	if x >= n && y >= m {
//		ans += nums[x][y]
//		return ans
//	}
//	xAns, yAns := math.MaxInt64, math.MaxInt64
//	if x < n {
//		xAns = dfs(nums, x+1, y, m, n, ans+nums[x][y])
//	}
//	if y < m {
//		yAns = dfs(nums, x, y+1, m, n, ans+nums[x][y])
//	}
//	if xAns < yAns {
//		return xAns
//	}
//	return yAns
//}
//
//func minPathSum(grid [][]int) int {
//	m := len(grid)
//	if m == 0 {
//		return 0
//	}
//	n := len(grid[0])
//	ans := dfs(grid, 0, 0, n-1, m-1, 0)
//	return ans
//}

func main() {
	testCases := [][][]int{
		{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
			//{1, 2},
			//{1, 2},
		},
	}
	for _, c := range testCases {
		r := minPathSum(c)
		fmt.Println("result:", r)
	}
}
