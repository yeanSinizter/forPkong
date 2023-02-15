package entities

type UserRepositories interface {
	InsertUser(*UserModel) (*UserModel, error)
}

type UserUsecase interface {
	AddUser(*UserModel) (*UserModel, error)
}

func (u *UserModel) Table() string {
	return "user"
}

type UserModel struct {
	Id       int    `json:"id" gorm:"increment"`
	Username string `json:"username"`
	Password string `json:"password"`
}
