package main

import (
	"gims/router"
)


func main() {
	r := router.Setup()
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
