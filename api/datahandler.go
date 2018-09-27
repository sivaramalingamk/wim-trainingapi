package api

import (
	"github.com/boniface/wim-trainingapi/domain"
	"github.com/kataras/iris"
)

func VehicleDataHandler(ctx iris.Context) {
	var data domain.VehicleData

	if err := ctx.ReadJSON(&data); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	ctx.Writef("Received: %#+v\n", data)
}

// VehicleDataCollectionHandler reads a collection of VehicleData from JSON post body.
func VehicleDataCollectionHandler(ctx iris.Context) {
	var dataCollection []domain.VehicleData
	err := ctx.ReadJSON(&dataCollection)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	ctx.Writef("Received: %#+v\n", dataCollection)
}
