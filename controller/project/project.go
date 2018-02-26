package project

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"sujor.com/leo/sujor-api/config"
	. "sujor.com/leo/sujor-api/model/project"
)

func GetProjectsApi(c *gin.Context)  {
	cLimit := c.Query("limit")
	limit, err := strconv.Atoi(cLimit)
	// 参数错误处理
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err,
				"projects": nil,
			})
			return
		}
	}()
	if err != nil || limit <= 0 {
		panic(config.ParamError)
	}
	cPage := c.Query("page")
	page, err := strconv.Atoi(cPage)

	if err != nil || page <= 0 {
		panic(config.ParamError)
	}

	var p Project
	projects, err := p.GetProjects(limit, page)

	// SQL错误 或 无查找到数据
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"msg": err,
				"projects": nil,
			})
			return
		}
	}()

	if err != nil {
		panic(config.SqlError)
	}

	c.JSON(http.StatusOK, gin.H{
		"projects": projects,
	})

}
