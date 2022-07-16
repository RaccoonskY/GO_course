package storage

import (
	"errors"
	"math"
)

type Employee struct {
	Id     int
	Name   string
	Age    string
	Salary int
}

type Storage interface {
	Insert(e Employee) error
	Get(id int) (Employee, error)
	Delete(id int) error
}

type MemoryStorage struct {
	Data map[int]Employee
}

func NewMemStorage() *MemoryStorage {
	return &MemoryStorage{
		Data: make(map[int]Employee),
	}
}

func (s *MemoryStorage) Insert(e Employee) error {
	s.Data[e.Id] = e
	return nil
}

func (s *MemoryStorage) Get(id int) (Employee, error) {

	e, exists := s.Data[id]
	if !exists {
		return Employee{}, errors.New("there is no Employee with such id! ")
	}
	return e, nil
}

func (s *MemoryStorage) Delete(id int) error {
	delete(s.Data, id)
	return nil
}

///Primitives start here

type IPrimitive interface {
	Get_area() float32
	Get_color() string
}

type Circle struct {
	_radius int
	_color  string
}

func NewCircle(rad int, color string) *Circle {
	return &Circle{
		_radius: rad,
		_color:  color,
	}
}

func (c *Circle) Get_area() float32 {
	return float32(math.Pow(float64(c._radius), 2) * math.Pi)
}

func (c *Circle) Get_color() string {
	return c._color
}
