package database

import (
	"time"

	"github.com/twinj/uuid"
)

// Product describe a ecom
type Product struct {
	ID          int64  `json:"id"`
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Photoid     string `json:"photoid"`
	Description string `json:"description"`

	// Unit price
	Price        float64   `json:"price"`
	ProductStock float64   `json:"stock"`
	UpdateDate   time.Time `json:"update_date"`
	//AvailableSizes  int       `json:"size"`
	//AvailableColor  string    `json:"color"`
	QuantityPerUnit int64 `json:"quantitypreunit"`
	CategoryID      int64 `json:"category"`
	Category        Category
}

// Category ..
type Category struct {
	ID          int64  `json:"id"`
	UUID        int64  `json:"uuid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
}

// AddCategory ....
func (c *Category) AddCategory() (err error) {
	uuid4 := uuid.NewV4()

	_, err = DB.Exec(`insert into category(uuid, name, description, picture) values(?,?,?,?)`, uuid4, c.Name, c.Description, c.Picture)
	return
}

func GetCategories() (categories []Category, err error) {
	rows, err := DB.Query(`select id, uuid, name, description, picture from category;`)
	if err != nil {
		return
	}

	for rows.Next() {
		category := new(Category)

		if err = rows.Scan(&category.ID, &category.UUID, &category.Name, &category.Description, &category.Picture); err != nil {
			return
		}

		categories = append(categories, *category)

	}
	rows.Close()
	return
}

func GetCategoryById(id string) (category Category, err error) {
	err = DB.QueryRow(`select id, uuid, name, description, picture from  category`).Scan(&category.ID,
		&category.UUID, &category.Name, &category.Description, &category.Picture)
	return
}

// GetProducts ...
func GetProducts() (products []Product, err error) {
	rows, err := DB.Query(`select products.id,products.uuid, .products.name, photoid, products.description,
							price,categoryid, category.id,category.name,
							category.description from  products INNER JOIN category ON products.categoryid = category.id;`)
	if err != nil {
		return
	}
	for rows.Next() {
		product := Product{}
		if err = rows.Scan(&product.ID, &product.UUID, &product.Name, &product.Photoid,
			&product.Description, &product.Price, &product.CategoryID, &product.Category.ID,
			&product.Category.Name, &product.Category.Description); err != nil {
			return
		}

		products = append(products, product)
	}
	rows.Close()
	return
}

// GetProduct ...
func GetProduct(uuid string) (product Product, err error) {
	err = DB.QueryRow(`select * from products where uuid = ?`, uuid).Scan(&product.ID,
		&product.UUID, &product.Name, &product.Photoid,
		&product.Description, &product.Price)
	return
}

// AddProduct ....
func AddProduct(product Product) error {
	_, err := DB.Exec(`insert into products(uuid, name, photoid, description,
						price, categoryid) values(?, ?, ? , ?, ?, ?);`, product.UUID,
		product.Name, product.Photoid, product.Description, product.Price, 1)
	return err
}
