package main

import (
	"github.com/boniface/wim-trainingapi/api"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	// Landing Page Display
	app.RegisterView(iris.HTML("./views", ".html"))
	app.Get("/", api.Index())
	app.Get("/user/{id:long}", api.ParamFunction())
	app.Post("/eucdata", api.VehicleDataHandler)
	app.Post("/eucdatacollection", api.VehicleDataCollectionHandler)
	// Start the server using a network address.
	app.Run(iris.Addr(":8080"))
}
