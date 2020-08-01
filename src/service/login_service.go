package service

import (
	"database/sql"
	"github.com/sunil206b/oauth_api/src/model"
	"github.com/sunil206b/oauth_api/src/repo"
	"github.com/sunil206b/oauth_api/src/utils/errors"
)

type loginService struct {
	repo repo.ILoginRepo
}

func NewLoginService(db *sql.DB) ILoginService{
	return &loginService{
		repo: repo.NewLoginRepo(db),
	}
}

func (ls *loginService) GetById(tokenId string) (*model.AccessToken, *errors.RestErr) {
	accessToken, errMsg := ls.repo.GetById(tokenId)
	if errMsg != nil {
		return nil, errMsg
	}
	return accessToken, nil
}