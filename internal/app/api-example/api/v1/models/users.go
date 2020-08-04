package models

type CreateUserResponse struct {
	ID string `json:"id"`
}

type ListUsersResponse []*UserResponse

type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type UserDetailResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
