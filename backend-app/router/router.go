package router

import (
	"backend-app/utils/di"
	"database/sql"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

func GetRouter(db *sql.DB) *gin.Engine {

	cr, err := di.InitializeControllers(db)
	if err != nil {
		slog.Error("Failed to initialize controllers: ", err)
		os.Exit(1)
	}

	router := gin.Default()

	router.GET("/admin/users", cr.IUserController.GetUsersList)

	return router
}
