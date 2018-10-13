package main

import (
	"github.com/Vorian-Atreides/24sessions/backend"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//
// Gin middleware overlay the application code to run a REST server with Gin
//

func ginCompile(ctx *gin.Context, response *backend.Response) {
	ctx.Header("Content-Type", response.ContentType)
	if response.Err != nil {
		logrus.WithError(response.Err).Error(response.ErrMessage)
		errResponse := backend.NewError(response.ErrMessage)
		ctx.AbortWithStatusJSON(response.StatusCode, errResponse)
		return
	}
	if response.Data == nil {
		ctx.Status(response.StatusCode)
		return
	}
	ctx.JSON(response.StatusCode, response.Data)
}

type geoController struct {
	controller *backend.GeolocationController
}

func (g *geoController) getByIP(ctx *gin.Context) {
	request := &backend.GetByIPRequest{
		IP: ctx.Param("ip"),
	}

	answer := g.controller.GetByIP(request)
	ginCompile(ctx, answer)
}
