package inputs

type UpdateUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	CPF   string `json:"cpf"`
	Phone string `json:"phone"`
}
