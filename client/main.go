package main

import (
    "fmt"
    "github.com/vasubabu/oops/payment"
)

func main() {
    credit := payment.CreateCreditAccount(
        "Axis Bank",
        "1111-2222-3333-4444",
        5,
        2022,
        123)

    fmt.Printf("Owner Name: %v\n", credit.OwnerName())
    fmt.Printf("Card Number: %v\n", credit.CardNumber())
    fmt.Println("Trying to change card number")
    err := credit.SetCardNumber("Invalid")
    if err != nil {
        fmt.Printf("That didn't work: %v\n", err)
    }
    fmt.Printf("Avaliable credit : %v\n", credit.AvaliableCredit())
}
