package main

import "fmt"

// OCP
// open for extension, closed for modification
// specification

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, product := range products {
		if product.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for i, product := range products {
		if product.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

type Specification interface {
	isSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (cs ColorSpecification) isSatisfied(p *Product) bool {
	return p.color == cs.color
}

type SizeSpecification struct {
	size Size
}

func (ss SizeSpecification) isSatisfied(p *Product) bool {
	return p.size == ss.size
}

type AndSpecification struct {
	first, second Specification
}

func (s AndSpecification) isSatisfied(p *Product) bool {
	return s.first.isSatisfied(p) && s.second.isSatisfied(p)
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, product := range products {
		if spec.isSatisfied(&product) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	filter := Filter{}
	for _, product := range filter.FilterByColor(products, green) {
		fmt.Println(product)
	}

	betterFilter := BetterFilter{}
	for _, product := range betterFilter.Filter(products, ColorSpecification{green}) {
		fmt.Println(product)
	}
}
