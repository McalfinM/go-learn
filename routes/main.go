package routes

import (
	authcontroller "learn/api/controllers/authController"
	database "learn/api/databases"
	"learn/api/domains/account"
	"learn/api/initializers"

	"github.com/gin-gonic/gin"
)

func Register() {
	Routes()
}
func Routes() {
	initializers.LoadEnv()
	database.ConnectDatabase()

	db := database.DB

	router := gin.Default()

	v1 := router.Group("/v1")

	accoutRepository := account.NewRepository(db)
	accountService := account.NewService(accoutRepository)
	accountHandler := authcontroller.NewAuthController(accountService)

	v1.GET("/", accountHandler.FindAll)
	v1.POST("/register", accountHandler.Register)

	router.Run(":3003")

}
