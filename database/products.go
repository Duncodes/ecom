package database

import (
	"log"
	"time"

	"github.com/twinj/uuid"
)

// Product describe a ecom
type Product struct {
	ID          int64            `json:"id"`
	UUID        string           `json:"uuid"`
	Name        string           `json:"name"`
	Images      []ProductsImages `json:"images"`
	Description string           `json:"description"`

	// Unit price
	Price        float64   `json:"price"`
	ProductStock float64   `json:"stock"`
	UpdateDate   time.Time `json:"update_date"`
	//AvailableSizes  int       `json:"size"`
	//AvailableColor  string    `json:"color"`
	QuantityPerUnit int64    `json:"quantitypreunit"`
	CategoryID      int64    `json:"categoryid"`
	Category        Category `json:"category"`
}

type ProductsImages struct {
	ID        int64  `json:"id"`
	UUID      string `json:"uuid"`
	ImageUlr  string `json:"image_url"`
	ProductID int64  `json:"productid"`
}

// Category ..
type Category struct {
	ID          int64  `json:"id"`
	UUID        string `json:"uuid"`
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
	rows, err := DB.Query(`select products.id,products.uuid, .products.name, products.description,
							price,categoryid, category.id,category.name,
							category.description from  products INNER JOIN category ON products.categoryid = category.id;`)
	if err != nil {
		return
	}
	for rows.Next() {
		product := Product{}
		if err = rows.Scan(&product.ID, &product.UUID, &product.Name,
			&product.Description, &product.Price, &product.CategoryID, &product.Category.ID,
			&product.Category.Name, &product.Category.Description); err != nil {
			return
		}
		r, err := DB.Query(`select * from productimage where productid =?`, product.ID)
		defer r.Close()
		for r.Next() {
			image := ProductsImages{}
			if err = r.Scan(&image.ID, &image.UUID, &image.ImageUlr, &image.ProductID); err != nil {
				log.Println(err)
				//return
			}

			product.Images = append(product.Images, image)
		}
		products = append(products, product)
	}
	rows.Close()
	return
}

// GetProduct ...
func GetProduct(uuid string) (product Product, err error) {
	err = DB.QueryRow(`select * from products where uuid = ?`, uuid).Scan(&product.ID,
		&product.UUID, &product.Name,
		&product.Description, &product.Price)

	rows, err := DB.Query(`select * from productimage where productid =?`, product.ID)
	defer rows.Close()
	for rows.Next() {
		image := ProductsImages{}

		if err = rows.Scan(&image.ID, &image.UUID, &image.ImageUlr, &image.ProductID); err != nil {
			return
		}

		product.Images = append(product.Images, image)
	}

	return
}

// AddProduct ....
// TODO update this to user images
func AddProduct(product Product) error {
	_, err := DB.Exec(`insert into products(uuid, name, description,
						price, categoryid) values(?, ?, ? , ?, ?, ?);`, product.UUID,
		product.Name, product.Description, product.Price, 1)
	return err
}
