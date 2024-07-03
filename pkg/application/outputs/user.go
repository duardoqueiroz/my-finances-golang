package outputs

type CreateUserOutput struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

type FindUserByIDOutput struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	CPF   string `json:"cpf"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
}

type FindAllUserOutput struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	CPF   string `json:"cpf"`
	Phone string `json:"phone"`
}
