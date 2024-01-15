package variable

import "github.com/nicelogic/tokendecoder/model"

var (
	JwtMapCliamsKeyUserInfo = "user_info"
	UserCtxKey  = &model.ContextKey{Name: "user"}
	ErrorCtxKey = &model.ContextKey{Name: "error"}
)

const (
	RoleNormal = "normal"
	RoleAnonymous = "anonymous"
)
