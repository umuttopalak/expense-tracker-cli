package repository

import "github.com/umuttopalak/expense-tracker-cli/internal/domain"

type ExpenseService interface {
	AddExpense(expense *domain.Expense) error
	UpdateExpense(id int, description string, amount float32) (string, error)
	DeleteExpenseByID(id int) error
	GetExpense(id int) (*domain.Expense, error)
	ListAllExpense() ([]*domain.Expense, error)
	ListAllExpenseByFilter(filter string) ([]*domain.Expense, error)
}
