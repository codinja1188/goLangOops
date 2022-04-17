package procedural

type CheckingAmount struct {
    AccountOwner    string
    RoutingNumber   string
    AccountNUmber   string
    Balance         float32
}

func PayWithCheck(account * CheckingAmount, amount float32) bool {
    if account.Balance > amount {
        account.Balance -= amount
        return true
    } else {
        return false
    }
}
