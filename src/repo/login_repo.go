package repo

import (
	"database/sql"
	"github.com/gocql/gocql"
	"github.com/sunil206b/oauth_api/src/driver"
	"github.com/sunil206b/oauth_api/src/model"
	"github.com/sunil206b/store_utils_go/errors"
)

const (
	queryGetToken = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpires = "UPDATE access_tokens SET expires=? WHERE access_token=?"
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
	var result model.AccessToken
	if err := driver.GetSession().Query(queryGetToken, token).Scan(&result.AccessToken,
		&result.UserId, &result.ClientId, &result.Expires); err != nil {
		if err == gocql.ErrNotFound{
			return nil, errors.NewNotFoundError("No access token with given token")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &result, nil
}

func (lr *loginRepo) CreateToken(at *model.AccessToken) *errors.RestErr {
	if err := driver.GetSession().Query(queryCreateToken, at.AccessToken,
		at.UserId, at.ClientId, at.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (lr *loginRepo) UpdateExpiration(at *model.AccessToken) *errors.RestErr {
	if err := driver.GetSession().Query(queryUpdateExpires, at.Expires, at.AccessToken).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}