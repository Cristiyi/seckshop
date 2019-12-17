package utils

import (
	"time"
	"github.com/kataras/iris/v12/sessions"
)

var (
	cookieNameForSessionID = "usersess"
	UserSess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID, Expires: time.Hour*time.Duration(1), AllowReclaim: true})
)
