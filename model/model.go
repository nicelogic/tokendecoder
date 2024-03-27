package model

import "encoding/json"

type User struct {
	Id    string   `json:"id"`
	Roles []string `json:"roles"`
	Token string   `json:"token"`
}

func (user *User) ToJson() (string, error) {
	userJson, err := json.Marshal(*user)
	return string(userJson), err
}
func (user *User) FromJson(userJson string) error {
	return json.Unmarshal([]byte(userJson), user)
}

type ContextKey struct {
	Name string
}
