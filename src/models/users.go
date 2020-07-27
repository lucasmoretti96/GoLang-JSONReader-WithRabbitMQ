package models

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Username  string `json:"username"`
	AccountId string `json:"accountId"`
}
