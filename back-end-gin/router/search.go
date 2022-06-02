package router

import (
	"fmt"
	"net/http"
	"search/model"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//搜索结果
func searchResult(c *gin.Context) {
	//连接服务器
	dsn := "root:888888@tcp(34.66.167.238:3306)/histories?charset=utf8mb4&parseTime=True&loc=Local"
	//打开服务器
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//检测服务器错误
	if err != nil {
		panic("can't connect to database")
	}
	//如果没有服务器，按照models创建服务器
	db.AutoMigrate(&model.History{})
	//query网址来获取搜索关键词
	result := c.Query("result")
	filter := c.Query("filter")

	//确认后端收到没有
	info := "backend received: result is " + result + " and filter is " + filter
	var userEmail model.UserEmail
	//接受用户email来添加历史记录
	err2 := c.ShouldBindJSON(&userEmail)
	fmt.Println(userEmail)
	//如果找不到用户email（未登录），不添加历史记录
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not found!"})
		return
	}
	//添加历史记录时间
	dt := time.Now()
	fmt.Println("Time is ")
	fmt.Println(dt.String())
	//转换时间
	dtConverted := dt.Format("2006-01-02 15:04:05")
	//存在对象record里
	record := model.History{
		Email:    userEmail.Email,
		Result:   result,
		Filter:   filter,
		CreateAt: dtConverted,
	}

	if len(record.Email) != 0 {
		//如果用户存在，存到数据库里
		db.Create(&record)
	} else {
		//如果用户不存在，什么都不存
		fmt.Println("Not registered!")
	}
	c.JSON(http.StatusOK, gin.H{"data": info})

	// searchResult = result
	// searchFilter = filter
}

//获取搜索结果（WIP）
func getResult(c *gin.Context) {

}
