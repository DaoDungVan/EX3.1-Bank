package main

import (
	"Bank/services"
	"fmt"
)

func main() {
	var username string
	var password string

	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("Password: ")
	fmt.Scanln(&password)

	account, err := services.Login(username, password)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	for {
		fmt.Println("\n1. Check balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Transfer")
		fmt.Println("5. Change password")
		fmt.Println("6. Find account")
		fmt.Println("7. Exit")

		var choice int
		fmt.Print("Choose: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Printf("Balance: %.2f\n", account.Balance)

		case 2:
			var amount float64
			fmt.Print("Deposit amount: ")
			fmt.Scanln(&amount)
			if err := services.Deposit(account, amount); err != nil {
				fmt.Println("ERROR:", err)
				continue
			}
			fmt.Println("Deposit successful.")

		case 3:
			var amount float64
			fmt.Print("Withdraw amount: ")
			fmt.Scanln(&amount)
			if err := services.Withdraw(account, amount); err != nil {
				fmt.Println("ERROR:", err)
				continue
			}
			fmt.Println("Withdraw successful.")

		case 4:
			var toUser string
			var amount float64
			fmt.Print("Receiver account: ")
			fmt.Scanln(&toUser)
			fmt.Print("Transfer amount: ")
			fmt.Scanln(&amount)
			if err := services.Transfer(account, toUser, amount); err != nil {
				fmt.Println("ERROR:", err)
				continue
			}
			fmt.Println("Transfer successful.")

		case 5:
			var newPassword string
			fmt.Print("New password: ")
			fmt.Scanln(&newPassword)
			if err := services.ChangePassword(account, newPassword); err != nil {
				fmt.Println("ERROR:", err)
				continue
			}
			fmt.Println("Password changed successfully.")

		case 6:
			var findUser string
			fmt.Print("Username to find: ")
			fmt.Scanln(&findUser)
			found, err := services.FindAccount(findUser)
			if err != nil {
				fmt.Println("ERROR:", err)
				continue
			}
			fmt.Println("Account found:", found.Username)

		case 7:
			fmt.Println("Program ended.")
			return

		default:
			fmt.Println("Invalid option.")
		}
	}
}
