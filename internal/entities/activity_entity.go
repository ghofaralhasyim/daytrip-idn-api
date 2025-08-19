package entities

type ActivityEntity struct {
	Id    int64
	Name  string
	Image string
}

func MakeActivityEntity(
	id int64,
	name, image string,
) *ActivityEntity {
	return &ActivityEntity{
		Id:    id,
		Name:  name,
		Image: image,
	}
}
