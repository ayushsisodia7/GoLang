package main

import (
	"jwt/controllers"
	"jwt/initializers"
	"jwt/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	//fmt.Println("step 1 done")
	initializers.ConnectToDB()
	//fmt.Println("step 2 done")
	initializers.SyncDatabase()
	//fmt.Println("step 3 done")
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.POST("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run()

}
