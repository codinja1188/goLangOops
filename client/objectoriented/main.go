package main

import (
    "fmt"
    "strings"
    "github.com/vasubabu/oops/payment/objectoriented"
)

func main() {
    const amount = 500

    fmt.Println("paying with cash")
    cash := &objectoriented.Cash{}
    cash.ProcessPayment(amount)
    fmt.Printf(strings.Repeat("*", 30) + "\n\n");

    credit := objectoriented.CreateCreditAccount (
        "Axis Bank",
        "1111-2222-3333-4444",
        5,
        2022,
        1234,
    )

    fmt.Println("Paying with creditcard")
    fmt.Printf("Initial Balance: $%.2f\n", credit.AvaliableCredit())
    credit.ProcessPayment(amount)
    fmt.Printf("Balance now: $%.2f\n", credit.AvaliableCredit())
    fmt.Printf(strings.Repeat("*", 30) + "\n\n");

    checking := objectoriented.CreateCheckingAccount(
        "Axis Bank",
        "01010011",
        "1234567891234567")
    fmt.Println("Paying with checking")
    fmt.Printf("Initial Balance: $%.2f\n", checking.Balance())
    checking.ProcessPayment(amount)
    fmt.Printf("Initial Balance: $%.2f\n", checking.Balance())

    fmt.Println("Hmm, not enough in the account. Nothing we can do!")
    fmt.Printf(strings.Repeat("*", 30) + "\n\n");
}
