package database

import (
	"time"

	"github.com/twinj/uuid"
)

type Order struct {
	ID             int64     `json:"id"`
	UUID           string    `json:"uuid"`
	ProductID      int64     `json:"productid"`
	CustomerID     int64     `json:"customer"`
	PaymentID      int64     `json:"payment"`
	ShippingAdress string    `json:"shippingadress"`
	Paid           bool      `json:"paid"`
	Fulfilled      bool      `json:"fulfilled"`
	TimePlaced     time.Time `json:"timestamp"`
	Price          float64   `json:"price"`
	Quantity       int       `json:"quantity"`
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

type CartItem struct {
	ProductID int64 `json:"productid"`
	Quantity  int   `json:"quantity"`
}

type Cart struct {
	Items          []CartItem `json:"items"`
	PaymentMethod  int        `json:"paymentid"`
	ShippingAdress string     `json:"shippingadress"`
}

func AddOrder(customerid int64, cart Cart) (orders []Order, err error) {
	// loop all the elemets of the cart and write them to aorder
	for _, item := range cart.Items {
		uuid4 := uuid.NewV4()
		_, err = DB.Exec(`insert into orders(uuid , customerid, paymentid,
				shippingadress ,quantity , productid) values(?,?,?,?,?,?)`, uuid4,
			1, cart.PaymentMethod, cart.ShippingAdress, item.Quantity, item.ProductID)
		if err != nil {
			return
		}

	}
	rows, err := DB.Query(`select id, uuid,productid ,customerid,paymentid,shippingadress,price,quantity  from orders where customerid = ?`, 1)
	if err != nil {
		return
	}

	for rows.Next() {
		order := new(Order)

		if err = rows.Scan(&order.ID, &order.UUID, &order.ProductID, &order.CustomerID, &order.PaymentID, &order.ShippingAdress, &order.Price, &order.Quantity); err != nil {
			return
		}
		orders = append(orders, *order)
	}
	return
}
