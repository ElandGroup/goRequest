package main

import "goRequest/api"

func RegisterUserApi() {
	apiGroup := e.Group("/api")
	v1 := apiGroup.Group("/v1")

	reqXml := v1.Group("/reqxml")
	reqXml.GET("", api.ReqXml)

	v1.POST("/acceptjson", api.AcceptJson)

}
