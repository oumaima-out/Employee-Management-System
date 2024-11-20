package routes

import (
	"GO_API/controllers"  
	"github.com/gin-gonic/gin"
)

func EmployeeRoute(router *gin.Engine) {
	router.POST("/employees", controllers.CreateEmployee()) 
	router.GET("/employees/:id", controllers.GetAEmployee()) 
	router.PUT("/employees/:id", controllers.EditAEmployee()) 
	router.DELETE("/employees/:id", controllers.DeleteAEmployee()) 
	router.GET("/employees", controllers.GetAllEmployees())
}
