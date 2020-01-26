package main

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

//
// dynamic program best-time-to-buy-and-sell-stock
//
func TestBestTimeToBuyAndSellStock(t *testing.T) {
	Convey("Init TestBestTimeToBuyAndSellStock", t, func() {
		testArr := []int{7, 1, 5, 3, 6, 4}
		Convey("Call maxProfit func", func() {
			ret := maxProfit(testArr)
			fmt.Printf("ret= %d\n", ret)

			Convey("result should be: true", func() {
				So(5, ShouldEqual, ret)
			})
		})
	})
}

func maxProfit(prices []int) int {
	maxNum := len(prices)
	if 0 == maxNum {
		return 0
	}
	minVal := prices[0]
	maxVal := minVal
	for _, curVal := range prices {
		if curVal < minVal {
			minVal = curVal
			maxVal = minVal
		} else if curVal > maxVal {
			maxVal = curVal
		}
	}
	return maxVal - minVal
}

//
// dynamic program https://leetcode.com/problems/minimum-path-sum
//
func TestMinimumPathSum(t *testing.T) {
	Convey("Init TestMinimumPathSum", t, func() {
		grid := [][]int{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
		}
		Convey("Call minPathSum func", func() {
			ret := minPathSum(grid)
			fmt.Printf("ret= %d\n", ret)

			Convey("result should be: true", func() {
				So(7, ShouldEqual, ret)
			})
		})
	})
}

func minPathSum(grid [][]int) int {
	maxRow := len(grid) - 1
	if maxRow <= 0 {
		return 0
	}
	maxColumn := len(grid[0]) - 1
	if maxColumn <= 0 {
		return 0
	}
	//fmt.Printf("grid[%d][%d] = %#v\n", maxRow, maxColumn, grid)
	dist := make([]int, maxColumn+1)
	for row := maxRow; row >= 0; row-- {
		for column := maxColumn; column >= 0; column-- {
			if row == maxRow && column == maxColumn {
				dist[column] = grid[row][column]
			} else if row == maxRow {
				// right
				dist[column] = grid[row][column] + dist[column+1]
			} else if column == maxColumn {
				// down
				dist[column] = grid[row][column] + dist[column]
			} else {
				/// TODO minus
				dist[column] = grid[row][column] + min(dist[column], dist[column+1])
			}
			//fmt.Printf("dist[%d] = %#v\n", column, dist)
		}
	}

	return dist[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
