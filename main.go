package main

import (
	"gin-blog/model"
	"gin-blog/routes"
)

func main(){
	model.InitDb()
	routes.InitRouter()
}