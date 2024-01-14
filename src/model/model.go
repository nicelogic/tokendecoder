package model

type User struct {
	Id    string   `json:"id"`
	Roles []string `json:"roles"`
}

type ContextKey struct {
	Name string
}

