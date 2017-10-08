package database

import "time"

type Order struct {
	ID         int64     `json:"id"`
	UUID       string    `json:"uuid"`
	CustomerID int64     `json:"customer"`
	PaymentID  int64     `json:"payment"`
	ShipAdress string    `json:"shipadress"`
	Paid       bool      `json:"paid"`
	Fulfilled  bool      `json:"fulfilled"`
	TimePlaced time.Time `json:"timestamp"`
}

type Payment struct {
	ID          int64 `json:"id"`
	UUID        int64 `json:"uuid"`
	PaymentType int64 `json:"paymenttype"`
	Allowed     bool  `json:"allowed"`
}

func GetPayment() (p Payment, err error) {

	return
}

type OrderDetails struct {
	ID        int64   `json:"id"`
	UUID      int64   `json:"uuid"`
	OrderID   int64   `json:"order"`
	ProductID int64   `json:"product"`
	Quantity  int64   `json:"quantity"`
	Price     float64 `json:"price"`
	Total     float64 `json:"total"`
}

type CartItem struct {
	ProductID int64 `json:"productid"`
	Quantity  int   `json:"quantity"`
}

type Cart struct {
	Items         []CartItem `json:"items"`
	PaymentMethod int        `json:"paymentid"`
}

func AddOrder(userid string, productid string, paymentid string, units int64) (err error) {

	return
}
