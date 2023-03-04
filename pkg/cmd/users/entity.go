package users

type UserListResponse struct {
	Users []User `json:"Users"`
}

type User struct {
	Username    string `json:"Username"`
	DisplayName string `json:"Displayname"`
}
