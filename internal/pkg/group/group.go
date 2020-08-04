package group

type Group struct {
	ID   string
	Name string
}

type GroupService struct{}

func NewGroupService() *GroupService {
	return &GroupService{}
}

func (us *GroupService) CreateGroup(group *Group) error {
	// Any code we need to create groups
	return nil
}

func (us *GroupService) ListGroups() ([]*Group, error) {
	// Get all our groups, this would probably have some kind of filter or paging on it
	return []*Group{
		{ID: "id", Name: "name"},
	}, nil
}

func (us *GroupService) GetGroupByID(id string) (*Group, error) {
	// Get the group by the ID, probably not a string but that's easiest
	// for this example because it's not really important
	return &Group{ID: "id", Name: "name"}, nil
}
