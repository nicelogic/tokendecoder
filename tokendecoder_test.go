package tokendecoder

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nicelogic/tokendecoder/model"
)

func TestUserFromJwt(t *testing.T) {
	decoder := &TokenDecoder{}
	err := decoder.Init("/etc/app-0/secret-jwt/jwt-publickey")
	if err != nil{
		t.Error(err)
	}
	user, err := decoder.UserFromToken("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjkzODU3MjQsInVzZXIiOnsiaWQiOiJ0ZXN0In19.RLaT0gmGtcm7S-4Fl958mzHBvU4yRwhyDtzX0MRT1dLKkL23CAyXqtIAJkUTXeu1a6OTyRswPPv2_K6QkBDErDRmM4R_pWvMtll_xEuYWazeZ-dwQUCHypz4iLMcxNr0VNhbPRMizB5AvXs19-oKUlLmMU0EDu3M7VVKQYMG3TzHrN5T0UdZ2rIho2wnKWOMkUszPcey52z2ui_pQgsMnmtQzAY2KCwugPTE9AoYo1MbfJihdUKW0K0ls3xABFGDP6Z6LR5QM8MVuYvEKtYetfUoalTRCl3_THeiriJtP7CZertx3nwb8ieOljY-ztrYW-k2lp1jIm82dAzq7PVsmw")
	if err != nil {
		if err.Error() == "Token is expired" {
			log.Println("token is expired")
		} else {
			t.Error(err)
		}
	}

	log.Printf("%#v\n", user)
	// if user.Id != "pdKBcAc7lKzSC9nyglCwQ"{
	// 	t.Errorf("userFromJwt want test, but: %s", user.Id)
	// }
}

func TestJwt(t *testing.T) {
	user := model.User{Id: "test"}
	mapClaims := make(jwt.MapClaims)
	mapClaims["user"] = user
	expireTime := time.Now().Add(60 * time.Second).Unix()
	mapClaims["exp"] = expireTime
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, mapClaims)
	privateKey, err := os.ReadFile("/etc/app-0/secret-jwt/jwt-privatekey")
	if err != nil {
		t.Errorf("error reading private key file: %v\n", err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		t.Errorf("error parsing RSA private key: %v\n", err)
	}
	tokenString, err := token.SignedString(key)
	if err != nil {
		t.Errorf("error signing token: %v\n", err)
	}
	log.Println(tokenString)
}
