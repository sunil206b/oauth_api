package repo

import (
	"database/sql"
	"github.com/sunil206b/oauth_api/src/model"
	"github.com/sunil206b/oauth_api/src/utils/errors"
)

type loginRepo struct {
	conn *sql.DB
}

func NewLoginRepo(db *sql.DB) ILoginRepo {
	return &loginRepo{
		conn: db,
	}
}

func (lr *loginRepo) GetById(token string) (*model.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("database connection not implemented yet")
}