package main

import (
	"Memorandum/config"
	"Memorandum/routes"
)

func main() {
	config.Init()
	r := routes.NewRouter()
	r.Spin()
}
