package database

// Product describe a ecom
type Product struct {
	ID          int     `json:"id"`
	UUID        string  `json:"uuid"`
	Name        string  `json:"name"`
	Photoid     string  `json:"photoid"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// GetProducts ...
func GetProducts() (products []Product, err error) {
	rows, err := DB.Query("select * from  products;")
	if err != nil {
		return
	}
	for rows.Next() {
		product := Product{}
		if err = rows.Scan(&product.ID, &product.UUID, &product.Name, &product.Photoid, &product.Description, &product.Price); err != nil {
			return
		}

		products = append(products, product)
	}
	rows.Close()
	return
}

// GetProduct ...
func GetProduct(uuid string) (product Product, err error) {
	err = DB.QueryRow("select * from products where uuid = ?", uuid).Scan(&product.ID, &product.UUID, &product.Name, &product.Photoid, &product.Description, &product.Price)
	return
}

// AddProduct ....
func AddProduct(product Product) error {
	_, err := DB.Exec("insert into products(uuid, name, photoid, description, price) values(?, ?, ? , ?, ?);", product.UUID, product.Name, product.Photoid, product.Description, product.Price)

	return err
}
