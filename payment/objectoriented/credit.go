package objectoriented

type CreditCard struct {
    ownerName           string
    cardNumber          string
    expirationMonth     int
    expirationYear      int
    securityCode        int
    avaliableCredit     float32
}

func CreateCreditAccount(ownerName, cardNumber string, expirationMonth, expirationYear, securityCode int) *CreditCard {
    return &CreditCard {
        ownerName:  ownerName,
        cardNumber: cardNumber,
        expirationMonth:    expirationMonth,
        expirationYear:     expirationYear,
        securityCode:       securityCode,
        avaliableCredit:    5000,
    }
}

func (c CreditCard) ProcessPayment(amount float32) bool {
    return true
}

func (c CreditCard) AvaliableCredit() float32 {
    return c.avaliableCredit
}
