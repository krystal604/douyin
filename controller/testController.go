package controller

import (
	"douyin/dao"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	id := dao.SelectUserById(1)
	fmt.Println(id)
}
