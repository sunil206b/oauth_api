package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/sunil206b/oauth_api/src/dto"
	"github.com/sunil206b/oauth_api/src/service"
	"github.com/sunil206b/oauth_api/src/utils/errors"
	"net/http"
	"strings"
)

type LoginController struct {
	ls service.ILoginService
}


func NewLoginController(db *sql.DB) *LoginController {
	return &LoginController{
		ls: service.NewLoginService(db),
	}
}

func (lc *LoginController) GetById(c *gin.Context) {
	accessTokenId := strings.TrimSpace(c.Param("access_token"))
	if accessTokenId == "" {
		errMsg := errors.NewBadRequest("invalid access token id")
		c.JSON(errMsg.StatusCode, errMsg)
		return
	}
	accessToken, errMsg := lc.ls.GetById(accessTokenId)
	if errMsg != nil {
		c.JSON(errMsg.StatusCode, errMsg)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (lc *LoginController) CreateToken(c *gin.Context) {
	var userDTO dto.LoginUserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		errMsg := errors.NewBadRequest("invalid json body")
		c.JSON(errMsg.StatusCode, errMsg)
		return
	}
	at, errMsg := lc.ls.CreateToken(&userDTO)
	if errMsg != nil {
		c.JSON(errMsg.StatusCode, errMsg)
		return
	}
	c.JSON(http.StatusCreated, at)
}