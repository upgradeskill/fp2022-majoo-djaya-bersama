package util

import (
	"github.com/gorilla/sessions"
)

var (
	key          = []byte(GlobalConfig.SESSION_SECRET)
	SESSION_ID   = "id"
	SessionStore = sessions.NewCookieStore(key)
)
