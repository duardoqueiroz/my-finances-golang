package repositories

import "github.com/duardoqueiroz/my-finances-golang/pkg/domain/entities"

type BankAccountRepository interface {
	FindAll() []entities.BankAccount
	FindById(id string) entities.BankAccount
	Create(bankAccount *entities.BankAccount) error
	Update(bankAccount *entities.BankAccount) error
	Delete(id string) error
}
