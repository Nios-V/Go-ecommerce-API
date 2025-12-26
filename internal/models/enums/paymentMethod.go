package enums

type PaymentMethod string

const (
	PaymentMethodCreditCard     PaymentMethod = "CREDIT_CARD"
	PaymentMethodPayPal         PaymentMethod = "PAYPAL"
	PaymentMethodBankTransfer   PaymentMethod = "BANK_TRANSFER"
	PaymentMethodCashOnDelivery PaymentMethod = "CASH_ON_DELIVERY"
)
