package main

import (
	"fmt"
	"go-dlv/consulclient"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	a := 2
	b := 3
	c := a + b

	fmt.Println(c)

	r.GET("/", func(c *gin.Context) {
		consulclient.GetClient()
		c.String(200, "pong")
	})

	r.POST("register", Register)

	if err := r.Run(":8089"); err != nil {
		fmt.Println(err)
	}
}

func Register(c *gin.Context) {
	var r RegisterRequest
	err := c.ShouldBindJSON(&r)
	if err != nil {
		fmt.Println("register failed")
		c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		return
	}
	//验证 存储操作省略.....
	fmt.Println("register success")
	c.JSON(http.StatusOK, "successful")
}

// RegisterRequest is
// boss: 这小子还不会使用validator库进行数据校验，开了～ https://mp.weixin.qq.com/s/rleIt9uVn0-OobCuOhnlqQ
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Age      uint8  `json:"age" binding:"gte=1,lte=120"`
}
