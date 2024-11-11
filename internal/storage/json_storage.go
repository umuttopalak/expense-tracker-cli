package storage

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/umuttopalak/expense-tracker-cli/internal/domain"
	"github.com/umuttopalak/expense-tracker-cli/internal/repository"
)

var expenseNotFound = errors.New("Expense Not Found")

type JSONExpenseRepository struct {
	FilePath string
	expenses []*domain.Expense
}

func (j *JSONExpenseRepository) AddExpense(expense *domain.Expense) error {
	if err := j.loadExpenses(); err != nil {
		return err
	}

	j.expenses = append(j.expenses, expense)
	return j.saveExpenses()
}

func (j *JSONExpenseRepository) DeleteExpenseByID(id int) error {
	if err := j.loadExpenses(); err != nil {
		return err
	}
	for i, exp := range j.expenses {
		if exp.ID == id {
			j.expenses = append(j.expenses[:i], j.expenses[i+1:]...)
			return j.saveExpenses()
		}
	}
	return expenseNotFound
}

// GetExpense implements repository.ExpenseService.
func (j *JSONExpenseRepository) GetExpense(id int) (*domain.Expense, error) {
	if err := j.loadExpenses(); err != nil {
		return nil, err
	}
	for _, exp := range j.expenses {
		if exp.ID == id {
			return exp, nil
		}
	}
	return nil, expenseNotFound
}

// ListAllExpense implements repository.ExpenseService.
func (j *JSONExpenseRepository) ListAllExpense() ([]*domain.Expense, error) {
	if err := j.loadExpenses(); err != nil {
		return nil, err
	}
	return j.expenses, nil
}

// ListAllExpenseByFilter implements repository.ExpenseService.
func (j *JSONExpenseRepository) ListAllExpenseByFilter(filter string) ([]*domain.Expense, error) {
	if err := j.loadExpenses(); err != nil {
		return nil, err
	}
	var filteredExpenses []*domain.Expense
	for _, exp := range j.expenses {
		if exp.Description == filter {
			filteredExpenses = append(filteredExpenses, exp)
		}
	}
	return filteredExpenses, nil
}

// UpdateExpense implements repository.ExpenseService.
func (j *JSONExpenseRepository) UpdateExpense(id int, description string, amount float32) (string, error) {
	if err := j.loadExpenses(); err != nil {
		return "", err
	}
	for _, exp := range j.expenses {
		if exp.ID == id {
			exp.Description = description
			exp.Amount = amount
			if err := j.saveExpenses(); err != nil {
				return "", err
			}
			return "Expense updated successfully", nil
		}
	}
	return "", expenseNotFound
}

// loadExpenses loads the expenses from the JSON file into memory.
func (j *JSONExpenseRepository) loadExpenses() error {
	file, err := os.Open(j.FilePath)
	if os.IsNotExist(err) {
		return nil // Dosya yoksa, baştan oluşturacağız
	} else if err != nil {
		return err
	}
	defer file.Close()

	var expenses []*domain.Expense
	if err := json.NewDecoder(file).Decode(&expenses); err != nil {
		return err
	}
	j.expenses = expenses
	return nil
}

// saveExpenses saves the expenses to the JSON file.
func (j *JSONExpenseRepository) saveExpenses() error {
	file, err := os.Create(j.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(j.expenses)
}

func NewJSONExpenseRepository(filePath string) repository.ExpenseService {
	return &JSONExpenseRepository{
		FilePath: filePath,
		expenses: []*domain.Expense{},
	}
}
