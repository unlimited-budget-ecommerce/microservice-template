package main

import "github.com/unlimited-budget-ecommerce/microservice-template/internal"

var version string // set at build

func main() {
	_ = internal.NewConfig()
}
