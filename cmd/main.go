package main

import "fsa-food-api/api"

func main() {
	server := api.Server{}
	server.Start()
}
