package users

type User struct {
	Name   string `bson:"name"`
	Gender string `bson:"gender"`
	Age    int    `bson:"age"`
}

type UserRequestDTO struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

func (u *User) RequestDtoToObject(requestDTO UserRequestDTO) {
	u.Name = requestDTO.Name
	u.Gender = requestDTO.Gender
	u.Age = requestDTO.Age
}
