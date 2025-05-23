package main

import (
	"context"
	"database/sql"
	"expenseTracker/internal/db"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

type Expense struct {
	id          int
	Description string
	Amount      int
}

func main() {
	ctx := context.Background()
	var rootCommand = &cobra.Command{}
	var Description, amount, id, month string
	// var listExpense string

	connStr := "password=root user=postgres dbname=postgres sslmode=disable"
	dbConn, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	defer dbConn.Close()
	queries := db.New(dbConn)

	var cmd = &cobra.Command{
		Use:   "criar",
		Short: "Adicione uma despesa",
		Run: func(cmd *cobra.Command, args []string) {

			if Description == "" {
				fmt.Println("Você precisa digitar uma descrição")
				return
			}

			if amount == "" {
				fmt.Println(" Você precisa digitar um valor")
				return
			}

			_, err := strconv.Atoi(amount)
			if err != nil {
				fmt.Println("Você precisa digitar um valor válido")
				return
			}

			description := Description

			expenseParams := db.CreateExpenseParams{
				Description: description,
				Amount:      amount,
			}

			expenseCreate, err := queries.CreateExpense(ctx, expenseParams)

			if err != nil {
				fmt.Println("Erro ao criar o banco de dados")
			}

			fmt.Printf("Despesa adicionada com sucesso (ID: %v)", expenseCreate.ID)
		},
	}

	var cmd2 = &cobra.Command{
		Use:   "listar",
		Short: "Listar despesas",
		Run: func(cmd *cobra.Command, args []string) {
			listExpense, err := queries.ListExpense(ctx)

			if err != nil {
				fmt.Println("Erro ao listar despesas")
			}

			fmt.Println("ID       DESCRIPTION       AMOUNT")
			for _, value := range listExpense {
				fmt.Printf("%v       %s       %v\n", value.ID, value.Description, value.Amount)
			}
		},
	}

	var cmd3 = &cobra.Command{
		Use:   "total",
		Short: "Total de despesas",
		Run: func(cmd *cobra.Command, args []string) {
			listExpense, err := queries.ListExpense(ctx)

			if err != nil {
				fmt.Println("Erro ao listar despesas")
			}

			var countTracker float64
			for _, value := range listExpense {
				amountConvert, err := strconv.ParseFloat(value.Amount, 64)

				if err != nil {
					fmt.Println("Erro ao converter valor %v", err)
				}

				countTracker = countTracker + amountConvert
			}

			fmt.Printf("Total de despesas: %v", countTracker)
		},
	}

	var cmd4 = &cobra.Command{
		Use:   "excluir",
		Short: "Excluir despesas",
		Run: func(cmd *cobra.Command, args []string) {
			idConvert, err := strconv.Atoi(id)

			if err != nil {
				fmt.Println("Erro ao converter id: %v", err)
			}

			err = queries.DeleteExpense(ctx, int32(idConvert))

			if err != nil {
				fmt.Println("Erro ao excluir despesa %v", err)
			}

			fmt.Println("Despesa excluída com sucesso")
		},
	}

	var cmd5 = &cobra.Command{
		Use:   "mes",
		Short: "Total de despesas do mês",
		Run: func(cmd *cobra.Command, args []string) {
			listExpense, err := queries.ListExpense(ctx)

			if err != nil {
				fmt.Println("Erro ao listar despesas")
			}

			monthInt, err := strconv.Atoi(month)
			if err != nil {
				fmt.Println("Erro ao converter mês: %v", err)
			}

			if monthInt < 1 || monthInt > 12 {
				fmt.Println("Mês inválido. O mês deve estar entre 1 e 12")
				return
			}

			monthIntTime := time.Month(monthInt)

			var countTracker float64
			for _, value := range listExpense {
				expenseMonth := value.CreatedAt.Month()

				if expenseMonth == monthIntTime {
					amountConvert, err := strconv.ParseFloat(value.Amount, 64)

					if err != nil {
						fmt.Println("Erro ao converter valor %v", err)
					}

					countTracker = countTracker + amountConvert
				}

				if countTracker == 0 {
					fmt.Printf("Não há despesas para o mês %v", monthInt)
				} else {
					fmt.Printf("Total de despesas do mês %v: %v", countTracker)
				}
			}

			fmt.Printf("Total de despesas do mês: %v", countTracker)
		},
	}

	cmd.Flags().StringVarP(&Description, "descrição", "d", "", "Descrição da despesa")
	cmd.Flags().StringVarP(&amount, "valor", "v", "", "Valor da despesa")
	cmd4.Flags().StringVarP(&id, "id", "i", "", "id da despesa")
	cmd5.Flags().StringVarP(&id, "mes", "m", "", "mês da despesa")

	rootCommand.AddCommand(cmd, cmd2, cmd3, cmd4)
	rootCommand.Execute()
}
