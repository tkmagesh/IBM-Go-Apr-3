package main

import (
	"fmt"
	"sort"
)

type Product struct {
	id       int
	name     string
	cost     float32
	units    int
	category string
}

type Products []Product

type ProductsComparer func(p1, p2 Product) bool

/* Methods for Stringer interface of the 'fmt' package */
func (p Product) String() string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%v, Units=%d, Category=%s\n", p.id, p.name, p.cost, p.units, p.category)
}

func (products Products) String() string {
	result := ""
	for _, product := range products {
		result += fmt.Sprint(product)
	}
	return result
}

var currentProductComparer ProductsComparer

var productComparers map[string]ProductsComparer = map[string]ProductsComparer{
	"id": func(p1, p2 Product) bool {
		return p1.id < p2.id
	},
	"name": func(p1, p2 Product) bool {
		return p1.name < p2.name
	},
	"cost": func(p1, p2 Product) bool {
		return p1.cost < p2.cost
	},
	"units": func(p1, p2 Product) bool {
		return p1.units < p2.units
	},
	"category": func(p1, p2 Product) bool {
		return p1.category < p2.category
	},
}

/* Methods for 'Interface' interface of 'sort' package  */
func (products Products) Len() int {
	return len(products)
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

func (products Products) Less(i, j int) bool {
	return currentProductComparer(products[i], products[j])
}

func (products Products) Sort(attrName string) {
	currentProductComparer = productComparers[attrName]
	sort.Sort(products)
}

func main() {

	products := Products{
		{id: 100, name: "Pen", cost: 10, units: 10, category: "stationary"},
		{id: 106, name: "Den", cost: 16, units: 50, category: "stationary"},
		{id: 107, name: "Ken", cost: 12, units: 20, category: "utencil"},
		{id: 102, name: "Zen", cost: 18, units: 70, category: "stationary"},
		{id: 104, name: "Ten", cost: 15, units: 30, category: "utencil"},
		{id: 103, name: "Len", cost: 14, units: 50, category: "stationary"},
	}
	fmt.Println("Initial List")
	fmt.Println(products)

	fmt.Println("Sort By id")
	products.Sort("id")
	fmt.Println(products)

	fmt.Println("Sort By name")
	products.Sort("name")
	fmt.Println(products)

	fmt.Println("Sort By cost")
	products.Sort("cost")
	fmt.Println(products)

	fmt.Println("Sort By units")
	products.Sort("units")
	fmt.Println(products)

	fmt.Println("Sort By category")
	products.Sort("category")
	fmt.Println(products)

}
