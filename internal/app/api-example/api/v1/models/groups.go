package models

type CreateGroupResponse struct {
	ID string `json:"id"`
}

type ListGroupsResponse []*GroupResponse

type GroupResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GroupDetailResponse struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Users []string `json:"users"`
}
