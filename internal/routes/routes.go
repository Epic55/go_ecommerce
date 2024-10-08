package routes

import (
	"github.com/Epic55/ecommerce/internal/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.AddProduct())
	incomingRoutes.GET("/users/showallproducts", controllers.ShowAllProducts())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
}
