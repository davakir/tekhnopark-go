package main

import (
	"fmt"
	"strconv"
	"sort"
)

func main() {
	mapForSort := map[int]string {
		40: "Осень",
		10: "Зима",
		30: "Лето",
		20: "Весна",
	}
	res := GetMapValuesSortedByKey(mapForSort)

	fmt.Println(res)
}

func ReturnInt() int {
	return 1
}

func ReturnFloat() float32 {
	return 1.1
}

func ReturnIntArray() [3]int {
	return [...]int{1, 3, 4}
}

func ReturnIntSlice() []int {
	return []int{1,2,3}
}

func IntSliceToString(slice []int) string {
	var resultString string

	for i := 0; i < len(slice); i++ {
		resultString += strconv.Itoa(slice[i])
	}

	return resultString
}

func MergeSlices(firstSlice []float32, secondSlice []int32) []int {
	result := make([]int, 0, 0)

	for i := 0; i < len(firstSlice); i++ {
		result = append(result, int(firstSlice[i]))
	}

	for i := 0; i < len(secondSlice); i++ {
		result = append(result, int(secondSlice[i]))
	}

	return result
}

func GetMapValuesSortedByKey(mapForSort map[int]string) []string {
	var keys []int
	for key := range mapForSort {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	var res []string
	for i := 0; i < len(keys); i++ {
		res = append(res, mapForSort[keys[i]])
	}

	return res
}