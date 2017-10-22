package database

import (
	"log"
	"time"

	"github.com/twinj/uuid"
)

type Order struct {
	ID             int64          `json:"id"`
	UUID           string         `json:"uuid"`
	ProductID      int64          `json:"productid"`
	CustomerID     int64          `json:"customer"`
	PaymentID      int64          `json:"payment"`
	ShippingAdress ShippingAdress `json:"shippingadress"`
	Paid           bool           `json:"paid"`
	Fulfilled      bool           `json:"fulfilled"`
	TimePlaced     time.Time      `json:"timestamp"`
	Price          float64        `json:"price"`
	Quantity       int            `json:"quantity"`
}

type Payment struct {
	ID          int64  `json:"id"`
	UUID        string `json:"uuid"`
	PaymentType string `json:"paymenttype"`
	Allowed     bool   `json:"allowed"`
}

type ShippingAdress struct {
	Contry string `json:"contry"`
	Zip    string `json:"zip"`
	Adress string `json:"adress"`
}

func GetPayment(id string) (p Payment, err error) {

	return
}

func GetPayments() (p []Payment, err error) {
	rows, err := DB.Query(`select id, uuid,paymenttype ,allowed from payment`)
	if err != nil {
		return
	}

	for rows.Next() {
		payment := new(Payment)

		if err = rows.Scan(&payment.ID, &payment.UUID, &payment.PaymentType, &payment.Allowed); err != nil {
			return
		}
		p = append(p, *payment)
	}
	return
}

type CartItem struct {
	ProductID int64 `json:"productid"`
	Quantity  int   `json:"quantity"`
}

type Cart struct {
	Items          []CartItem     `json:"items"`
	PaymentMethod  int            `json:"paymentid"`
	ShippingAdress ShippingAdress `json:"shippingadress"`
}

func AddOrder(customerid int64, cart Cart) (orders []Order, err error) {

	// DEBUG
	log.Println(customerid)
	// loop all the elemets of the cart and write them to aorder
	for _, item := range cart.Items {
		uuid4 := uuid.NewV4()
		_, err = DB.Exec(`insert into orders(uuid , customerid, paymentid,
				shippingadress, shippingcontry, shippingzip ,quantity , productid) values(?,?,?,?,?,?)`, uuid4,
			customerid, cart.PaymentMethod, cart.ShippingAdress.Adress, cart.ShippingAdress.Contry, cart.ShippingAdress.Zip, item.Quantity, item.ProductID)
		if err != nil {
			return
		}

	}

	orders, err = GetUserOrders(customerid)
	return
}

func GetUserOrders(userid int64) (orders []Order, err error) {
	rows, err := DB.Query(`select id, uuid,productid ,customerid,paymentid,shippingadress, shippingcontry, shippingzip,price,quantity  from orders where customerid = ? and fulfilled=false`, userid)
	if err != nil {
		return
	}

	for rows.Next() {
		order := new(Order)

		if err = rows.Scan(&order.ID, &order.UUID, &order.ProductID, &order.CustomerID, &order.PaymentID, &order.ShippingAdress.Adress, &order.ShippingAdress.Contry, &order.ShippingAdress.Zip, &order.Price, &order.Quantity); err != nil {
			return
		}
		orders = append(orders, *order)
	}
	return

}
