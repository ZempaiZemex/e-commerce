package main

import "e-commerce/productService/api/server"

func main() {
	server.InitServer()
	server.Start()
}
