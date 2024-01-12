package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	r.Use(MiddleWare())

	v2 := r.Group("v2")
	{
		v2.POST("loginJSON", loginJSON)
	}
	r.Run(":8083")
}

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		c.Set("request", "中间件")
		// 执行路由对应的函数
		//c.Next()
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func loginJSON(c *gin.Context) {
	var loginInfo Login
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if loginInfo.User != "root" || loginInfo.Password != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}

type Login struct {
	User     string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
