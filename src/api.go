package tokendecoder

import (
	"context"

	"github.com/nicelogic/tokendecoder/model"
	"github.com/nicelogic/tokendecoder/variable"
)

func GetUser(ctx context.Context) (*model.User, error) {
	userInfo, _ := ctx.Value(variable.UserCtxKey).(*model.User)
	err, _ := ctx.Value(variable.ErrorCtxKey).(error)
	return userInfo, err
}
