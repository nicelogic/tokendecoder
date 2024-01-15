package tokendecoder

import (
	"context"
	"errors"
	tokendecodererror "github.com/nicelogic/tokendecoder/error"

	"github.com/nicelogic/tokendecoder/model"
	"github.com/nicelogic/tokendecoder/variable"
)

func GetUser(ctx context.Context) (*model.User, error) {
	userInfo, _ := ctx.Value(variable.UserCtxKey).(*model.User)
	err, _ := ctx.Value(variable.ErrorCtxKey).(error)
	return userInfo, err
}

func IsTokenExpired(err error) bool {
	return errors.Is(err, tokendecodererror.TokenDecoderError{Err: tokendecodererror.TokenExpired})
}
func IsTokenInvalid(err error) bool {
	return errors.Is(err, tokendecodererror.TokenDecoderError{Err: tokendecodererror.TokenInvalid})
}
