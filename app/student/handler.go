package student

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//@Summary 获取全部学生数据接口
//@Description 查询全部学生数据
//@Tags 用户的增删查改
//@Produce json
//@Success 200 {object} student.StudentTest	"All Student"
//@Router /findAll [get]
func FindAll(ctx *gin.Context) {
	all := findAll()
	fmt.Println(all)
	ctx.JSON(200, gin.H{"data": all})
}

//@Summary 获取某学生信息接口
//@Description 查询一个学生数据
//@Tags 用户的增删查改
//@Accept json
//@Produce json
//@Param id query int true "ID"
//@Success 200 {object} student.StudentTest	"Student"
//@Router /findOne [get]
func FindOne(ctx *gin.Context) {
	id := ctx.Query("id")
	one := findOne(id)
	ctx.JSON(http.StatusOK, gin.H{"data": one})
}

//@Summary 通过名字获取某学生信息接口
//@Description 通过名字查询学生数据
//@Tags 用户的增删查改
//@Accept json
//@Produce json
//@Param name query string true "Student Name"
//@Success 200 {object} student.StudentTest	"Student"
//@Router /findByName [get]
func FindByName(ctx *gin.Context) {
	name := ctx.Query("name")
	byName := findByName(name)
	ctx.JSON(http.StatusOK, gin.H{"data": byName})
}

//@Summary 新增学生信息接口
//@Description 新增学生数据
//@Tags 用户的增删查改
//@Accept json
//@Produce json
//@Param student body student.StudentTest true "Student"
//@Success 200 {int} int "RowsAffected"
//@Router /create [post]
func Create(ctx *gin.Context) {
	student := StudentTest{}
	ctx.BindJSON(&student)
	i := create(student)
	ctx.JSON(http.StatusOK, gin.H{"data": i})
}

//@Summary 更新学生信息接口
//@Description 更新学生数据
//@Tags 用户的增删查改
//@Accept json
//@Produce json
//@Param student body student.StudentTest true "Student"
//@Success 200 {int} int "RowsAffected"
//@Router /update [put]
func Update(ctx *gin.Context)  {
	stMap := make(map[string]interface{}) //注意该结构接受的内容
	ctx.BindJSON(&stMap)
	i := update(stMap)
	ctx.JSON(http.StatusOK, gin.H{"data": i})
}

//@Summary 删除学生信息接口
//@Description 删除学生数据
//@Tags 用户的增删查改
//@Accept json
//@Produce json
//@Param id query int true "Student ID"
//@Success 200 {int} int "RowsAffected"
//@Router /delete [delete]
func Delete(ctx *gin.Context) {
	id := ctx.Query("id")
	i := deleteSt(id)
	ctx.JSON(http.StatusOK, gin.H{"data": i})
}

//@Summary 给kafka发送消息
//@Description 给kafka的cfltest找个topic发送消息
//@Tags kafka测试student
//@Accept json
//@Produce json
//@Param test body student.StudentTest true "send student to kafka"
//@Success 200 {object} student.StudentTest "student info"
//@Router /sendStudentToKafka [post]
func SendStudentToKafka(ctx *gin.Context) {
	content := make(map[string]interface{})
	ctx.BindJSON(&content)
	contestBytes, _ := json.Marshal(content)
	err := sendToKafka("cfltest", contestBytes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "发送kafka失败，error=" + err.Error()})
	}
	ctx.JSON(200, gin.H{"message": "发送kafka成功"})

}