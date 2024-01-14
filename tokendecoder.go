package tokendecoder

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	tokendecodererror "github.com/nicelogic/tokendecoder/error"
	"github.com/nicelogic/tokendecoder/model"
	"github.com/nicelogic/tokendecoder/variable"
)

type TokenDecoder struct {
	publicKey *rsa.PublicKey
}

func (decoder *TokenDecoder) Init(publicKeyFilePath string) error {
	publicKey, err := os.ReadFile(publicKeyFilePath)
	if err != nil {
		return fmt.Errorf("read public key file error(%v)", err)
	}
	decoder.publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return fmt.Errorf("ParseRSAPublicKeyFromPEM error(%v)", err)
	}
	return nil
}

func (decoder *TokenDecoder) Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			reqToken := request.Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer ")
			var ctx context.Context
			if len(splitToken) != 2 {
				ctx = context.WithValue(request.Context(), variable.ErrorCtxKey, tokendecodererror.TokenDecoderError{Err: tokendecodererror.TokenInvalid})
			} else {
				jwtToken := splitToken[1]
				user, err := decoder.UserFromToken(jwtToken)
				if err != nil {
					ctx = context.WithValue(request.Context(), variable.ErrorCtxKey, fmt.Errorf("userFromToken.eror(%w)", err))
				} else {
					ctx = context.WithValue(request.Context(), variable.UserCtxKey, user)
				}
			}
			request = request.WithContext(ctx)
			next.ServeHTTP(writer, request)
		})
	}
}

func (decoder *TokenDecoder) UserFromToken(tokenString string) (*model.User, error) {
	var err error
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("recovered error(%w)", err)
		}
	}()

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method(%v)", token.Header["alg"])
		}
		return decoder.publicKey, nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}
	userInfo, ok := claims[variable.JwtMapCliamsKeyUserInfo].(string)
	if !ok {
		return nil, fmt.Errorf("claims{%v} is not type string", variable.JwtMapCliamsKeyUserInfo)
	}
	var user model.User
	err = json.Unmarshal([]byte(userInfo), &user)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal(userInfo).error(%w)", err)
	}
	return &user, nil
}
