package usecases

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	CPF      string `json:"cpf"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type CreateUserOutput struct {
	ID string
}

type UpdateUserInput struct {
	Name  string
	Email string
	CPF   string
	Phone string
}

type FindAllUserOutput struct {
	ID    string
	Name  string
	Email string
	CPF   string
	Phone string
}

type FindUserByIDOutput struct {
	ID    string
	Name  string
	Email string
	CPF   string
	Phone string
}

type UserUseCase interface {
	Create(input CreateUserInput) (*CreateUserOutput, error)
	Update(id string, input UpdateUserInput) error
	FindAll() ([]FindAllUserOutput, error)
	FindByID(id string) (*FindUserByIDOutput, error)
	Delete(id string) error
}
