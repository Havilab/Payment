package model

type Payment struct {
	PaymentId      int32  `json:"paymentid"gorm:"primarykey"`
	UserID         int32  `json:"userid"`
	CardHolderName string `json:"cardholdername"`
	CardNumber     string `json:"cardnumber"`
	ExpiryDate     string `json:"expiredate"`
	Cvv            string `json:"cvv"`
	CreateDate     string `json:"createdate"`
}
