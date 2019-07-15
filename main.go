package main

import (
	app "tournament/pkg"
	"tournament/pkg/router"
)

func main() {
	h := router.Route()
	app.Init(h)

}
