package domain

import (
	"fmt"
	"time"
)

type Expense struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Amount      float32   `json:"amount"`
	Date        time.Time `json:"date"`
}

func NewExpense(description string, amount float32) *Expense {
	return &Expense{
		Description: description,
		Amount:      amount,
		Date:        time.Now(),
	}
}

func (e *Expense) String() string {
	return fmt.Sprintf("ID: %d, Date: %s, Description: %s, Amount: %f", e.ID, e.Date.Format(time.RFC3339), e.Description, e.Amount)
}
