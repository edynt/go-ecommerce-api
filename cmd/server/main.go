package main

import "github.com/edynt/go-ecommerce-api/internal/routers"

func main() {
	r := routers.NewRouter()
	r.Run()
}
