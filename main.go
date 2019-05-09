package main

import (
	"apigin/router"
)

func main() {
	r := router.InitRouter()
	r.Run("8080")
}
