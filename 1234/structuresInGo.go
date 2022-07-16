package main

import (
	"fmt"
	"math"
)

func ArrayFiller(arr []string) {
	for index := range arr {
		(arr)[index] = fmt.Sprint(index)
		fmt.Printf("%d. new item has arrived\n", index)
	}

}
func ArraySwitcher(_arr []string) {

	var arr = _arr[0:]
	for left, right := 0, len(arr)-1; left < right; left, right = left+1, right-1 { //this func is from github "sliceTricks"
		arr[left], arr[right] = arr[right], arr[left]
	}
}

func nExponentShower(n float64) {
	if n > 0 {
		for i := float64(1); i <= n; i++ {
			fmt.Print(math.Pow(2, i), ", ")
		}
	}
}

// func main() {
// 	var toDoList = [3]string{
// 		"wash_the_dishes",
// 		"write_some_code",
// 		"find_job",
// 	}

// 	fmt.Println("Array Switcher: ")
// 	fmt.Println(toDoList)
// 	ArraySwitcher(toDoList[:])
// 	fmt.Println(toDoList)
// 	nExponentShower(5)
// 	// for index := range toDoList {
// 	// 	fmt.Printf("%d. %s\n", index, toDoList[index])
// 	// }

// 	// var newSlice = toDoList[0:]

// 	// ArrayFiller(newSlice)
// 	// fmt.Print("LIST:\n")
// 	// fmt.Printf("%s\n", toDoList)
// 	// fmt.Print("SLICE:\n")
// 	// fmt.Print(newSlice)

// }
