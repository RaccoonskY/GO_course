package main

import (
	"fmt"
)

func ArrayFiller(arr []string) {
	for index := range arr {
		(arr)[index] = fmt.Sprint(index)
		fmt.Printf("%d. new item has arrived\n", index)
	}

}
func ArraySwitcher(_arr []string) {
	_begin := 0
	_tail := len(_arr) - 1
	var _temp string

	var arr = _arr[0:]
	for arr[_begin] != arr[_tail] {
		_temp = arr[_begin]
		arr[_begin] = arr[_tail]
		arr[_tail] = _temp
		_tail--
		_begin++
	}
}

func main() {
	var toDoList = [3]string{
		"wash_the_dishes",
		"write_some_code",
		"find_job",
	}

	fmt.Println("Array Switcher: ")
	fmt.Println(toDoList)
	ArraySwitcher(toDoList[:])
	fmt.Print(toDoList)
	// for index := range toDoList {
	// 	fmt.Printf("%d. %s\n", index, toDoList[index])
	// }

	// var newSlice = toDoList[0:]

	// ArrayFiller(newSlice)
	// fmt.Print("LIST:\n")
	// fmt.Printf("%s\n", toDoList)
	// fmt.Print("SLICE:\n")
	// fmt.Print(newSlice)

}
