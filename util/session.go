package util

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

var (
	key          = []byte(GlobalConfig.SESSION_SECRET)
	SESSION_ID   = "id"
	SessionStore = sessions.NewCookieStore(key)
)

func GetSessionByName(ctx echo.Context, name string) interface{} {
	session, _ := SessionStore.Get(ctx.Request(), SESSION_ID)
	return session.Values[name]
}
