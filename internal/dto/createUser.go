package dto


type CreateUser struct{
	Username string `json:"username"`
	Age int `json:"age"`
	Hobbies []string `json:"hobbies"`
}
