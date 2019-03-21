package main

import (
	"voting_system/app/routers"
)

func main() {
	router := routers.InitRoutes()

	router.Run()
}
