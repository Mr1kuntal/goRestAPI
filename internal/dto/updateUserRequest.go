package dto

type UpdateUserRequest struct {
	Username *string `json:"username"`
	Age *int `json:"age"`
	Hobbies *[]string `json:"hobbies"`
}