package main

import "fmt"

// Liskov Substitution Principle

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	setHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) setHeight(height int) {
	r.height = height
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	s := &Square{}
	s.width = size
	s.height = size
	return s
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

type Square2 struct {
	size int
}

func (s Square2) Rectangle() Rectangle {
	return Rectangle{s.size, s.size}
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.setHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetHeight() * sized.GetWidth()
	fmt.Println(expectedArea, actualArea)
}

func main() {
	rc := &Rectangle{2, 3}
	UseIt(rc)

	sq := NewSquare(5)
	UseIt(sq)
}
