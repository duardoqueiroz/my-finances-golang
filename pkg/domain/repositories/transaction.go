package repositories

import "github.com/duardoqueiroz/my-finances-golang/pkg/domain/entities"

type TransactionRepository interface {
	FindAll() []entities.Transaction
	FindById(id string) entities.Transaction
	Create(transaction *entities.Transaction) error
	Update(transaction *entities.BankAccount) error
	Delete(id string) error
}
