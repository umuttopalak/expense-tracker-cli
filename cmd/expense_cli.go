package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/umuttopalak/expense-tracker-cli/internal/cli"
	"github.com/umuttopalak/expense-tracker-cli/internal/service"
	"github.com/umuttopalak/expense-tracker-cli/internal/storage"
)

func Execute() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error: Unable to get the home directory")
		os.Exit(1)
	}

	expensesFilePath := filepath.Join(homeDir, ".expense-tracker.json")

	repo := storage.NewJSONExpenseRepository(expensesFilePath)

	expenseService := service.NewExpenseService(repo)

	cliHandler := cli.NewCLIHandler(expenseService)

	cliHandler.Run(os.Args)
}
