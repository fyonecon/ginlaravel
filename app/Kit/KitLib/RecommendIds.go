package KitLib
// 生成权重id集合

import (
	"fmt"
	"sort"
)

type Recommend struct {
	//Array []int
}

// RecommendArraysInt 分词ID推荐
// 这个算法用在按分词在数据库中查询的id集合，最后输出带有权重的id数组。类似的思路是Elasticsearch中的中文检索算法。
// 多组int数组
// 重复int组成优先级第一的，不重复int各数组跟在后面
// 调用：rec := KitLib.Recommend{}
//      rec.RecommendArraysId(arrays)
func (kit *Recommend) RecommendArraysInt(arrays map[string][]int) []int {

	// 合并所有数组为一个数组，只输出value
	var newArray []int
	for _, v := range arrays {
		newArray = append(newArray, v...)
	}

	rec := Recommend{}
	alikeArray := rec.OrderHasArray(newArray) // 按出现次数排序

	return alikeArray
}

// OrderHasArray 按元素出现次数排名数组，出现次数多的排前面
func (kit *Recommend) OrderHasArray(array []int) []int {

	// 获取最大出现次数
	max := SameNum(array)
	// 按出现次数排名数组值
	recoverArray := array
	alikeArray := array
	for m:=0; m<max-1; m++ {
		recoverArray, _ = SameArray(recoverArray)
		alikeArray = append(alikeArray, recoverArray...)
	}
	alikeArray = ReverseArray(alikeArray)      // 反转数组
	alikeArray = RemoveRepeatArray(alikeArray) // 去重

	return alikeArray
}

//======

// SameNum 获取数组中重复最多的值的次数
func SameNum(array []int) int {
	fmt.Println(array)

	m1 := make(map[int]int)
	var s2 []int
	var max int
	for _, v := range array {
		if m1[v] != 0 {
			m1[v]++
		} else {
			m1[v] = 1
		}
	}
	for _, v := range m1 {
		s2 = append(s2, v)
	}
	//fmt.Println(s2)

	if len(s2) > 0 {
		max = s2[0]
		for i := 0; i < len(s2); i++ {
			if max < s2[i] {
				max = s2[i]
			}
		}
		//fmt.Println(max)
		return max
	}else {
		return len(array)
	}
}

// ReverseArray 数组反转
func ReverseArray(array []int) []int {
	x := array
	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}
	return x
}

// RemoveRepeatArray 数组去重
func RemoveRepeatArray(array []int) []int {
	newArray := make([]int, 0)

	for _, i := range array {
		if len(newArray) == 0 {
			newArray = append(newArray, i)
		} else {
			for k, v := range newArray {
				if i == v {
					break
				}
				if k == len(newArray)-1 {
					newArray = append(newArray,i)
				}
			}
		}
	}
	return newArray
}

//SameArray 获取一个数组中相等的键的集合（出现2次）
func SameArray(nums []int) ([]int, []int) {
	sort.Ints(nums)
	var i int
	var alike []int // 相同值数组
	var unlike []int // 异同值数组
	// 出现2次
	for i = 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			alike = append(alike, nums[i])
		}else {
			unlike = append(unlike, nums[i])
		}
	}
	return alike, unlike
}