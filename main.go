package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("UCCU...")
	fmt.Println("Reminder: Omit dollar signs")
	var sharedSavings float64
	fmt.Printf("Shared Savings available balance: ")
	fmt.Scanln(&sharedSavings)
	var sharedChecking float64
	fmt.Printf("Shared Checking available balance: ")
	fmt.Scanln(&sharedChecking)
	var sharedOverdraft float64
	fmt.Printf("Shared Overdraft balance: ")
	fmt.Scanln(&sharedOverdraft)
	fmt.Println()
	fmt.Println("Bank Balances as of", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("\nUCCU Shared Account:")
	fmt.Printf("$%.0f Shared Savings (1211627-0)\n", sharedSavings)
	fmt.Printf("$%.0f Shared Checking (1211627-9)\n", sharedChecking)
	if sharedOverdraft > 0 {
		fmt.Printf("$%.0f Shared Overdraft (1211627-99)", sharedOverdraft)
	}
	fmt.Println()
}
