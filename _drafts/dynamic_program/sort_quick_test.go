package main

import (
	"fmt"
	"sort"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestQuickSort(t *testing.T) {
	Convey("Init TestQuickSort", t, func() {
		Convey("order elements", func() {
			datas := []int{1, 2, 3, 4, 5}
			So(tSort(datas, quickSort1), ShouldResemble, stdSort(datas))
			So(tSort(datas, quickSort2), ShouldResemble, stdSort(datas))
		})

		Convey("reversed order elements", func() {
			datas := []int{10, 3, 0, -1, -333}
			So(tSort(datas, quickSort1), ShouldResemble, stdSort(datas))
			So(tSort(datas, quickSort2), ShouldResemble, stdSort(datas))
		})

		Convey("duplicate", func() {
			datas := []int{3, 3, 3, 3, 3}
			So(tSort(datas, quickSort1), ShouldResemble, stdSort(datas))
			So(tSort(datas, quickSort2), ShouldResemble, stdSort(datas))
		})

		Convey("empty elements", func() {
			datas := []int{}
			So(tSort(datas, quickSort1), ShouldResemble, stdSort(datas))
			So(tSort(datas, quickSort2), ShouldResemble, stdSort(datas))
		})

		Convey("one elements", func() {
			datas := []int{100}
			So(tSort(datas, quickSort1), ShouldResemble, stdSort(datas))
			So(tSort(datas, quickSort2), ShouldResemble, stdSort(datas))
		})

		Convey("common", func() {
			datas := []int{2, 3, 2, 4, 3, 2, 343, -1, 0}
			So(tSort(datas, quickSort1), ShouldResemble, stdSort(datas))
			So(tSort(datas, quickSort2), ShouldResemble, stdSort(datas))
		})
	})
}

func stdSort(nums []int) []int {
	nNums := make([]int, len(nums))
	copy(nNums, nums)
	sort.IntSlice(nNums).Sort()
	return nNums
}

func tSort(nums []int, doSort func([]int, int, int)) []int {
	nNums := make([]int, len(nums))
	copy(nNums, nums)
	doSort(nNums, 0, len(nNums)-1)
	return nNums
}

//
// privot|------->LastLowerThen privot-----------<LastGreaterThen privot--------|end
//
// 从数组两侧向中间遍历数据。
// 从左面找到第一个大于 privot 的，从右面找到第一个小于等于 privot 的，然后交换两个数值。
// 最后要把 privot 放置到中间的位置。
//
func quickSort2(nums []int, start, end int) {
	if start < 0 {
		panic("start < 0")
		return
	}
	if end >= len(nums) {
		panic("end >= len(nums)")
		return
	}
	if start > end {
		return
	}
	if end-start == 0 {
		return
	}
	if end-start == 1 {
		return
	}
	middle := partition(nums, start, end)
	if middle > start {
		quickSort2(nums, start, middle)
	}
	middle++
	if end > middle {
		quickSort2(nums, middle, end)
	}
}

//
// start end 之间至少两个元素，否则此函数执行无意义
// 调用方要检测异常，保证传参数正确。
// 本函数只要发现异常，应该直接 panic
//
func partition(nums []int, start, end int) int {
	if end-start <= 1 {
		panic(fmt.Sprintf("at least two elem between (start=%d,end=%d, nums=%#v)", start, end, nums))
	}
	privot := start
	left := start + 1
	right := end
	for left < right {
		for left < right && nums[left] <= nums[privot] {
			// start - left 之间是小于等于 privot 的数值
			left++
		}
		for left < right && nums[right] > nums[privot] {
			// right － end之间是大于 privot 的数值
			right--
		}
		if left < right {
			swap(nums, left, right)
		}
	}
	if nums[right] < nums[privot] {
		swap(nums, right, privot)
		return right
	}
	left-- // 此时  left==right ，所以要让left向左退回一位
	//fmt.Printf("\nstart=%d, end=%d, left=%d, right=%d, nums=%#v\n", start, end, left, right, nums)
	swap(nums, left, privot)
	return left
}

//
// [privot]
// [store]                                   [index]
// [2]                                       []         [3 1 4 0 5]
//                                                                                //
//                                                                                //
// [2]                                       [3]        [1 4 0 5]
//            <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<[index]                         cause: 3 > privot(2)
// [2]       [3]                                        [1 4 0 5]
//                                                                                 //
//                                                                                 //
// [2]       [3]                             [1]        [4 0 5]
//           [++store]<--------swap--------->[index]                         cause: 1 <= privot(2)
// [2]       [1]                             [3]        [4 0 5]
// [2]       [1]         [3]                            [4 0 5]
//                                                                                 //
//                                                                                 //
// [2]       [1]         [3]                 [4]        [0 5]
//                                  <<<<<<<<<[index]                         cause: 4 > privot(2)
// [2]       [1]         [3         4]                  [0 5]
//                                                                                 //
//                                                                                 //
// [2]       [1]         [3]       [4]       [0]        [5]
//                       [++store]<---swap-->[index]                         cause: 0 <= privot(2)
// [2]       [1]         [0]       [4]       [3]        [5]
// [2]       [1]         [0]       [4 3]     []         [5]
//                                                                                 //
//                                                                                 //
// [2]       [1]         [0]       [4 3]     [5]        []
//                                      <<<<<[index]                         cause: 4 > privot(2)
// [2]       [1]         [0]       [4 3 5]              []
//                                                                                 //
//                                                                                 //
// [privot]<---swap----->[store]                                             cause: end
// [0]       [1]         [2]       [4 3 5]   []        []
//
// 从左向右依次处理数据。
// 主要分为这几个区域：privot 低数区 高数区 未处理区域。
// store 指向低数区最后一个数据，index 指向未处理区域第一个数据。
// 从 index 找到的每个`小于等于` privot 的数据都放到 store 。
// 最后把 privot 交换到中间位置，所以让 privot 与 store 交换。
//
func quickSort1(nums []int, start, end int) {
	if len(nums) == 0 {
		return
	}
	if start >= end {
		return
	}
	privot := start // privot 的选择直接影响快速排序的效率
	storeIndex := privot
	for index := start + 1; index <= end; index++ {
		if nums[index] <= nums[privot] {
			storeIndex++
			swap(nums, index, storeIndex)
		}
	}
	swap(nums, privot, storeIndex)
	quickSort1(nums, start, storeIndex-1)
	quickSort1(nums, storeIndex+1, end)
	return
}

func swap(nums []int, a, b int) {
	tmp := nums[a]
	nums[a] = nums[b]
	nums[b] = tmp
}
