package main

import "book-management/pkg/server/pkg/rest"

func main() {
	s := rest.NewBookServer()
	s.Start()
}
