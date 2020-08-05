package model

import (
	"fmt"
	"github.com/sunil206b/oauth_api/src/utils/errors"
	"github.com/sunil206b/users_api/utils/crypt"
	"strings"
	"time"
)

const (
	expirationTime = 24
)
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId int64 `json:"user_id"`
	ClientId int64 `json:"client_id"`
	Expires int64 `json:"expires"`
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId: userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequest("invalid access token id")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequest("invalid user id")
	}
	if at.ClientId <= 0 {
		return errors.NewBadRequest("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequest("invalid expiration time")
	}
	return nil
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypt.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}