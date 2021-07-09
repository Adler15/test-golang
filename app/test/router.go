package test

import "github.com/gin-gonic/gin"

func Routers(router *gin.Engine) {
	router.GET("findOneTest", FindOneTest)
	router.GET("findAllTest", FindAllTest)
	router.GET("findByWhere", FindByWhere)
	router.POST("createTest", CreateTest)
	router.PUT("updateTest", UpdateTest)
	router.DELETE("deleteTest", DeleteTest)
	router.POST("sendToKafka",SendToKafka)
}
