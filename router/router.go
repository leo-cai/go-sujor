package router

import (
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"time"
	indexController "sujor.com/leo/sujor-api/controller/index"
	userController "sujor.com/leo/sujor-api/controller/user"
	permissionController "sujor.com/leo/sujor-api/controller/permission"
	projectController "sujor.com/leo/sujor-api/controller/project"
	mottoController "sujor.com/leo/sujor-api/controller/motto"
	articleController "sujor.com/leo/sujor-api/controller/article"
)

// Initialize Router
func InitRouter() *gin.Engine {

	// 获得Gin路由实例
	router := gin.Default()

	// AJAX请求跨域
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, PATCH, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	// Index Router 测试api进程是否启用
	router.GET("/", indexController.IndexApi)

	// Router Group 群组
	// 版本v1
	v1 := router.Group("v1")
	{
		// User Auth 鉴权
		v1.GET("/users", userController.GetUsersApi) // 获取用户组
		v1.GET("/user/by/id/:id", userController.GetUserByIdApi) // 根据id获取用户
		v1.GET("/user/by/name/:username", userController.GetUserByNameApi) // 根据username获取用户
		v1.GET("/user/permissions/by/name/:username", permissionController.GetPermissionsByNameApi) // 根据username获取用户权限
		v1.POST("user/signup", userController.PostSignUpApi) // 用户注册
		v1.POST("user/signin", userController.PostSignInApi) // 用户登录
		v1.POST("user/signout", userController.PostSignOutApi) // 用户注销

		// Project 项目 api
		v1.GET("/projects", projectController.GetProjectsApi)

		// Motto 格言
		v1.GET("/mottos", mottoController.GetMottosApi)

		// Article 文章
		v1.GET("/articles", articleController.GetArticlesApi)
	}
	return router
}
