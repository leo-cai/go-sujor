package main

import (
	db "sujor.com/leo/sujor-api/database"
	rt "sujor.com/leo/sujor-api/router"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	//gin.SetMode(gin.DebugMode)
	defer db.SqlDB.Close()
	router := rt.InitRouter()
	router.Run(":8000")
}
