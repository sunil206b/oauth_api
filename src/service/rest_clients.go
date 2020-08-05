package service

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/sunil206b/oauth_api/src/dto"
	"github.com/sunil206b/oauth_api/src/utils/errors"
	"time"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8080",
		Timeout: 100 * time.Millisecond,
	}
)

func (ls *loginService) LoginUser(loginUserDTO *dto.LoginUserDTO) (*dto.UserDTO, *errors.RestErr) {
	res := usersRestClient.Get(fmt.Sprintf("/login/users/%s", loginUserDTO.Email))
	if res == nil || res.Response == nil {
		return nil, errors.NewInternalServerError("invalid rest client response when trying to login user")
	}
	if res.StatusCode > 299 {
		var errMsg errors.RestErr
		err := json.Unmarshal(res.Bytes(), &errMsg)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login")
		}
		return nil, &errMsg
	}

	var userDTO dto.UserDTO
	if err := json.Unmarshal(res.Bytes(), &userDTO); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal users response")
	}
	return &userDTO, nil
}