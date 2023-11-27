package main

import "fmt"

// the OCP is open for extension but close for modification
// enterprise pattern - Specification

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

/*
	small Size = iota: This line declares a constant named small of type Size and
assigns it the value of iota.
The iota starts with a value of 0 and increments by 1 for each subsequent line
in the constant block.
*/

const (
	small  Size = iota // 0
	medium             // 1
	large              // 2
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
	//
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}

	return result
}

func (f *Filter) FilterByColorAndSize(products []Product, color Color, size Size) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color && v.size == size {
			result = append(result, &products[i])
		}
	}

	return result
}

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (c SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == c.size
}

type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

type BetterFilter struct{}

func (bf *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
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

	fmt.Println("Green Products (OLD):")
	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	fmt.Printf("\n\nGreen Products (NEW):\n")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	fmt.Printf("\n\nLarge Products (NEW):\n")
	largeSpec := SizeSpecification{large}
	for _, v := range bf.Filter(products, largeSpec) {
		fmt.Printf(" - %s is large\n", v.name)
	}

	fmt.Printf("\n\nLarge Green Products:\n")
	lgSpec := AndSpecification{greenSpec, largeSpec}
	for _, v := range bf.Filter(products, lgSpec) {
		fmt.Printf(" - %s is large and green\n", v.name)
	}
}
