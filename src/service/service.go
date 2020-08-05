package service

import (
	"github.com/sunil206b/oauth_api/src/dto"
	"github.com/sunil206b/oauth_api/src/model"
	"github.com/sunil206b/oauth_api/src/utils/errors"
)

type ILoginService interface {
	GetById(string) (*model.AccessToken, *errors.RestErr)
	CreateToken(userDTO *dto.LoginUserDTO) (*model.AccessToken,*errors.RestErr)
	UpdateExpiration(*model.AccessToken) *errors.RestErr
	LoginUser(loginUserDTO *dto.LoginUserDTO) (*dto.UserDTO, *errors.RestErr)
}
