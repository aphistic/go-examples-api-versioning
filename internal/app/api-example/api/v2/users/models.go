package users

type ListUsersResponse []*UserResponse

type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email_address"`
}

type UserDetailResponse struct {
	ID    string `json:"id"`
	Email string `json:"email_address"`
	Name  string `json:"name"`
}
