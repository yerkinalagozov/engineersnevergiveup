package main

import (
	"fmt"
	"sort"
)

type Users struct {
	Name string
	Age  int
}

func main() {

	//users := []Users{
	//	{
	//		Name: "a",
	//		Age:  1,
	//	},
	//}
	//fmt.Println(users)
	//
	//var u []Users
	//u = append(u, users...)

	//O(n)

	//arr11 := [10]int{}
	//fmt.Println(arr11, "--arr11 origin--")
	//arr12 := make([]int, 10)
	//fmt.Println(arr12, "--arr12 origin--")

	//////////////////////////////////////

	//slice02 := []int{}
	//fmt.Println(slice02, "--slice02 origin--")
	//
	//slice03 := make([]int, 10)
	//fmt.Println(slice03, "--slice03 origin--")

	//////////////////////////////////////

	//arr01 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//slice04 := arr01[0:10]
	//fmt.Println(slice04, "--slice04 origin--")
	//slice05 := arr01[5:10]
	//fmt.Println(slice05, "--slice05 origin--")
	//arr01[2] = 100
	//fmt.Println(slice04, "--slice04 after--")
	//fmt.Println(slice05, "--slice05 after--")
	//arr01[7] = 222
	//fmt.Println(slice04, "--slice04 after 2--")
	//fmt.Println(slice05, "--slice05 after 2--")

	//////////////////////////////////////

	//slice01 := []int{1, 2, 3, 4, 5}
	//fmt.Println(slice01, "--slice01 origin--")
	//fmt.Printf("======== Адрес слайса: %p\n", &slice01)
	//fmt.Printf("======== Адрес базового массива: %p\n", &slice01[0])

	//addNum2Slice(slice01)
	//fmt.Println(slice01, "--addNum2Slice--")
	//fmt.Printf("--addNum2Slice--  Адрес среза: %p\n", &slice01)
	//fmt.Printf("--addNum2Slice--  Адрес базового массива: %p\n\n", &slice01[0])

	//result1 := addNum2Slice2(slice01)
	//fmt.Println(result1, "--addNum2Slice2--")
	//fmt.Printf("--addNum2Slice2--  Адрес среза: %p\n", &result1)
	//fmt.Printf("--addNum2Slice2--  Адрес базового массива: %p\n", &result1[0])

	//addNum2Slice3(&slice01)
	//fmt.Println(slice01, "--addNum2Slice3--")
	//result2 := addNum2Slice4(&slice01)
	//fmt.Println(result2, "--addNum2Slice4--")
	//changeNumIntoSlice(slice01)
	//fmt.Println(slice01, "--changeNumIntoSlice--")
	//changeNumIntoSlice2(&slice01)
	//fmt.Println(slice01, "--changeNumIntoSlice2--")
	//result3 := changeNumIntoSlice3(slice01)
	//fmt.Println(result3, "--changeNumIntoSlice3--")
	//result4 := changeNumIntoSlice4(&slice01)
	//fmt.Println(result4, "--changeNumIntoSlice4--")

	sl01 := []int{1, 5, 3, 1, 6}
	sl02 := []int{6, 7, 3, 9, 1}

	result5 := append2SliceAndSort(sl01, sl02)
	fmt.Println(result5, "--append2Slice--")
}

func addNum2Slice(slice01 []int) {
	fmt.Printf("Before: Адрес среза: %p\n", &slice01)
	fmt.Printf("Before: Адрес базового массива: %p\n", &slice01[0])
	slice01[0] = 100
	fmt.Println(slice01, "--slice01 origin+++++--")
	slice01 = append(slice01, 6)
	slice01[1] = 666
	slice01 = append(slice01, 7)
	fmt.Printf("After: Адрес среза: %p\n", &slice01)
	fmt.Printf("After: Адрес базового массива: %p\n", &slice01[0])
}

func addNum2Slice2(slice01 []int) []int {
	slice02 := append(slice01, 6)
	return slice02
}

func addNum2Slice3(slice01 *[]int) {
	*slice01 = append(*slice01, 6)
}

func addNum2Slice4(slice01 *[]int) *[]int {
	*slice01 = append(*slice01, 6)
	return slice01
}

func changeNumIntoSlice(slice01 []int) {
	slice01[0] = 100
}

func changeNumIntoSlice2(slice01 *[]int) {
	(*slice01)[0] = 100
}

func changeNumIntoSlice3(slice01 []int) []int {
	slice01[0] = 100
	return slice01
}

func changeNumIntoSlice4(slice01 *[]int) *[]int {
	(*slice01)[0] = 100
	return slice01
}

func append2Slice(sl01 []int, sl02 []int) []int {
	result := append(sl01, sl02...)
	return result
}

func append2SliceAndSort(sl01 []int, sl02 []int) []int {
	result := append(sl01, sl02...)
	sort.Ints(result) //quick sort - быстрая сортировка
	return result
}
