package cli

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/umuttopalak/expense-tracker-cli/internal/domain"
)

type CLIHandler struct {
	expenseService domain.ExpenseService
}

func NewCLIHandler(expenseService domain.ExpenseService) *CLIHandler {
	return &CLIHandler{
		expenseService: expenseService,
	}
}

func (h *CLIHandler) Run(args []string) {
	if len(args) < 2 {
		h.printUsage()
		return
	}

	command := args[1]
	switch command {
	case "add":
		if err := h.handleAddExpense(args); err != nil {
			fmt.Println("Error adding expense:", err)
		}
	case "get":
		if err := h.handleGetExpense(args); err != nil {
			fmt.Println("Error getting expense:", err)
		}
	case "update":
		if err := h.handleUpdateExpense(args); err != nil {
			fmt.Println("Error updating expense:", err)
		}
	case "delete":
		if err := h.handleDeleteExpenseByID(args); err != nil {
			fmt.Println("Error deleting expense:", err)
		}
	case "list":
		if err := h.handleListAllExpense(args); err != nil {
			fmt.Println("Error listing expenses:", err)
		}
	case "list-filter":
		if err := h.handleListAllExpenseByFilter(args); err != nil {
			fmt.Println("Error listing expenses by filter:", err)
		}
	default:
		fmt.Println("Unknown command:", command)
		h.printUsage()
	}
}

func (c *CLIHandler) handleAddExpense(args []string) error {
	description := ""
	amount := 0.0
	for i, arg := range args {
		switch arg {
		case "--description":
			if i+1 < len(args) {
				description = args[i+1]
			}
		case "--amount":
			if i+1 < len(args) {
				amountVal, err := strconv.ParseFloat(args[i+1], 64)
				if err != nil {
					return errors.New("invalid amount")
				}
				amount = amountVal
			}
		}
	}
	if description == "" || amount <= 0 {
		return errors.New("description and amount are required")
	}

	expense, err := c.expenseService.AddExpense(description, float32(amount))
	if err != nil {
		return err
	}

	fmt.Printf("Expense added successfully (ID: %d)\n", expense.ID)
	return nil
}

func (c *CLIHandler) handleUpdateExpense(args []string) error {
	if len(args) < 4 {
		return errors.New("update requires an ID, description, and amount")
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		return errors.New("invalid ID")
	}

	description := ""
	amount := 0.0
	for i, arg := range args[3:] {
		switch arg {
		case "--description":
			if i+1 < len(args[3:]) {
				description = args[i+4]
			}
		case "--amount":
			if i+1 < len(args[3:]) {
				amountVal, err := strconv.ParseFloat(args[i+4], 64)
				if err != nil {
					return errors.New("invalid amount")
				}
				amount = amountVal
			}
		}
	}

	if description == "" || amount <= 0 {
		return errors.New("description and amount are required")
	}

	result, err := c.expenseService.UpdateExpense(id, description, float32(amount))
	if err != nil {
		return err
	}

	fmt.Println("Expense updated successfully:", result)
	return nil
}

func (c *CLIHandler) handleDeleteExpenseByID(args []string) error {
	if len(args) < 3 {
		return errors.New("delete requires an ID")
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		return errors.New("invalid ID")
	}

	err = c.expenseService.DeleteExpenseByID(id)
	if err != nil {
		return err
	}

	fmt.Printf("Expense with ID %d deleted successfully\n", id)
	return nil
}

func (c *CLIHandler) handleGetExpense(args []string) error {
	if len(args) < 3 {
		return errors.New("get requires an ID")
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		return errors.New("invalid ID")
	}

	expense, err := c.expenseService.GetExpense(id)
	if err != nil {
		return err
	}

	fmt.Printf("Expense Details (ID: %d):\nDescription: %s\nAmount: %.2f\n", expense.ID, expense.Description, expense.Amount)
	return nil
}

func (c *CLIHandler) handleListAllExpense(args []string) error {
	expenses, err := c.expenseService.ListAllExpense()
	if err != nil {
		return err
	}

	fmt.Println("All Expenses:")
	for _, expense := range expenses {
		fmt.Printf("ID: %d, Description: %s, Amount: %.2f\n", expense.ID, expense.Description, expense.Amount)
	}
	return nil
}

func (c *CLIHandler) handleListAllExpenseByFilter(args []string) error {
	if len(args) < 3 {
		return errors.New("list-filter requires a filter argument")
	}

	filter := args[2]
	expenses, err := c.expenseService.ListAllExpenseByFilter(filter)
	if err != nil {
		return err
	}

	fmt.Printf("Filtered Expenses (Filter: %s):\n", filter)
	for _, expense := range expenses {
		fmt.Printf("ID: %d, Description: %s, Amount: %.2f\n", expense.ID, expense.Description, expense.Amount)
	}
	return nil
}

func (h *CLIHandler) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  expense-cli add --description '' --amount <float32>")
	fmt.Println("  expense-cli get <ID>")
	fmt.Println("  expense-cli update <ID> --description '' --amount <float32>")
	fmt.Println("  expense-cli delete <ID>")
	fmt.Println("  expense-cli list")
	fmt.Println("  expense-cli list-filter <filter>")
}
