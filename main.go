package main

import (
	"github.com/gin-gonic/gin"
	"github.com/humbinal/gin-openapi/openapi"
	"github.com/swaggo/swag/example/celler/controller"
	"log"
)

// @title			Swagger Example API
// @version		1.0
// @description	This is a sample server api docs.
//
// @servers.url	/api/v1
// @securityDefinitions.basic	BasicAuth
func main() {
	r := gin.Default()

	c := controller.NewController()

	v1 := r.Group("/api/v1")
	{
		examples := v1.Group("/examples")
		{
			examples.GET("ping", c.PingExample)
			examples.GET("calc", c.CalcExample)
			examples.GET("groups/:group_id/accounts/:account_id", c.PathParamsExample)
			examples.GET("header", c.HeaderExample)
			examples.GET("securities", c.SecuritiesExample)
			examples.GET("attribute", c.AttributeExample)
		}
	}
	openapi.RegisterOpenApiRoute(r)
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
		return
	}
}
