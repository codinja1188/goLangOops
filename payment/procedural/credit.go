package procedural

type CreditCard struct {
    OwnerName           string
    CardNumber          string
    ExpirationMonth     int
    ExpirationYear      int
    SecurityCode        int
    AvaliableCredit     float32
}

func PayWithCredit(card *CreditCard, amount float32) bool {
    if card.AvaliableCredit > amount {
        card.AvaliableCredit -= amount
        return true
    } else {
        return false
    }
}
