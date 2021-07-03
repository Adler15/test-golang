package student

import "github.com/gin-gonic/gin"

func Routers(route *gin.Engine) {
	route.GET("findAll", FindAll)
	route.GET("findOne", FindOne)
	route.GET("findByName", FindByName)
	route.POST("create", Create)
	route.PUT("update", Update)
	route.DELETE("delete", Delete)
}
