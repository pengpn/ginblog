package main

import (
	"gin-use/model"
	"gin-use/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()

}
