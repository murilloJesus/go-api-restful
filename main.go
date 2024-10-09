package main

import "restful/routes"

func main() {
	router := routes.SetupRouter()
	router.Run(":8080")
}
