package main

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

type Expense struct {
	Description string
	Amount      int
}

var expenses []Expense

func main() {
	var rootCommand = &cobra.Command{}
	var Description, amount string
	// var listExpense string

	var cmd = &cobra.Command{
		Use:   "criar",
		Short: "Adicione uma despesa",
		Run: func(cmd *cobra.Command, args []string) {
			// validations
			if Description == "" {
				fmt.Println("Você precisa digitar uma descrição")
				return
			}
			if amount == "" {
				fmt.Println(" Você precisa digitar um valor")
				return
			}

			amountValue, err := strconv.Atoi(amount)
			if err != nil {
				fmt.Println("Você precisa digitar um valor válido")
				return
			}

			add(Description, amountValue)
			fmt.Println("Despesa adicionada com sucesso")
		},
	}

	cmd.Flags().StringVarP(&Description, "descrição", "d", "", "Descrição da despesa")
	cmd.Flags().StringVarP(&amount, "valor", "v", "", "Valor da despesa")

	rootCommand.AddCommand(cmd)
	rootCommand.Execute()
}

func add(description string, amount int) {
	count := Expense{
		Description: description,
		Amount:      amount,
	}

	expenses = append(expenses, count)

	fmt.Println("Descrição: %s", count.Description)
	fmt.Println("Valor: %d", count.Amount)

	println("Despesa salva com sucesso")
}
