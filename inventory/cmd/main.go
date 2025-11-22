package main

import "time"

type Category int

const (
	UNKNOWN Category = iota
	ENGINE
	FUEL
	PORTHOLE
	WING
)

type Dimensions struct {
	length float64
	width  float64
	height float64
	weight float64
}

type Manufacturer struct {
	name    string
	country string
	website string
}

type Part struct {
	uuid           string
	name           string
	description    string
	price          float64
	stock_quantity int64
	category       Category
	dimensions     Dimensions
	manufacturer   Manufacturer
	tags           []string
	metadata       map[string]any
	created_at     time.Time
	updated_at     time.Time
}

func GetPart() {

}
func main() {

}
