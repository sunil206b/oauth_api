package app

import (
	"database/sql"
	"github.com/sunil206b/oauth_api/src/controller"
)

func mapUrl(db *sql.DB) {
	loginHandler := controller.NewLoginController(db)
	router.GET("/login/access_token/:access_token", loginHandler.GetById)
}
