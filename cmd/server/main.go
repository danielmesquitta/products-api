package main

import "github.com/danielmesquitta/products-api/internal/app/http"

// @title Products API
// @version 1.0
// @description This is a CRUD API for products.
// @contact.name Daniel Mesquita
// @contact.email danielmesquitta123@gmail.com
// @BasePath /api/v1
func main() {
	http.Start()
}
