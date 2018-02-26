package user

import (
	. "sujor.com/leo/sujor-api/model/user"
	"sujor.com/leo/sujor-api/config"
	"sujor.com/leo/sujor-api/middleware/authJWT"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"crypto/sha1"
	"fmt"
)

// GetUsers 查询用户组 API
func GetUsersApi(c *gin.Context) {
	// 初始化users
	users := make([]User, 0)
	// 获取limit参数
	cLimit := c.Query("limit")
	limit, err := strconv.Atoi(cLimit)
	// 统一错误处理
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg": err,
				"bean": users,
			})
			return
		}
	}()
	// limit只能是正整数
	if err != nil || limit <= 0 {
		panic(config.ParamError + "limit: " + cLimit)
	}
	// 获取page参数
	cPage := c.Query("page")
	page, err := strconv.Atoi(cPage)
	// page只能是正整数
	if err != nil || page <= 0 {
		panic(config.ParamError + "page: " + cPage)
	}
	// 调用GetUsers 获取users
	var u User
	users, err = u.GetUsers(limit, page)
	// 错误处理
	if err != nil {
		panic(config.SqlError)
	}
	// 成功返回数据
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg": config.RequestSuccess,
		"bean": users,
	})
}

// GetUserById 查询用户 API
func GetUserByIdApi(c *gin.Context) {
	// 获取id参数
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
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
	// id参数错误处理
	if err != nil {
		panic(config.ParamError + "id: " + cid)
	}
	// 调用GetUserById 获取user
	u := User{Id: id}
	user, err := u.GetUserById()
	// 错误处理
	if err != nil {
		panic(config.SqlError)
	}
	// 成功返回数据
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg": config.RequestSuccess,
		"bean": user,
	})
}

// GetUserByName 查询用户 API
func GetUserByNameApi(c *gin.Context) {
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
	// 调用GetUserByName 获取user
	u := User{Username: username}
	user, err := u.GetUserByName()
	// 错误处理
	if err != nil {
		panic(config.SqlError)
	}
	// 成功返回数据
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg": config.RequestSuccess,
		"bean": user,
	})
}

// PostSignUpApi 注册用户 API
func PostSignUpApi(c *gin.Context) {
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
	// bind JSON错误处理
	var jsonUser User
	if c.BindJSON(&jsonUser) != nil {
		panic(config.ParamError)
	}
	// 查重 是否已注册
	u := User{Username: jsonUser.Username}
	user, err := u.GetUserByName()
	// 错误处理
	if user.Id != 0 && err == nil {
		panic(config.SignUpError)
	}

	// 密码 sha1 加密
	h := sha1.New()
	h.Write([]byte(jsonUser.Password))
	jsonUser.Password = fmt.Sprintf("%x", h.Sum(nil))

	// 调用SignUp注册 返回user
	user, err = u.SignUp(jsonUser.Username, jsonUser.Password)
	if err != nil {
		panic(config.SqlError)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg": config.SignUpSuccess,
		"bean": user,
	})
}

// PostSignInApi 登录用户 API
func PostSignInApi(c *gin.Context) {
	// 统一错误处理
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg": err,
				"bean": nil,
			})
			return
		}
	}()
	// bind JSON错误处理
	var jsonUser User
	if c.BindJSON(&jsonUser) != nil {
		panic(config.ParamError)
	}

	// 密码 sha1 加密
	h := sha1.New()
	h.Write([]byte(jsonUser.Password))
	jsonUser.Password = fmt.Sprintf("%x", h.Sum(nil))

	// 调用SignIn登录 返回user
	var u User
	user, err := u.SignIn(jsonUser.Username, jsonUser.Password)
	if err != nil {
		panic(config.SignInError)
	}
	// 生成jwt 错误处理
	tokenString, err := authJWT.GenerateJWT()
	if err != nil {
		panic(config.InternalServerError)
	}
	// 添加到 header
	c.Writer.Header().Add("Authorization", "Bearer " + tokenString)
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg": config.SignInSuccess,
		"bean": user,
		"token": tokenString,
		"expires": 1, // 1天后失效
	})
}

func PostSignOutApi(c *gin.Context)  {
	// 统一错误处理
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg": err,
				"bean": nil,
			})
			return
		}
	}()

	// bind JSON错误处理
	var jsonUser User
	if c.BindJSON(&jsonUser) != nil {
		panic(config.ParamError)
	}
	// 调用SignOut 更新上次登录时间
	var u User
	err := u.SignOut(jsonUser.Username)
	if err != nil {
		panic("时间戳错误！")
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg": config.SignOutSuccess,
		"bean": nil,
	})
}