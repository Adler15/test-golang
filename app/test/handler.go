package test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

//@Summary 新增测试类
//@Description 新增测试类
//@Tags 测试类的增删查改
//@Accept json
//@Produce json
//@Param test body test.Test true "Test"
//@Success 200 {int} int "成功的条数"
//@Router /createTest [post]
func CreateTest(ctx *gin.Context) {
	test := Test{}
	ctx.BindJSON(&test)
	i := createTest(test)
	ctx.JSON(http.StatusOK, gin.H{"data": i})
}

//@Summary 更新测试类
//@Description 更新测试类
//@Tags 测试类的增删查改
//@Accept json
//@Produce json
//@Param test body test.Test true "Test"
//@Success 200 {int} int "成功的条数"
//@Router /updateTest [put]
func UpdateTest(ctx *gin.Context) {
	test := make(map[string]interface{})
	ctx.BindJSON(&test)
	i := updateTest(test)
	ctx.JSON(200, gin.H{"data": i})
}

//@Summary 删除测试类
//@Description 删除测试类
//@Tags 测试类的增删查改
//@Accept json
//@Produce json
//@Param id query string true "Test ID"
//@Success 200 {int} int "成功的条数"
//@Router /deleteTest [delete]
func DeleteTest(ctx *gin.Context) {
	id := ctx.Query("id")
	i := deleteTest(id)
	ctx.JSON(200, gin.H{"data": i})
}

//@Summary 查询全部测试类
//@Description 查询全部测试类
//@Tags 测试类的增删查改
//@Accept json
//@Produce json
//@Success 200 {object} test.Test "All Test"
//@Router /findAllTest [get]
func FindAllTest(ctx *gin.Context) {
	allTest := findAllTest()
	ctx.JSON(http.StatusOK, gin.H{"data": allTest})
}

//@Summary 查询一条测试类
//@Description 查询一条测试类
//@Tags 测试类的增删查改
//@Accept json
//@Produce json
//@Param id query string true "Test ID"
//@Success 200 {object} test.Test "Test"
//@Router /findOneTest [get]
func FindOneTest(ctx *gin.Context) {
	id := ctx.Query("id")
	oneTest := findOneTest(id)
	ctx.JSON(200, gin.H{"data": oneTest})
}

//@Summary 通过条件查询测试类
//@Description 通过条件查询测试类
//@Tags 测试类的增删查改
//@Accept json
//@Produce json
//@Param where body test.Test true "Where of Test"
//@Success 200 {object} test.Test "Where Test"
//@Router /findByWhere [get]
func FindByWhere(ctx *gin.Context) {
	where := make(map[string]interface{})
	ctx.BindJSON(&where)
	byWhere := findByWhere(where)
	ctx.JSON(200, gin.H{"data:": byWhere})
}

//@Summary 给kafka发送消息
//@Description 给kafka的cfltest找个topic发送消息
//@Tags kafka测试
//@Accept json
//@Produce json
//@Param test body test.Test true "send test to kafka"
//@Success 200 {object} test.Test "info"
//@Router /sendToKafka [post]
func SendToKafka(ctx *gin.Context) {
	content := make(map[string]interface{})
	ctx.BindJSON(&content)
	contestBytes, _ := json.Marshal(content)
	err := sendToKafka("cfltest", contestBytes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "发送kafka失败，error=" + err.Error()})
	}
	ctx.JSON(200, gin.H{"message": "发送kafka成功"})

}
