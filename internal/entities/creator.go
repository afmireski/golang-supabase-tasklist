package entities

type Creator struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewCreator(id int32, name string, email string) *Creator {
	return &Creator{
		Id:    id,
		Name:  name,
		Email: email,
	}
}
