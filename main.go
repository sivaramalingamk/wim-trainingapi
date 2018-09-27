package main

import (

	"github.com/kataras/iris"
	"github.com/boniface/wim-trainingapi/api"
)

func main() {
	app := iris.New()
	// Landing Page Display
	app.RegisterView(iris.HTML("./views", ".html"))
	app.Get("/", api.Index())
	app.Get("/user/{id:long}", api.ParamFunction())
	app.Post("/ecudata", api.VehicleDataHandler)
	app.Post("/bulkecudata", api.VehicleDataCollectionHandler)
	// Start the server using a network address.
	app.Run(iris.Addr(":8080"))
}
