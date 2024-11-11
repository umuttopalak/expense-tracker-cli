package service

import (
	"fmt"

	"github.com/umuttopalak/expense-tracker-cli/internal/domain"
	"github.com/umuttopalak/expense-tracker-cli/internal/repository"
)

type ExpenseService struct {
	repo repository.ExpenseService
}

func NewExpenseService(repo repository.ExpenseService) domain.ExpenseService {
	return &ExpenseService{
		repo: repo,
	}
}

func (e *ExpenseService) AddExpense(description string, amount float32) (*domain.Expense, error) {
	expense := domain.NewExpense(description, amount)
	err := e.repo.AddExpense(expense)
	if err != nil {
		return nil, err
	}

	return expense, nil
}

func (e *ExpenseService) DeleteExpenseByID(id int) error {
	if err := e.repo.DeleteExpenseByID(id); err != nil {
		return err
	}

	return nil
}

func (e *ExpenseService) GetExpense(id int) (*domain.Expense, error) {
	expense, err := e.repo.GetExpense(id)
	if err != nil {
		return nil, err
	}
	return expense, nil
}

func (e *ExpenseService) ListAllExpense() ([]*domain.Expense, error) {
	var expenseList []*domain.Expense
	expenseList, err := e.repo.ListAllExpense()
	if err != nil {
		return nil, err
	}
	return expenseList, nil

}

func (e *ExpenseService) ListAllExpenseByFilter(filter string) ([]*domain.Expense, error) {
	panic("unimplemented")
}

func (e *ExpenseService) UpdateExpense(id int, description string, amount float32) (string, error) {
	expense, err := e.repo.UpdateExpense(id, description, amount)
	if err != nil {
		return "", err
	}

	return fmt.Sprint(expense), err
}
