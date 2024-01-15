package variable

import "github.com/nicelogic/tokendecoder/model"

var (
	JwtMapCliamsKeyUser = "user"
	UserCtxKey  = &model.ContextKey{Name: "user"}
	ErrorCtxKey = &model.ContextKey{Name: "error"}
)

const (
	RoleNormal = "normal"
	RoleAnonymous = "anonymous"
)
