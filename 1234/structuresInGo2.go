package main

import (
	"fmt"
)

type MyStruct struct {
	_name   string
	_salary float32
	_next   *MyStruct
}

func (m MyStruct) getName() string {
	return m._name
}

func (m MyStruct) getSalary() float32 {
	return m._salary
}

func (m MyStruct) getInfo() string {
	return fmt.Sprintf("Worker: %s\nSalary: %f", m._name, m._salary)
}

func newMyStruct(name string, salary float32, next *MyStruct) MyStruct {
	return MyStruct{
		_name:   name,
		_salary: salary,
		_next:   next,
	}
}

func (m *MyStruct) setName(name string) {
	m._name = name
}

// func main() {
// 	var mybuddy MyStruct = newMyStruct("Johnny", 30000, nil)
// 	var newName string
// 	scanner := bufio.NewScanner(os.Stdin) //std::cin analog,

// 	fmt.Printf("%s\n", mybuddy.getInfo())
// 	scanner.Scan() //std::cin analog, but it's object
// 	newName = scanner.Text()
// 	mybuddy.setName(newName)

// 	fmt.Printf("%s\n", mybuddy.getInfo())

// }
