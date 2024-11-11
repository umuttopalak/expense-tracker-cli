package domain

type ExpenseService interface {
	AddExpense(description string, amount float32) (*Expense, error)
	UpdateExpense(id int, description string, amount float32) (string, error)
	DeleteExpenseByID(id int) error
	GetExpense(id int) (*Expense, error)
	ListAllExpense() ([]*Expense, error)
	ListAllExpenseByFilter(filter string) ([]*Expense, error)
}
