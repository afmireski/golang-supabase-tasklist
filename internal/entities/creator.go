package entities

type Creator struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewCreator(id string, name string, email string) *Creator {
	return &Creator{
		Id:    id,
		Name:  name,
		Email: email,
	}
}
