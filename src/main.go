package main

import (
	"database/sql"
	"github.com/sunil206b/oauth_api/src/app"
)

func main() {
	app.StartApp(&sql.DB{})
}
