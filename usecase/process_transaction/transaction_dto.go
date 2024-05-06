package process_transaction

type TransactionalDTOInput struct {
	ID	   string `json:"id"`
	AccontId string `json:"account_id"`
	TestCreditCardNumber string `json:"test_credit_card_number"`
	CreditCardName string `json:"credit_card_name"`
	CreditCardExpirationMonth int `json:"credit_card_expiration_month"`
	CreditCardExpirationYear int `json:"credit_card_expiration_year"`
	CreditCardCvv int `json:"credit_card_cvv"`
	Amount float64 `json:"amount"`

}

type TransactionalDTOOutput struct {
	ID	   string `json:"id"`
	Status string `json:"status"`
	ErrorMessage string `json:"error_message"`
}