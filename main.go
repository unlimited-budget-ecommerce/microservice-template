package main

import "github.com/unlimited-budget-ecommerce/microservice-template/config"

var version string // set at build

func main() {
	_ = config.New()
}
