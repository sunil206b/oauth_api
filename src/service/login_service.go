package service

import (
	"database/sql"
	"github.com/sunil206b/oauth_api/src/dto"
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

func (ls *loginService) CreateToken(userDTO *dto.LoginUserDTO) (*model.AccessToken,*errors.RestErr) {
	user, err := ls.LoginUser(userDTO)
	if err != nil {
		return nil, err
	}

	at := model.GetNewAccessToken(user.Id)
	at.Generate()
	errMsg := at.Validate()
	if errMsg != nil {
		return nil, errMsg
	}

	if err := ls.repo.CreateToken(&at); err != nil {
		return nil, err
	}
	return &at, err
}

func (ls *loginService) UpdateExpiration(at *model.AccessToken) *errors.RestErr {
	return ls.repo.UpdateExpiration(at)
}