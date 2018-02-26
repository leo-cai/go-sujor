package permission

import (
	"github.com/gin-gonic/gin"
	. "sujor.com/leo/sujor-api/model/permission"
	"sujor.com/leo/sujor-api/config"
	"log"
	"net/http"
)

func GetPermissionsByNameApi(c *gin.Context)  {
	// 初始化users
	permissions := make([]Permission, 0)
	username := c.Param("username")
	// 统一错误处理
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg": err,
				"bean": nil,
			})
			return
		}
	}()

	// 调用 GetPermissionsByName 返回 permissions
	var p Permission
	permissions, err := p.GetPermissionsByName(username)

	// 错误处理
	if err != nil {
		panic(config.SqlError)
	}
	// 成功返回数据
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg": config.RequestSuccess,
		"bean": permissions,
	})
}