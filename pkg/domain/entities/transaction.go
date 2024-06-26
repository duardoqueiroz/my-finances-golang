package entities

import (
	"fmt"
	"time"

	"github.com/duardoqueiroz/my-finances-golang/pkg/domain/keyvalues"
	"github.com/google/uuid"
)

type Transaction struct {
	id                string
	value             *keyvalues.TransactionValue
	date              *keyvalues.Date
	description       *keyvalues.Description
	paymentMethod     *PaymentMethod
	paymentConditions *PaymentCondition
	installments      *keyvalues.Installments
	recurrence        *keyvalues.Recurrence
}

func NewTransaction(value float64, date time.Time, description, paymentMethod, paymentConditions string, installments, recurrence int) (*Transaction, error) {
	transactionValue, err := keyvalues.NewTransactionValue(value)
	if err != nil {
		return nil, fmt.Errorf("error creating transaction: %w", err)
	}

	dateValue, err := keyvalues.NewDate(date)
	if err != nil {
		return nil, fmt.Errorf("error creating transaction: %w", err)
	}

	descriptionValue, err := keyvalues.NewDescription(description)
	if err != nil {
		return nil, fmt.Errorf("error creating transaction: %w", err)
	}

	paymentMethodValue, err := NewPaymentMethod(paymentMethod)
	if err != nil {
		return nil, fmt.Errorf("error creating transaction: %w", err)
	}

	paymentConditionsValue, err := NewPaymentConditions(paymentConditions)
	if err != nil {
		return nil, fmt.Errorf("error creating transaction: %w", err)
	}

	installmentsValue, err := keyvalues.NewInstallments(installments)
	if err != nil {
		return nil, fmt.Errorf("error creating transaction: %w", err)
	}

	recurrenceValue, err := keyvalues.NewRecurrence(recurrence)
	if err != nil {
		return nil, fmt.Errorf("error creating transaction: %w", err)
	}

	id := uuid.New().String()

	return &Transaction{
		id:                id,
		value:             transactionValue,
		date:              dateValue,
		description:       descriptionValue,
		paymentMethod:     paymentMethodValue,
		paymentConditions: paymentConditionsValue,
		installments:      installmentsValue,
		recurrence:        recurrenceValue,
	}, nil
}

func (t Transaction) Id() string {
	return t.id
}

func (t Transaction) Description() string {
	return t.description.Value()
}

func (t Transaction) Date() time.Time {
	return t.date.Value()
}

func (t Transaction) Value() float64 {
	return t.value.Value()
}

func (t Transaction) PaymentMethod() *PaymentMethod {
	return t.paymentMethod
}

func (t Transaction) PaymentConditions() *PaymentConditions {
	return t.paymentConditions
}

func (t Transaction) Recurrence() int {
	return t.recurrence.Value()
}

func (t Transaction) Installments() int {
	return t.installments.Value()
}
