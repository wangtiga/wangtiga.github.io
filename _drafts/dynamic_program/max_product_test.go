package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

//
// https://leetcode-cn.com/problems/maximum-product-subarray/
//
func TestMaxProductSubarray(t *testing.T) {
	Convey("Init TestMaxProductSubarray", t, func() {
		Convey("common", func() {
			So(
				maxProductSubarray([]int{2, 3, 2, 4}),
				ShouldEqual,
				3*16,
			)
		})
		Convey("single minus", func() {
			So(
				maxProductSubarray([]int{-2}),
				ShouldEqual,
				-2,
			)
			So(
				maxProductSubarray([]int{2, 3, -2, 4}),
				ShouldEqual,
				6,
			)
		})
		Convey("double minus", func() {
			So(
				maxProductSubarray([]int{2, 3, -2, 4, -3}),
				ShouldEqual,
				16*9,
			)
		})
		Convey("with zero", func() {
			So(
				maxProductSubarray([]int{2, 9, 0, -2, 4, -3}),
				ShouldEqual,
				24,
			)
		})
	})
}

func maxProductSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	min := 1
	max := 1
	var ret *int = nil
	for _, item := range nums {
		minN := Min(Min(item*min, item*max), item)
		maxN := Max(Max(item*min, item*max), item)
		min = minN
		max = maxN
		if nil == ret {
			nVal := max
			ret = &nVal
		} else {
			// 关键：如果没有 ret ，那么含有奇数个负数时，前面的部分大乘积可能会丢失
			*ret = Max(*ret, maxN)
		}
	}

	return *ret
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
