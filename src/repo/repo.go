package repo

import (
	"github.com/sunil206b/oauth_api/src/model"
	"github.com/sunil206b/oauth_api/src/utils/errors"
)

type ILoginRepo interface {
	GetById(string) (*model.AccessToken, *errors.RestErr)
}
