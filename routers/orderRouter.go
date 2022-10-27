package routers

import (
	"rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders/:orderId", controllers.GetOrderById)
	router.PUT("/orders/:orderId", controllers.UpdateOrder)
	router.DELETE("/orders/:orderId", controllers.DeleteOrderById)

	return router
}
