package database

// Item describe a ecom
type Itemstruct struct {
	Id          int     `json:"id"`
	Uuid        string  `json:"uuid"`
	Name        string  `json:"name"`
	Photoid     string  `json:"photoid"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func GetItems() (items []Item, err error) {
	rows, err := DB.Query("select * from  items;")
	if err != nil {
		return
	}
	for rows.Next() {
		item := Item{}
		if err = rows.Scan(&item.Id, &item.Uuid, &item.Name, &item.Photoid, &item.Description, &item.Price); err != nil {
			return
		}

		items = append(items, item)
	}
	rows.Close()
	return
}

func GetItem(uuid string) (item Item, err error) {
	err = DB.QueryRow("select * from items where uuid = ?", uuid).Scan(&item.Id, &item.Uuid, &item.Name, &item.Photoid, &item.Description, &item.Price)
	return
}

func AddItem(item Item) error {
	_, err := DB.Exec("insert into items(uuid, name, photoid, description, price) values(?, ?, ? , ?, ?);", item.Uuid, item.Name, item.Photoid, item.Description, item.Price)

	return err
}

type Category struct {
	ID          int
	Name        string
	Description string
}
