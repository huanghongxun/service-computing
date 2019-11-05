package main

import (
	"github.com/gin-gonic/gin"
	"github.com/huanghongxun/cloudgo-io/errors"
	"github.com/huanghongxun/cloudgo-io/web"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = make([]User, 0)

func JsonHandler(c *gin.Context) {
	web.ResSuccessJSON(c, users)
}

func UnknownHandler(c *gin.Context) {
	web.ResError(c, errors.ErrNotImplemented)
}

func FormHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	users = append(users, User{
		Username: username,
		Password: password,
	})

	c.Redirect(http.StatusFound, "/user-list.html")
}
