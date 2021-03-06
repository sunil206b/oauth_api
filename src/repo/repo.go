package repo

import (
	"github.com/sunil206b/oauth_api/src/model"
	"github.com/sunil206b/store_utils_go/errors"
)

type ILoginRepo interface {
	GetById(string) (*model.AccessToken, *errors.RestErr)
	CreateToken(*model.AccessToken) *errors.RestErr
	UpdateExpiration(*model.AccessToken) *errors.RestErr
}
