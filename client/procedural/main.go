package main

import (
    "fmt"
    "strings"
    "github.com/vasubabu/oops/payment/procedural"
)

func main() {
    const amount = 500
    fmt.Println("paying with cash")
    procedural.PayWithCash(amount)
    fmt.Printf(strings.Repeat("*", 30) + "\n\n");

    credit := &procedural.CreditCard {
        OwnerName:          "Axis Bank",
        CardNumber:         "1111-2222-3333-4444",
        ExpirationMonth:    5,
        ExpirationYear:     23,
        SecurityCode:       123,
        AvaliableCredit:    5000,
    }

    fmt.Println("Paying with creditcard")
    fmt.Printf("Initial Balance: $%.2f\n", credit.AvaliableCredit)
    procedural.PayWithCredit(credit, amount)
    fmt.Printf("Balance now: $%.2f\n", credit.AvaliableCredit)
    fmt.Printf(strings.Repeat("*", 30) + "\n\n");

    checking := &procedural.CheckingAmount {
        AccountOwner:       "Axis bank",
        RoutingNumber:      "01010011",
        AccountNUmber:      "1234567891234567",
        Balance:            250,
    }

    fmt.Println("Paying with check")
    fmt.Printf("Initial Balance: $%.2f\n", checking.Balance)
    procedural.PayWithCheck(checking, amount)
    fmt.Printf("Balance now: $%.2f\n", checking.Balance)
    fmt.Printf(strings.Repeat("*", 30) + "\n\n");

    fmt.Printf("Hmm, not enough in the account. We can fix that!")
    fmt.Println("Changing account balance")
    checking.Balance = 5000

    fmt.Println("Paying with check")
    fmt.Printf("Initial Balance: $%.2f\n", checking.Balance)
    procedural.PayWithCheck(checking, amount)
    fmt.Printf("Balance now: $%.2f\n", checking.Balance)
    fmt.Printf(strings.Repeat("*", 30) + "\n\n");
}
