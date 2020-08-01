package app

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/sunil206b/users_api/logger"
)

var (
	router = gin.Default()
)

func StartApp(db *sql.DB) {
	mapUrl(db)
	logger.Info("about to start the application...")
	router.Run(":8081")
}
