package api

import (
	"github.com/kataras/iris"
	"github.com/sivaramalingamk/wim-trainingapi/domain"
)

func VehicleDataHandler(ctx iris.Context) {
	var data domain.RawInputData

	if err := ctx.ReadJSON(&data); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	ctx.Writef("Received: %#+v\n", data)
}

// VehicleDataCollectionHandler reads a collection of VehicleData from JSON post body.
func VehicleDataCollectionHandler(ctx iris.Context) {
	var dataCollection []domain.RawInputData
	err := ctx.ReadJSON(&dataCollection)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	ctx.Writef("Received: %#+v\n", dataCollection)
}
