package main

import "fmt"

/*

<?php
function binarySearch($array, $val) {
    $count = count($array);
    $low = 0;
    $high = $count - 1;
    while ($low <= $high) {
        $mid = intval(($low + $high) / 2);
        if ($array[$mid] == $val) {
            return $mid;
        }
        if ($array[$mid] < $val) {
            $low = $mid + 1;
        } else {
            $high = $mid - 1;
        }
    }
    return false;
}

*/

// BinarySearch 二分法查找
//切片s是升序的
//k为待查找的整数
//如果查到有就返回对应角标,
//没有就返回-1
func BinarySearch(s []int, k int) int {
	lo, hi := 0, len(s)-1
	for lo <= hi {
		m := (lo + hi) >> 1 // 位运算获取数组元素，这样就不需要新生成数组了
		if s[m] < k {
			lo = m + 1
		} else if s[m] > k {
			hi = m - 1
		} else {
			return m
		}
	}
	return -1
}

// Search 二分法查找
func Search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l+(r-l) >> 1 // 位运算获取数组元素，这样就不需要新生成数组了
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid -1
		}
	}
	return -1
}

func main()  {

	back1 := Search([]int{0, 1, 6, 7, 9}, 7)
	back2 := Search([]int{0, 1, 6, 7, 9}, 3)

	fmt.Println(back1, back2) // 3, -1

}
