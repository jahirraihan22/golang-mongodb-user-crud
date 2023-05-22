package users

type User struct {
	Name     string `bson:"name"`
	Gender   string `bson:"gender"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Age      int    `bson:"age"`
}

type UserRequestDTO struct {
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

func (u *User) RequestDtoToObject(requestDTO UserRequestDTO) {
	u.Name = requestDTO.Name
	u.Gender = requestDTO.Gender
	u.Email = requestDTO.Email
	u.Password = requestDTO.Password
	u.Age = requestDTO.Age
}

type UserResponse struct {
	ID       string `bson:"_id"`
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Gender   string `bson:"gender"`
	Password string `bson:"password"`
	Age      int    `bson:"age"`
}

type UserAuthData struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

type UserAuthDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
